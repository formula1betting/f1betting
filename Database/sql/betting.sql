-- name: CreatePodiumBet
INSERT INTO podium_bets (user_id, race_id, first_position, second_position, third_position, amount, status)
VALUES ($1, $2, $3, $4, $5, $6, 'PENDING');

-- name: CreatePolePositionBet
INSERT INTO pole_position_bets (user_id, race_id, driver_id, amount, status)
VALUES ($1, $2, $3, $4, 'PENDING');

-- name: CreateRainBet
INSERT INTO rain_bets (user_id, race_id, rain_prediction, amount, status)
VALUES ($1, $2, $3, $4, 'PENDING');

-- name: CreateRetirementBet
INSERT INTO retirement_bets (user_id, race_id, driver_id, amount, status)
VALUES ($1, $2, $3, $4, 'PENDING');

-- name: CreateFastestLapBet
INSERT INTO fastest_lap_bets (user_id, race_id, driver_id, amount, status)
VALUES ($1, $2, $3, $4, 'PENDING');

-- name: CreateLapTimingBet
INSERT INTO lap_timing_bets (user_id, race_id, lap_number, driver_id, amount, status)
VALUES ($1, $2, $3, $4, $5, 'PENDING');

-- name: GetUserActiveBets
SELECT * FROM (
    SELECT 'podium' as bet_type, * FROM podium_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'pole' as bet_type, * FROM pole_position_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'rain' as bet_type, * FROM rain_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'retirement' as bet_type, * FROM retirement_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'fastest_lap' as bet_type, * FROM fastest_lap_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'lap_timing' as bet_type, * FROM lap_timing_bets WHERE user_id = $1 AND status = 'PENDING'
) active_bets;

-- name: UpdateBetStatus
UPDATE %s SET status = $1 WHERE id = $2 AND user_id = $3;

-- name: GetRaceBetsByType
SELECT * FROM %s WHERE race_id = $1 AND status = $2;

-- name: GetUserBetHistory
SELECT * FROM (
    SELECT 'podium' as bet_type, * FROM podium_bets WHERE user_id = $1
    UNION ALL
    SELECT 'pole' as bet_type, * FROM pole_position_bets WHERE user_id = $1
    UNION ALL
    SELECT 'rain' as bet_type, * FROM rain_bets WHERE user_id = $1
    UNION ALL
    SELECT 'retirement' as bet_type, * FROM retirement_bets WHERE user_id = $1
    UNION ALL
    SELECT 'fastest_lap' as bet_type, * FROM fastest_lap_bets WHERE user_id = $1
    UNION ALL
    SELECT 'lap_timing' as bet_type, * FROM lap_timing_bets WHERE user_id = $1
) bet_history
ORDER BY created_at DESC;

-- name: GetPendingBetsForRace
SELECT * FROM (
    SELECT 'podium' as bet_type, * FROM podium_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'pole' as bet_type, * FROM pole_position_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'rain' as bet_type, * FROM rain_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'retirement' as bet_type, * FROM retirement_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'fastest_lap' as bet_type, * FROM fastest_lap_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'lap_timing' as bet_type, * FROM lap_timing_bets WHERE race_id = $1 AND status = 'PENDING'
) pending_bets;
