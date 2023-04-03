package user

type RefreshTokenReq struct {
	Uid   int64  `json:"uid" binding:"min=1"`
	Token string `json:"token" binding:"required"`
}

type RefreshTokenReply struct {
	Token string `json:"token"`
}
