# Claude Project Configuration

This file contains configuration and notes for the Claude Code CLI.

## Project Structure
This is a Go project with the following structure:
- `cmd/` - Command line applications
- `internal/` - Internal packages
- `pkg/` - Public packages
- `config/` - Configuration files

## Commands
- Build: `make build` or `go build ./...`
- Test: `make test` or `go test ./...`
- Run: `make run` or `go run ./cmd/...`

## Notes
- Project initialized with Claude Code CLI
- Uses Go modules (go.mod present)
- Standard Go project layout

## 需求
- 基于gin框架封装一个通用的layout，能够完成以下功能
- 配置文件加载 mustLoad()
- 路由中间件
- 日志追踪
- orm 操作mysql，以及redis client操作redis
- 热加载方便调试
- 快速打包交付等功能