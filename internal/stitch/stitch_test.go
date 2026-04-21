package stitch

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"testing"
)

func TestStitchVertical_CenterAlign(t *testing.T) {
	img1 := solidImage(image.Rect(0, 0, 4, 2), color.NRGBA{R: 255, A: 255})
	img2 := solidImage(image.Rect(0, 0, 2, 3), color.NRGBA{G: 255, A: 255})

	out, err := StitchVertical([]image.Image{img1, img2}, DefaultOptions())
	if err != nil {
		t.Fatalf("StitchVertical returned error: %v", err)
	}

	if got, want := out.Bounds().Dx(), 4; got != want {
		t.Fatalf("width = %d, want %d", got, want)
	}
	if got, want := out.Bounds().Dy(), 5; got != want {
		t.Fatalf("height = %d, want %d", got, want)
	}

	// The second image should be centered at x=1 when output width=4 and src width=2.
	if c := color.NRGBAModel.Convert(out.At(1, 2)).(color.NRGBA); c.G != 255 {
		t.Fatalf("expected green pixel at centered location, got %#v", c)
	}
	if c := color.NRGBAModel.Convert(out.At(0, 2)).(color.NRGBA); c.R != 255 || c.G != 255 || c.B != 255 {
		t.Fatalf("expected white background at left padding, got %#v", c)
	}
}

func TestStitchVertical_EmptyInput(t *testing.T) {
	_, err := StitchVertical(nil, DefaultOptions())
	if err == nil {
		t.Fatal("expected error for empty input")
	}
}

func TestStitchVertical_OutputLimit(t *testing.T) {
	img := solidImage(image.Rect(0, 0, 10, 10), color.NRGBA{B: 255, A: 255})
	opts := DefaultOptions()
	opts.MaxOutputPixels = 50

	_, err := StitchVertical([]image.Image{img}, opts)
	if err == nil {
		t.Fatal("expected output limit error")
	}
	if err != ErrOutputTooLarge {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLoadImages(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.png")
	img := solidImage(image.Rect(0, 0, 3, 3), color.NRGBA{R: 10, G: 20, B: 30, A: 255})
	if err := SavePNG(path, img); err != nil {
		t.Fatalf("SavePNG error: %v", err)
	}

	loaded, err := LoadImages([]string{path}, DefaultMaxInputPixels)
	if err != nil {
		t.Fatalf("LoadImages error: %v", err)
	}
	if len(loaded) != 1 {
		t.Fatalf("len = %d, want 1", len(loaded))
	}
	if got, want := loaded[0].Bounds().Dx(), 3; got != want {
		t.Fatalf("width = %d, want %d", got, want)
	}
}

func TestLoadImages_UnsupportedFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.txt")
	if err := os.WriteFile(path, []byte("not image"), 0o644); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	_, err := LoadImages([]string{path}, DefaultMaxInputPixels)
	if err == nil {
		t.Fatal("expected error")
	}
}

func solidImage(rect image.Rectangle, c color.NRGBA) image.Image {
	img := image.NewNRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.SetNRGBA(x, y, c)
		}
	}
	return img
}
