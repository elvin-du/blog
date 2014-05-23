## blog

> A private blog with golang and html5 on beego frame.

## 编译安装说明：

获取源代码，下载完成后会自动编译为blog可执行文件
	
>$ go get github.com/macs-du/blog

### 修改配置文件
	
>$ vim github.com/macs-du/blogconf/app.conf

> appname = blog

> httpport = 8080

> runmode = dev

> mysqluser = "yourusername"

> mysqlpass = "yourpassword"

> mysqlurls = "127.0.0.1"

> mysqldb   = "blog"


### 导入MySQL

> 登录mysql后，执行如下命令：

> source  /yourpath/init.sql

### 运行
	
> $ ./blog

> 或

> $ nohup ./blog 2>&1 > blog.log &

> 设为后台运行

### 访问： 

>http://localhost:8080

### 后台地址：

> http://localhost:8080/admin/login

> 帐号：root

> 密码：root


