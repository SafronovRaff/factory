package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}
type StationPhone interface {
	ButtonsCount() int
}
type Smartphone interface {
	OS() string
}
type features struct {
	brand     string
	model     string
	phoneType string
	os        string
}

type applePhone struct {
	features
}

type androidPhone struct {
	features
}

type radioPhone struct {
	buttonsCount int
	features
}

// ios
func (a *applePhone) OS() string {
	return a.os
}

func (a *applePhone) Brand() string {
	return a.brand
}

func (a *applePhone) Model() string {
	return a.model
}
func (a *applePhone) Type() string {
	return a.phoneType
}

// vedroid
func (an *androidPhone) OS() string {
	return an.os
}

func (an *androidPhone) Brand() string {
	return an.brand
}

func (an *androidPhone) Model() string {
	return an.model
}
func (an *androidPhone) Type() string {
	return an.phoneType
}

// tapok
func (r *radioPhone) ButtonsCount() int {
	return r.buttonsCount
}

func (r *radioPhone) Brand() string {
	return r.brand
}

func (r *radioPhone) Model() string {
	return r.model
}
func (r *radioPhone) Type() string {
	return r.phoneType
}

func DesignerApplePhone(b string) *applePhone {
	return &applePhone{features{
		brand:     "aplle",
		model:     b,
		phoneType: "smartphone",
		os:        "IOS",
	}}
}

func DesignerAndroidPhone(b, m string) *androidPhone {
	return &androidPhone{features{
		brand:     b,
		model:     m,
		phoneType: "smartphone",
		os:        "android",
	}}

}
func DesignerRadioPhone(b, m string) *radioPhone {
	return &radioPhone{
		buttonsCount: 15,
		features: features{
			brand:     b,
			model:     m,
			phoneType: "radioPhone",
		},
	}

}
