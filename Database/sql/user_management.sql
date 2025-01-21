-- name: CreateUsersTable
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    full_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    phone_number TEXT,
    government_id TEXT NOT NULL,
    address TEXT NOT NULL,
    tax_id TEXT,
    account_status TEXT NOT NULL DEFAULT 'ACTIVE',
    registration_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    role TEXT NOT NULL DEFAULT 'USER',
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    last_password_change TIMESTAMP,
    deleted_at TIMESTAMP,
    country TEXT,
    preferred_currency TEXT,
    favorite_team TEXT,
    profile_picture_url TEXT,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
    
    CONSTRAINT valid_account_status CHECK (account_status IN ('ACTIVE', 'SUSPENDED', 'BANNED', 'DELETED')),
    CONSTRAINT valid_role CHECK (role IN ('USER', 'ADMIN')),
    CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

-- name: CreateUser
INSERT INTO users (
    full_name,
    email,
    username,
    password_hash,
    date_of_birth,
    phone_number,
    government_id,
    address,
    tax_id,
    account_status,
    registration_date,
    role,
    email_verified,
    country,
    preferred_currency,
    favorite_team,
    profile_picture_url,
    balance
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, 
    'ACTIVE',                  -- default account status
    CURRENT_TIMESTAMP,         -- registration date
    'USER',                    -- default role
    false,                     -- email_verified
    $10, $11, $12, $13, $14
) RETURNING id;

-- name: GetUserByID
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername
SELECT * FROM users WHERE username = $1;

-- name: UpdateUserProfile
UPDATE users SET
    full_name = $1,
    phone_number = $2,
    address = $3,
    country = $4,
    preferred_currency = $5,
    favorite_team = $6
WHERE id = $7;

-- name: UpdateUserEmail
UPDATE users SET 
    email = $1,
    email_verified = false
WHERE id = $2;

-- name: UpdateUserPassword
UPDATE users SET 
    password_hash = $1,
    last_password_change = CURRENT_TIMESTAMP
WHERE id = $2;

-- name: UpdateUserStatus
UPDATE users SET 
    account_status = $1
WHERE id = $2;

-- name: UpdateUserBalance
UPDATE users SET 
    balance = $1
WHERE id = $2;

-- name: DeleteUser
UPDATE users SET 
    account_status = 'DELETED',
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetUsersByStatus
SELECT * FROM users WHERE account_status = $1;

-- name: VerifyUserEmail
UPDATE users SET 
    email_verified = true 
WHERE id = $1;

-- name: UpdateUserTable
ALTER TABLE users
	ADD COLUMN balance DECIMAL(10, 2) DEFAULT 0.00

-- name: AddToUserBalance
UPDATE users SET 
    balance = balance + $1
WHERE id = $2;