package main

import "fmt"

var boilingPointWaterKelvin float64 =  373.15

func main() {
	boilingPointWaterCelsius := boilingPointWaterKelvin - 273.15

	fmt.Println("O valor do ponto de ebulição da água em Celsius (°C) é:",boilingPointWaterCelsius)
}