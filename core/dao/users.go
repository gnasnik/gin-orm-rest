package dao

import (
	"context"
	"github.com/gnasnik/gin-orm-rest/generated/model"
	"github.com/gnasnik/gin-orm-rest/generated/query"
)

func CreateUser(ctx context.Context, user *model.User) error {
	return query.User.WithContext(ctx).Create(user)
}

func GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return query.User.WithContext(ctx).Where(query.User.Username.Eq(username)).Take()
}
