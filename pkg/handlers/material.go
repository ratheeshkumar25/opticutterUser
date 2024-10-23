package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

func (u *UserHandler) FindMaterialByID(ctx context.Context, p *pb.UserMaterialID) (*pb.UserMaterial, error) {
	response, err := u.SVC.FindMaterialByIDService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (u *UserHandler) FindAllMaterial(ctx context.Context, p *pb.NoParam) (*pb.UserMaterialList, error) {
	response, err := u.SVC.FindAllMaterialService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
