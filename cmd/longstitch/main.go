package main

import (
	"fmt"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"github.com/freesialuo/longstitch-go/internal/stitch"
)

func main() {
	application := app.NewWithID("io.github.longstitch")
	window := application.NewWindow("LongStitch - 竖向拼接长图")
	window.Resize(fyne.NewSize(820, 560))

	var imagePaths []string
	selectedIndex := -1
	outputPath := ""

	statusLabel := widget.NewLabel("请选择图片并设置输出路径。")
	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder("输出文件路径（PNG）")
	outputEntry.Disable()

	list := widget.NewList(
		func() int {
			return len(imagePaths)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(fmt.Sprintf("%d. %s", id+1, imagePaths[id]))
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		selectedIndex = id
	}
	list.OnUnselected = func(widget.ListItemID) {
		selectedIndex = -1
	}

	addImage := func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()

			path := reader.URI().Path()
			for _, existing := range imagePaths {
				if existing == path {
					statusLabel.SetText("该图片已在列表中。")
					return
				}
			}
			imagePaths = append(imagePaths, path)
			list.Refresh()
			statusLabel.SetText(fmt.Sprintf("已添加：%s", filepath.Base(path)))
		}, window)
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg", ".gif"}))
		fileDialog.Show()
	}

	chooseOutput := func() {
		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			if writer.URI() == nil {
				return
			}
			outputPath = writer.URI().Path()
			outputEntry.SetText(outputPath)
			statusLabel.SetText("已设置输出路径。")
		}, window)
		saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png"}))
		saveDialog.SetFileName("stitched-long-image.png")
		saveDialog.Show()
	}

	build := func() {
		if len(imagePaths) == 0 {
			dialog.ShowInformation("提示", "请先添加至少一张图片。", window)
			return
		}
		if outputPath == "" {
			dialog.ShowInformation("提示", "请先选择输出路径。", window)
			return
		}

		statusLabel.SetText("正在加载图片...")
		images, err := stitch.LoadImages(imagePaths, stitch.DefaultMaxInputPixels)
		if err != nil {
			dialog.ShowError(err, window)
			statusLabel.SetText("加载失败。")
			return
		}

		statusLabel.SetText("正在拼接...")
		result, err := stitch.StitchVertical(images, stitch.DefaultOptions())
		if err != nil {
			dialog.ShowError(err, window)
			statusLabel.SetText("拼接失败。")
			return
		}

		statusLabel.SetText("正在保存...")
		if err := stitch.SavePNG(outputPath, result); err != nil {
			dialog.ShowError(err, window)
			statusLabel.SetText("保存失败。")
			return
		}

		statusLabel.SetText(fmt.Sprintf("完成：%s", outputPath))
		dialog.ShowInformation("完成", "长图已生成。", window)
	}

	addBtn := widget.NewButton("添加图片", addImage)
	removeBtn := widget.NewButton("移除选中", func() {
		if selectedIndex < 0 || selectedIndex >= len(imagePaths) {
			return
		}
		removed := imagePaths[selectedIndex]
		imagePaths = append(imagePaths[:selectedIndex], imagePaths[selectedIndex+1:]...)
		selectedIndex = -1
		list.UnselectAll()
		list.Refresh()
		statusLabel.SetText(fmt.Sprintf("已移除：%s", filepath.Base(removed)))
	})
	clearBtn := widget.NewButton("清空列表", func() {
		imagePaths = nil
		selectedIndex = -1
		list.UnselectAll()
		list.Refresh()
		statusLabel.SetText("已清空图片列表。")
	})
	outputBtn := widget.NewButton("选择输出路径", chooseOutput)
	buildBtn := widget.NewButton("生成长图", build)

	controls := container.NewHBox(addBtn, removeBtn, clearBtn)
	outputRow := container.NewBorder(nil, nil, nil, outputBtn, outputEntry)
	mainContent := container.NewBorder(
		container.NewVBox(
			controls,
			widget.NewSeparator(),
		),
		container.NewVBox(
			widget.NewSeparator(),
			outputRow,
			buildBtn,
			statusLabel,
		),
		nil,
		nil,
		list,
	)

	window.SetContent(container.NewPadded(mainContent))
	window.ShowAndRun()
}
