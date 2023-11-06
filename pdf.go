package main

import (
	"bytes"
	_ "embed"

	"github.com/signintech/gopdf"
)

//go:embed "font/RobotoMono-Regular.ttf"
var fontBytes []byte

type LabelPDF struct {
	pdf gopdf.GoPdf
}

func NewLabelPDF() (*LabelPDF, error) {
	this := &LabelPDF{}
	this.pdf = gopdf.GoPdf{}
	this.pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeLetter,
		Unit:     gopdf.UnitIN,
	})

	fontRd := bytes.NewReader(fontBytes)
	err := this.pdf.AddTTFFontByReaderWithOption("font", fontRd, gopdf.TtfOption{})
	if err != nil {
		return nil, err
	}

	err = this.pdf.SetFont("font", "", 8)
	if err != nil {
		return nil, err
	}
	this.pdf.SetFillColor(10, 255, 0)

	return this, nil
}

func (lp *LabelPDF) NewPage() {
	lp.pdf.AddPage()
}

func (lp *LabelPDF) BottomCenteredText(x, y, w, h float64, text string) error {
	lp.pdf.SetXY(x, y)
	rect := gopdf.Rect{
		W: w,
		H: h,
	}
	return lp.pdf.CellWithOption(&rect, text, gopdf.CellOption{
		Align: gopdf.Bottom | gopdf.Center,
	})
}

func (lp *LabelPDF) Save(filename string) error {
	return lp.pdf.WritePdf(filename)
}

func (lp *LabelPDF) HollowRect(x, y, w, h float64) {
	lp.pdf.RectFromUpperLeftWithStyle(x, y, w, h, "D")
}

func (lp *LabelPDF) FilledRect(x, y, w, h float64) {
	lp.pdf.RectFromUpperLeftWithStyle(x, y, w, h, "F")
}
