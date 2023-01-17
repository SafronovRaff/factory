package Auto

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		if u.T == CM {
			value = value / 2.54
		} else {
			value = value * 2.54
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type auto struct {
	brand       string
	model       string
	dimensions  Dimensions
	maxSpeed    int
	enginePower int
}

type dimensions struct {
	length   float64
	width    float64
	height   float64
	unitType UnitType
}

func (a *auto) Brand() string {
	return a.brand
}
func (a *auto) Model() string {
	return a.model
}
func (a *auto) MaxSpeed() int {
	return a.maxSpeed
}
func (a *auto) EnginePower() int {
	return a.enginePower
}

func (d *dimensions) Length() Unit {
	return Unit{
		Value: d.length,
		T:     d.unitType,
	}
}

func (d *dimensions) Width() Unit {
	return Unit{
		Value: d.width,
		T:     d.unitType,
	}
}
func (d *dimensions) Height() Unit {
	return Unit{
		Value: d.height,
		T:     d.unitType,
	}
}

func ConstructAuto(b, m string, len, wid, hei float64, ms, en int) *auto {
	return &auto{
		brand: b,
		model: m,
		dimensions: &dimensions{
			length:   len,
			width:    wid,
			height:   hei,
			unitType: "",
		},
		maxSpeed:    ms,
		enginePower: en,
	}

}
