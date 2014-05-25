DROP DATABASE IF EXISTS `blog`;
CREATE DATABASE `blog`;
USE `blog`;

#
#**************************CREATE TABLE****************************************************
#

DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`(
	`id` int(10) PRIMARY KEY AUTO_INCREMENT,
	`name` VARCHAR(100) NOT NULL, 
	`password` VARCHAR(100) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `admin_sessions`;
CREATE TABLE `admin_sessions`(
	`id` int(10) AUTO_INCREMENT,
	`uid` int(10) NOT NULL,
	`session` varchar(50) NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY (`session`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `infos`;
CREATE TABLE `infos`(
	`id`   int(10) PRIMARY KEY AUTO_INCREMENT, 
	`name` VARCHAR(512) NOT NULL, 
	`value` TEXT
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`(
	`id` 	int(10) PRIMARY KEY AUTO_INCREMENT,
	`title`	 VARCHAR(512), 
	`excerpt` TEXT NOT NULL,
	`content` TEXT NOT NULL, 
	`ctime` 	 datetime NOT NULL,
	`tag_id`	 int(10) DEFAULT 1
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`(
	`id` int(10) PRIMARY KEY AUTO_INCREMENT,
	`ip` VARCHAR(10) NOT NULL,
	`comment` TEXT NOT NULL,
	`nick`  VARCHAR(32) NOT NULL,
	`email` VARCHAR(32) NOT NULL,
	`site` VARCHAR(64) NOT NULL,
	`ctime` datetime NOT NULL,
	`article_id` int(10)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`(
	`id`  int(10) PRIMARY KEY AUTO_INCREMENT, 
	`tag` VARCHAR(512) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
#*******************************INSERT*****************************************************
#

INSERT `admins` VALUES(1,'root', '63a9f0ea7bb98050796b649e85481845');
INSERT `infos` VALUES(1,'email', 'macs130828@gmail.com');
INSERT `infos` VALUES(2,'nick', 'Macs.Du');
INSERT `infos` VALUES(3,'about', 
'本网站是为了学习golang和网站开发所建成的，纯属娱乐之用，如果你也喜欢本网站，请移步<a href="https://github.com/macs-du/blog" target="_blank">GitHub</a>');
INSERT `tags`(id, tag) VALUES(1,'原创');
INSERT articles(title,excerpt,content,ctime) VALUES(
	"网站初始化成功",
	'当你看到这篇的时候，就代表着您网站初始化成功',
	'当你看到这篇的时候，就代表着您网站初始化成功。</br>即将开始您的个人博客之路。',
	now()
);