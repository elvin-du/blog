drop DATABASE if exists `blog`;

CREATE DATABASE blog;

USE blog;

drop table if exists `admins`;

CREATE TABLE admins(id int(10) PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL);

INSERT INTO admins(name, password) values("root", "63a9f0ea7bb98050796b649e85481845");

drop table if exists `blogs`;

CREATE TABLE articles(id int(10) PRIMARY KEY AUTO_INCREMENT,title VARCHAR(512), content TEXT NOT NULL, ctime datetime NOT NULL, tag_id int(10));

drop table if exists `comments`;

CREATE TABLE comments(id int(10) PRIMARY KEY AUTO_INCREMENT, ip VARCHAR(10) NOT NULL, comment TEXT NOT NULL, ctime datetime NOT NULL, article_id int(10));

drop table if exists `tags`;

CREATE TABLE tags(id int(10) PRIMARY KEY AUTO_INCREMENT, tag VARCHAR(512) NOT NULL);

INSERT INTO articles(title,content,ctime) VALUES(
	"目前采用首先把目录下所有的文件进行缓存，所以用户还可以通过类似这样",
	"对于一个复杂的 LayoutContent，其中可能包括有javascript脚本、CSS 引用等，根据惯例，通常 css 会放到 Head 元素中，javascript 脚本需要放到 body 元素的末尾，而其它内容则根据需要放在合适的位置。在 Layout 页中仅有一个 LayoutContent 是不够的。所以在 Controller 中增加了一个 LayoutSections属性，可以允许 Layout 页中设置多个 section，然后每个 section 可以分别包含各自的子模板页。",
	now()
);