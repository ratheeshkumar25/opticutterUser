package services

import (
	"context"
	"fmt"

	materialpb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

// UserPaymentService implements interfaces.UserServiceInter.
func (u *UserService) UserPaymentService(p *pb.UserOrder) (*pb.UserPaymentResponse, error) {
	ctx := context.Background()

	order := &materialpb.Order{
		User_ID:  p.User_ID,
		Order_ID: p.Order_ID,
		Amount:   p.Amount,
		Status:   p.Status,
	}

	result, err := u.MaterialClient.CreatePayment(ctx, order)
	if err != nil {
		return nil, err
	}
	return &pb.UserPaymentResponse{
		PaymentId:    result.PaymentId,
		ClientSecret: result.ClientSecret,
		OrderId:      result.OrderId,
		Amount:       result.Amount,
	}, nil
}

// UserPaymentSuccessService implements interfaces.UserServiceInter.
func (u *UserService) UserPaymentSuccessService(p *pb.UserPayment) (*pb.UserPaymentStatusResponse, error) {
	ctx := context.Background()
	fmt.Println(p.Order_ID)

	payment := &materialpb.Payment{
		User_ID:    p.User_ID,
		Payment_ID: p.Payment_ID,
		Order_ID:   p.Order_ID,
		Amount:     p.Amount,
	}

	result, err := u.MaterialClient.PaymentSuccess(ctx, payment)
	if err != nil {
		return nil, err
	}

	return &pb.UserPaymentStatusResponse{
		Status:  pb.UserPaymentStatusResponse_SUCCESS,
		Message: result.Message,
	}, nil
}
