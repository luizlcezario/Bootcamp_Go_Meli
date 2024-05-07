1. Como é chamado um JOIN em um banco de dados?
    Join
2. Nomeie e explique 2 tipos de JOIN.
    Inner Join
    Esse é um formato comum de join, que retorna dados apenas quando as duas tabelas tem chaves correspondentes na cláusula ON do join.
    Left Join
    É um dos formatos mais usados de join, que retorna a Tabela A inteira e apenas os registros que coincidirem com a igualdade do join na TabelaB (ou campos nulos para os campos sem correspondência).
3. Para que serve o GROUP BY?
    Uma forma de agrupar tabelas por campos 

4. Para que serve o HAVING?
    E uma forma de filtrar como Where porem usada depois de um group by
5. Dados os diagramas a seguir, indique a qual tipo de JOIN cada um corresponde:
    Left Outer Join
6. Escreva uma consulta genérica para cada um dos diagramas abaixo:
    imagem 1:
    SELECT TabelaA.*, TabelaB.* FROM TabelaA RIGHT JOIN TabelaB ON TabelaA.Chave = TabelaB.Chave
    imagem 2:
    SELECT TabelaA.*, TabelaB.* FROM TabelaA FULL OUTER JOIN TabelaB ON TabelaA.Chave
