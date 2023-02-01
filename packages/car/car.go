package car

type Car struct {
	Name  string
	Color string
}

func (c Car) GetColor() string {
	return c.Color
}
