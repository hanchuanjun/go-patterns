package main

import "fmt"
type Speed float64

const (
    MPH Speed = 1
    KPH Speed = 1.60934
)

type Color string

const (
    BlueColor  Color = "blue"
    GreenColor Color = "green"
    RedColor   Color = "red"
)

type Wheels string

const (
    SportsWheels Wheels = "sports"
    SteelWheels   Wheels = "steel"
)

type Builder interface {
    Color(Color) Builder
    Wheels(Wheels) Builder
    TopSpeed(Speed) Builder
    Build() Interface
}

type Interface interface {
    Drive() error
    Stop() error
}

type builder struct {
    color  Color
    wheels Wheels 
    speed  Speed
}

func NewBuilder() Builder {
	return &builder{}
}

func (b *builder) Color(color Color) Builder {
	b.color = color
	return b
}
func (b *builder) Wheels(wheels Wheels) Builder {
	b.wheels = wheels
	return b
}	
func (b *builder) TopSpeed(speed Speed) Builder {
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

type car struct {
    color  Color
    wheels Wheels
    speed  Speed
}

func (c *car) Drive() error {
    return nil
}

func (c *car) Stop() error {
    return nil
}

func main() {
	assembly := NewBuilder().Color(RedColor)

	familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()
	
	sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
	fmt.Println("Sports car built successfully")
}

