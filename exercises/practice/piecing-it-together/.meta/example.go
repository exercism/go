package piecingittogether

import (
	"errors"
	"math"
)

type PuzzleDetails struct {
	Pieces      int
	Border      int
	Inside      int
	Rows        int
	Columns     int
	AspectRatio float64
	Format      string
}

var contradictoryData = errors.New("Contradictory data")

func (p *PuzzleDetails) detailsKnown() int {
	c := 0
	if p.Pieces != 0 {
		c++
	}
	if p.Border != 0 {
		c++
	}
	if p.Inside != 0 {
		c++
	}
	if p.Rows != 0 {
		c++
	}
	if p.Columns != 0 {
		c++
	}
	if p.AspectRatio != 0 {
		c++
	}
	if p.Format != "" {
		c++
	}
	return c
}

func (p *PuzzleDetails) setPiecesFromRowsColumns() error {
	if p.Rows != 0 && p.Columns != 0 {
		d := p.Rows * p.Columns
		if p.Pieces != 0 && p.Pieces != d {
			return contradictoryData
		}
		p.Pieces = d
	}
	return nil
}

func (p *PuzzleDetails) setRowsFromPiecesColumns() error {
	if p.Pieces != 0 && p.Columns != 0 {
		r := p.Pieces / p.Columns
		if p.Rows != 0 && p.Rows != r {
			return contradictoryData
		}
		p.Rows = r
	}
	return nil
}

func (p *PuzzleDetails) setColumsFromPiecesRows() error {
	if p.Pieces != 0 && p.Rows != 0 {
		d := p.Pieces / p.Rows
		if p.Columns != 0 && p.Columns != d {
			return contradictoryData
		}
		p.Columns = d
	}
	return nil
}

func (p *PuzzleDetails) setSquareRatio() error {
	if p.Format == "square" {
		d := 1.0
		if p.AspectRatio != 0 && p.AspectRatio != d {
			return contradictoryData
		}
		if p.Rows != 0 && p.Columns != 0 && p.Rows != p.Columns {
			return contradictoryData
		}
		p.AspectRatio = d
	}
	return nil
}

func (p *PuzzleDetails) setFormatFromRatio() error {
	if p.AspectRatio != 0 {
		if p.AspectRatio < 1 {
			d := "portrait"
			if p.Format != "" && p.Format != d {
				return contradictoryData
			}
			p.Format = d
		} else if p.AspectRatio > 1 {
			d := "landscape"
			if p.Format != "" && p.Format != d {
				return contradictoryData
			}
			p.Format = d
		} else {
			d := "square"
			if p.Format != "" && p.Format != d {
				return contradictoryData
			}
			p.Format = d
		}
	}
	return nil
}

func (p *PuzzleDetails) setDetailsForSquare() error {
	// Square: can tell values from one of rows, columns, pieces
	if p.AspectRatio != 1 {
		return nil
	}
	if p.Columns != 0 {
		d := p.Columns
		if p.Rows != 0 && p.Rows != d {
			return contradictoryData
		}
		p.Rows = d
	}
	if p.Rows != 0 {
		d := p.Rows
		if p.Columns != 0 && p.Columns != d {
			return contradictoryData
		}
		p.Columns = d
		d = p.Rows * p.Columns
		if p.Pieces != 0 && p.Pieces != d {
			return contradictoryData
		}
		p.Pieces = d
	}
	if p.Pieces != 0 {
		d := int(math.Sqrt(float64(p.Pieces)))
		if p.Rows != 0 && p.Rows != d {
			return contradictoryData
		}
		p.Rows = d
		d = int(math.Sqrt(float64(p.Pieces)))
		if p.Columns != 0 && p.Columns != d {
			return contradictoryData
		}
		p.Columns = d
	}
	if p.Inside != 0 && p.Rows == 0 {
		for i := 1; i <= p.Inside; i++ {
			pieces := i * i
			border := (i + i - 2) * 2
			if p.Inside == pieces-border {
				d := i
				if p.Rows != 0 && p.Rows != d {
					return contradictoryData
				}
				p.Rows = d
				d = i
				if p.Columns != 0 && p.Columns != d {
					return contradictoryData
				}
				p.Columns = d
			}
		}
	}
	return nil
}

