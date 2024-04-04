package price_calculator

type DefaultPriceCalculator struct {
}

func NewDefaultPriceCalculator() *DefaultPriceCalculator {
	return &DefaultPriceCalculator{}
}

func (d *DefaultPriceCalculator) Calculate(distance float64) float64 {
	return distance * 2.1
}
