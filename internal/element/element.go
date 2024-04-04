package element

type Element uint8

const (
	Unknown Element = iota
	Fire
	Water
	Thunder
	Ice
	Dragon
)

func (e Element) String() string {
	switch e {
	case Fire:
		return "Fire"
	case Water:
		return "Water"
	case Thunder:
		return "Thunder"
	case Ice:
		return "Ice"
	case Dragon:
		return "Dragon"
	default:
		return "Unknown"
	}
}
