package db

import (
	"context"
	"github.com/sakasikai/GoFive/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

func QueryUserByName(ctx context.Context, uname string) ([]*User, error) {
	users := make([]*User, 0)

	if err := DB.WithContext(ctx).Where("user_name = ?", uname).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func QueryUserByID(ctx context.Context, userID int64) ([]*User, error) {
	users := make([]*User, 0)

	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
