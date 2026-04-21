package stitch

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const (
	// DefaultMaxInputPixels limits individual input image size to avoid OOM.
	DefaultMaxInputPixels = 80_000_000
	// DefaultMaxOutputPixels limits stitched output size.
	DefaultMaxOutputPixels = 300_000_000
)

var (
	ErrNoImages       = errors.New("no images provided")
	ErrImageTooLarge  = errors.New("image exceeds configured pixel limit")
	ErrOutputTooLarge = errors.New("stitched image exceeds configured pixel limit")
)

type Options struct {
	Background      color.Color
	CenterAlign     bool
	MaxInputPixels  int
	MaxOutputPixels int64
}

func DefaultOptions() Options {
	return Options{
		Background:      color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		CenterAlign:     true,
		MaxInputPixels:  DefaultMaxInputPixels,
		MaxOutputPixels: DefaultMaxOutputPixels,
	}
}

func LoadImages(paths []string, maxInputPixels int) ([]image.Image, error) {
	if len(paths) == 0 {
		return nil, ErrNoImages
	}

	if maxInputPixels <= 0 {
		maxInputPixels = DefaultMaxInputPixels
	}

	images := make([]image.Image, 0, len(paths))
	for _, p := range paths {
		img, err := loadSingleImage(p, maxInputPixels)
		if err != nil {
			return nil, fmt.Errorf("load %q: %w", p, err)
		}
		images = append(images, img)
	}
	return images, nil
}

func StitchVertical(images []image.Image, opts Options) (*image.NRGBA, error) {
	if len(images) == 0 {
		return nil, ErrNoImages
	}

	if opts.MaxOutputPixels <= 0 {
		opts.MaxOutputPixels = DefaultMaxOutputPixels
	}
	if opts.Background == nil {
		opts.Background = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	}

	maxWidth := 0
	totalHeight := int64(0)
	for _, img := range images {
		b := img.Bounds()
		if b.Dx() > maxWidth {
			maxWidth = b.Dx()
		}
		totalHeight += int64(b.Dy())
	}

	if maxWidth <= 0 || totalHeight <= 0 {
		return nil, errors.New("invalid image dimensions")
	}
	if int64(maxWidth)*totalHeight > opts.MaxOutputPixels {
		return nil, ErrOutputTooLarge
	}
	if totalHeight > int64(^uint(0)>>1) {
		return nil, errors.New("stitched image height overflows int")
	}

	canvas := image.NewNRGBA(image.Rect(0, 0, maxWidth, int(totalHeight)))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: opts.Background}, image.Point{}, draw.Src)

	y := 0
	for _, img := range images {
		b := img.Bounds()
		x := 0
		if opts.CenterAlign {
			x = (maxWidth - b.Dx()) / 2
		}
		dst := image.Rect(x, y, x+b.Dx(), y+b.Dy())
		draw.Draw(canvas, dst, img, b.Min, draw.Over)
		y += b.Dy()
	}

	return canvas, nil
}

func SavePNG(dstPath string, img image.Image) error {
	if img == nil {
		return errors.New("nil image")
	}
	if dstPath == "" {
		return errors.New("empty output path")
	}

	if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	file, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("encode png: %w", err)
	}
	return nil
}

func loadSingleImage(path string, maxInputPixels int) (image.Image, error) {
	cfgFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	cfg, _, cfgErr := image.DecodeConfig(cfgFile)
	_ = cfgFile.Close()
	if cfgErr != nil {
		return nil, cfgErr
	}

	if cfg.Width <= 0 || cfg.Height <= 0 {
		return nil, errors.New("invalid image dimensions")
	}
	if int64(cfg.Width)*int64(cfg.Height) > int64(maxInputPixels) {
		return nil, ErrImageTooLarge
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
