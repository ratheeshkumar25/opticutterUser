package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ratheeshkumar25/opti_cut_userservice/config"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/model"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/utils"
	"gorm.io/gorm"
)

// SignupService method receive the data in proto messages and start the verfification process.
func (u *UserService) SignupService(p *pb.Signup) (*pb.Response, error) {
	hashedPass, err := utils.HashPassword(p.Password)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in hasing password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("unable to hashpassword")
	}
	user := &model.User{
		FirstName: p.First_Name,
		LastName:  p.Last_Name,
		Phone:     p.Phone,
		Email:     p.Email,
		Password:  hashedPass,
	}
	existingUser, err := u.Repo.FindUserByEmail(user.Email)
	if err == nil || existingUser != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "email already exists",
		}, errors.New("email already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "data base error",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	resp, err := u.twilio.SendTwilioOTP(p.Phone)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in sending otp using twillio",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	if resp.Status != nil {
		fmt.Println(*resp.Status)
	} else {
		fmt.Println(resp.Status)
	}

	userData, err := json.Marshal(&user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in marshaling data",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error while marshalling data")
	}
	key := fmt.Sprintf("user_%v", p.Email)
	err = u.redis.SetDataInRedis(key, userData, time.Minute*3)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in setting data in redis",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error setting data in redis")
	}
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Go to verification page",
	}, nil

}

// VerificationService implements interfaces.UserServiceInter.
func (u *UserService) VerificationService(p *pb.OTP) (*pb.Response, error) {
	key := fmt.Sprintf("user_%v", p.Email)
	userData, err := u.redis.GetFromRedis(key)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error on getting data from redis",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	//Unmarashal the data from redis

	var user model.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in unmarshalling data",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	err = u.twilio.VerifyTwilioOTP(user.Phone, p.Otp)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in verifying twillio otp",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	userID, err := u.Repo.CreateUser(&user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in creating user in database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("unable to create user")
	}
	log.Printf("User created with ID: %v", userID)

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "OTP verifed successfully. Login to contionue",
	}, nil

}

// LoginService implements interfaces.UserServiceInter.
func (u *UserService) LoginService(p *pb.Login) (*pb.Response, error) {
	user, err := u.Repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}

	log.Println("user", user)
	if !utils.CheckPassword(p.Password, user.Password) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Password incorrect try again with newpassword",
		}, errors.New("password incorrect")
	}

	if user.IsBlocked {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "user is blocked by Admin",
		}, errors.New("you are blocked by Admin")
	}

	jwtToken, err := utils.GenerateToken(config.LoadConfig().SECERETKEY, user.Email, user.ID)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in generating token",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error in generating token")
	}
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Login successful",
		Payload: &pb.Response_Data{Data: jwtToken},
	}, nil

}

// // GoogleSignInService implements interfaces.UserServiceInter.
// func (u *UserService) GoogleSignInService(p *pb.GoogleSignInRequest) (*pb.UserSignUpResponse, error) {
// 	// Find or create user using Google ID
// 	user, err := u.Repo.FindorCreateUserByGoogleID(
// 		p.GoogleId,
// 		p.Email,
// 		p.Name,
// 		p.AccessToken,
// 		p.RefreshToken,
// 		p.Tokenexpiry,
// 	)

// 	// Handle any errors that might occur during user creation
// 	if err != nil {
// 		return &pb.UserSignUpResponse{
// 			Profile:      &pb.GoogleSignInResponse{},
// 			AccessToken:  p.AccessToken,
// 			RefreshToken: p.RefreshToken,
// 		}, err
// 	}

// 	// If user creation is successful, return the user's details
// 	return &pb.UserSignUpResponse{
// 		Profile: &pb.GoogleSignInResponse{
// 			Id:       user.UUID,
// 			GoogleId: user.GoogleID,
// 			Fullname: user.FullName,
// 			Email:    user.Email,
// 		},
// 		AccessToken:  user.AccessToken,  // Access token from the user
// 		RefreshToken: user.RefreshToken, // Refresh token from the user
// 	}, nil

// }

// // GetUserGoogleDetaisbyID implements interfaces.UserServiceInter.
// func (u *UserService) GetUserGoogleDetaisbyID(p *pb.ID) (*pb.GoogleUserDetails, error) {
// 	user, err := u.Repo.GetUserGoogleDetailsByID(string(p.ID))
// 	if err != nil {
// 		return &pb.GoogleUserDetails{
// 			Googleid:     user.GoogleID,
// 			Email:        user.GoogleEmail,
// 			Accesstoken:  user.AccessToken,
// 			Refreshtoken: user.RefreshToken,
// 		}, nil
// 	}
// 	return &pb.GoogleUserDetails{}, nil
// }

// // UpdateUserGoogleToken implements interfaces.UserServiceInter.
// func (u *UserService) UpdateUserGoogleToken(p *pb.UpdateGoogleTokenReq) (*pb.UpdateGoogleTokenRes, error) {
// 	err := u.Repo.UpdateUserGoogleToken(p.GoogleID, p.AccessToken, p.RefreshToken, p.TokenExpiry)
// 	if err != nil {
// 		return &pb.UpdateGoogleTokenRes{}, err
// 	}
// 	return &pb.UpdateGoogleTokenRes{}, nil
// }
