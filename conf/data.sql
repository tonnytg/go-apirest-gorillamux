create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES ('teste1', 'historia1'),('teste2', 'historia2');
