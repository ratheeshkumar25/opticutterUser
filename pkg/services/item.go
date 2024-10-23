package services

import (
	"context"
	"fmt"

	materialpb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

func (u *UserService) AddItemService(p *pb.Item) (*pb.ItemResponse, error) {
	ctx := context.Background()

	newItem := &materialpb.Item{
		Item_Name:     p.Item_Name,
		Material_ID:   p.Material_ID,
		Length:        p.Length,
		Width:         p.Width,
		Fixed_Size_ID: p.Fixed_Size_ID,
		Is_Custom:     p.Is_Custom,
	}

	itemID, err := u.MaterialClient.AddItem(ctx, newItem)
	if err != nil {
		return &pb.ItemResponse{
			Status:  pb.ItemResponse_ERROR,
			Message: "Failed to create product",
			Payload: &pb.ItemResponse_Error{Error: err.Error()},
		}, err
	}
	// Return success response with the new item ID
	return &pb.ItemResponse{
		Status:  pb.ItemResponse_OK,
		Message: "Item created successfully",
		Payload: &pb.ItemResponse_Data{
			Data: fmt.Sprintf("ItemID:%s", itemID),
		},
	}, nil
}

// EditItemService implements interfaces.UserServiceInter.
func (u *UserService) EditItemService(p *pb.Item) (*pb.Item, error) {
	ctx := context.Background()

	// Map pb.Item to materialpb.Item for sending to MaterialClient
	updatedItem := &materialpb.Item{
		Item_ID:       p.Item_ID,
		Item_Name:     p.Item_Name,
		Material_ID:   p.Material_ID,
		Length:        p.Length,
		Width:         p.Width,
		Fixed_Size_ID: p.Fixed_Size_ID,
		Is_Custom:     p.Is_Custom,
	}

	// Call the MaterialClient's EditItem method
	_, err := u.MaterialClient.EditItem(ctx, updatedItem)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// FindAllItem implements interfaces.UserServiceInter.
func (u *UserService) FindAllItem(p *pb.ItemNoParams) (*pb.ItemList, error) {
	ctx := context.Background()

	// Call the MaterialClient's FindAllItem method
	result, err := u.MaterialClient.FindAllItem(ctx, &materialpb.ItemNoParams{})
	if err != nil {
		return nil, err
	}

	// Convert materialpb.ItemList to pb.ItemList
	var items []*pb.Item
	for _, item := range result.Items {
		pbItem := &pb.Item{
			Item_ID:         item.Item_ID,
			Item_Name:       item.Item_Name,
			Length:          item.Length,
			Width:           item.Width,
			Fixed_Size_ID:   item.Fixed_Size_ID,
			Is_Custom:       item.Is_Custom,
			Estimated_Price: item.Estimated_Price,
		}
		items = append(items, pbItem)
	}

	return &pb.ItemList{
		Items: items,
	}, nil
}

// FindItemByID implements interfaces.UserServiceInter.
func (u *UserService) FindItemByID(p *pb.ItemID) (*pb.Item, error) {
	ctx := context.Background()

	// Call the MaterialClient's FindItemByID method
	item, err := u.MaterialClient.FindItemByID(ctx, &materialpb.ItemID{ID: p.ID})
	if err != nil {
		return nil, err
	}

	// Convert materialpb.Item to pb.Item
	return &pb.Item{
		Item_ID:         item.Item_ID,
		Item_Name:       item.Item_Name,
		Length:          item.Length,
		Width:           item.Width,
		Fixed_Size_ID:   item.Fixed_Size_ID,
		Is_Custom:       item.Is_Custom,
		Estimated_Price: item.Estimated_Price,
	}, nil
}

// RemoveItemService implements interfaces.UserServiceInter.
func (u *UserService) RemoveItemService(p *pb.ItemID) (*pb.ItemResponse, error) {
	ctx := context.Background()

	// Call the MaterialClient's RemoveItem method
	_, err := u.MaterialClient.RemoveItem(ctx, &materialpb.ItemID{ID: p.ID})
	if err != nil {
		return &pb.ItemResponse{
			Status:  pb.ItemResponse_ERROR,
			Message: "Failed to remove item",
			Payload: &pb.ItemResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.ItemResponse{
		Status:  pb.ItemResponse_OK,
		Message: "Item removed successfully",
	}, nil
}
