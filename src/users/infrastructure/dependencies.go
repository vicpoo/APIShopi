//dependencies.go
package infrastructure

import (
	"github.com/vicpoo/apiShop/src/users/application"
)

func InitUserDependencies() (
	*RegisterUserController,
	*LoginUserController,
	*CreateUserController,
	*UpdateUserController,
	*DeleteUserController,
	*GetUserByIDController,
	*GetAllUsersController,
) {
	// Repositorio implementado en infraestructura
	repo := NewMySQLUserRepository() 

	// Casos de uso
	registerUseCase := application.NewRegisterUserUseCase(repo)
	loginUseCase := application.NewLoginUserUseCase(repo)
	createUseCase := application.NewCreateUserUseCase(repo)
	updateUseCase := application.NewUpdateUserUseCase(repo)
	deleteUseCase := application.NewDeleteUserUseCase(repo)
	getByIDUseCase := application.NewGetUserByIDUseCase(repo)
	getAllUseCase := application.NewGetAllUsersUseCase(repo)

	// Controladores
	registerController := NewRegisterUserController(registerUseCase)
	loginController := NewLoginUserController(loginUseCase)
	createController := NewCreateUserController(createUseCase)
	updateController := NewUpdateUserController(updateUseCase)
	deleteController := NewDeleteUserController(deleteUseCase)
	getByIDController := NewGetUserByIDController(getByIDUseCase)
	getAllController := NewGetAllUsersController(getAllUseCase)

	return registerController, loginController, createController, updateController,
		deleteController, getByIDController, getAllController
}