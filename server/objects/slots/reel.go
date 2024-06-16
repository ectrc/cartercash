package slots

import "math/rand"

type Reel struct {
	ActiveIndex int
	Slots       []ClassicSlot
}

func NewReel(slots ...ClassicSlot) *Reel {
	return &Reel{
		Slots: slots,
	}
}

func (r *Reel) Spin() *Reel {
	rand.Shuffle(len(r.Slots), func(i, j int) {
		r.Slots[i], r.Slots[j] = r.Slots[j], r.Slots[i]
	})

	r.ActiveIndex = rand.Intn(len(r.Slots))
	return r
}

func (r *Reel) GetActiveSlot() ClassicSlot {
	return r.Slots[r.ActiveIndex]
}