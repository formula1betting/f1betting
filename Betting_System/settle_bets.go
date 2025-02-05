package betting_system

import (
	"context"
	"f1betting/race_info"
	"math"

	"github.com/jackc/pgx/v5"
)

// updatePodiumBetStatus updates the status of podium bets for a race to 'SETTLED'.
func updatePodiumBetStatus(ctx context.Context, tx pgx.Tx, SessionID int) error {
	_, err := tx.Exec(ctx, (*bettingQueries)["UpdatePodiumBetStatusBySession"], SessionID)
	return err
}

// updatePolePositionBetStatus updates the status of pole position bets for a race to 'SETTLED'.
func updatePolePositionBetStatus(ctx context.Context, tx pgx.Tx, SessionID int) error {
	_, err := tx.Exec(ctx, (*bettingQueries)["UpdatePolePositionBetStatusBySession"], SessionID)
	return err
}

// updateRainBetStatus updates the status of rain bets for a race to 'SETTLED'.
func updateRainBetStatus(ctx context.Context, tx pgx.Tx, SessionID int) error {
	_, err := tx.Exec(ctx, (*bettingQueries)["UpdateRainBetStatusBySession"], SessionID)
	return err
}

// updateRetirementBetStatus updates the status of retirement bets for a race to 'SETTLED'.
func updateRetirementBetStatus(ctx context.Context, tx pgx.Tx, SessionID int) error {
	_, err := tx.Exec(ctx, (*bettingQueries)["UpdateRetirementBetStatusBySession"], SessionID)
	return err
}

// updateFastestLapBetStatus updates the status of fastest lap bets for a race to 'SETTLED'.
func updateFastestLapBetStatus(ctx context.Context, tx pgx.Tx, SessionID int) error {
	_, err := tx.Exec(ctx, (*bettingQueries)["UpdateFastestLapBetStatusBySession"], SessionID)
	return err
}

// updateLapTimingBetStatus updates the status of lap timing bets for a race to 'SETTLED'.
func updateLapTimingBetStatus(ctx context.Context, tx pgx.Tx, SessionID int) error {
	_, err := tx.Exec(ctx, (*bettingQueries)["UpdateLapTimingBetStatusBySession"], SessionID)
	return err
}

// SettleBetsAndUpdateBalanceFastestLap settles bets for the fastest lap and updates the user balance.
func SettleBetsAndUpdateBalanceFastestLap(ctx context.Context, conn *pgx.Conn, SessionID int) error {

	GetFastestLapBetsByRace(ctx, conn, SessionID, "PENDING")

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	winners, err := GetFastestLapWinningBetsAndPayouts(ctx, conn, SessionID)
	if err != nil {
		return err
	}

	for _, winner := range winners {
		_, err := tx.Exec(ctx, (*bettingQueries)["AddToUserBalance"], winner.Payout, winner.UserID)
		if err != nil {
			return err
		}
	}

	if err := updateFastestLapBetStatus(ctx, tx, SessionID); err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	fastestLapBets = fastestLapBets[:0] // Clear the global variable after settling bets

	return nil
}

// GetWinningBetsAndPayouts is a placeholder function to get the winning bets and their payouts.
func GetFastestLapWinningBetsAndPayouts(ctx context.Context, conn *pgx.Conn, sessionID int) ([]Winner, error) {
	// Get the fastest lap info from the race
	fastestLap, err := race_info.GetFastestDriverWholeRace(sessionID)
	if err != nil {
		return nil, err
	}

	// Use the global variable to get pending bets
	var totalPool float64
	var winningBets []FastestLapBet
	var losingBetsAmount float64

	// Calculate total pool and separate winning/losing bets
	for _, bet := range fastestLapBets {
		if bet.SessionID == sessionID && bet.Status == "PENDING" {
			totalPool += bet.Amount
			if bet.DriverID == fastestLap.DriverNumber {
				winningBets = append(winningBets, bet)
			} else {
				losingBetsAmount += bet.Amount
			}
		}
	}

	// If no winning bets, return empty result
	if len(winningBets) == 0 {
		return []Winner{}, nil
	}

	// Calculate payout for each winner
	winners := make([]Winner, 0, len(winningBets))
	payoutPerWinner := math.Floor(losingBetsAmount / float64(len(winningBets)))

	for _, bet := range winningBets {
		winners = append(winners, Winner{
			UserID: bet.UserID,
			Payout: payoutPerWinner + bet.Amount, // Return original bet amount plus winnings
		})
	}

	return winners, nil
}

// Winner represents a winning bet and its payout.
type Winner struct {
	UserID int64
	Payout float64
}
