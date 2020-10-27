DROP DATABASE IF EXISTS DB_Notes;
CREATE DATABASE IF NOT EXISTS DB_Notes;
USE DB_Notes;

CREATE TABLE IF NOT EXISTS tb_users (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY ,
    name VARCHAR( 20 ) NOT NULL ,
    last_name VARCHAR( 30 ) NOT NULL ,
    nick_name VARCHAR( 10 ) NOT NULL ,
    profile LONGBLOB,
    email VARCHAR( 40 ) NOT NULL ,
    password VARCHAR( 40 ) NOT NULL
);

CREATE TABLE IF NOT EXISTS tb_typeOfnote (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY ,
    type VARCHAR( 20 ) NOT NULL
);

CREATE TABLE IF NOT EXISTS tb_notes (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY ,
    title VARCHAR( 30 ) NOT NULL ,
    body TEXT NOT NULL ,
    id_user INT NOT NULL,
    id_type INT NOT NULL ,
    FOREIGN KEY ( id_user ) REFERENCES tb_users ( id ) ,
    FOREIGN KEY ( id_type ) REFERENCES tb_typeOfnote ( id )
);


