drop DATABASE if exists `blog`;

CREATE DATABASE blog;

USE blog;

drop table if exists `admins`;

CREATE TABLE admins(id int(10) PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL);

INSERT INTO admins(name, password) values("root", "63a9f0ea7bb98050796b649e85481845");

drop table if exists `blogs`;

CREATE TABLE articles(id int(10) PRIMARY KEY AUTO_INCREMENT,title VARCHAR(512), content TEXT NOT NULL, ctime datetime NOT NULL, tag_id int(10) NOT NULL);

drop table if exists `comments`;

CREATE TABLE comments(id int(10) PRIMARY KEY AUTO_INCREMENT, ip VARCHAR(10) NOT NULL, comment TEXT NOT NULL, ctime datetime NOT NULL, article_id int(10) NOT NULL);

drop table if exists `tags`;

CREATE TABLE tags(id int(10) PRIMARY KEY AUTO_INCREMENT, tag VARCHAR(512) NOT NULL);
