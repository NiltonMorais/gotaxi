package service

import (
	"time"

	"github.com/NiltonMorais/gotaxi/internal/domain/service/price_calculator"
)

type PriceCalculatorService interface {
	Calculate(distance float64) float64
}

type PriceCalculatorServiceFactory struct {
}

func NewPriceCalculatorServiceFactory() *PriceCalculatorServiceFactory {
	return &PriceCalculatorServiceFactory{}
}

func (p *PriceCalculatorServiceFactory) NewPriceCalculatorService(date time.Time) PriceCalculatorService {
	if date.Weekday() == time.Sunday {
		return price_calculator.NewSundayPriceCalculator()
	}
	if date.Hour() >= 22 || date.Hour() <= 6 {
		return price_calculator.NewNightlyPriceCalculator()
	}
	return price_calculator.NewDefaultPriceCalculator()
}
