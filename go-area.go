package go-area

import (
	"math"
)

// PI é uma constante matemática
const PI = 3.1416

// Circ calcula a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

// Rect calcula a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// _TrianguloEq não é visível fora do pacote
func _TrianguloEq(base, altura float64) float64 {
	return Rect(base, altura) / 2
}
