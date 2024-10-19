package interfaces

import (
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

type UserServiceInter interface {
	SignupService(p *pb.Signup) (*pb.Response, error)
	// GoogleSignInService(p *pb.GoogleSignInRequest) (*pb.UserSignUpResponse, error)
	VerificationService(p *pb.OTP) (*pb.Response, error)
	LoginService(p *pb.Login) (*pb.Response, error)
	ViewProfileSevice(p *pb.ID) (*pb.Profile, error)
	EditProfileServive(p *pb.Profile) (*pb.Profile, error)
	ChangePassword(p *pb.Password) (*pb.Response, error)
	BlockedUserService(p *pb.ID) (*pb.Response, error)
	UnblockUserService(p *pb.ID) (*pb.Response, error)
	AddAddressService(p *pb.Address) (*pb.Response, error)
	EditAddressService(p *pb.Address) (*pb.Address, error)
	ViewAllAddress(p *pb.ID) (*pb.AddressList, error)
	RemoveAddressService(p *pb.IDs) (*pb.Response, error)
	ViewUserList(p *pb.NoParam) (*pb.UserListResponse, error)
	// GetUserGoogleDetaisbyID(p *pb.ID) (*pb.GoogleUserDetails, error)
	// UpdateUserGoogleToken(p *pb.UpdateGoogleTokenReq) (*pb.UpdateGoogleTokenRes, error)
}
