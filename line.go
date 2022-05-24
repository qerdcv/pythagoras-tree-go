package main

import "math"

type line struct {
	x1, y1, x2, y2 float64
}

func (l *line) move(x, y float64) {
	l.x1 += x
	l.x2 += x

	l.y1 += x
	l.y2 += x
}

func (l *line) flip(deg float64) {
	dx := l.x1 * math.Sin(deg)
	// dy := l.y1 * math.Sin(deg)

	l.x1 = l.x1 + dx
	l.x2 = l.x2 - dx

	// l.y1
}