func (p *PuzzleDetails) setDetailsFromAspect() error {
	if p.AspectRatio == 0 {
		return nil
	}
	if p.Pieces != 0 {
		for i := 1; i <= p.Pieces; i++ {
			if p.Pieces%i == 0 && p.AspectRatio == float64(i)/float64(p.Pieces/i) {
				d := i
				if p.Columns != 0 && p.Columns != d {
					return contradictoryData
				}
				p.Columns = d
				d = p.Pieces / i
				if p.Rows != 0 && p.Rows != d {
					return contradictoryData
				}
				p.Rows = d
				break
			}
		}
	}
	if p.Columns != 0 {
		d := int(float64(p.Columns) / p.AspectRatio)
		if p.Rows != 0 && p.Rows != d {
			return contradictoryData
		}
		p.Rows = d
	}
	if p.Rows != 0 {
		d := int(float64(p.Rows) * p.AspectRatio)
		if p.Columns != 0 && p.Columns != d {
			return contradictoryData
		}
		p.Columns = d
	}
	return nil
}

func (p *PuzzleDetails) setAspect() error {
	if p.Rows != 0 && p.Columns != 0 {
		d := float64(p.Columns) / float64(p.Rows)
		if p.AspectRatio != 0 && p.AspectRatio != d {
			return contradictoryData
		}
		p.AspectRatio = d
	}
	return nil
}

func (p *PuzzleDetails) setBorderAndInside() error {
	if p.Rows != 0 && p.Columns != 0 {
		d := (p.Rows + p.Columns - 2) * 2
		if p.Border != 0 && p.Border != d {
			return contradictoryData
		}
		p.Border = d
	}
	if p.Pieces != 0 && p.Border != 0 {
		d := p.Pieces - p.Border
		if p.Inside != 0 && p.Inside != d {
			return contradictoryData
		}
		p.Inside = d
	}
	if p.Pieces != 0 && p.Inside != 0 {
		d := p.Pieces - p.Inside
		if p.Border != 0 && p.Border != d {
			return contradictoryData
		}
		p.Border = d
	}
	if p.Border != 0 && p.Inside != 0 {
		d := p.Border + p.Inside
		if p.Pieces != 0 && p.Pieces != d {
			return contradictoryData
		}
		p.Pieces = d
	}
	return nil
}

func (p *PuzzleDetails) setRowsColumnsFromBorder() error {
	if p.Pieces != 0 && p.Border != 0 && p.Format == "portrait" {
		for c := 1; c <= p.Pieces; c++ {
			r := p.Pieces / c
			if (r+c-2)*2 == p.Border && c < r {
				d := c
				if p.Columns != 0 && p.Columns != d {
					return contradictoryData
				}
				p.Columns = d
				break
			}
		}
	}
	if p.Pieces != 0 && p.Border != 0 && p.Format == "landscape" {
		for c := 1; c <= p.Pieces; c++ {
			r := p.Pieces / c
			if (r+c-2)*2 == p.Border && c > r {
				d := c
				if p.Columns != 0 && p.Columns != d {
					return contradictoryData
				}
				p.Columns = d
				break
			}
		}
	}
	return nil
}

func (p *PuzzleDetails) updateDetails() error {
	var err error
	err = errors.Join(err, p.setPiecesFromRowsColumns())
	err = errors.Join(err, p.setRowsFromPiecesColumns())
	err = errors.Join(err, p.setColumsFromPiecesRows())
	err = errors.Join(err, p.setSquareRatio())
	err = errors.Join(err, p.setFormatFromRatio())
	err = errors.Join(err, p.setDetailsForSquare())
	err = errors.Join(err, p.setDetailsFromAspect())
	err = errors.Join(err, p.setAspect())
	err = errors.Join(err, p.setBorderAndInside())
	err = errors.Join(err, p.setRowsColumnsFromBorder())
	return err
}

func JigsawData(details PuzzleDetails) (PuzzleDetails, error) {
	p := &details
	for p.detailsKnown() != 7 {
		had := p.detailsKnown()
		if err := p.updateDetails(); err != nil {
			return details, err
		}
		has := p.detailsKnown()
		if had == has {
			return details, errors.New("Insufficient data")
		}
	}
	return details, nil
}
