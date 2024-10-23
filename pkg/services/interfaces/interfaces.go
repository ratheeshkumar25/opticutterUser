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

	FindMaterialByIDService(p *pb.MaterialID) (*pb.Material, error)
	FindAllMaterialService(p *pb.MaterialNoParams) (*pb.MaterialList, error)

	// Service to handle item management
	AddItemService(p *pb.Item) (*pb.ItemResponse, error)
	EditItemService(p *pb.Item) (*pb.Item, error)
	RemoveItemService(p *pb.ItemID) (*pb.ItemResponse, error)
	FindItemByID(p *pb.ItemID) (*pb.Item, error)
	FindAllItem(p *pb.ItemNoParams) (*pb.ItemList, error)

	// Service to handle orders
	PlaceOrderService(p *pb.Order) (*pb.OrderResponse, error)
	FindAllOrdersSvc(p *pb.ItemNoParams) (*pb.OrderList, error)
	FindOrderSvc(p *pb.ItemID) (*pb.Order, error)
}
