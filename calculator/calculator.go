package calculator

type calculator struct{}

func NewCalculator(expression string) Calculator {
	return &calculator{}
}

func (c *calculator) Evaluate() (float64, error) {
	return .0, nil
}
