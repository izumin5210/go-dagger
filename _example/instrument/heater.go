package instrument

type Heater interface {
	On()
	Off()
	IsHot() bool
}
