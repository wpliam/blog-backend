package mdb

import "blog-backend/model"

func (cli *MysqlClient) GetUserInfo(userID int64) (*model.User, error) {
	var user *model.User
	if err := cli.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (cli *MysqlClient) GetAccountInfo(username string) (*model.Account, error) {
	var account *model.Account
	if err := cli.First(&account, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return account, nil
}

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
