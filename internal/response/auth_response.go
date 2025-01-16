package response

import (
	"github.com/google/uuid"
	"github.com/rizalarfiyan/be-tilik-jalan/constant"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/model"
	"github.com/rizalarfiyan/be-tilik-jalan/utils"
)

type AuthMe struct {
	Id     uuid.UUID         `json:"id"`
	Email  string            `json:"email"`
	Name   string            `json:"name"`
	Role   constant.AuthRole `json:"role"`
	Avatar string            `json:"avatar"`
}

func (a *AuthMe) FromUser(user *model.User) {
	a.Id = user.Id
	a.Email = user.Email
	a.Name = user.Name
	a.Role = user.Role
	a.Avatar = utils.GetGravatar(user.Email)
}
