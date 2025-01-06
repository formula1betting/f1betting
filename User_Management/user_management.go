package user_management

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID                 int64
	FullName           string
	Email              string
	Username           string
	PasswordHash       string
	DateOfBirth        time.Time
	PhoneNumber        *string
	GovernmentID       string
	Address            string
	TaxID              *string
	AccountStatus      string
	RegistrationDate   time.Time
	Role               string
	EmailVerified      bool
	LastPasswordChange *time.Time
	DeletedAt          *time.Time
	Country            *string
	PreferredCurrency  *string
	FavoriteTeam       *string
	ProfilePictureURL  *string
}

func CreateUser(ctx context.Context, conn *pgx.Conn, user User) (int64, error) {
	var userID int64
	err := conn.QueryRow(ctx, userManagementQueries["CreateUser"],
		user.FullName,
		user.Email,
		user.Username,
		user.PasswordHash,
		user.DateOfBirth,
		user.PhoneNumber,
		user.GovernmentID,
		user.Address,
		user.TaxID,
		user.Country,
		user.PreferredCurrency,
		user.FavoriteTeam,
		user.ProfilePictureURL,
	).Scan(&userID)
	return userID, err
}

func GetUserByID(ctx context.Context, conn *pgx.Conn, id int64) (*User, error) {
	var user User
	err := conn.QueryRow(ctx, userManagementQueries["GetUserByID"], id).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Username, &user.PasswordHash,
		&user.DateOfBirth, &user.PhoneNumber, &user.GovernmentID, &user.Address,
		&user.TaxID, &user.AccountStatus, &user.RegistrationDate, &user.Role,
		&user.EmailVerified, &user.LastPasswordChange, &user.DeletedAt,
		&user.Country, &user.PreferredCurrency, &user.FavoriteTeam,
		&user.ProfilePictureURL,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func GetUserByEmail(ctx context.Context, conn *pgx.Conn, email string) (*User, error) {
	var user User
	err := conn.QueryRow(ctx, userManagementQueries["GetUserByEmail"], email).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Username, &user.PasswordHash,
		&user.DateOfBirth, &user.PhoneNumber, &user.GovernmentID, &user.Address,
		&user.TaxID, &user.AccountStatus, &user.RegistrationDate, &user.Role,
		&user.EmailVerified, &user.LastPasswordChange, &user.DeletedAt,
		&user.Country, &user.PreferredCurrency, &user.FavoriteTeam,
		&user.ProfilePictureURL,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func GetUserByUsername(ctx context.Context, conn *pgx.Conn, username string) (*User, error) {
	var user User
	err := conn.QueryRow(ctx, userManagementQueries["GetUserByUsername"], username).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Username, &user.PasswordHash,
		&user.DateOfBirth, &user.PhoneNumber, &user.GovernmentID, &user.Address,
		&user.TaxID, &user.AccountStatus, &user.RegistrationDate, &user.Role,
		&user.EmailVerified, &user.LastPasswordChange, &user.DeletedAt,
		&user.Country, &user.PreferredCurrency, &user.FavoriteTeam,
		&user.ProfilePictureURL,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func UpdateUserProfile(ctx context.Context, conn *pgx.Conn, userID int64, updates User) error {
	_, err := conn.Exec(ctx, userManagementQueries["UpdateUserProfile"],
		updates.FullName, updates.PhoneNumber, updates.Address,
		updates.Country, updates.PreferredCurrency, updates.FavoriteTeam,
		userID,
	)
	return err
}

func UpdateUserEmail(ctx context.Context, conn *pgx.Conn, userID int64, newEmail string) error {
	_, err := conn.Exec(ctx, userManagementQueries["UpdateUserEmail"], newEmail, userID)
	return err
}

func UpdateUserPassword(ctx context.Context, conn *pgx.Conn, userID int64, newPasswordHash string) error {
	_, err := conn.Exec(ctx, userManagementQueries["UpdateUserPassword"], newPasswordHash, userID)
	return err
}

func UpdateUserStatus(ctx context.Context, conn *pgx.Conn, userID int64, newStatus string) error {
	_, err := conn.Exec(ctx, userManagementQueries["UpdateUserStatus"], newStatus, userID)
	return err
}

func DeleteUser(ctx context.Context, conn *pgx.Conn, userID int64) error {
	_, err := conn.Exec(ctx, userManagementQueries["DeleteUser"], userID)
	return err
}

func GetUsersByStatus(ctx context.Context, conn *pgx.Conn, status string) ([]User, error) {
	rows, err := conn.Query(ctx, userManagementQueries["GetUsersByStatus"], status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID, &user.FullName, &user.Email, &user.Username, &user.PasswordHash,
			&user.DateOfBirth, &user.PhoneNumber, &user.GovernmentID, &user.Address,
			&user.TaxID, &user.AccountStatus, &user.RegistrationDate, &user.Role,
			&user.EmailVerified, &user.LastPasswordChange, &user.DeletedAt,
			&user.Country, &user.PreferredCurrency, &user.FavoriteTeam,
			&user.ProfilePictureURL,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func VerifyUserEmail(ctx context.Context, conn *pgx.Conn, userID int64) error {
	_, err := conn.Exec(ctx, userManagementQueries["VerifyUserEmail"], userID)
	return err
}

func CreateUsersTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, userManagementQueries["CreateUsersTable"])
	return err
}
