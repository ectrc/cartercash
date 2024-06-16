package slots

type PayoutRuleType int

const (
	PayoutRuleTypeThreeOfAKind PayoutRuleType = iota
	PayoutRuleTypeTwoOfAKind
	PayoutRuleTypeOneOfAKind
)

type PayoutRuleMultiplierFor int

const (
	PayoutRuleMultiplierForBet PayoutRuleMultiplierFor = iota
	PayoutRuleMultiplierForTotal
)

type SlotUnion = ClassicSlot
type PayoutRule[T SlotUnion] struct {
	Type          PayoutRuleType
	For           T
	Multiplier    float32
	MultiplierFor PayoutRuleMultiplierFor
	Priority      int
}

func NewPayoutRule[T SlotUnion](t PayoutRuleType, slot T, multiplier float32, priority int, stackable bool) *PayoutRule[T] {
	return &PayoutRule[T]{Type: t, For: slot, Multiplier: multiplier, Priority: priority, MultiplierFor: PayoutRuleMultiplierForBet}
}

func (pr *PayoutRule[T]) appliesTo(slots []T) bool {
	switch pr.Type {
	case PayoutRuleTypeThreeOfAKind:
		{
			if slots[0] == pr.For && slots[1] == pr.For && slots[2] == pr.For {
				return true
			}
		}
	case PayoutRuleTypeTwoOfAKind:
		{
			if slots[0] == pr.For && slots[1] == pr.For {
				return true
			}

			if slots[1] == pr.For && slots[2] == pr.For {
				return true
			}

			if slots[0] == pr.For && slots[2] == pr.For {
				return true
			}
		}
	}

	return false
}

func (pr *PayoutRule[T]) CalculatePayout(slots []T, total, bet uint64) uint64 {
	if !pr.appliesTo(slots) {
		return 0
	}

	var payout uint64 = 0

	if pr.MultiplierFor == PayoutRuleMultiplierForTotal {
		payout = uint64(pr.Multiplier * float32(total+bet))
	}

	if pr.MultiplierFor == PayoutRuleMultiplierForBet {
		payout = uint64(pr.Multiplier * float32(bet))
	}

	return payout
}