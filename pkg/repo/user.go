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

	if err := u.DB.Model(&model.User{}).Where("email = ?", email).First(&user).Error; err != nil {
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

// GetuserList fetech the userlist with DB
func (u *UserRepository) GetUserList() ([]*model.User, error) {
	var user []*model.User
	if err := u.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// // FindorCreateUserByGoogleID implements interfaces.UserRepoInter.
// func (u *UserRepository) FindorCreateUserByGoogleID(googleID string, email string, name string, accessToken string, refreshToken string, tokenExpiry string) (*model.GoogleSignupdetailResponse, error) {
// 	var user *model.User
// 	if err := u.DB.Where(&model.User{GoogleID: googleID}).First(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			// Create a new user with a UUID
// 			user = &model.User{
// 				UUID:         generateShortUUID(), // UUID stored in the separate field
// 				GoogleID:     googleID,
// 				Email:        email,
// 				FirstName:    name,
// 				AccessToken:  accessToken,
// 				RefreshToken: refreshToken,
// 				TokenExpiry:  tokenExpiry,
// 			}
// 			// Save the new user record in the database
// 			if err := u.DB.Create(&user).Error; err != nil {
// 				return &model.GoogleSignupdetailResponse{}, err
// 			}
// 		} else {
// 			return &model.GoogleSignupdetailResponse{}, err
// 		}
// 	}

// 	// Return the response with no errors
// 	return &model.GoogleSignupdetailResponse{
// 		Id:       strconv.FormatUint(uint64(user.ID), 10), // Convert uint ID to string
// 		Email:    user.Email,
// 		FullName: user.FirstName,
// 		GoogleID: user.GoogleID,
// 	}, nil
// }

// // GetUserGoogleDetailsByID implements interfaces.UserRepoInter.
// func (u *UserRepository) GetUserGoogleDetailsByID(userID string) (*model.GoogleUserDetails, error) {
// 	var user model.User
// 	// Convert the string userID to uint
// 	id, err := strconv.ParseUint(userID, 10, 32)
// 	if err != nil {
// 		return nil, errors.New("invalid user ID")
// 	}

// 	// Find the user by ID
// 	if err := u.DB.First(&user, uint(id)).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, errors.New("user not found")
// 		}
// 		return nil, err
// 	}

// 	// Return Google user details
// 	return &model.GoogleUserDetails{
// 		GoogleID:     user.GoogleID,
// 		GoogleEmail:  user.Email,
// 		AccessToken:  user.AccessToken,
// 		RefreshToken: user.RefreshToken,
// 		TokenExpiry:  user.TokenExpiry,
// 	}, nil
// }

// // UpdateUserGoogleToken implements interfaces.UserRepoInter.
// func (u *UserRepository) UpdateUserGoogleToken(googleID string, accessToken string, refreshToken string, tokenExpiry string) error {
// 	var user model.User

// 	// Find the user by GoogleID
// 	if err := u.DB.Where("google_id = ?", googleID).First(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return errors.New("user not found")
// 		}
// 		return err
// 	}

// 	// Update the user's tokens and token expiry
// 	user.AccessToken = accessToken
// 	user.RefreshToken = refreshToken
// 	user.TokenExpiry = tokenExpiry

// 	// Save the changes
// 	if err := u.DB.Save(&user).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func generateShortUUID() string {
// 	uuid := uuid.New()
// 	return base64.RawURLEncoding.EncodeToString(uuid[:])
// }
