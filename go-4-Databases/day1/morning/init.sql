-- This script is used to create the tables and the data for the first day of the course
	CREATE TABLE IF NOT EXISTS employes (
		id VARCHAR(36) default(uuid_to_bin(uuid())) NOT NULL Primary key,
		file Int NOT NULL AUTO_INCREMENT unique,
		sour_name VARCHAR(255),
		name VARCHAR(255),
		rolef VARCHAR(255),
		salary INT,
		birthday  DATE,
		entry_day DATE,
	);
	
	CREATE TABLE IF NOT EXISTS department (
		id VARCHAR(36) default(uuid_to_bin(uuid())) NOT NULL Primary key,
		name VARCHAR(255) NOT NULL unique,
		endereço VARCHAR(255),
		cidade VARCHAR(255)
	);
	

	CREATE TABLE IF NOT EXISTS roles (
		id VARCHAR(36) default(uuid_to_bin(uuid())) NOT NULL Primary key,
		name VARCHAR(255) NOT NULL unique,
		salary_limit INT,
		salary_base INt
	);
	
	ALTER TABLE employes  ADD COLUMN dep_fk varchar(36);
	ALTER TABLE employes  DROP COLUMN rolef;
	ALTER TABLE employes  ADD COLUMN role_fk varchar(36);
	ALTER TABLE employes ADD FOREIGN KEY (dep_fk) REFERENCES department(id);
	ALTER TABLE employes ADD FOREIGN KEY (role_fk) REFERENCES roles(id);

    -- Ex 3
 
    INSERT INTO department VALUES 
    ("de61092c-0bc0-11ef-bcb0-0242c0a82002","TI","Avenida Anchieta", "Santo Andre"),
    ("de61092c-0bc0-11ef-bcb0-0242c0a82002","MKT", "Avenida Paulista", "São Paulo"),
    ("de6131b5-0bc0-11ef-bcb0-0242c0a82002","RH", "Avenida Faria Lima", "São Paulo"),
    ("de613288-0bc0-11ef-bcb0-0242c0a82002","Financeiro", "Avenida Brigadeiro Faria Lima", "São Paulo"),(uuid(),"Comercial", "Avenida Paulista", "São Paulo");

    INSERT INTO roles VALUES 
    ("95f0ffa2-0bc1-11ef-bcb0-0242c0a82002", "Analista", 10000, 5000),
    ("95f115eb-0bc1-11ef-bcb0-0242c0a82002", "Desenvolvedor", 15000, 7000),
    ("95f11806-0bc1-11ef-bcb0-0242c0a82002", "Gerente", 20000, 10000),
    ("95f1191f-0bc1-11ef-bcb0-0242c0a82002", "Diretor", 30000, 15000),
    ("95f119c8-0bc1-11ef-bcb0-0242c0a82002", "Estagiario", 5000, 2000);

	INSERT Into employes VALUES
    (uuid(), 1, "João", "João", 10000, "1990-01-01", "2022-01-01", "de61092c-0bc0-11ef-bcb0-0242c0a82002", "95f115eb-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 2, "Maria", "Maria", 15000, "1990-01-01", "2022-01-01", "de61092c-0bc0-11ef-bcb0-0242c0a82002", "95f115eb-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 3, "Pedro", "Pedro", 20000, "1990-01-01", "2022-01-01", "de612f84-0bc0-11ef-bcb0-0242c0a82002", "95f115eb-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 4, "Lucas", "Lucas", 30000, "1990-01-01", "2022-01-01", "de612f84-0bc0-11ef-bcb0-0242c0a82002", "95f115eb-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 5, "Ana", "Ana", 5000, "1990-01-01", "2022-01-01", "de612f84-0bc0-11ef-bcb0-0242c0a82002", "95f119c8-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 6, "Carlos", "Carlos", 10000, "1990-01-01", "2022-01-01", "de612f84-0bc0-11ef-bcb0-0242c0a82002", "95f119c8-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 7, "Julia", "Julia", 15000, "1990-01-01", "2022-01-01", "de6131b5-0bc0-11ef-bcb0-0242c0a82002", "95f119c8-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 8, "Fernanda", "Fernanda", 20000, "1990-01-01", "2022-01-01", "de6131b5-0bc0-11ef-bcb0-0242c0a82002", "95f119c8-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 9, "Rafael", "Rafael", 30000, "1990-01-01", "2022-01-01", "de6131b5-0bc0-11ef-bcb0-0242c0a82002", "95f0ffa2-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 10, "Mariana", "Mariana", 5000, "1990-01-01", "2022-01-01", "de613288-0bc0-11ef-bcb0-0242c0a82002", "95f0ffa2-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 11, "Gustavo", "Gustavo", 10000, "1990-01-01", "2022-01-01", "de613288-0bc0-11ef-bcb0-0242c0a82002", "95f11806-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 12, "Larissa", "Larissa", 15000, "1990-01-01", "2022-01-01", "de613288-0bc0-11ef-bcb0-0242c0a82002", "95f11806-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 13, "Roberto", "Roberto", 20000, "1990-01-01", "2022-01-01", "de61339b-0bc0-11ef-bcb0-0242c0a82002", "95f11806-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 14, "Patricia", "Patricia", 30000, "1990-01-01", "2022-01-01", "de61339b-0bc0-11ef-bcb0-0242c0a82002", "95f1191f-0bc1-11ef-bcb0-0242c0a82002");
    (uuid(), 15, "Ricardo", "Ricardo", 5000, "1990-01-01", "2022-01-01", "de61339b-0bc0-11ef-bcb0-0242c0a82002", "95f0ffa2-0bc1-11ef-bcb0-0242c0a82002");
