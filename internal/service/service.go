package service

import (
	"context"
	"stakeway/config"
	"stakeway/internal/store/pg"
	"stakeway/pkg/logger"
)

type Service struct {
	Ctx    context.Context
	Cfg    config.Config
	DB     pg.Repository
	Logger *logger.Logger
}

func New(ctx context.Context, cfg config.Config, db pg.Repository, log *logger.Logger) Service {
	return Service{ctx, cfg, db, log}
}
