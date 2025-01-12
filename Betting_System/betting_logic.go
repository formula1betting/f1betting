package betting_system



func UpdatePodiumBetOdds(bets []PodiumBet, firstPosition string, secondPosition string, thirdPosition string, raceId int64) (float64) {
	correct := 0
	totalBets := 0

	for _, bet := range bets {
		totalBets++
		if bet.FirstPosition == firstPosition && bet.SecondPosition == secondPosition && bet.ThirdPosition == thirdPosition && bet.RaceID == raceId {
			correct++
	
	}
}

	return float64(correct)/float64(totalBets) 
}

func UpdateLapTimingBetOdds(bets []LapTimingBet, fastestDriver int64, lapNum int64, raceId int64) (float64) {
	correct := 0
	totalBets := 0

	for _, bet := range bets {
		totalBets++
		if bet.DriverID == fastestDriver && bet.LapNumber == lapNum && bet.RaceID == raceId {
			correct++
	
	}
}

	return float64(correct)/float64(totalBets) 
}

func UpdatePolePositionBetOdds(bets []PolePositionBet, poleDriver int64, raceId int64) (float64) {
	correct := 0
	totalBets := 0

	for _, bet := range bets {
		totalBets++
		if bet.DriverID == poleDriver && bet.RaceID == raceId {
			correct++
	
	}
}

	return float64(correct)/float64(totalBets) 
}

func UpdateRetirementBetOdds(bets []RetirementBet, retireDriver int64, raceId int64) (float64) {
	correct := 0
	totalBets := 0

	for _, bet := range bets {
		totalBets++
		if bet.DriverID == retireDriver && bet.RaceID == raceId {
			correct++
	
	}
}

	return float64(correct)/float64(totalBets) 
}

func UpdateRainBetOdds(bets []RainBet, raceId int64, rainOutcome bool) (float64) {
	correct := 0
	totalBets := 0

	for _, bet := range bets {
		totalBets++
		if bet.RaceID == raceId && bet.RainPrediction == rainOutcome {
			correct++	
	}
}

	return float64(correct)/float64(totalBets)
}

//TODO: Reduce the platform fees before distributing the winnings
//Check the case what if one guy bets 10000 and another bets 100 and the guy betting 10000 wins?