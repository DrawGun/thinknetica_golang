package postgres

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"
	"thinknetica_golang/task_21/pkg/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	pg  *DB
	ctx = context.Background()
)

func TestMain(m *testing.M) {
	testDB, err := pgxpool.Connect(ctx, "postgres://postgres:example@localhost:5432/golang_test")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	pg = New(testDB)

	defer testDB.Close()
	os.Exit(m.Run())
}

func TestDB_integration(t *testing.T) {
	// Тест на SelectMovies
	movies, err := pg.SelectMovies(ctx, 0)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}

	got := len(movies)
	want := 11
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}

	newMovies := []model.Movie{
		{ID: 12, Title: "Movie 111", ReleaseYear: 2020, RatingID: 2, Fees: 1_000_000, StudioID: 1},
		{ID: 13, Title: "Movie 212", ReleaseYear: 2021, RatingID: 3, Fees: 2_000_000, StudioID: 2},
	}

	// Тест на InsertMovies
	err = pg.InsertMovies(ctx, newMovies)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}

	movies, err = pg.SelectMovies(ctx, 0)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}

	var singleMovie model.Movie
	found := []model.Movie{}
	for _, movie := range movies {
		if movie.Title == "Movie 111" {
			found = append(found, movie)
		}

		if movie.Title == "Movie 212" {
			found = append(found, movie)
		}

		if movie.StudioID == 5 {
			singleMovie = movie
		}
	}

	if !reflect.DeepEqual(newMovies, found) {
		t.Errorf("got %v; want %v", newMovies, found)
	}

	w1 := 13
	if len(movies) != w1 {
		t.Errorf("got %d; want %d", len(movies), w1)
	}

	// Тест на SelectMovies с параметром студии
	movies, err = pg.SelectMovies(ctx, singleMovie.StudioID)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}

	w2 := []model.Movie{
		singleMovie,
	}

	if !reflect.DeepEqual(movies, w2) {
		t.Errorf("got %v; want %v", movies, w2)
	}

	// Тест на UpdateMovie
	singleMovie.Title = "New Title"
	err = pg.UpdateMovie(ctx, singleMovie)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}

	movies, err = pg.SelectMovies(ctx, 0)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}

	var updatedMovie model.Movie
	for _, movie := range movies {
		if movie.StudioID == 5 {
			updatedMovie = movie
		}
	}

	if !reflect.DeepEqual(singleMovie, updatedMovie) {
		t.Errorf("got %v; want %v", singleMovie, updatedMovie)
	}

	// Тест на DeleteMovie
	err = pg.DeleteMovie(ctx, singleMovie.ID)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}

	movies, err = pg.SelectMovies(ctx, singleMovie.StudioID)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}

	w3 := []model.Movie{}
	if !reflect.DeepEqual(movies, w3) {
		t.Errorf("got %v; want %v", movies, w3)
	}
}
