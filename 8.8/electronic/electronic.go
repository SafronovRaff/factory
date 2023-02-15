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

type ApplePhone struct {
	features
}

type AndroidPhone struct {
	features
}

type RadioPhone struct {
	buttonsCount int
	features
}

// ios
func (a *ApplePhone) OS() string {
	return a.os
}

func (a *ApplePhone) Brand() string {
	return a.brand
}

func (a *ApplePhone) Model() string {
	return a.model
}
func (a *ApplePhone) Type() string {
	return a.phoneType
}

// vedroid
func (an *AndroidPhone) OS() string {
	return an.os
}

func (an *AndroidPhone) Brand() string {
	return an.brand
}

func (an *AndroidPhone) Model() string {
	return an.model
}
func (an *AndroidPhone) Type() string {
	return an.phoneType
}

// tapok
func (r *RadioPhone) ButtonsCount() int {
	return r.buttonsCount
}

func (r *RadioPhone) Brand() string {
	return r.brand
}

func (r *RadioPhone) Model() string {
	return r.model
}
func (r *RadioPhone) Type() string {
	return r.phoneType
}

func NewDesignerApplePhone(b string) *ApplePhone {
	return &ApplePhone{features{
		brand:     "aplle",
		model:     b,
		phoneType: "smartphone",
		os:        "IOS",
	}}
}

func NewDesignerAndroidPhone(b, m string) *AndroidPhone {
	return &AndroidPhone{features{
		brand:     b,
		model:     m,
		phoneType: "smartphone",
		os:        "android",
	}}

}
func NewDesignerRadioPhone(b, m string) *RadioPhone {
	return &RadioPhone{
		buttonsCount: 15,
		features: features{
			brand:     b,
			model:     m,
			phoneType: "radioPhone",
		},
	}

}
