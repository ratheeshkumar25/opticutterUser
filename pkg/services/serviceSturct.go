package services

import (
	"github.com/ratheeshkumar25/opti_cut_userservice/config"
	inter "github.com/ratheeshkumar25/opti_cut_userservice/pkg/repo/interfaces"
	"github.com/ratheeshkumar25/opti_cut_userservice/pkg/services/interfaces"
)

type UserService struct {
	Repo   inter.UserRepoInter
	twilio *config.TwilioService
	redis  *config.RedisService
}

func NewUserService(repo inter.UserRepoInter, redis *config.RedisService, twilio *config.TwilioService) interfaces.UserServiceInter {
	return &UserService{
		Repo:   repo,
		twilio: twilio,
		redis:  redis,
	}
}
