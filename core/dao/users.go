package dao

import (
	"context"
	"github.com/gnasnik/titan-operator/generated/model"
	"github.com/gnasnik/titan-operator/generated/query"
)

func CreateUser(user *model.User) error {
	return query.User.Create(user)
}

func GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return query.User.WithContext(ctx).Where(query.User.Username.Eq(username)).Take()
}
