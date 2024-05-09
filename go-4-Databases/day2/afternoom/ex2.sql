-- 1. Explique o conceito de normalização e por que ele é usado.
-- A normalização é o processo de organização de um banco de dados em tabelas para evitar a redundância de dados e garantir a integridade dos dados. Ela é usada para evitar a duplicação de informações e garantir que cada dado seja armazenado apenas uma vez, o que facilita a manutenção e a atualização dos dados.
-- 2. Adicione um filme à tabela de filmes.
	INSERT INTO movies VALUES (0, now() ,NULL,'Avatar 2',7.9,3,'2024-10-04 00:00:00', 120,5)
-- 3. Adicione um gênero à tabela de gêneros.
	INSERT INTO genres VALUES (0, now() ,NULL,'Tests', 13 , 1)
-- 4. Associe o filme do Ex 2. ao gênero criado no Ex 3.
	UPDATE movies SET genre_id = 13 WHERE title = 'Avatar 2';
-- 5. Modifique a tabela de atores para que pelo menos um ator tenha o filme adicionado no Ex.2 como favorito.
	INSERT INTO actor_movie VALUES (0, 1, );
-- 6. Crie uma cópia de tabela temporária da tabela de filmes.
	CREATE TEMPORARY TABLE temp_movie as select * from movies;
-- 7. Elimine dessa tabela temporária todos os filmes que ganharam menos de 5 prêmios.
	DELETE from temp_movie where awards < 5;
-- 8. Obtenha a lista de todos os gêneros que possuem pelo menos um filme
	select distinct g.* from genres g join movies m on g.id = m.genre_id;
-- 9. Obtenha a lista de atores cujo filme favorito ganhou mais de 3 prêmios.
	select distinct a.* from actors a join actor_movie am on a.id = am.actor_id join movies m on am.movie_id = m.id where m.awards > 3;
-- 10. Use o plano de explicação para analisar as consultas em Ex.6 e 7.
	EXPLAIN SELECT * FROM temp_movie;
	EXPLAIN DELETE FROM temp_movie WHERE awards < 5;
-- 11. O que são índices? Para que servem?
	Sao valores de chaves primarias para serem usados em filtros e outras situacoes devem ser unicos
-- 12. Crie um índice no nome na tabela de filmes.
	CREATE INDEX idx_name ON movies (title);
-- 13. Verifique se o índice foi criado corretamente.
	HELP INDEX movies;

