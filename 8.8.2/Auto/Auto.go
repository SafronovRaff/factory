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

type AllAuto struct {
	auto
	dimensions
}

type auto struct {
	brand       string
	model       string
	maxSpeed    int
	enginePower int
}

/*func (a *auto) Dimensions() Dimensions {

	panic("implement me")
}*/

type dimensions struct {
	length   float64
	width    float64
	height   float64
	unitType UnitType
}

func (a *AllAuto) Brand() string {
	return a.brand
}
func (a *AllAuto) Model() string {
	return a.model
}
func (a *AllAuto) MaxSpeed() int {
	return a.maxSpeed
}
func (a *AllAuto) EnginePower() int {
	return a.enginePower
}

func (a *AllAuto) Length() Unit {
	return Unit{
		Value: a.length,
		T:     a.unitType,
	}
}

func (a *AllAuto) Width() Unit {
	return Unit{
		Value: a.width,
		T:     a.unitType,
	}
}
func (a *AllAuto) Height() Unit {
	return Unit{
		Value: a.height,
		T:     a.unitType,
	}
}

func NewConstructAuto(b, m string, len, wid, hei float64, ms, en int, unitType UnitType) *AllAuto {
	return &AllAuto{
		auto{
			brand:       b,
			model:       m,
			maxSpeed:    ms,
			enginePower: en,
		},
		dimensions{
			length:   len,
			width:    wid,
			height:   hei,
			unitType: unitType,
		},
	}

}
