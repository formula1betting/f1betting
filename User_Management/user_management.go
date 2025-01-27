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
	Balance            float64
}

func CreateUser(ctx context.Context, conn *pgx.Conn, user User) (int64, error) {
	var userID int64
	err := conn.QueryRow(ctx, (*UserManagementQueries)["CreateUser"],
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
		user.Balance,
	).Scan(&userID)
	return userID, err
}

func GetUserByID(ctx context.Context, conn *pgx.Conn, id int64) (*User, error) {
	var user User
	err := conn.QueryRow(ctx, (*UserManagementQueries)["GetUserByID"], id).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Username, &user.PasswordHash,
		&user.DateOfBirth, &user.PhoneNumber, &user.GovernmentID, &user.Address,
		&user.TaxID, &user.AccountStatus, &user.RegistrationDate, &user.Role,
		&user.EmailVerified, &user.LastPasswordChange, &user.DeletedAt,
		&user.Country, &user.PreferredCurrency, &user.FavoriteTeam,
		&user.ProfilePictureURL, &user.Balance,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func GetUserByEmail(ctx context.Context, conn *pgx.Conn, email string) (*User, error) {
	var user User
	err := conn.QueryRow(ctx, (*UserManagementQueries)["GetUserByEmail"], email).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Username, &user.PasswordHash,
		&user.DateOfBirth, &user.PhoneNumber, &user.GovernmentID, &user.Address,
		&user.TaxID, &user.AccountStatus, &user.RegistrationDate, &user.Role,
		&user.EmailVerified, &user.LastPasswordChange, &user.DeletedAt,
		&user.Country, &user.PreferredCurrency, &user.FavoriteTeam,
		&user.ProfilePictureURL, &user.Balance,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func GetUserByUsername(ctx context.Context, conn *pgx.Conn, username string) (*User, error) {
	var user User
	err := conn.QueryRow(ctx, (*UserManagementQueries)["GetUserByUsername"], username).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Username, &user.PasswordHash,
		&user.DateOfBirth, &user.PhoneNumber, &user.GovernmentID, &user.Address,
		&user.TaxID, &user.AccountStatus, &user.RegistrationDate, &user.Role,
		&user.EmailVerified, &user.LastPasswordChange, &user.DeletedAt,
		&user.Country, &user.PreferredCurrency, &user.FavoriteTeam,
		&user.ProfilePictureURL, &user.Balance,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func UpdateUserProfile(ctx context.Context, conn *pgx.Conn, userID int64, updates User) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["UpdateUserProfile"],
		updates.FullName, updates.PhoneNumber, updates.Address,
		updates.Country, updates.PreferredCurrency, updates.FavoriteTeam,
		userID,
	)
	return err
}

func UpdateUserEmail(ctx context.Context, conn *pgx.Conn, userID int64, newEmail string) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["UpdateUserEmail"], newEmail, userID)
	return err
}

func UpdateUserPassword(ctx context.Context, conn *pgx.Conn, userID int64, newPasswordHash string) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["UpdateUserPassword"], newPasswordHash, userID)
	return err
}

func UpdateUserStatus(ctx context.Context, conn *pgx.Conn, userID int64, newStatus string) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["UpdateUserStatus"], newStatus, userID)
	return err
}

func UpdateUserBalance(ctx context.Context, conn *pgx.Conn, userID int64, newBalance int64) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["UpdateUserBalance"], newBalance, userID)
	return err
}

func DeleteUser(ctx context.Context, conn *pgx.Conn, userID int64) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["DeleteUser"], userID)
	return err
}

func UpdateUserTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["UpdateUserTable"])
	return err
}

func GetUsersByStatus(ctx context.Context, conn *pgx.Conn, status string) ([]User, error) {
	rows, err := conn.Query(ctx, (*UserManagementQueries)["GetUsersByStatus"], status)
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
			&user.ProfilePictureURL, &user.Balance,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func VerifyUserEmail(ctx context.Context, conn *pgx.Conn, userID int64) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["VerifyUserEmail"], userID)
	return err
}

func CreateUsersTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, (*UserManagementQueries)["CreateUsersTable"])
	return err
}
