package services

import (
	"context"

	materialpb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

// FindAllMaterialService implements interfaces.UserServiceInter.
func (u *UserService) FindAllMaterialService(p *pb.NoParam) (*pb.UserMaterialList, error) {
	ctx := context.Background()

	result, err := u.MaterialClient.FindAllMaterial(ctx, &materialpb.MaterialNoParams{})
	if err != nil {
		return nil, err
	}

	// Prepare a slice to hold the  materials
	var materials []*pb.UserMaterial

	for _, material := range result.Materials {
		pbMaterial := &pb.UserMaterial{
			Material_ID:   uint32(material.Material_ID),
			Material_Name: material.Material_Name,
			Description:   material.Description,
			Stock:         int32(material.Stock),
			Price:         material.Price,
		}
		materials = append(materials, pbMaterial)
	}

	return &pb.UserMaterialList{
		Materials: materials,
	}, nil
}

// FindMaterialByIDService implements interfaces.UserServiceInter.
func (u *UserService) FindMaterialByIDService(p *pb.UserMaterialID) (*pb.UserMaterial, error) {
	// Create a new context
	ctx := context.Background()

	// Call the MaterialClient to fetch the material by ID
	result, err := u.MaterialClient.FindMaterialByID(ctx, &materialpb.MaterialID{
		ID: p.ID,
	})
	if err != nil {
		return nil, err // Return error if something goes wrong
	}

	// If the material is found, map the response to the pb.Material type
	pbMaterial := &pb.UserMaterial{
		Material_Name: result.Material_Name,
		Description:   result.Description,
		Stock:         int32(result.Stock),
		Price:         result.Price,
	}

	// Return the pb.Material object and no error
	return pbMaterial, nil

}
