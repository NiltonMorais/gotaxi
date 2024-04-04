package price_calculator

type NightlyPriceCalculator struct {
}

func NewNightlyPriceCalculator() *NightlyPriceCalculator {
	return &NightlyPriceCalculator{}
}

func (d *NightlyPriceCalculator) Calculate(distance float64) float64 {
	return distance * 3.9
}
