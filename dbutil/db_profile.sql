DROP DATABASE IF EXISTS db_profile;
CREATE DATABASE db_profile;
USE db_profile;

CREATE TABLE IF NOT EXISTS tb_user (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY ,
    name VARCHAR ( 60 ) NOT NULL ,
    last_name VARCHAR ( 60 ) NOT NULL ,
    phone VARCHAR ( 30 ) NOT NULL UNIQUE ,
    email VARCHAR ( 40 ) NOT NULL UNIQUE ,
    password VARCHAR ( 50 ) NOT NULL ,
    profile LONGBLOB
);

INSERT INTO tb_user ( name, last_name, phone, email, password, profile )
VALUES
       (?, ?, ?, ?, ?, ?);

SELECT
       id, name, last_name, phone,
       email, password, profile
FROM tb_user;

SELECT
       id, name, last_name,
       phone, email, profile
FROM tb_user
WHERE
      email = ?
AND
      password = ?;

DELETE FROM tb_user
WHERE
      email = ?
  AND
      password = ?;

UPDATE tb_user
SET
    name = ?, last_name = ?,
    phone = ?, profile = ?
WHERE id = ?;