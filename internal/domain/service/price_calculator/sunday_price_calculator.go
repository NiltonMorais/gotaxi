package price_calculator

type SundayPriceCalculator struct {
}

func NewSundayPriceCalculator() *SundayPriceCalculator {
	return &SundayPriceCalculator{}
}

func (d *SundayPriceCalculator) Calculate(distance float64) float64 {
	return distance * 2.9
}
