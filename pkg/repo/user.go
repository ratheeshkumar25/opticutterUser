package repo

import (
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/model"
)

// CreateUser implements interfaces.UserRepoInter.
func (u *UserRepository) CreateUser(user *model.User) (uint, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// FindUserByEmail implements interfaces.UserRepoInter.
func (u *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User

	if err := u.DB.Model(&model.User{}).Where("email = ?", &email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByID implements interfaces.UserRepoInter.
func (u *UserRepository) FindUserByID(userID uint) (*model.User, error) {
	var user model.User

	if err := u.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser implements interfaces.UserRepoInter.

func (u *UserRepository) UpdateUser(user *model.User) error {
	if err := u.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// func (r *UserRepository) UpdateUser(user *model.User) error {
// 	// Use Save to ensure the entire record is updated
// 	return r.DB.Save(user).Error
// }

// GetuserList fetech the userlist with DB
func (u *UserRepository) GetUserList() ([]*model.User, error) {
	var user []*model.User
	if err := u.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
