# longstitch-go

**English** · [简体中文](README.zh-CN.md)

[![CI](https://github.com/freesialuo/longstitch-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/freesialuo/longstitch-go/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/freesialuo/longstitch-go)](https://github.com/freesialuo/longstitch-go/releases/latest)
[![License: MIT](https://img.shields.io/badge/license-MIT-green)](LICENSE)

A small desktop app that stitches several images, top to bottom, into one long PNG. It is written in Go with [Fyne](https://fyne.io/).

## Features

- Add images one at a time (PNG, JPEG or GIF), remove the selected one, or clear the list.
- Images are stacked in list order. A narrower image is centered horizontally.
- The result is saved as a single PNG.
- Size guards stop a job that is too large: at most 80 million pixels per input image and 300 million pixels for the output. An oversized job ends with an error instead of exhausting memory.

The interface text is in Chinese.

## Download

Prebuilt binaries for macOS on Apple silicon and for 64-bit Windows are on the [latest release](https://github.com/freesialuo/longstitch-go/releases/latest).

## Run from source

You need Go 1.25 or later and a C compiler, because Fyne uses cgo.

```bash
go run ./cmd/longstitch
```

`make run` does the same.

## Test

```bash
make test
```

The tests cover the stitched image size and centered padding, the empty-input error, the output pixel limit, and loading both valid images and unsupported files.

## Build

```bash
make build
```

The binary is written to `bin/longstitch`.

## CI and releases

- **CI** runs `go test ./internal/...` on Linux, macOS and Windows for every push to `main` and every pull request.
- **Release** runs when a tag such as `v0.1.0` is pushed. It builds `macos-arm64` and `windows-amd64.exe` and attaches them to a GitHub Release.

```bash
git tag v0.1.0
git push origin v0.1.0
```

## Project layout

```text
.
├── cmd/longstitch/main.go          # GUI entry point
├── internal/stitch/stitch.go       # Image loading, stitching and PNG export
├── internal/stitch/stitch_test.go  # Unit tests for the core logic
├── .github/workflows/ci.yml        # Continuous integration
├── .github/workflows/release.yml   # Build and release
├── docs/REFERENCES.md              # Documentation the implementation follows
└── Makefile
```

## License

[MIT](LICENSE)
