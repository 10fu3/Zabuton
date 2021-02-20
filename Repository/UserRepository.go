package Repository

import "../Entity"

type userRepository struct {

}

type UserRepository interface {
	CreateUser(user Entity.User)
	GetUserByID(uuid string) (Entity.User,error)
	DeleteUserByID(uuid string) error
	UpdateUser(user Entity.User) error
}

