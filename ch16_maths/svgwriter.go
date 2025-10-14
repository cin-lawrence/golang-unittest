package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	defaultSecondHandLength = 90.0
	defaultMinuteHandLength = 80.0
	defaultHourHandLength   = 50.0
	defaultCenterX          = 150.0
	defaultCenterY          = 150.0
)

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

type ClockHand struct {
	Second float64
	Minute float64
	Hour   float64
}

type SVGWriter struct {
	HandLength ClockHand
	Center     Point
}

func DefaultSVGWriter() *SVGWriter {
	return &SVGWriter{
		ClockHand{
			defaultSecondHandLength,
			defaultMinuteHandLength,
			defaultHourHandLength,
		},
		Point{defaultCenterX, defaultCenterY},
	}
}

func (s *SVGWriter) writeSecondHand(w io.Writer, t time.Time) {
	p := s.makeHand(secondHandPoint(t), s.HandLength.Second)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
func (s *SVGWriter) writeMinuteHand(w io.Writer, t time.Time) {
	p := s.makeHand(minuteHandPoint(t), s.HandLength.Minute)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func (s *SVGWriter) writeHourHand(w io.Writer, t time.Time) {
	p := s.makeHand(hourHandPoint(t), s.HandLength.Hour)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func (s *SVGWriter) makeHand(p Point, length float64) Point {
	return Point{
		s.Center.X + p.X*length,
		s.Center.Y - p.Y*length,
	}
}

func (s *SVGWriter) WriteSVG(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	s.writeSecondHand(w, t)
	s.writeMinuteHand(w, t)
	s.writeHourHand(w, t)
	io.WriteString(w, svgEnd)
}
