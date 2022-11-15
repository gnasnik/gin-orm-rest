package core

import (
	"context"
	"github.com/gnasnik/titan-operator/config"
	"github.com/gnasnik/titan-operator/core/logger"
	"github.com/gnasnik/titan-operator/generated/query"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) error {
	if cfg.DatabaseURL == "" {
		return errors.New("database url not setup")
	}

	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		return err
	}

	query.SetDefault(db)
	logger.Subscribe(context.Background())
	return nil
}
