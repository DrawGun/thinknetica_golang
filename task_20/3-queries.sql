-- выборка фильмов с названием студии; 
SELECT movies.title, pc.name as company
FROM movies
    JOIN studios pc on movies.studio_id = pc.id;

-- выборка фильмов для нескольких режиссёров из списка (подзапрос)
SELECT DISTINCT movies.title
FROM movies JOIN directors_movies AS dm on movies.id = dm.movie_id
WHERE dm.director_id IN (
    SELECT d.id
    FROM directors AS d
    WHERE d.first_name IN ('Director')
    AND d.second_name IN ('Two', 'Four', 'Five')
)
ORDER BY movies.title;

-- подсчёт количества фильмов со сборами больше 1000
SELECT count(*) FROM movies
WHERE fees > 1000;

-- подсчитать количество режиссёров, фильмы которых собрали больше 1000
SELECT count(DISTINCT dm.director_id)
FROM directors_movies AS dm
WHERE movie_id IN (
    SELECT m.id
    FROM movies AS m
    WHERE m.fees > 1000
);

-- подсчёт количества фильмов, имеющих дубли по названию
SELECT count(*)
FROM (
    SELECT mt.title
    FROM movies AS mt
    GROUP BY title
    HAVING count(*) > 1
) AS mc;