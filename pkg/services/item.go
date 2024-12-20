package services

import (
	"context"
	"fmt"

	materialpb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

func (u *UserService) AddItemService(p *pb.UserItem) (*pb.Response, error) {
	ctx := context.Background()

	newItem := &materialpb.Item{
		Item_Name:     p.Item_Name,
		Material_ID:   p.Material_ID,
		Length:        p.Length,
		Width:         p.Width,
		Fixed_Size_ID: p.Fixed_Size_ID,
		Is_Custom:     p.Is_Custom,
		User_ID:       p.User_ID,
	}

	itemID, err := u.MaterialClient.AddItem(ctx, newItem)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to create product",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	// Return success response with the new item ID
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Item created successfully",
		Payload: &pb.Response_Data{
			Data: fmt.Sprintf("ItemID:%s", itemID),
		},
	}, nil
}

// EditItemService implements interfaces.UserServiceInter.
func (u *UserService) EditItemService(p *pb.UserItem) (*pb.UserItem, error) {
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
		User_ID:       p.User_ID,
	}

	// Call the MaterialClient's EditItem method
	_, err := u.MaterialClient.EditItem(ctx, updatedItem)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// FindAllItem implements interfaces.UserServiceInter.
func (u *UserService) FindAllItem(p *pb.NoParam) (*pb.UserItemList, error) {
	ctx := context.Background()

	// Call the MaterialClient's FindAllItem method
	result, err := u.MaterialClient.FindAllItem(ctx, &materialpb.ItemNoParams{})
	if err != nil {
		return nil, err
	}

	// Convert materialpb.ItemList to pb.ItemList
	var items []*pb.UserItem
	for _, item := range result.Items {
		pbItem := &pb.UserItem{
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

	return &pb.UserItemList{
		Items: items,
	}, nil
}

// FindItemByID implements interfaces.UserServiceInter.
func (u *UserService) FindItemByID(p *pb.UserItemID) (*pb.UserItem, error) {
	ctx := context.Background()

	// Call the MaterialClient's FindItemByID method
	item, err := u.MaterialClient.FindItemByID(ctx, &materialpb.ItemID{ID: p.ID})
	if err != nil {
		return nil, err
	}

	// Convert materialpb.Item to pb.Item
	return &pb.UserItem{
		Item_ID:         item.Item_ID,
		Item_Name:       item.Item_Name,
		Length:          item.Length,
		Width:           item.Width,
		Fixed_Size_ID:   item.Fixed_Size_ID,
		Is_Custom:       item.Is_Custom,
		Estimated_Price: item.Estimated_Price,
		User_ID:         item.User_ID,
	}, nil
}

// RemoveItemService implements interfaces.UserServiceInter.
func (u *UserService) RemoveItemService(p *pb.UserItemID) (*pb.Response, error) {
	ctx := context.Background()

	// Call the MaterialClient's RemoveItem method
	_, err := u.MaterialClient.RemoveItem(ctx, &materialpb.ItemID{ID: p.ID})
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to remove item",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Item removed successfully",
	}, nil
}

// FindAllItemByUser implements interfaces.UserServiceInter.
func (u *UserService) FindAllItemByUser(p *pb.UserItemID) (*pb.UserItemList, error) {
	ctx := context.Background()

	result, err := u.MaterialClient.FindAllItemByUser(ctx, &materialpb.ItemID{ID: p.ID})
	if err != nil {
		return nil, err
	}

	var items []*pb.UserItem

	for _, itm := range result.Items {
		pbItems := &pb.UserItem{
			Item_ID:         itm.Item_ID,
			Item_Name:       itm.Item_Name,
			Material_ID:     itm.Material_ID,
			Length:          itm.Length,
			Width:           itm.Width,
			Fixed_Size_ID:   itm.Fixed_Size_ID,
			Estimated_Price: itm.Estimated_Price,
			User_ID:         itm.User_ID,
		}
		items = append(items, pbItems)
	}
	return &pb.UserItemList{
		Items: items,
	}, nil
}
