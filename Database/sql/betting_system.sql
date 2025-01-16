-- name: CreatePodiumBetsTable
CREATE TABLE IF NOT EXISTS podium_bets (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    race_id INT NOT NULL,
    first_position SMALLINT NOT NULL,
    second_position SMALLINT NOT NULL,
    third_position SMALLINT NOT NULL,
    amount FLOAT NOT NULL,
    status TEXT NOT NULL,
    betting_pool INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- name: CreatePolePositionBetsTable
CREATE TABLE IF NOT EXISTS pole_position_bets (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    race_id INT NOT NULL,
    driver_id SMALLINT NOT NULL,
    amount FLOAT NOT NULL,
    status TEXT NOT NULL,
    betting_pool INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- name: CreateRainBetsTable
CREATE TABLE IF NOT EXISTS rain_bets (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    race_id INT NOT NULL,
    rain_prediction BOOLEAN NOT NULL,
    amount FLOAT NOT NULL,
    status TEXT NOT NULL,
    betting_pool INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- name: CreateRetirementBetsTable
CREATE TABLE IF NOT EXISTS retirement_bets (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    race_id INT NOT NULL,
    driver_id SMALLINT NOT NULL,
    amount FLOAT NOT NULL,
    status TEXT NOT NULL,
    betting_pool INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- name: CreateFastestLapBetsTable
CREATE TABLE IF NOT EXISTS fastest_lap_bets (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    race_id INT NOT NULL,
    driver_id SMALLINT NOT NULL,
    amount FLOAT NOT NULL,
    status TEXT NOT NULL,
    betting_pool INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- name: CreateLapTimingBetsTable
CREATE TABLE IF NOT EXISTS lap_timing_bets (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    race_id INT NOT NULL,
    lap_number INT NOT NULL,
    driver_id SMALLINT NOT NULL,
    amount FLOAT NOT NULL,
    status TEXT NOT NULL,
    betting_pool INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- name: CreatePodiumBet
INSERT INTO podium_bets (user_id, race_id, first_position, second_position, third_position, amount, status, betting_pool)
VALUES ($1, $2, $3, $4, $5, $6, 'PENDING', $7)
RETURNING id;

-- name: CreatePolePositionBet
INSERT INTO pole_position_bets (user_id, race_id, driver_id, amount, status, betting_pool)
VALUES ($1, $2, $3, $4, 'PENDING', $5)
RETURNING id;

-- name: CreateRainBet
INSERT INTO rain_bets (user_id, race_id, rain_prediction, amount, status, betting_pool)
VALUES ($1, $2, $3, $4, 'PENDING', $5)
RETURNING id;

-- name: CreateRetirementBet
INSERT INTO retirement_bets (user_id, race_id, driver_id, amount, status, betting_pool)
VALUES ($1, $2, $3, $4, 'PENDING', $5)
RETURNING id;

-- name: CreateFastestLapBet
INSERT INTO fastest_lap_bets (user_id, race_id, driver_id, amount, status, betting_pool)
VALUES ($1, $2, $3, $4, 'PENDING', $5)
RETURNING id;

-- name: CreateLapTimingBet
INSERT INTO lap_timing_bets (user_id, race_id, lap_number, driver_id, amount, status, betting_pool)
VALUES ($1, $2, $3, $4, $5, 'PENDING', $6)
RETURNING id;

-- name: GetUserActiveBets
SELECT * FROM (
    SELECT 'podium' as bet_type, id, user_id, race_id, first_position, second_position, third_position, amount, status, betting_pool, created_at FROM podium_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'pole' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM pole_position_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'rain' as bet_type, id, user_id, race_id, rain_prediction, amount, status, betting_pool, created_at FROM rain_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'retirement' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM retirement_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'fastest_lap' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM fastest_lap_bets WHERE user_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'lap_timing' as bet_type, id, user_id, race_id, lap_number, driver_id, amount, status, betting_pool, created_at FROM lap_timing_bets WHERE user_id = $1 AND status = 'PENDING'
) active_bets;

-- name: UpdatePodiumBetStatus
UPDATE podium_bets SET status = $1 WHERE id = $2 AND user_id = $3
RETURNING id;

-- name: UpdatePolePositionBetStatus
UPDATE pole_position_bets SET status = $1 WHERE id = $2 AND user_id = $3
RETURNING id;

-- name: UpdateRainBetStatus
UPDATE rain_bets SET status = $1 WHERE id = $2 AND user_id = $3
RETURNING id;

-- name: UpdateRetirementBetStatus
UPDATE retirement_bets SET status = $1 WHERE id = $2 AND user_id = $3
RETURNING id;

-- name: UpdateFastestLapBetStatus
UPDATE fastest_lap_bets SET status = $1 WHERE id = $2 AND user_id = $3
RETURNING id;

-- name: UpdateLapTimingBetStatus
UPDATE lap_timing_bets SET status = $1 WHERE id = $2 AND user_id = $3
RETURNING id;

-- name: GetRacePodiumBets
SELECT id, user_id, race_id, first_position, second_position, third_position, amount, status, betting_pool, created_at FROM podium_bets WHERE race_id = $1 AND status = $2;

-- name: GetRacePolePositionBets
SELECT id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM pole_position_bets WHERE race_id = $1 AND status = $2;

-- name: GetRaceRainBets
SELECT id, user_id, race_id, rain_prediction, amount, status, betting_pool, created_at FROM rain_bets WHERE race_id = $1 AND status = $2;

-- name: GetRaceRetirementBets
SELECT id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM retirement_bets WHERE race_id = $1 AND status = $2;

-- name: GetRaceFastestLapBets
SELECT id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM fastest_lap_bets WHERE race_id = $1 AND status = $2;

-- name: GetRaceLapTimingBets
SELECT id, user_id, race_id, lap_number, driver_id, amount, status, betting_pool, created_at FROM lap_timing_bets WHERE race_id = $1 AND status = $2;

-- name: GetUserBetHistory
SELECT * FROM (
    SELECT 'podium' as bet_type, id, user_id, race_id, first_position, second_position, third_position, amount, status, betting_pool, created_at FROM podium_bets WHERE user_id = $1
    UNION ALL
    SELECT 'pole' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM pole_position_bets WHERE user_id = $1
    UNION ALL
    SELECT 'rain' as bet_type, id, user_id, race_id, rain_prediction, amount, status, betting_pool, created_at FROM rain_bets WHERE user_id = $1
    UNION ALL
    SELECT 'retirement' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM retirement_bets WHERE user_id = $1
    UNION ALL
    SELECT 'fastest_lap' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM fastest_lap_bets WHERE user_id = $1
    UNION ALL
    SELECT 'lap_timing' as bet_type, id, user_id, race_id, lap_number, driver_id, amount, status, betting_pool, created_at FROM lap_timing_bets WHERE user_id = $1
) bet_history
ORDER BY created_at DESC;

-- name: GetPendingBetsForRace
SELECT * FROM (
    SELECT 'podium' as bet_type, id, user_id, race_id, first_position, second_position, third_position, amount, status, betting_pool, created_at FROM podium_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'pole' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM pole_position_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'rain' as bet_type, id, user_id, race_id, rain_prediction, amount, status, betting_pool, created_at FROM rain_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'retirement' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM retirement_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'fastest_lap' as bet_type, id, user_id, race_id, driver_id, amount, status, betting_pool, created_at FROM fastest_lap_bets WHERE race_id = $1 AND status = 'PENDING'
    UNION ALL
    SELECT 'lap_timing' as bet_type, id, user_id, race_id, lap_number, driver_id, amount, status, betting_pool, created_at FROM lap_timing_bets WHERE race_id = $1 AND status = 'PENDING'
) pending_bets;
