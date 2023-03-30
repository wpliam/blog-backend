package mdb

import (
	"blog-backend/model"
)

// GetUserInfo 获取用户信息
func (cli *MysqlClient) GetUserInfo(uid int64) (*model.User, error) {
	var user *model.User
	if err := cli.First(&user, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetAccountInfo 获取账号信息 主要用于登录验证
func (cli *MysqlClient) GetAccountInfo(username string) (*model.Account, error) {
	var account *model.Account
	if err := cli.First(&account, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return account, nil
}

// BatchGetUserInfo 批量获取用户信息
func (cli *MysqlClient) BatchGetUserInfo(userIDs []int64) (map[int64]*model.User, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	var users []*model.User
	where := make(map[string]interface{})
	where["id"] = userIDs
	if err := cli.Where(where).Find(&users).Error; err != nil {
		return nil, err
	}
	userInfo := make(map[int64]*model.User)
	for _, user := range users {
		userInfo[user.ID] = user
	}
	return userInfo, nil
}

// UpdateUserInfo 更新用户登录信息
func (cli *MysqlClient) UpdateUserInfo(uid int64, field map[string]interface{}) error {
	return cli.Model(&model.User{}).Where("id = ?", uid).Updates(field).Error
}
