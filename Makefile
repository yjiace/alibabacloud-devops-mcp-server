# 项目名称
PROJECT_NAME := devops-mcp-server
# 主程序入口
MAIN_FILE := main.go
# 目标输出目录
BUILD_DIR := ./bin
# 版本信息
VERSION := $(shell git describe --tags --always --dirty || echo "unknown")
# Git commit hash
COMMIT_SHA := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
# 编译时间
BUILD_TIME := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# Go ldflags选项
LD_FLAGS := -ldflags "-s -w -X main.Version=$(VERSION) -X main.CommitSha=$(COMMIT_SHA) -X main.BuildTime=$(BUILD_TIME)"

# 当前系统信息
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOENV := CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH)

# 支持的系统和架构列表
PLATFORMS := windows/amd64 windows/arm64 darwin/amd64 darwin/arm64 linux/amd64 linux/arm64

# 确保目录存在
$(shell mkdir -p $(BUILD_DIR))

# 定义颜色输出
RED=\033[0;31m
GREEN=\033[0;32m
BLUE=\033[0;34m
YELLOW=\033[0;33m
NC=\033[0m # No Color

# 默认目标，仅构建当前平台
.PHONY: default
default: build

# 清理构建产物
.PHONY: clean
clean:
	@printf "$(BLUE)Cleaning...$(NC)\n"
	@rm -rf $(BUILD_DIR)
	@go clean
	@printf "$(GREEN)Cleaned build directory$(NC)\n"

# 安装依赖
.PHONY: deps
deps:
	@printf "$(BLUE)Installing dependencies...$(NC)\n"
	@go mod tidy
	@printf "$(GREEN)Dependencies installed$(NC)\n"

# 只构建当前平台
.PHONY: build
build:
	@printf "$(BLUE)Building for current platform...$(NC)\n"
	@$(GOENV) go build $(LD_FLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)$(if $(filter windows,$(GOOS)),.exe,) $(MAIN_FILE)
	@printf "$(GREEN)Build complete: $(BUILD_DIR)/$(PROJECT_NAME)$(if $(filter windows,$(GOOS)),.exe,)$(NC)\n"

# 为所有平台构建
.PHONY: build-all
build-all: clean
	@$(MAKE) $(PLATFORMS)
	@printf "$(GREEN)All platforms built successfully!$(NC)\n"

# 为特定平台构建的模式规则
.PHONY: $(PLATFORMS)
$(PLATFORMS):
	@printf "$(BLUE)Building for $@...$(NC)\n"
	@$(eval PLATFORM_GOOS := $(word 1,$(subst /, ,$@)))
	@$(eval PLATFORM_GOARCH := $(word 2,$(subst /, ,$@)))
	@$(eval PLATFORM_SUFFIX := $(PLATFORM_GOOS)_$(PLATFORM_GOARCH)$(if $(filter windows,$(PLATFORM_GOOS)),.exe,))
	@CGO_ENABLED=0 GOOS=$(PLATFORM_GOOS) GOARCH=$(PLATFORM_GOARCH) go build $(LD_FLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_$(PLATFORM_SUFFIX) $(MAIN_FILE)
	@printf "$(GREEN)Build complete: $(BUILD_DIR)/$(PROJECT_NAME)_$(PLATFORM_SUFFIX)$(NC)\n"

# 构建Linux平台
.PHONY: build-linux
build-linux:
	@printf "$(BLUE)Building for linux...$(NC)\n"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LD_FLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_linux_amd64 $(MAIN_FILE)
	@GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build $(LD_FLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_linux_arm64 $(MAIN_FILE)
	@printf "$(GREEN)Build complete: Linux builds$(NC)\n"

# 构建Windows平台
.PHONY: build-windows
build-windows:
	@printf "$(BLUE)Building for windows...$(NC)\n"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LD_FLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_windows_amd64.exe $(MAIN_FILE)
	@GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build $(LD_FLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_windows_arm64.exe $(MAIN_FILE)
	@printf "$(GREEN)Build complete: Windows builds$(NC)\n"

# 构建macOS平台
.PHONY: build-darwin
build-darwin:
	@printf "$(BLUE)Building for macOS...$(NC)\n"
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(LD_FLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_darwin_amd64 $(MAIN_FILE)
	@GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build $(LD_FLAGS) -o $(BUILD_DIR)/$(PROJECT_NAME)_darwin_arm64 $(MAIN_FILE)
	@printf "$(GREEN)Build complete: macOS builds$(NC)\n"

# 帮助信息
.PHONY: help
help:
	@echo "$(BLUE)Available targets:$(NC)"
	@echo "  $(GREEN)build$(NC)       - Build for current platform"
	@echo "  $(GREEN)build-all$(NC)   - Build for all platforms"
	@echo "  $(GREEN)build-linux$(NC) - Build for Linux"
	@echo "  $(GREEN)build-windows$(NC) - Build for Windows"
	@echo "  $(GREEN)build-darwin$(NC) - Build for macOS"
	@echo "  $(GREEN)clean$(NC)       - Remove build artifacts"
	@echo "  $(GREEN)deps$(NC)        - Install dependencies"
	@echo "  $(GREEN)help$(NC)        - Show this help message" 