package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

// AddAddress will pass the details to service layer to add address to database.
func (u *UserHandler) AddAddress(ctx context.Context, p *pb.Address) (*pb.Response, error) {
	response, err := u.SVC.AddAddressService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// EditAddress handles the request to edit an existing address for a user.
func (u *UserHandler) EditAddress(ctx context.Context, p *pb.Address) (*pb.Address, error) {
	response, err := u.SVC.EditAddressService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// ViewAllAddress handler the fetch the address from the db
func (u *UserHandler) ViewAllAddress(ctx context.Context, p *pb.ID) (*pb.AddressList, error) {
	response, err := u.SVC.ViewAllAddress(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// RemoveAddress handles the request to remove an existing address for a user.
func (u *UserHandler) RemoveAddress(ctx context.Context, p *pb.IDs) (*pb.Response, error) {
	response, err := u.SVC.RemoveAddressService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
