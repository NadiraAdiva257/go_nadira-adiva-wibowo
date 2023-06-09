package repo

import "clean-architecture/entity/user"

func (ur *userRepo) GetAllUsers(users *[]user.User) error {
	if err := ur.db.Find(&users).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) CreateUser(user *user.User) error {
	if err := ur.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) LoginUser(user *user.User) error {
	if err := ur.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}

	return nil
}
