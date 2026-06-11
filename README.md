<div align="center">
  <h1>ve-blog-gin</h1>
  <p>基于 Go + Gin + GORM 的博客后端，适合新手学习</p>

  <img src="https://img.shields.io/badge/Go-1.26-blue?logo=go" />
  <img src="https://img.shields.io/badge/Gin-1.11-brightgreen?logo=go" />
  <img src="https://img.shields.io/badge/GORM-1.31-blue" />
  <img src="https://img.shields.io/badge/MySQL-8.0-orange?logo=mysql" />
  <img src="https://img.shields.io/badge/Redis-7.0-red?logo=redis" />
  <img src="https://img.shields.io/badge/License-MIT-green" />
</div>

<br/>

<div align="center">

  <a target="_blank" href="https://blog.veweiyi.cn">
    <img src="https://img.shields.io/badge/在线预览-博客前台-3498db?style=for-the-badge" />
  </a>
  &nbsp;
  <a target="_blank" href="https://admin.veweiyi.cn">
    <img src="https://img.shields.io/badge/在线预览-管理后台-e74c3c?style=for-the-badge" />
  </a>
</div>

<br/>

## 项目简介

ve-blog-gin 是 ve-blog 博客系统的 Gin 框架版后端，基于 Go + Gin + GORM + MySQL + Redis 构建。覆盖认证鉴权、内容管理、权限控制、日志审计等完整后端功能。

