package repo

import "github.com/ratheeshkumar25/opti_cut_userservice/pkg/model"

// CreateAddrees implements interfaces.UserRepoInter.
func (u *UserRepository) CreateAddress(address *model.Address) error {
	if err := u.DB.Create(&address).Error; err != nil {
		return err
	}
	return nil
}

// EditAddress implements interfaces.UserRepoInter.
func (u *UserRepository) EditAddress(address *model.Address) error {
	if err := u.DB.Save(&address).Error; err != nil {
		return err
	}
	return nil
}

// FindAddress implements interfaces.UserRepoInter.
func (u *UserRepository) FindAddress(userID uint) (*model.Address, error) {
	var address model.Address
	if err := u.DB.Where("userid =?", userID).First(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

// Get All address from database
func (u *UserRepository) GetAllAddresses(userID uint) (*[]model.Address, error) {
	var address []model.Address
	// Filter addresses by user ID
	if err := u.DB.Where("user_id = ?", userID).Find(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil

}

// RemoveAddress implements interfaces.UserRepoInter.
func (u *UserRepository) RemoveAddress(addressID uint, userID uint) error {
	if err := u.DB.Where("id=? AND user_id = ?", addressID, userID).Delete(&model.Address{}).Error; err != nil {
		return err
	}
	return nil
}
