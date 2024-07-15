package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	start := time.Now()
	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, elementScreenshot("https://pkg.go.dev/", `img.Homepage-logo`, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("elementScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png, took: %s", time.Since(start).String())

	start = time.Now()

	if err := chromedp.Run(ctx, fullScreenshot(`https://brank.as/`, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote fullScreenshot.png, took: %s", time.Since(start).String())
}

func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
