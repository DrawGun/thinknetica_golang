package storage

import (
	"context"
	"thinknetica_golang/task_21/pkg/model"
)

// Interface определяет контракт хранилища данных.
type Interface interface {
	InsertMovies(context.Context, []model.Movie) error
	DeleteMovie(context.Context, int) error
	UpdateMovie(context.Context, model.Movie) error
	SelectMovies(context.Context, int) ([]model.Movie, error)
}
