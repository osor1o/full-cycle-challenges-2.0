CREATE DATABASE IF NOT EXISTS homestead;

USE homestead;

CREATE TABLE IF NOT EXISTS people (
  id      INT             NOT NULL  AUTO_INCREMENT,
  name    VARCHAR(255)    NOT NULL,

  PRIMARY KEY(id)
);
