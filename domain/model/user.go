package model

import (
	"regexp"
	"time"

	"github.com/mahiro72/go-api-sample/utils"

	"github.com/google/uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.NewString())
}

func (u UserID) String() string {
	return string(u)
}

func (u UserID) IsSame(other UserID) bool {
	return u.String() == other.String()
}

func ParseUserID(userIDString string) (UserID, error) {
	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return "", ErrUserIDValidation
	}
	return UserID(userID.String()), nil
}

type User struct {
	//ユーザーを一意に識別するID
	ID UserID
	//ユーザーの名前
	Name string
	//ユーザーのパスワード
	Password string
	//ユーザーの作成日
	CreatedAt time.Time
}

func NewUser(name, password string) (User, error) {
	if !(4 <= len(name) && len(name) <= 20) {
		return User{}, ErrUserFieldsValidation
	}

	if !(4 <= len(password) && len(password) <= 20) {
		return User{}, ErrUserFieldsValidation
	}

	// 大文字小文字アルファベット, 数字のみを許可
	regexPattern := `^[A-Za-z0-9]+$`

	// nameのvalidation
	ok, err := regexp.MatchString(regexPattern, name)
	if err != nil {
		return User{}, err
	}
	if !ok {
		return User{}, ErrUserFieldsValidation
	}

	// passwordのvalidation
	ok, err = regexp.MatchString(regexPattern, password)
	if err != nil {
		return User{}, err
	}
	if !ok {
		return User{}, ErrUserFieldsValidation
	}

	// nameをsaltとして使いhash化する
	hashPwd := utils.CreateHashFromString(name + password)

	return User{
		ID:        NewUserID(),
		Name:      name,
		Password:  hashPwd,
		CreatedAt: time.Now(),
	}, nil
}
