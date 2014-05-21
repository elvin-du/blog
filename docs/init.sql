DROP DATABASE IF EXISTS `blog`;
CREATE DATABASE blog;
USE blog;

DROP TABLE IF EXISTS `admins`;
CREATE TABLE admins(
	id int(10) PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL, 
	password VARCHAR(100) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `info`;
CREATE TABLE info(
	id int(10) PRIMARY KEY AUTO_INCREMENT, 
	name VARCHAR(512) NOT NULL, 
	info TEXT
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `admin_sessions`;
CREATE TABLE `admin_sessions`(
	`id` int(10) AUTO_INCREMENT,
	`session` varchar(50) NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY (`session`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`(
	id int(10) PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(512), 
	content TEXT NOT NULL, 
	ctime datetime NOT NULL,
	tag_id int(10) DEFAULT 1
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`(
	id int(10) PRIMARY KEY AUTO_INCREMENT,
	ip VARCHAR(10) NOT NULL,
	comment TEXT NOT NULL,
	ctime datetime NOT NULL,
	article_id int(10)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`(
	id int(10) PRIMARY KEY AUTO_INCREMENT, 
	tag VARCHAR(512) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO admins(name, password) values("root", "63a9f0ea7bb98050796b649e85481845");

INSERT tags(id, tag) VALUES(1,"原创");

INSERT info(name, info) VALUES("email", "macs130828@gmail.com");
INSERT info(name, info) VALUES("nick", "Macs.Du");
INSERT info(name, info) VALUES(
"about", 
'本网站是为了学习golang和网站开发所建成的，纯属娱乐之用，如果你也喜欢本网站，请移步<a href="https://github.com/macs-du/blog" target="_blank">GitHub</a>'
);

INSERT INTO articles(title,content,ctime) VALUES(
	"网站初始化成功",
	"当你看到这篇的时候，就代表着您网站初始化成功。即将开始您的个人博客之路。",
	now()
);