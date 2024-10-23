package services

import (
	"context"

	materialpb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

// FindAllOrdersSvc implements interfaces.UserServiceInter.
func (u *UserService) FindAllOrdersSvc(p *pb.ItemNoParams) (*pb.OrderList, error) {
	ctx := context.Background()

	response, err := u.MaterialClient.OrderHistory(ctx, &materialpb.ItemNoParams{})
	if err != nil {
		return nil, err
	}
	var orders []*pb.Order

	for _, order := range response.Orders {

		orders = append(orders, &pb.Order{
			Order_ID:   uint32(order.Order_ID),
			User_ID:    uint32(order.User_ID),
			Item_ID:    uint32(order.Item_ID),
			Status:     order.Status,
			Amount:     order.Amount,
			Is_Custom:  order.Is_Custom,
			Payment_ID: order.Payment_ID,
		})
	}

	return &pb.OrderList{
		Orders: orders,
	}, nil

}

// FindOrderSvc implements interfaces.UserServiceInter.
func (u *UserService) FindOrderSvc(p *pb.ItemID) (*pb.Order, error) {
	ctx := context.Background()

	order, err := u.MaterialClient.FindOrder(ctx, &materialpb.ItemID{ID: p.ID})
	if err != nil {
		return nil, err
	}
	return &pb.Order{
		Order_ID:   uint32(order.Order_ID),
		User_ID:    uint32(order.User_ID),
		Item_ID:    uint32(order.Item_ID),
		Status:     order.Status,
		Amount:     order.Amount,
		Is_Custom:  order.Is_Custom,
		Payment_ID: order.Payment_ID,
	}, nil

}

// PlaceOrderService implements interfaces.UserServiceInter.
func (u *UserService) PlaceOrderService(p *pb.Order) (*pb.OrderResponse, error) {

	ctx := context.Background()
	// Construct the gRPC request for the MaterialClient
	materialOrder := &materialpb.Order{
		Order_ID:   p.Order_ID,
		User_ID:    p.User_ID,
		Item_ID:    p.Item_ID,
		Quantity:   p.Quantity,
		Status:     p.Status,
		CustomCut:  p.CustomCut,
		Is_Custom:  p.Is_Custom,
		Amount:     p.Amount,
		Payment_ID: p.Payment_ID,
	}

	// Call the PlaceOrder method on the MaterialClient
	response, err := u.MaterialClient.PlaceOrder(ctx, materialOrder)
	if err != nil {
		return &pb.OrderResponse{
			Status:  pb.OrderResponse_ERROR,
			Message: "failed to place order",
			Payload: &pb.OrderResponse_Error{Error: err.Error()},
		}, err
	}

	// Return the response with the status and message from the Material service
	return &pb.OrderResponse{
		Status:  pb.OrderResponse_OK,
		Message: "Order placed successfully",
		Payload: &pb.OrderResponse_Data{
			Data: response.GetPayload().(*materialpb.OrderResponse_Data).Data, // Copy the data from material service response
		},
	}, nil

}