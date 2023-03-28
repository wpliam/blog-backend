package mdb

import "blog-backend/model"

// GetBannerList 获取banner
func (cli *MysqlClient) GetBannerList() ([]*model.Banner, error) {
	var banners []*model.Banner
	if err := cli.Limit(3).Order("weight desc").Find(&banners).Error; err != nil {
		return nil, err
	}
	return banners, nil
}
