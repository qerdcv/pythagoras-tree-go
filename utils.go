package main

const (
	radInDeg = 0.017453
)

func degToRad(deg float64) float64 {
	return (deg + 90) * radInDeg
}
