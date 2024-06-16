package games

import (
	"github.com/ectrc/cartercash/objects/slots"
)

type ClassicSlotMachine struct {
	Reels []*slots.Reel
	Rules []*slots.PayoutRule[slots.ClassicSlot] 
}

func NewClassicSlotMachine() *ClassicSlotMachine {
	machine := &ClassicSlotMachine{
		Reels: []*slots.Reel{},
		Rules: []*slots.PayoutRule[slots.ClassicSlot]{
			slots.NewPayoutRule(slots.PayoutRuleTypeThreeOfAKind, slots.ClassicSlotCherry, 15.0, 1, true),
			slots.NewPayoutRule(slots.PayoutRuleTypeThreeOfAKind, slots.ClassicSlotSeven, 20.0, 2, true),
			slots.NewPayoutRule(slots.PayoutRuleTypeThreeOfAKind, slots.ClassicSlotBar, 30.0, 3, true),
			slots.NewPayoutRule(slots.PayoutRuleTypeThreeOfAKind, slots.ClassicSlotWatermelon, 50.0, 4, true),
			slots.NewPayoutRule(slots.PayoutRuleTypeTwoOfAKind, slots.ClassicSlotCherry, 1.0, 5, true),
			slots.NewPayoutRule(slots.PayoutRuleTypeTwoOfAKind, slots.ClassicSlotSeven, 2.0, 6, true),
			slots.NewPayoutRule(slots.PayoutRuleTypeTwoOfAKind, slots.ClassicSlotBar, 3.0, 7, true),
			slots.NewPayoutRule(slots.PayoutRuleTypeTwoOfAKind, slots.ClassicSlotWatermelon, 4.0, 8, true),
		},
	}

	for i := 0; i < 3; i++ {
		machine.Reels = append(machine.Reels, slots.NewReel(
			slots.ClassicSlotLemon, slots.ClassicSlotOrange, slots.ClassicSlotCherry, 
			slots.ClassicSlotCherry, slots.ClassicSlotCherry, slots.ClassicSlotCherry, 
			slots.ClassicSlotPlum, slots.ClassicSlotSeven, slots.ClassicSlotSeven,
			slots.ClassicSlotSeven, slots.ClassicSlotBar, slots.ClassicSlotBar, 
			slots.ClassicSlotWatermelon, slots.ClassicSlotPlum, slots.ClassicSlotLemon, 
			slots.ClassicSlotOrange, slots.ClassicSlotLemon, slots.ClassicSlotPlum, 
			slots.ClassicSlotOrange, slots.ClassicSlotCherry,
		).Spin())
	}

	return machine
}

func (m *ClassicSlotMachine) spin() {
	for _, reel := range m.Reels {
		reel.Spin()
	}
}

func (m *ClassicSlotMachine) payout(bet uint64) uint64 {
	var total uint64 = 0
	winLine := []slots.ClassicSlot{
		m.Reels[0].GetActiveSlot(),
		m.Reels[1].GetActiveSlot(),
		m.Reels[2].GetActiveSlot(),
	}

	for _, rule := range m.Rules {
		total += rule.CalculatePayout(winLine, total, bet)
	}

	return total
}

func (m *ClassicSlotMachine) Play(bet uint64) uint64 {
	m.spin()
	payout := m.payout(bet)
	return payout
}