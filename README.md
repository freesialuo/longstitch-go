# longstitch-go

一个简单的 Go GUI 桌面工具：把多张图片按顺序竖向拼接成一张长图（PNG 输出）。

推荐仓库名：`longstitch-go`

## 功能

- 图形界面操作（Fyne）
- 支持逐张添加图片（`png/jpg/jpeg/gif`）
- 支持移除选中、清空列表
- 输出为一张竖向长图（PNG）
- 短图会在大图中自动水平居中
- 内置尺寸保护，避免超大图导致内存风险

## 运行

```bash
go mod tidy
make run
```

或：

```bash
go run ./cmd/longstitch
```

## 测试

```bash
make test
```

测试覆盖核心拼接逻辑：

- 拼接后尺寸是否正确
- 居中填充是否正确
- 空输入错误
- 输出像素上限错误
- 加载图片与异常文件处理

## 本地构建

```bash
make build
```

## CI/CD 说明

- `CI` 工作流：
  - 触发：`push main`、`pull_request`
  - 平台：`ubuntu-latest` / `macos-latest` / `windows-latest`
  - 执行：`go test ./internal/...`

- `Release` 工作流：
  - 触发：推送 tag（例如 `v0.1.0`）
  - 构建：`macos-arm64` 和 `windows-amd64.exe`
  - 发布：自动上传到 GitHub Release（`softprops/action-gh-release`）

发布示例：

```bash
git tag v0.1.0
git push origin v0.1.0
```

## 项目结构

```text
.
├── cmd/longstitch/main.go          # GUI 入口
├── internal/stitch/stitch.go       # 图像加载、拼接、导出
├── internal/stitch/stitch_test.go  # 核心单测
├── .github/workflows/ci.yml        # 持续集成
├── .github/workflows/release.yml   # 构建与发布
└── Makefile
```

## 实现依据（文档参考）

- Fyne 基础 App 与对话框 API：
  - https://github.com/fyne-io/docs.fyne.io/blob/main/content/docs/api/v2/fyne/pkg.md
  - https://github.com/fyne-io/docs.fyne.io/blob/main/content/docs/api/v2/dialog/pkg.md
  - https://github.com/fyne-io/docs.fyne.io/blob/main/content/docs/api/v2/storage/filefilter.md
- Go 图像与绘制：
  - https://go.dev/pkg/image
  - https://go.dev/pkg/image/draw
- GitHub Actions 工件与权限：
  - https://github.com/actions/upload-artifact/blob/v4/README.md
  - https://docs.github.com/en/actions/using-workflows/storing-workflow-data-as-artifacts
  - https://docs.github.com/en/actions/using-jobs/assigning-permissions-to-jobs
- GitHub Release Action：
  - https://github.com/softprops/action-gh-release
