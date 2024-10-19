package handlers

import (
	pb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/pb"
	inter "github.com/ratheeshkumar25/opti_cut_userservice/pkg/services/interfaces"
)

type UserHandler struct {
	SVC inter.UserServiceInter
	pb.UserServiceServer
}

func NewUserHandler(svc inter.UserServiceInter) *UserHandler {
	return &UserHandler{
		SVC: svc,
	}
}
