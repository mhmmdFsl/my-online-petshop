package model

import "time"

type User struct {
	Id            int           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string        `json:"name"`
	Status        string        `json:"status"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updated_at"`
	PrincipalUser PrincipalUser `json:"principalUser"`
	PasswordUser  PasswordUser  `json:"-"`
}

type PasswordUser struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId       int       `json:"userId"`
	HashPassword string    `json:"hashPassword"`
	CreatedAt    time.Time `json:"createdAt"`
}

type PrincipalUser struct {
	Id             int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId         int       `json:"userId"`
	PrincipalType  string    `json:"principalType"`
	PrincipalValue string    `json:"principalValue"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type UserRefreshToken struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId       int       `json:"userId"`
	RefreshToken string    `json:"refreshToken"`
	CreatedAt    time.Time `json:"createdAt"`
	ExpiredAt    time.Time `json:"expiredAt"`
}
