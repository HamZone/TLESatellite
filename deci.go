package satellite

import "github.com/shopspring/decimal"

//Deci2

//DecimalMul2 d1 * d2
func DecimalMul2(d1, d2 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Mul(decimal.NewFromFloat(d2)).Float64()
	return res
}

//DecimalAdd2 d1 + d2
func DecimalAdd2(d1, d2 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Add(decimal.NewFromFloat(d2)).Float64()
	return res
}

//DecimalAdd2 d1 + d2 + d3
func DecimalAdd3(d1, d2, d3 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Add(decimal.NewFromFloat(d2)).
		Add(decimal.NewFromFloat(d3)).Float64()
	return res
}

//DecimalSub2 d1 - d2
func DecimalSub2(d1, d2 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Sub(decimal.NewFromFloat(d2)).Float64()
	return res
}

//DecimalSub3 d1 - d2 - d3
func DecimalSub3(d1, d2, d3 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Sub(decimal.NewFromFloat(d2)).Sub(decimal.NewFromFloat(d3)).
		Float64()
	return res
}

//DecimalSub4 d1 - d2 - d3 - d4
func DecimalSub4(d1, d2, d3, d4 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Sub(decimal.NewFromFloat(d2)).
		Sub(decimal.NewFromFloat(d3)).Sub(decimal.NewFromFloat(d4)).
		Float64()
	return res
}

//DecimalDiv2 d1 / d2
func DecimalDiv2(d1, d2 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Div(decimal.NewFromFloat(d2)).Float64()
	return res
}

//DecimalMul3 d1 * d2 * d3
func DecimalMul3(d1, d2, d3 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Mul(decimal.NewFromFloat(d2)).Mul(decimal.NewFromFloat(d3)).Float64()
	return res
}

//DecimalMul3 d1 * d2 * d3 * d4
func DecimalMul4(d1, d2, d3, d4 float64) float64 {
	res, _ := decimal.NewFromFloat(d1).Mul(decimal.NewFromFloat(d2)).
		Mul(decimal.NewFromFloat(d3)).Mul(decimal.NewFromFloat(d4)).
		Float64()
	return res
}
