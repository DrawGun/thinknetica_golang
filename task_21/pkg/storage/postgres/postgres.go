package postgres

import (
	"context"
	"thinknetica_golang/task_21/pkg/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// DB - хранилище данных
type DB struct {
	pool *pgxpool.Pool
}

// New - создает новый экземпляр типа DB
func New(pool *pgxpool.Pool) *DB {
	db := DB{
		pool: pool,
	}
	return &db
}

// InsertMovies добавляет в БД массив фильмов одной транзакцией.
func (db *DB) InsertMovies(ctx context.Context, movies []model.Movie) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	sql := "INSERT INTO movies (title, release_year, rating_id, fees, studio_id) VALUES ($1, $2, $3, $4, $5)"
	batch := new(pgx.Batch)
	for _, movie := range movies {
		batch.Queue(sql, movie.Title, movie.ReleaseYear, movie.RatingID, movie.Fees, movie.StudioID)
	}

	res := tx.SendBatch(ctx, batch)
	err = res.Close()
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// SelectMovies возвращает список фильмов
func (db *DB) SelectMovies(ctx context.Context, StudioID int) ([]model.Movie, error) {
	movies := []model.Movie{}

	sql := "SELECT * from MOVIES WHERE studio_id = $1 OR $1 = 0"
	rows, err := db.pool.Query(ctx, sql, StudioID)
	if err != nil {
		return movies, err
	}
	defer rows.Close()

	for rows.Next() {
		var m model.Movie
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.ReleaseYear,
			&m.RatingID,
			&m.Fees,
			&m.StudioID,
		)
		if err != nil {
			return movies, err
		}
		movies = append(movies, m)
	}

	err = rows.Err()
	if err != nil {
		return movies, err
	}
	return movies, nil
}

// UpdateMovie обновляет фильм
func (db *DB) UpdateMovie(ctx context.Context, movie model.Movie) error {
	sql := "UPDATE movies SET title = $1, release_year = $2, rating_id = $3, fees = $4, studio_id = $5 WHERE id = $6"
	_, err := db.pool.Exec(ctx, sql, movie.Title, movie.ReleaseYear, movie.RatingID, movie.Fees, movie.StudioID, movie.ID)
	return err
}

// DeleteMovie выполняет удаление фильма
func (db *DB) DeleteMovie(ctx context.Context, id int) error {
	sql := "DELETE FROM movies WHERE id = $1"
	_, err := db.pool.Exec(ctx, sql, id)
	return err
}
