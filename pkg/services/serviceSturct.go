package services

import (
	"github.com/ratheeshkumar25/opti_cut_userservice/config"
	materialpb "github.com/ratheeshkumar25/opti_cut_userservice/pkg/client/material/pb"
	inter "github.com/ratheeshkumar25/opti_cut_userservice/pkg/repo/interfaces"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/services/interfaces"
)

type UserService struct {
	Repo           inter.UserRepoInter
	twilio         *config.TwilioService
	redis          *config.RedisService
	MaterialClient materialpb.MaterialServiceClient
}

func NewUserService(repo inter.UserRepoInter, redis *config.RedisService, twilio *config.TwilioService, materialClient materialpb.MaterialServiceClient) interfaces.UserServiceInter {
	return &UserService{
		Repo:           repo,
		twilio:         twilio,
		redis:          redis,
		MaterialClient: materialClient,
	}
}
