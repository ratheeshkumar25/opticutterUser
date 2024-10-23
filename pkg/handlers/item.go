package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

func (u *UserHandler) AddItem(ctx context.Context, p *pb.UserItem) (*pb.Response, error) {
	response, err := u.SVC.AddItemService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) FindItemByID(ctx context.Context, p *pb.UserItemID) (*pb.UserItem, error) {
	response, err := u.SVC.FindItemByID(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) FindAllItem(ctx context.Context, p *pb.NoParam) (*pb.UserItemList, error) {
	response, err := u.SVC.FindAllItem(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) EditItem(ctx context.Context, p *pb.UserItem) (*pb.UserItem, error) {
	response, err := u.SVC.EditItemService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) RemoveItem(ctx context.Context, p *pb.UserItemID) (*pb.Response, error) {
	response, err := u.SVC.RemoveItemService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
