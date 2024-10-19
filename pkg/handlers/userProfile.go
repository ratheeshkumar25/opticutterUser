package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

// ViewProfile fetch the user profile from database
func (u *UserHandler) ViewProfile(ctx context.Context, p *pb.ID) (*pb.Profile, error) {
	response, err := u.SVC.ViewProfileSevice(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// EditProfile update the user profile in database
func (u *UserHandler) EditProftle(ctx context.Context, p *pb.Profile) (*pb.Profile, error) {
	response, err := u.SVC.EditProfileServive(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// ChangePassword update the user profile in database
func (u *UserHandler) ChangePassword(ctx context.Context, p *pb.Password) (*pb.Response, error) {
	response, err := u.SVC.ChangePassword(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// BlockUser update the user as blocked in database
func (u *UserHandler) BlockUser(ctx context.Context, p *pb.ID) (*pb.Response, error) {
	response, err := u.SVC.BlockedUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// UnblockUser updates the user as unblocked in the database
func (u *UserHandler) UnblockUser(ctx context.Context, p *pb.ID) (*pb.Response, error) {
	// Call the service function that handles the unblocking logic
	response, err := u.SVC.UnblockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// UserList fetech the userdata from database
func (u *UserHandler) UserList(ctx context.Context, p *pb.NoParam) (*pb.UserListResponse, error) {
	response, err := u.SVC.ViewUserList(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
