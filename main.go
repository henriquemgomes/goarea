package goarea

import "math"

//Pi é uma proporção numérica
const Pi = 3.141516

// Circulo calcula a area do circulo
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Retangulo calcula a area do retangulo
func Retangulo(base, altura float64) float64 {
	return base * altura
}

//Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
