package main

import "fmt"

// Month constants
const (
	JANUARY = 1
	FREBRUARY = 2
	MARCH = 3
	APRIL = 4
	MAY = 5
	JUNE = 6
	JULY = 7
	AUGUST = 8
	SEPTEMBER = 9
	OCTOBY = 10
	NOVEMBER = 11
	DECEMBER = 12
)
// Zodiac constants
const (
	ARIES = "Aries"
	TAURIES = "Tauro"
	GEMIN = "Geminis"
	CANCER = "Cancer"
	LEO = "Leo"
	VIRG = "Virgo"
	LIBRA = "Libra"
	SCORPIO = "Escorpio"
	SAGITTAURUS = "Sagitario"
	CAPRICORN = "Capricornio"
	AQUARIUS = "Acuario"
	PISCES = "Piscis"
)

func main() {
	var day int
	var month int
	var sign string

	fmt.Scan(&day)
	fmt.Scan(&month)

	switch {
	case month == MARCH:
		if day >= 21 && day <= 31 {
			sign = ARIES
		} else if day >= 1 && day < 21{
			sign = PISCES
		} else{
			sign = "Error"
		}
	case month == APRIL:
		if day >= 1 && day <= 20 {
			sign = ARIES
		} else if day >= 21 && day <= 30 {
			sign = TAURIES
		} else{
			sign = "Error"
		}
	case month == MAY:
		if day >= 1 && day <= 21 {
			sign = TAURIES
		} else if day >= 22 && day <= 31 {
			sign = GEMIN
		} else {
			sign = "Error"
		}
	case month == JUNE:
		if day >= 1 && day <= 21 {
			sign = GEMIN
		} else if day >= 22 && day <= 30 {
			sign = CANCER
		} else {
			sign = "Error"
		}
	case month == JULY:
		if day >= 1 && day <= 22 {
			sign = CANCER
		} else if day >= 23 && day <= 31 {
			sign = LEO
		} else{
			sign = "Error"
		}
	case month == AUGUST:
		if day >= 1 && day <= 23 {
			sign = LEO
		} else if day >= 24 && day <= 31 {
			sign = VIRG
		} else{
			sign = "Error"
		}
	case month == SEPTEMBER:
		if day >= 1 && day <= 23 {
			sign = VIRG
		} else if day >= 24 && day <= 30 {
			sign = LIBRA
		} else{
			sign = "Error"
		}
	case month == OCTOBY:
		if day >= 1 && day <= 23 {
			sign = LIBRA
		} else if day >= 24 && day <= 31 {
			sign = SCORPIO
		} else {
			sign = "Error"
		}
	case month == NOVEMBER:
		if day >= 1 && day <= 22 {
			sign = SCORPIO
		} else if day >= 23 && day <= 30 {
			sign = SAGITTAURUS
		} else {
			sign = "Error"
		}
	case month == DECEMBER:
		if day >= 1 && day <= 21 {
			sign = SAGITTAURUS
		} else if day >= 22 && day <= 31 {
			sign = CAPRICORN
		} else{
			sign = "Error"
		}
	case month == JANUARY:
		if day >= 1 && day <= 20 {
			sign = CAPRICORN
		} else if day >= 21 && day <= 31 {
			sign = AQUARIUS
		} else{
			sign = "Error"
		}
	case month == FREBRUARY:
		if day >= 1 && day <= 18 {
			sign = AQUARIUS
		} else if day >= 19 && day <= 29 {
			sign = PISCES
		} else {
			sign = "Error"
		}
	default:
		sign = "Error"
	}

	fmt.Println(sign)
}