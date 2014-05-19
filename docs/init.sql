drop DATABASE if exists `blog`;
CREATE DATABASE blog;
USE blog;

drop table if exists `admins`;
CREATE TABLE admins(id int(10) PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL);
INSERT INTO admins(name, password) values("root", "63a9f0ea7bb98050796b649e85481845");

DROP TABLE IF EXISTS `admin_sessions`;
CREATE TABLE `admin_sessions`(
	`id` int(10) AUTO_INCREMENT,
	`session` varchar(50) NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE KEY (`session`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists `blogs`;
CREATE TABLE articles(id int(10) PRIMARY KEY AUTO_INCREMENT,title VARCHAR(512), content TEXT NOT NULL, ctime datetime NOT NULL, tag_id int(10));

drop table if exists `comments`;
CREATE TABLE comments(id int(10) PRIMARY KEY AUTO_INCREMENT, ip VARCHAR(10) NOT NULL, comment TEXT NOT NULL, ctime datetime NOT NULL, article_id int(10));

drop table if exists `tags`;
CREATE TABLE tags(id int(10) PRIMARY KEY AUTO_INCREMENT, tag VARCHAR(512) NOT NULL);

drop TABLE IF EXISTS `info`;
CREATE TABLE info(id int(10) PRIMARY KEY AUTO_INCREMENT, name VARCHAR(512) NOT NULL, info TEXT);

INSERT info(name, info) VALUES("email", "macs130828@gmail.com");
INSERT info(name, info) VALUES("nick", "Macs.Du");
INSERT info(name, info) VALUES("about", 
'本网站是为了学习golang和网站开发所建成的，纯属娱乐之用，如果你也喜欢本网站，请移步<a href="http://www.github.com/">GitHub</a>'
);

INSERT INTO articles(title,content,ctime) VALUES(
	"目前采用首先把目录下所有的文件进行缓存，所以用户还可以通过类似这样",
	"对于一个复杂的 LayoutContent，其中可能包括有javascript脚本、CSS 引用等，根据惯例，通常 css 会放到 Head 元素中，javascript 脚本需要放到 body 元素的末尾，而其它内容则根据需要放在合适的位置。在 Layout 页中仅有一个 LayoutContent 是不够的。所以在 Controller 中增加了一个 LayoutSections属性，可以允许 Layout 页中设置多个 section，然后每个 section 可以分别包含各自的子模板页。",
	now()
);

INSERT INTO articles(title,content,ctime) VALUES(
	"Fxh.Go的新计划",
	"原先使用JSON存储，管理和备份方便，但是查询复杂。尤其是排序和关联查询，很是费劲。虽然说博客数据都缓存在内存中，但是每次关联查询都要拆分成单条数据的操作，增加许多的代码。现在是知道SQL数据库的好处啦。
SQL数据库问题也很多。Fxh.Go一直是纯go程序，不需要安装任何依赖（除了Go语言运行的依赖比如libc）。但SQL数据库有名的MySQL，需要另行安装，SQLite又依赖cgo，非常郁闷。后来发现了ql——纯Go的SQL实现。感觉可以尝试一下。
既然用了数据库，ORM几乎是必须的，否则Rows.Scan这种API是要人命的。xorm的作者lunny也是gogits的开发者之一。他也打算让xorm支持ql。这个很值得期待",
	now()
);

INSERT INTO articles(title,content,ctime) VALUES(
	"Gogs：用二进制才是真正的部署",
	"首先，请允许我代表开发团队感谢所有在 GitHub 上支持 Gogs 的同学。要知道，v0.2.0 是 Gogs 的首个公开发布版本，而在这之前一周的时间里，该项目已经获得了超过 650 个 Star。然后，我代表个人致以开发团队所有成员最诚挚的敬意，每个成员都为首个版本的发布做出了非常大的努力，是我们的团结一心和默契配合成就了 Gogs 这个项目的建立与成长。项目概述既然是首个公开发布版本，那么自然有必要对 Gogs 这个项目进行一定的说明，让大家更好地了解我们为什么开发 Gogs、是如何进行开发的，以及目前的开发状况。",
	now()
);