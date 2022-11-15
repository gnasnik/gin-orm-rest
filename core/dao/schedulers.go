package dao

import (
	"github.com/gnasnik/titan-operator/generated/model"
	"github.com/gnasnik/titan-operator/generated/query"
)

func GetSchedulers() ([]*model.Scheduler, error) {
	return query.Scheduler.Find()
}
