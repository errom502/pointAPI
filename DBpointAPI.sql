
create database pointapi;
create table client (id SERIAL PRIMARY KEY, login varchar(100), password varchar(100));
create table points (id_point SERIAL PRIMARY KEY, point_name TEXT, point_url TEXT,info TEXT, id_user int REFERENCES client(id));

--insert into client (login, password) values ('Loginmega1','sdkwdk2');
--insert into client (login, password) values ('AndUtoo','12345');
--insert into points (point_name, point_url, id_user) values ('Точка 1','https://yandex.ru/maps/-/CCUVEPDEKA',1);

--select * from client, points;



--CREATE TABLE users(id INTEGER PRIMARY KEY, username TEXT NOT NULL);

--CREATE TABLE requests (id_request SERIAL PRIMARY KEY,request TEXT,id_user INTEGER,FOREIGN KEY (id_user) REFERENCES users(id));