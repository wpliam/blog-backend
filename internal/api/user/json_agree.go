package user

import (
	"blog-backend/model"
)

type GetUserInfoReply struct {
	User     *model.User `json:"user"`
	IsFollow bool        `json:"isFollow"`
}
