package interfaces

import "github.com/ratheeshkumar25/opti_cut_userservice/pkg/model"

type UserRepoInter interface {
	CreateUser(user *model.User) (uint, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(userID uint) (*model.User, error)
	UpdateUser(user *model.User) error
	GetUserList() ([]*model.User, error)
	CreateAddress(address *model.Address) error
	EditAddress(address *model.Address) error
	FindAddress(userID uint) (*model.Address, error)
	GetAllAddresses(userID uint) (*[]model.Address, error)
	RemoveAddress(addressID, userID uint) error

	// FindorCreateUserByGoogleID(googleID, email, name, accessToken, refreshToken, tokenExpiry string) (*model.GoogleSignupdetailResponse, error)
	// GetUserGoogleDetailsByID(userID string) (*model.GoogleUserDetails, error)
	// UpdateUserGoogleToken(googleID, accessToken, refreshToken, tokenExpiry string) error
}
