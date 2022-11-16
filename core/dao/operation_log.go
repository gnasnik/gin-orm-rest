package dao

import (
	"context"
	"github.com/gnasnik/titan-operator/generated/model"
	"github.com/gnasnik/titan-operator/generated/query"
)

func AddOperationLog(ctx context.Context, log *model.OperationLog) error {
	return query.OperationLog.WithContext(ctx).Create(log)
}

func ListOperationLog(ctx context.Context, offset, limit int) ([]*model.OperationLog, int64, error) {
	return query.OperationLog.WithContext(ctx).FindByPage(offset, limit)
}
