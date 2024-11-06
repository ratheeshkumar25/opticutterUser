package services

import (
	"errors"
	"fmt"

	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/utils"
)

// ViewProfileSevice implements interfaces.UserServiceInter.
func (u *UserService) ViewProfileSevice(p *pb.ID) (*pb.Profile, error) {
	user, err := u.Repo.FindUserByID(uint(p.ID))
	if err != nil {
		return nil, err
	}
	userModel := &pb.Profile{
		First_Name: user.FirstName,
		Last_Name:  user.LastName,
		Email:      user.Email,
		Phone:      user.Phone,
		Wallet:     float32(user.Wallet),
	}
	return userModel, nil
}

// EditProfileServive implements interfaces.UserServiceInter.
func (u *UserService) EditProfileServive(p *pb.Profile) (*pb.Profile, error) {
	user, err := u.Repo.FindUserByID(uint(p.User_ID))
	if err != nil {
		return nil, err
	}
	user.FirstName = p.First_Name
	user.LastName = p.Last_Name
	user.Phone = p.Phone
	err = u.Repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// BlockedUserService implements interfaces.UserServiceInter.
func (u *UserService) BlockedUserService(p *pb.ID) (*pb.Response, error) {
	user, err := u.Repo.FindUserByID(uint(p.ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting user from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	user.IsBlocked = true

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in updating user in database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "user blocked successfully",
	}, nil
}

// ChangePassword implements interfaces.UserServiceInter.
func (u *UserService) ChangePassword(p *pb.Password) (*pb.Response, error) {
	user, err := u.Repo.FindUserByID(uint(p.User_ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting user from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	fmt.Println(p.Old_Password, user.Password)
	if !utils.CheckPassword(p.Old_Password, user.Password) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "old password is incorrect",
		}, errors.New("old password mismatch")
	}

	if p.New_Password != p.Confirm_Password {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "password is incorrect, passwords not matching",
		}, errors.New("new password mismatch")
	}

	newPassword, err := utils.HashPassword(p.New_Password)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error while hashing new password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	user.Password = newPassword

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error while updating password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Password changed successfully",
	}, nil
}

// UnblockUserService implements interfaces.UserServiceInter.
func (u *UserService) UnblockUserService(p *pb.ID) (*pb.Response, error) {

	user, err := u.Repo.FindUserByID(uint(p.ID))
	fmt.Println("userrr data", user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting user from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	user.IsBlocked = false
	fmt.Println("userrr", user.IsBlocked)

	err = u.Repo.UpdateUser(user)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in updating user in database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "user unblocked successfully",
	}, nil
}

// ViewUserList implements interfaces.UserServiceInter.
func (u *UserService) ViewUserList(p *pb.NoParam) (*pb.UserListResponse, error) {
	// Fetch the user list from the repository to interact with db
	users, err := u.Repo.GetUserList()
	if err != nil {
		return nil, err
	}

	//Create a response message to hold the user profiles
	userListResponse := &pb.UserListResponse{}

	//Map the retrieved user data to the gRPC profile format
	for _, user := range users {
		// Construct each Profile from the user data
		profile := &pb.Profile{
			User_ID:    uint32(user.ID),
			First_Name: user.FirstName,
			Last_Name:  user.LastName,
			Phone:      user.Phone,
			Email:      user.Email,
			Wallet:     float32(user.Wallet),
			Is_Blocked: user.IsBlocked,
		}

		// Append each Profile to the UserListResponse
		userListResponse.Profiles = append(userListResponse.Profiles, profile)
	}

	//  Return the populated UserListResponse and nil error
	return userListResponse, nil
}
