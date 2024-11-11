package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

func (u *UserHandler) UserCreatePayment(ctx context.Context, p *pb.UserOrder) (*pb.UserPaymentResponse, error) {
	response, err := u.SVC.UserPaymentService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (u *UserHandler) UserPaymentSuccess(ctx context.Context, p *pb.UserPayment) (*pb.UserPaymentStatusResponse, error) {
	response, err := u.SVC.UserPaymentSuccessService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
