package jsonagree

import "blog-backend/model"

type GetBannerCard struct {
}

type GetBannerReply struct {
	Banners []*model.Banner `json:"banners"`
}
