package services

import (
	"context"
	"log"

	materialpb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

// GetCuttingResService interacts with MaterialService to fetch the data from DB
func (u *UserService) GetCuttingResService(p *pb.UserItemID) (*pb.UserCuttingResultResponse, error) {
	ctx := context.Background()

	// Make the gRPC call to MaterialService
	result, err := u.MaterialClient.GetCuttingResult(ctx, &materialpb.ItemID{
		ID: p.ID,
	})
	if err != nil {
		log.Printf("Failed to fetch cutting result: %v", err)
		return nil, err
	}

	// Prepare the response object
	userCuttingResultResponse := &pb.UserCuttingResultResponse{
		Status:  pb.UserCuttingResultResponse_OK,
		Message: "Cutting result fetched successfully",
		// Map the components from MaterialService response to UserCuttingResultResponse
		CuttingResult: &pb.UsercuttingResult{
			Item_ID: uint32(p.ID),
		},
	}

	// Map the components from MaterialService response
	var userComponents []*pb.UserComponent
	for _, component := range result.CuttingResult.Components {
		userComponents = append(userComponents, &pb.UserComponent{
			Material_ID:   component.Material_ID,
			DoorPanel:     component.DoorPanel,
			BackSidePanel: component.BackSidePanel,
			SidePanel:     component.SidePanel,
			TopPanel:      component.TopPanel,
			BottomPanel:   component.BottomPanel,
			ShelvesPanel:  component.ShelvesPanel,
			Panel_Count:   component.Panel_Count,
		})
	}

	// Add components to the response
	userCuttingResultResponse.CuttingResult.Components = userComponents

	return userCuttingResultResponse, nil
}
