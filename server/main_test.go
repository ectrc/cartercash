package main

import (
	"fmt"
	"testing"

	"github.com/ectrc/cartercash/games"
)

func paypotPercentage() float64 {
	var bet uint64 = 100 // Example bet amount
	var total uint64 = 20000
	var totalBets uint64 = 0
	var totalWinnings uint64 = 0
	var numGames uint64 = 100000

	for i := uint64(0); i < numGames; i++ {
		total -= bet
		totalBets += bet
		
		classic := games.NewClassicSlotMachine()
		winnings := classic.Play(bet)
		total += winnings
		totalWinnings += winnings
	}
	payoutPercentage := func (totalBets, totalWinnings uint64) float64 {
		if totalBets == 0 {
			return 0.0
		}
		return (float64(totalWinnings) / float64(totalBets)) * 100
	}(totalBets, totalWinnings)
	fmt.Printf("[CLASSIC] %.2f%%.\n", payoutPercentage)
	
	return payoutPercentage
}

func TestSingleClassicSlotMachinePayout(t *testing.T) {
	if paypotPercentage() > 85.0 {
		t.Fatalf("Payout percentage is too high. Expected less than 85%%, got %.2f%%.", paypotPercentage())
	}
}

func TestThousandsOfClassicSlotMachinePayouts(t *testing.T) {
	var num float64 = 10
	averagePayout := 0.0
	for i := 0; i < int(num); i++ {
		averagePayout += paypotPercentage()
	}
	averagePayout /= num
	fmt.Printf("[CLASSIC AVERAGE] %.2f%%.\n", averagePayout)

	if averagePayout > 85.0 {
		t.Fatalf("Average payout percentage is too high. Expected less than 85%%, got %.2f%%.", averagePayout)
	}
}