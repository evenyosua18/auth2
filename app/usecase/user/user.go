package user

type IUserService interface {
}

type ServiceUser struct {
}

func NewUserService() IUserService {
	return &ServiceUser{}
}
