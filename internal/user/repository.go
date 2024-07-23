package user

import "github.com/wasupalonely/reservapp/pkg/db"

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetByIdentifier(identifier string) (*User, error) {
	var user User
	if err := db.DB.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	return db.DB.Create(user).Error
}
