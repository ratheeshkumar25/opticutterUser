package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
)

func (u *UserHandler) UserGetCuttingResult(ctx context.Context, p *pb.UserItemID) (*pb.UserCuttingResultResponse, error) {
	response, err := u.SVC.GetCuttingResService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