与 [ve-blog-golang](https://github.com/ve-weiyi/ve-blog-golang)（go-zero 微服务版）共用相同数据库和前端，适合对照学习两套框架的架构差异。

### 核心亮点

- **新手友好** — 标准三层架构（Handler → Logic → Model），完全手写，适合学习 Gin 框架
- **分层清晰** — `api/` 封装 HTTP 层，`infra/` 放基础设施，依赖方向单向流动
- **完整鉴权体系** — JWT 认证 + RBAC 角色权限 + 按钮级鉴权，中间件链清晰可追踪
- **功能完备** — 内容管理、社交互动、通知推送、日志审计、数据统计，与 go-zero 版功能对齐
- **前后端统一** — 完全兼容 ve-blog-naive 和 ve-admin-element 两个前端项目的 API 接口

## 技术栈

| 技术 | 说明 |
|------|------|
| Go 1.26 | 开发语言 |
| Gin | HTTP 框架，路由 + 中间件 |
| GORM | ORM 框架，MySQL 驱动 |
| MySQL | 关系型数据库 |
| Redis | 缓存 + Token 存储 + 限流 |
| RabbitMQ | 消息队列 |
| JWT | 身份认证 |
| Viper | 配置管理 |
| Cobra | CLI 命令行 |
| Swagger | API 文档生成 |

## 后端功能

> **当前状态**：项目已完成框架搭建（路由、中间件、配置、基础设施），业务逻辑（logic 层）尚未实现，功能模块待开发。

| 模块 | 功能 | 状态 |
|------|------|:--:|
| 认证鉴权 | 登录注册、OAuth2.0、JWT Token、验证码 | 🚧 |
| 用户管理 | 用户 CRUD、角色分配、个人信息 | 🚧 |
| 权限控制 | RBAC 角色权限、菜单管理、API 管理 | 🚧 |
| 内容管理 | 文章 CRUD、分类标签、回收站、置顶 | 🚧 |
| 评论消息 | 评论审核、留言管理 | 🚧 |
| 资源管理 | 相册管理、文件上传 | 🚧 |
| 社交管理 | 友链审核、说说管理 | 🚧 |
| 网站管理 | 站点配置、关于页面、统计看板 | 🚧 |
| 日志审计 | 登录/操作/访问/文件日志 | 🚧 |
| 消息通知 | 通知模板、消息记录、站内推送 | 🚧 |

## 项目源码

| 项目                                                               | 说明 |
|------------------------------------------------------------------|------|
| [ve-blog-golang](https://github.com/ve-weiyi/ve-blog-golang)     | go-zero 微服务版后端 |
| [ve-blog-gin](https://github.com/ve-weiyi/ve-blog-gin)           | Gin 单体版后端（本项目） |
| [ve-blog-naive](https://github.com/ve-weiyi/ve-blog-naive)       | 博客前台 |
| [ve-admin-element](https://github.com/ve-weiyi/ve-admin-element) | 博客后台管理 |

## 项目结构

```
ve-blog-gin/
├── main.go                       # 入口：加载配置 → 初始化服务 → 启动 HTTP
├── cmd/                          # Cobra CLI 命令
├── config/                       # 配置结构体 + 加载（file/nacos）
│   ├── config.go
│   ├── file.go
│   └── nacos.go
├── api/                          # HTTP API 层
│   ├── register.go               # 路由注册入口
│   ├── app/                      # 前台 API
│   │   ├── handler/              # 解析参数 → 调用 logic → 写响应
│   │   ├── logic/                # 业务逻辑编排
│   │   ├── router/               # 路由注册
│   │   └── types/                # 请求/响应类型
│   └── admin/                    # 后台 API
│       ├── handler/
│       ├── logic/
│       ├── router/
│       └── types/
├── infra/                        # 基础设施层
│   ├── biz/                      # 业务常量（错误码、请求头）
│   ├── middleware/               # 中间件（auth、cors、limit、trace、permission 等）
│   ├── request/                  # 请求工具
│   └── response/                 # 响应封装
├── initialize/                   # 初始化（GORM、Redis、RabbitMQ）
├── svctx/                        # 服务上下文（依赖注入容器）
├── config.yaml                   # 默认配置文件
├── Makefile
└── README.md
```

## 架构分层

```
main.go
  │
  ├── config.LoadConfigFromFile()   → yaml → Config
  ├── logz.SetLog()                 → 初始化 zap 日志
  │
  ├── svctx.NewServiceContext()     → 依赖注入容器
  │     ├── initialize.Gorm()       → MySQL 连接
  │     ├── initialize.Redis()      → Redis 连接
  │     ├── initialize.RabbitMQ()   → RabbitMQ 连接
  │     └── tokenstore              → JWT Token 存储
  │
  ├── gin.Default()                 → 创建 Gin 引擎
  ├── api.RegisterRouters()         → 注册路由 + 中间件
  │     ├── app/   (前台)          → Meta → UserAuth/RateLimit
  │     └── admin/ (后台)          → Meta → AdminAuth → Permission → OperationLog
  │
  └── http.Server.ListenAndServe()  → 启动 HTTP 服务
```

| 层级 | 职责 | 依赖方向 |
|------|------|:--:|
| `handler/` | 绑定请求参数、调用 logic、返回响应 | → `logic` |
| `logic/` | 业务逻辑编排、事务管理、缓存策略 | → `model` / DB |
| `middleware/` | 鉴权、权限、限流、日志 | → `svctx` |

## 环境要求

- **Go**: >= 1.26
- **MySQL**: >= 8.0
- **Redis**: >= 7.0

## 快速开始

```bash
# 1. 克隆
git clone https://github.com/ve-weiyi/ve-blog-gin.git && cd ve-blog-gin

# 2. 启动依赖服务（MySQL、Redis）
docker compose -f deploy/docker-compose/docker-compose.yaml up -d mysql redis

# 3. 修改配置
cp config.yaml config.yaml.local
# 编辑 config.yaml 中的数据库连接信息

# 4. 启动应用
go run main.go -f config.yaml
```

启动参数：

```bash
# 本地配置文件
go run main.go -f config.yaml

# Nacos 配置中心
go run main.go -nacos-host ve-weiyi.cn -nacos-port 8848 -nacos-namespace dev
```

## Docker 部署

```bash
# Docker 构建 & 运行
docker build -f deploy/docker/Dockerfile -t ve-blog-gin:latest .
docker run -d --name ve-blog-gin -p 8080:8080 -v ./config.yaml:/app/config.yaml ve-blog-gin:latest

# 或使用 Docker Compose
cd deploy/docker-compose && docker compose up -d
```

## 参与贡献

1. Fork 本仓库
2. 创建分支：`git checkout -b feature/your-feature`
3. 提交：`git commit -m 'feat: 添加某功能'`
4. 推送：`git push origin feature/your-feature`
5. 提交 Pull Request

提交规范遵循 [Conventional Commits](https://www.conventionalcommits.org/)：
`feat:` / `fix:` / `docs:` / `refactor:` / `style:` / `test:` / `chore:`

## 开源协议

MIT License — 可自由使用、修改和分发。

---

<div align="center">
  <p>如果这个项目对你有帮助，请给个 ⭐ Star 支持一下！</p>
  <p>Made with ❤️ by <a href="https://github.com/ve-weiyi">ve-weiyi</a></p>
</div>
