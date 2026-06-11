package permissionx

import "strings"

func normalizeMethod(method string) string {
	method = strings.TrimSpace(method)
	if method == "" {
		return ""
	}
	method = strings.ToUpper(method)
	if method == "*" || method == "ALL" {
		return "*"
	}
	return method
}

func normalizePath(path string) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return ""
	}
	if idx := strings.IndexByte(path, '?'); idx >= 0 {
		path = path[:idx]
	}
	if path != "/" {
		path = strings.TrimSuffix(path, "/")
	}
	if path != "" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}
