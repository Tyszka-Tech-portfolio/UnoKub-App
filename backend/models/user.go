package models

import "errors" // Imported errors package

type User struct {
	UserID     string `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	WalletID   string `json:"wallet_id"`
	IsVerified bool   `json:"is_verified"`
}

func NewUser(userID, username, email, password, walletID string) *User {
	return &User{
		UserID:     userID,
		Username:   username,
		Email:      email,
		Password:   password,
		WalletID:   walletID,
		IsVerified: false,
	}
}

func (u *User) VerifyUser() bool {
	// Logic to verify the user
	u.IsVerified = true
	return u.IsVerified
}

func VerifiedUser(userID string, allUsers []User) (*User, error) {
	for _, user := range allUsers {
		if user.UserID == userID {
			if user.IsVerified {
				return &user, nil
			}
			return nil, errors.New("User is not verified")
		}
	}
	return nil, errors.New("User not found")
}

func FetchUnfinishedGamesForVerifiedUser(userID string, allGames []Game, allUsers []User) ([]Game, error) {
	verifiedUser, err := VerifiedUser(userID, allUsers)
	if err != nil {
		return nil, err
	}
	return ContinueGame(verifiedUser.UserID, allGames)
}
