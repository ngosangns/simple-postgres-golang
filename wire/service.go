package wire

import (
	"infra/domain/user"
)

var UserServiceSingleton = NewSingleton(func() *user.UserService {
	return &user.UserService{PgxPool: ProvidePgxPoolSingleton()}
})
