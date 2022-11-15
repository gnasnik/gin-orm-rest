package dao

import (
	"context"
	"github.com/gnasnik/titan-operator/generated/model"
	"github.com/gnasnik/titan-operator/generated/query"
)

func AddLoginLog(ctx context.Context, log *model.LoginLog) error {
	return query.LoginLog.Create(log)
}

func ListLoginLog(ctx context.Context, offset, limit int) ([]*model.LoginLog, int64, error) {
	return query.LoginLog.FindByPage(offset, limit)
}
