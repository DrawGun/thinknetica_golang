DROP TABLE IF EXISTS movies, actors, directors, studios, ratings, actors_movies, directors_movies;

CREATE TABLE studios (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL DEFAULT 'Без рейтинга',
    UNIQUE(title)
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    release_year INTEGER NOT NULL DEFAULT 1800 CHECK ( release_year >= 1800 ),
    rating_id INTEGER REFERENCES ratings(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 1,
    fees BIGINT NOT NULL DEFAULT 0,
    studio_id INTEGER REFERENCES studios(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 1,
    UNIQUE(release_year, title)
);

CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    second_name TEXT NOT NULL,
    birthday DATE NOT NULL
);

CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    second_name TEXT NOT NULL,
    birthday DATE NOT NULL
);

CREATE TABLE actors_movies (
    id BIGSERIAL PRIMARY KEY,
    actor_id INTEGER REFERENCES actors(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 1,
    movie_id INTEGER REFERENCES movies(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 1,
    UNIQUE(actor_id, movie_id)
);

CREATE TABLE directors_movies (
    director_id INTEGER REFERENCES directors(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 1,
    movie_id INTEGER REFERENCES movies(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 1,
    UNIQUE(director_id, movie_id)
);
