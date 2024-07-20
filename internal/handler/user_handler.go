package handler

import (
	"fmt"

	"github.com/nadiannis/libry/internal/domain/input"
	"github.com/nadiannis/libry/internal/usecase"
	"github.com/nadiannis/libry/internal/utils"
)

type UserHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(usecase usecase.IUserUsecase) IUserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) GetAllUsers(parts []string) {
	if len(parts) != 1 {
		fmt.Println(`input should be \lu`)
		return
	}

	users := h.usecase.GetAllUsers()
	total := len(users)

	fmt.Printf("total: %d\n", total)

	if total == 0 {
		fmt.Println("there are no users available")
		return
	}

	utils.UserTable(users)
}

func (h *UserHandler) AddUser(input *input.UserInput) {
	if input.Username == "" {
		fmt.Println("username is required")
		return
	}

	savedUser, err := h.usecase.AddUser(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("('%s' is saved with ID '%s')\n", savedUser.Username, savedUser.ID)
}
