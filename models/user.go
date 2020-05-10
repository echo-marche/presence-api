package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id                  int64          `db:"id, primarykey, autoincrement"`
	Name                sql.NullString `db:"name"`
	Email               string         `db:"email"`
	EncryptedPassword   string         `db:"encrypted_password"`
	RememberToken       sql.NullString `db:"remember_token"`
	RememberCreatedAt   sql.NullTime   `db:"remember_created_at"`
	LastSignInAt        sql.NullTime   `db:"last_sign_in_at"`
	ConfirmationToken   sql.NullString `db:"confirmation_token"`
	ConfirmedAt         sql.NullTime   `db:"confirmed_at"`
	ConfirmationSentAt  sql.NullTime   `db:"confirmation_sent_at"`
	ResetPasswordToken  sql.NullString `db:"reset_password_token"`
	ResetPasswordSentAt sql.NullTime   `db:"reset_password_sent_at"`
	DeletedAt           sql.NullTime   `db:"deleted_at"`
	CreatedAt           time.Time      `db:"created_at"`
	UpdatedAt           time.Time      `db:"updated_at"`
}

type Users []User
