DROP DATABASE IF EXISTS `blog`;
CREATE DATABASE `blog`;
USE `blog`;

################CREATE TABLE#############################
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`(
	`id` int(16) PRIMARY KEY AUTO_INCREMENT,
	`name` VARCHAR(128) NOT NULL, 
	`password` VARCHAR(128) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `admin_sessions`;
CREATE TABLE `admin_sessions`(
	`id` int(16) AUTO_INCREMENT,
	`uid` int(16) NOT NULL,
	`session` varchar(64) NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY (`session`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `infos`;
CREATE TABLE `infos`(
	`id`   int(16) PRIMARY KEY AUTO_INCREMENT, 
	`name` VARCHAR(512) NOT NULL, 
	`value` TEXT
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`(
	`id` 	int(16) PRIMARY KEY AUTO_INCREMENT,
	`title`	 VARCHAR(512), 
	`excerpt` TEXT NOT NULL,
	`content` TEXT NOT NULL, 
	`ctime` 	 datetime NOT NULL,
	`read_count` int(64) DEFAULT 0,
	`comment_count` int(64) DEFAULT 0
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`(
	`id` int(16) PRIMARY KEY AUTO_INCREMENT,
	`ip` VARCHAR(64) NOT NULL,
	`comment` TEXT NOT NULL,
	`nick`  VARCHAR(64) NOT NULL,
	`email` VARCHAR(32) NOT NULL,
	`site` VARCHAR(64) NOT NULL,
	`ctime` datetime NOT NULL,
	`article_id` int(16)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`(
	`id`  int(16) PRIMARY KEY AUTO_INCREMENT, 
	`tag` VARCHAR(32) NOT NULL UNIQUE,
	`count` int(64) DEFAULT 0
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `article_tag_relation`;
CREATE TABLE `article_tag_relation`(
	`id` int(16) PRIMARY KEY AUTO_INCREMENT,
	`tag_id` int(16) NOT NULL,
	`article_id` int(16) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `visitors`;
CREATE TABLE `visitors`(
	`id` int(16) PRIMARY KEY AUTO_INCREMENT,
	`ip` VARCHAR(64) NOT NULL,
	`atime` datetime NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `read_history`;
CREATE TABLE `read_history`(
	`id` int(16) PRIMARY KEY AUTO_INCREMENT,
	`article_id` int(16) NOT NULL,
	`ip` VARCHAR(64) NOT NULL,
	`atime` datetime NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

######################INSERT##################################
INSERT `admins` VALUES(1,'root', '63a9f0ea7bb98050796b649e85481845');
INSERT `infos` VALUES(1,'email', 'macs130828@gmail.com');
INSERT `infos` VALUES(2,'nick', 'Macs.Du');
INSERT `infos` VALUES(3,'about', 
	'本网站是为了学习golang和网站开发所建成的，纯属娱乐之用，如果你也喜欢本网站，
	请移步<a href="https://github.com/macs-du/blog" target="_blank">GitHub</a>'
	);
INSERT `tags` VALUES(1,'其他',0);

######################TRIGGER##################################
DELIMITER |
CREATE TRIGGER read_count_trigger AFTER INSERT ON read_history FOR EACH ROW
BEGIN
UPDATE articles SET read_count = read_count+1 WHERE id = NEW.article_id;
END|
DELIMITER ;

DELIMITER |
CREATE TRIGGER comment_count_trigger AFTER INSERT ON comments FOR EACH ROW
BEGIN
UPDATE articles SET comment_count = comment_count +1 WHERE id = NEW.article_id;
END|
DELIMITER ;

DELIMITER |
CREATE TRIGGER article_tag_relation_trigger AFTER INSERT ON article_tag_relation FOR EACH ROW
BEGIN
UPDATE tags SET count=count+1 WHERE id=NEW.tag_id;
END|
DELIMITER ;