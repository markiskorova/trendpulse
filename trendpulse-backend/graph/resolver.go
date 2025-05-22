package graph

import (
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type Resolver struct {
	DB    *gorm.DB
	Asynq *asynq.Client
}
