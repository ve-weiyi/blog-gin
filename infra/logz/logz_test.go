package logz

import (
	"os"
	"path/filepath"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// restoreDefaultLogger 保存并在用例结束后还原进程级 logger，避免用例间互相污染。
func restoreDefaultLogger(t *testing.T) {
	t.Helper()
	saved := defaultLogger.Load()
	t.Cleanup(func() { defaultLogger.Store(saved) })
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name    string
		level   string
		want    zapcore.Level
		wantErr bool
	}{
		{name: "debug", level: "debug", want: zapcore.DebugLevel},
		{name: "info", level: "info", want: zapcore.InfoLevel},
		{name: "warn", level: "warn", want: zapcore.WarnLevel},
		{name: "error", level: "error", want: zapcore.ErrorLevel},
		{name: "空值按 info", level: "", want: zapcore.InfoLevel},
		{name: "无法识别的取值报错", level: "verbose", wantErr: true},
		{name: "大小写敏感", level: "INFO", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLevel(tt.level)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for level %q, got nil", tt.level)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestNew(t *testing.T) {
	t.Run("console mode", func(t *testing.T) {
		logger, err := New(&Config{Level: "info", Mode: "console"})
		if err != nil {
			t.Fatalf("New: %v", err)
		}
		defer func() { _ = logger.Sync() }()
		if logger == nil {
			t.Fatal("expected non-nil logger")
		}
	})

	t.Run("console mode has no file side effect", func(t *testing.T) {
		target := filepath.Join(t.TempDir(), "nested", "app.log")

		logger, err := New(&Config{Level: "info", Mode: "console", Filename: target})
		if err != nil {
			t.Fatalf("New: %v", err)
		}
		defer func() { _ = logger.Sync() }()

		if _, err := os.Stat(filepath.Dir(target)); !os.IsNotExist(err) {
			t.Fatalf("console 模式不应创建日志目录，stat err = %v", err)
		}
	})

	t.Run("file mode creates log dir", func(t *testing.T) {
		target := filepath.Join(t.TempDir(), "nested", "app.log")

		logger, err := New(&Config{Level: "info", Mode: "file", Filename: target})
		if err != nil {
			t.Fatalf("New: %v", err)
		}
		defer func() { _ = logger.Sync() }()

		info, err := os.Stat(filepath.Dir(target))
		if err != nil {
			t.Fatalf("日志目录未创建: %v", err)
		}
		if !info.IsDir() {
			t.Fatalf("%s 不是目录", filepath.Dir(target))
		}
	})

	t.Run("combined mode", func(t *testing.T) {
		target := filepath.Join(t.TempDir(), "app.log")

		logger, err := New(&Config{Level: "debug", Mode: "console|file", Filename: target})
		if err != nil {
			t.Fatalf("New: %v", err)
		}
		defer func() { _ = logger.Sync() }()
	})

	t.Run("empty mode returns error", func(t *testing.T) {
		if _, err := New(&Config{Level: "info", Mode: ""}); err == nil {
			t.Fatal("expected error for empty mode")
		}
	})

	t.Run("unknown level returns error", func(t *testing.T) {
		if _, err := New(&Config{Level: "verbose", Mode: "console"}); err == nil {
			t.Fatal("expected error for unknown level")
		}
	})
}

func TestL(t *testing.T) {
	t.Run("unset returns disabled logger", func(t *testing.T) {
		restoreDefaultLogger(t)
		defaultLogger.Store(nil)

		logger := L()
		if logger == nil {
			t.Fatal("expected non-nil logger")
		}
		if logger.Core().Enabled(zapcore.InfoLevel) {
			t.Fatal("未 SetLogger 时应返回丢弃日志的 logger")
		}
	})

	t.Run("returns installed logger", func(t *testing.T) {
		restoreDefaultLogger(t)

		core, _ := observer.New(zapcore.InfoLevel)
		want := zap.New(core)
		SetLogger(want)

		if got := L(); got != want {
			t.Fatal("L() 应返回 SetLogger 安装的 logger")
		}
	})
}

func TestS(t *testing.T) {
	t.Run("forwards to installed logger", func(t *testing.T) {
		restoreDefaultLogger(t)

		core, logs := observer.New(zapcore.InfoLevel)
		SetLogger(zap.New(core))

		S().Infof("hello %s", "world")

		entries := logs.All()
		if len(entries) != 1 {
			t.Fatalf("日志条数 = %d, want 1", len(entries))
		}
		if entries[0].Message != "hello world" {
			t.Fatalf("message = %q, want %q", entries[0].Message, "hello world")
		}
	})

	t.Run("unset does not panic", func(t *testing.T) {
		restoreDefaultLogger(t)
		defaultLogger.Store(nil)

		S().Infof("dropped %d", 1)
		L().Info("dropped")
	})
}
