SELECT * from movies;
SELECT first_name, last_name, rating FROM actors Order By rating;
SELECT title as titulo from movies as películas;
SELECT first_name, last_name FROM actors WHERE rating >= 7.5;
SELECT title, rating, awards from movies WHERE rating >= 7.5 and awards > 1;
SELECT title from movies Order by rating;
SELECT title from movies limit 3;
SELECT * from movies Order By rating DESC limit 5;
SELECT * from movies Order By rating DESC limit 5,5;
SELECT * from actors limit 10;
SELECT * from actors limit 20,10;
SELECT * from actors limit 40,10;
SELECT title, rating from movies where title like "Toy Story%";
SELECT * from actors WHERE first_name like "Sam%" ;
SELECT * from movies WHERE release_date BETWEEN '2004-01-01' And '2008-12-31';
SELECT title from movies WHERE rating>=3 and awards >= 2 and release_date  BETWEEN '1988-01-01' And '2009-12-31';
SELECT title from movies WHERE rating>=3 and awards >= 2 and release_date  BETWEEN '1988-01-01' And '2009-12-31' limit 3;


