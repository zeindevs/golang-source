package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"time"
)

func main() {
  start := time.Now()
	imgb, err := os.Open("image.jpg")
	if err != nil {
		panic(err)
	}
	img, err := jpeg.Decode(imgb)
	if err != nil {
		panic(err)
	}
	defer imgb.Close()

  wmb, err := os.Open("watermark.png")
  watermark, err := png.Decode(wmb)
  defer wmb.Close()

  offset := image.Pt(20, 20)
  b := img.Bounds()
  m := image.NewRGBA(b)
  draw.Draw(m, b, img, image.Point{}, draw.Src)
  draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.Point{}, draw.Over)

  imgw, err := os.Create("watermarked.jpg")
  jpeg.Encode(imgw, m, &jpeg.Options{Quality: jpeg.DefaultQuality})
  defer imgw.Close()

  log.Println("done, took", time.Since(start).String())
}
