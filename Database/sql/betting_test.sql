-- Test data setup
INSERT INTO users (id, full_name, email) VALUES 
(1, 'Test User', 'test@example.com');

INSERT INTO races (id, name, start_time) VALUES 
(1, 'Test Race', NOW() + INTERVAL '1 day');

-- Test CreatePodiumBet
BEGIN;
SELECT plan(1);
SELECT lives_ok(
    $$INSERT INTO podium_bets (user_id, race_id, first_position, second_position, third_position, amount, status)
    VALUES (1, 1, 'HAM', 'VER', 'BOT', 100, 'PENDING')$$,
    'Should create podium bet'
);
SELECT * FROM finish();
ROLLBACK;

-- Test CreatePolePositionBet
BEGIN;
SELECT plan(1);
SELECT lives_ok(
    $$INSERT INTO pole_position_bets (user_id, race_id, driver_id, amount, status)
    VALUES (1, 1, 'HAM', 100, 'PENDING')$$,
    'Should create pole position bet'
);
SELECT * FROM finish();
ROLLBACK;

-- Test GetUserActiveBets
BEGIN;
SELECT plan(2);
INSERT INTO podium_bets (user_id, race_id, first_position, second_position, third_position, amount, status)
VALUES (1, 1, 'HAM', 'VER', 'BOT', 100, 'PENDING');
SELECT results_eq(
    $$SELECT COUNT(*) FROM (
        SELECT 'podium' as bet_type, * FROM podium_bets WHERE user_id = 1 AND status = 'PENDING'
    ) active_bets$$,
    $$VALUES (1::bigint)$$,
    'Should return one active bet'
);
SELECT results_eq(
    $$SELECT status FROM podium_bets WHERE user_id = 1$$,
    $$VALUES ('PENDING'::text)$$,
    'Should have correct status'
);
SELECT * FROM finish();
ROLLBACK;
