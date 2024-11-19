package interfaces

import (
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

type UserServiceInter interface {
	SignupService(p *pb.Signup) (*pb.Response, error)
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

	FindMaterialByIDService(p *pb.UserMaterialID) (*pb.UserMaterial, error)
	FindAllMaterialService(p *pb.NoParam) (*pb.UserMaterialList, error)

	// Service to handle item management
	AddItemService(p *pb.UserItem) (*pb.Response, error)
	EditItemService(p *pb.UserItem) (*pb.UserItem, error)
	RemoveItemService(p *pb.UserItemID) (*pb.Response, error)
	FindItemByID(p *pb.UserItemID) (*pb.UserItem, error)
	FindAllItem(p *pb.NoParam) (*pb.UserItemList, error)
	FindAllItemByUser(p *pb.UserItemID) (*pb.UserItemList, error)

	// Service to handle orders
	PlaceOrderService(p *pb.UserOrder) (*pb.Response, error)
	FindAllOrdersSvc(p *pb.NoParam) (*pb.UserOrderList, error)
	FindOrderSvc(p *pb.UserItemID) (*pb.UserOrder, error)
	FindOrdersByUser(p *pb.UserItemID) (*pb.UserOrderList, error)

	//Service to handle payment
	UserPaymentService(p *pb.UserOrder) (*pb.UserPaymentResponse, error)
	UserPaymentSuccessService(p *pb.UserPayment) (*pb.UserPaymentStatusResponse, error)
	GetCuttingResService(p *pb.UserItemID) (*pb.UserCuttingResultResponse, error)
}
