package main

import (
	"flag"
	"fmt"
)

var prefix = flag.String("prefix", "not set", "barcode prefix")
var next = flag.Int("next", -1, "next barcode number")
var pages = flag.Int("pages", 1, "number of pages to generate")
var drawOutline = flag.Bool("debug-outline", false, "draw outline of label")
var output = flag.String("output", "not set", "output filename")

func main() {
	flag.Parse()
	if *prefix == "not set" {
		fmt.Println("Must specify barcode prefix")
		flag.Usage()
		return
	}
	if *next == -1 {
		fmt.Println("Must specify next barcode number")
		flag.Usage()
		return
	}

	if *output == "not set" {
		fmt.Println("Please specify an output filename")
		flag.Usage()
		return
	}

	// HARDCODED to the label sheets I have. All measurements in inches.
	page := LabelPage{
		LabelWidth:          1.5,
		LabelHeight:         0.5,
		LabelBarcodePadding: [4]float64{0.05, 0.05, 0.2, 0.05},
		Rows:                20,
		Columns:             4,
		Margins:             [4]float64{0.5, 0.5, 0.5, 0.5},
		PageWidth:           8.5,
		PageHeight:          11,
	}

	lp, err := NewLabelPDF()
	if err != nil {
		panic(err)
	}

	labels, err := page.LabelPositions()
	if err != nil {
		panic(err)
	}

	nextBarcodeNum := *next
	for pageNum := 0; pageNum < *pages; pageNum++ {
		lp.NewPage()
		for _, label := range labels {
			if *drawOutline {
				lp.HollowRect(label.X, label.Y, label.W, label.H)
			}

			barcodeVal := fmt.Sprintf("%s%06d", *prefix, nextBarcodeNum)
			if err != nil {
				panic(err)
			}

			err := lp.BottomCenteredText(label.X, label.Y, label.W, label.H, barcodeVal)
			if err != nil {
				panic(err)
			}

			bcX := label.X + page.LabelBarcodePadding[3]
			bcY := label.Y + page.LabelBarcodePadding[0]
			bcW := label.W - page.LabelBarcodePadding[1] - page.LabelBarcodePadding[3]
			bcH := label.H - page.LabelBarcodePadding[0] - page.LabelBarcodePadding[2]

			unitBc, err := NewBarcode(barcodeVal)
			bc := unitBc.Project(bcX, bcY, bcW, bcH)
			for _, bar := range bc.bars {
				lp.FilledRect(bar.X, bar.Y, bar.W, bar.H)
			}

			nextBarcodeNum++
		}
	}

	lp.Save(*output)
}
