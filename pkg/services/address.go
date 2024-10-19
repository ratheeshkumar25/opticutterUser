package services

import (
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/model"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
	"gorm.io/gorm"
)

// AddAddressService implements interfaces.UserServiceInter.
func (u *UserService) AddAddressService(p *pb.Address) (*pb.Response, error) {
	address := &model.Address{
		House:  p.House,
		Street: p.Street,
		City:   p.City,
		ZIP:    uint(p.Zip),
		State:  p.State,
		UserID: uint(p.User_ID),
	}

	err := u.Repo.CreateAddress(address)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error adding address",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Address added successfully",
	}, nil
}

// EditAddressService implements interfaces.UserServiceInter.
func (u *UserService) EditAddressService(p *pb.Address) (*pb.Address, error) {
	address := &model.Address{
		Model: gorm.Model{
			ID: uint(p.ID),
		},
		House:  p.House,
		Street: p.Street,
		City:   p.City,
		ZIP:    uint(p.Zip),
		State:  p.State,
		UserID: uint(p.User_ID),
	}
	err := u.Repo.EditAddress(address)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// ViewAllAddress implements interfaces.UserServiceInter.
func (u *UserService) ViewAllAddress(p *pb.ID) (*pb.AddressList, error) {
	// Call the repository method to fetch all addresses
	userID := p.ID
	addresses, err := u.Repo.GetAllAddresses(uint(userID))
	if err != nil {
		return nil, err
	}

	// Convert the address models to protobuf messages
	addressList := &pb.AddressList{
		Addresses: make([]*pb.Address, 0, len(*addresses)),
	}

	for _, addr := range *addresses {
		addressList.Addresses = append(addressList.Addresses, &pb.Address{
			ID:      uint32(addr.ID),
			House:   addr.House,
			Street:  addr.Street,
			City:    addr.City,
			Zip:     uint32(addr.ZIP),
			State:   addr.State,
			User_ID: uint32(addr.UserID),
		})
	}

	return addressList, nil
}

// RemoveAddressService implements interfaces.UserServiceInter.
func (u *UserService) RemoveAddressService(p *pb.IDs) (*pb.Response, error) {
	err := u.Repo.RemoveAddress(uint(p.ID), uint(p.User_ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error deleting address",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Address removed successfully",
	}, nil
}
