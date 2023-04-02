package util

import "fmt"

// GetUserLikeKey 用户点赞集合的key
func GetUserLikeKey(uid int64) string {
	return fmt.Sprintf("user_like_%d", uid)
}

// GetUserCollectKey 用户收藏集合的key
func GetUserCollectKey(uid int64) string {
	return fmt.Sprintf("user_collect_%d", uid)
}

// GetUserFollowKey 用户关注列表的key
func GetUserFollowKey(uid int64) string {
	return fmt.Sprintf("user_follow_%d", uid)
}

// GetUserFansKey 用户粉丝列表的key
func GetUserFansKey(uid int64) string {
	return fmt.Sprintf("user_fans_%d", uid)
}

// GetArticleIPKey 获取文章ip key
func GetArticleIPKey(ip string, articleID int64) string {
	return fmt.Sprintf("addr_%s_%d", ip, articleID)
}
