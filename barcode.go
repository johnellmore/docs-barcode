package main

import (
	"github.com/boombuler/barcode/code128"
)

type Barcode struct {
	bars  []BarcodeBar
	width int
}

type BarcodeBar struct {
	X float64
	Y float64
	W float64
	H float64
}

func NewBarcode(value string) (*Barcode, error) {
	bc, err := code128.Encode(value)
	if err != nil {
		return nil, err
	}
	w := float64(bc.Bounds().Dx())
	bars := make([]BarcodeBar, 0, bc.Bounds().Dx()/2)
	var thisBar *BarcodeBar = nil
	for i := 0; i < bc.Bounds().Dx(); i++ {
		v, _, _, _ := bc.At(i, 0).RGBA()
		if v == 0 {
			// black; start or continue a bar
			if thisBar == nil {
				thisBar = &BarcodeBar{
					X: float64(i) / w,
					Y: 0,
					W: 0,
					H: 1,
				}
			}
		} else {
			// white; end a bar (or ignore)
			if thisBar != nil {
				thisBar.W = (float64(i) / w) - thisBar.X
				bars = append(bars, *thisBar)
				thisBar = nil
			}
		}
	}
	// close the last bar (if any)
	if thisBar != nil {
		thisBar.W = (1.0 - thisBar.X)
		bars = append(bars, *thisBar)
	}
	return &Barcode{
		bars:  bars,
		width: bc.Bounds().Dx(),
	}, nil
}

func (bb *Barcode) Project(x, y, w, h float64) *Barcode {
	result := &Barcode{
		bars:  make([]BarcodeBar, len(bb.bars)),
		width: bb.width,
	}
	for i, bar := range bb.bars {
		result.bars[i] = BarcodeBar{
			X: x + bar.X*w,
			Y: y,
			W: bar.W * w,
			H: h,
		}
	}
	return result
}
