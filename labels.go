package main

import (
	"errors"
)

// All values in inches. Margins are top, right, bottom, left.
type LabelPage struct {
	LabelWidth          float64
	LabelHeight         float64
	LabelBarcodePadding [4]float64
	Rows                int
	Columns             int
	Margins             [4]float64
	PageWidth           float64
	PageHeight          float64
}

// Position on the page, relative to the top-right, in inches.
type LabelPosition struct {
	X float64
	Y float64
	W float64
	H float64
}

func (s LabelPage) LabelPositions() ([]LabelPosition, error) {
	availableWidth := s.PageWidth - s.Margins[3] - s.Margins[1]
	availableHeight := s.PageHeight - s.Margins[0] - s.Margins[2]

	hSpace := (availableWidth - float64(s.Columns)*s.LabelWidth) / float64(s.Columns-1)
	vSpace := (availableHeight - float64(s.Rows)*s.LabelHeight) / float64(s.Rows-1)
	if hSpace < 0 || vSpace < 0 {
		return nil, errors.New("Not enough space for labels")
	}

	labels := make([]LabelPosition, s.Rows*s.Columns)
	for row := 0; row < s.Rows; row++ {
		for col := 0; col < s.Columns; col++ {
			index := row*s.Columns + col
			labels[index].X = s.Margins[3] + float64(col)*(s.LabelWidth+hSpace)
			labels[index].Y = s.Margins[0] + float64(row)*(s.LabelHeight+vSpace)
			labels[index].W = s.LabelWidth
			labels[index].H = s.LabelHeight
		}
	}
	return labels, nil
}
