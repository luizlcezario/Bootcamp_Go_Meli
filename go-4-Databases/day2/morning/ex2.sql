-- Mostre o título e o nome do gênero de todas as séries.
SELECT movies.title , genres.name from movies  inner Join genres on movies.genre_id = genres.id;

-- Mostre o título dos episódios, o nome e sobrenome dos atores que trabalham em cada um deles.
SELECT episodes.title , actors.first_name , actors.last_name from episodes inner join actor_episode 
on episodes.`number` =  actor_episode.episode_id inner Join actors on actor_episode.actor_id = actors.id;

-- Mostre o título de todas as séries e o número total de temporadas que cada uma delas possui.
SELECT series.title, COUNT(episodes.id)
FROM episodes
INNER JOIN seasons ON episodes.season_id = seasons.id
INNER JOIN series ON seasons.serie_id = series.id
GROUP BY series.title;
-- Mostre o nome de todos os gêneros e o número total de filmes de cada um, desde que seja maior ou igual a 3.
SELECT genres.name, COUNT(movies.id)
FROM genres
INNER JOIN movies ON genres.id  = movies.genre_id 
GROUP BY genres.name HAVING  COUNT(movies.id) >= 3;
	
	-- Mostre apenas o nome e sobrenome dos atores que atuam em todos os filmes de Star Wars e que estes não se repitam.

SELECT DISTINCT actors.first_name, actors.last_name
FROM actors
INNER JOIN actor_movie ON actors.id = actor_movie.actor_id 
INNER JOIN movies ON actor_movie.movie_id = movies.id
WHERE movies.title like "La Guerra de las galaxias%"
GROUP BY actors.id
HAVING COUNT(movies.id) = (SELECT COUNT(*) FROM movies WHERE title LIKE 'La Guerra de las galaxias%');
