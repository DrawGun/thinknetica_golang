INSERT INTO studios (name) VALUES
    ('Неизвестная Studio'),
    ('Studio 1'),
    ('Studio 2'),
    ('Studio 3'),
    ('Studio 4'),
    ('Studio 5'),
    ('Studio 6'),
    ('Studio 7'),
    ('Studio 8'),
    ('Studio 9'),
    ('Studio 10');

INSERT INTO ratings (title) VALUES
    ('Без рейтинга'),
    ('PG-10'),
    ('PG-13'),
    ('PG-18');

INSERT INTO actors (first_name, second_name, birthday) VALUES
    ('Какой-то', 'Актер', '1978-01-1'),
    ('Actor', 'One', '1978-01-1'),
    ('Actor', 'Two', '1959-11-26'),
    ('Actor', 'Three', '1993-10-18'),
    ('Actor', 'Four', '1988-05-11'),
    ('Actor', 'Five', '1991-12-07');

INSERT INTO directors (first_name, second_name, birthday) VALUES
    ('Какой-то', 'Режиссёр', '1978-01-1'),
    ('Director', 'One', '1978-01-1'),
    ('Director', 'Two', '1959-11-26'),
    ('Director', 'Three', '1993-10-18'),
    ('Director', 'Four', '1988-05-11'),
    ('Director', 'Five', '1991-12-07');

INSERT INTO movies (title, release_year, rating_id, fees, studio_id) VALUES
    ('Неизвестное кино', '1994', 1, 0, 1),
    ('Movie 1', '1994', 1, 2000, 2),
    ('Movie 2', '1999', 1, 500, 3),
    ('Movie 3', '2000', 2, 4000, 5),
    ('Movie 4', '2001', 2, 1000, 8),
    ('Movie 5', '1990', 3, 990, 4),
    ('Movie 6', '1945', 1, 1500, 7),
    ('Movie 7', '2019', 3, 2500, 9),
    ('Movie 8', '2020', 2, 307, 3),
    ('Movie 9', '2020', 1, 2003, 6),
    ('Movie 10', '2020', 2, 2022, 10);

INSERT INTO actors_movies (actor_id, movie_id) VALUES
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 4),
    (5, 5);

INSERT INTO directors_movies (director_id, movie_id) VALUES
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 4),
    (5, 5),
    (1, 6),
    (2, 7),
    (3, 8),
    (4, 9),
    (5, 10);