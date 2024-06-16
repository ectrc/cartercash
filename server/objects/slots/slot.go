package slots

type ClassicSlot int

const (
	ClassicSlotCherry ClassicSlot = iota
	ClassicSlotSeven
	ClassicSlotBar
	ClassicSlotWatermelon
	ClassicSlotLemon
	ClassicSlotPlum
	ClassicSlotOrange
)

func (s ClassicSlot) String() string {
	return [...]string{"Cherry", "Seven", "Bar", "Watermelon", "Lemon", "Plum", "Orange"}[s]
}