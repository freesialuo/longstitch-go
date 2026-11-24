# longstitch-go

[English](README.md) · **简体中文**

[![CI](https://github.com/freesialuo/longstitch-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/freesialuo/longstitch-go/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/freesialuo/longstitch-go)](https://github.com/freesialuo/longstitch-go/releases/latest)
[![License: MIT](https://img.shields.io/badge/license-MIT-green)](LICENSE)

一个桌面小工具：把多张图片按顺序竖向拼接成一张长图，输出为 PNG。用 Go 和 [Fyne](https://fyne.io/) 编写。

## 功能

- 逐张添加图片（PNG、JPEG、GIF），可移除选中项或清空列表。
- 按列表顺序从上到下拼接，较窄的图片自动水平居中。
- 结果保存为一张 PNG。
- 内置尺寸保护：单张输入图片最多 8000 万像素，输出最多 3 亿像素。超出时直接报错，避免耗尽内存。

界面文字为中文。

## 下载

macOS（Apple 芯片）和 64 位 Windows 的预编译版本见[最新发布](https://github.com/freesialuo/longstitch-go/releases/latest)。

## 从源码运行

需要 Go 1.25 或更高版本，以及 C 编译器（Fyne 使用 cgo）。

```bash
go run ./cmd/longstitch
```

也可以用 `make run`。

## 测试

```bash
make test
```

测试覆盖拼接后的尺寸与居中填充、空输入错误、输出像素上限，以及加载正常图片和不支持的文件。

## 构建

```bash
make build
```

可执行文件输出到 `bin/longstitch`。

## CI 与发布

- **CI**：每次推送到 `main` 和每个 PR，都在 Linux、macOS、Windows 上运行 `go test ./internal/...`。
- **Release**：推送形如 `v0.1.0` 的标签时触发，构建 `macos-arm64` 和 `windows-amd64.exe` 并上传到 GitHub Release。

```bash
git tag v0.1.0
git push origin v0.1.0
```

## 项目结构

```text
.
├── cmd/longstitch/main.go          # GUI 入口
├── internal/stitch/stitch.go       # 图像加载、拼接、导出
├── internal/stitch/stitch_test.go  # 核心逻辑单元测试
├── .github/workflows/ci.yml        # 持续集成
├── .github/workflows/release.yml   # 构建与发布
├── docs/REFERENCES.md              # 实现所依据的文档
└── Makefile
```

## 许可证

[MIT](LICENSE)
