package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

func (u *UserHandler) PlaceOrder(ctx context.Context, p *pb.UserOrder) (*pb.Response, error) {
	response, err := u.SVC.PlaceOrderService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) OrderHistory(ctx context.Context, p *pb.NoParam) (*pb.UserOrderList, error) {
	response, err := u.SVC.FindAllOrdersSvc(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) FindOrder(ctx context.Context, p *pb.UserItemID) (*pb.UserOrder, error) {
	response, err := u.SVC.FindOrderSvc(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
