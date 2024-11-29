package main

import (
	"github.com/your-project/car"
)

// Builder interface defines the methods for constructing a car
type Builder interface {
	Paint(color car.Color) Builder
	Wheels(wheels car.Wheels) Builder
	TopSpeed(speed car.Speed) Builder
	Build() Interface
}

type builder struct {
	color  car.Color
	wheels car.Wheels
	speed  car.Speed
}

func (b *builder) Paint(color car.Color) Builder {
	b.color = color
	return b
}

func (b *builder) Wheels(wheels car.Wheels) Builder {
	b.wheels = wheels
	return b
}

func (b *builder) TopSpeed(speed car.Speed) Builder {
	b.speed = speed
	return b
}

func (b *builder) Build() Interface {
	return &car{
		color:  b.color,
		wheels: b.wheels,
		speed:  b.speed,
	}
}

func NewBuilder() Builder {
	return &builder{}
}

func main() {
	assembly := NewBuilder().Paint(car.RedColor)
	
	familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()
	
	sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	sportsCar.Drive()
} 