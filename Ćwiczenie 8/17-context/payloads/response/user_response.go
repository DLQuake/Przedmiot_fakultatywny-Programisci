package response

import (
	"github.com/jacky-htg/go-services/models"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"username"`
	IsActive  bool   `json:"is_active"`
}

func (u *UserResponse) Transform(user *models.User) {
	u.ID = user.ID
	u.Name = user.Name
	u.IsActive = user.IsActive
}