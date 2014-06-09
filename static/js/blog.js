function login(){
	var name = document.login_form.name.value;
	var passwd = document.login_form.passwd.value;
	if (name.length==0 || passwd.length==0) {
		alert("用户名或密码为空！");
		return false;
	};
	return true;
};

function add(){
	var title = document.add_form.title.value;
	KindEditor.instances[0].sync(); 
	if (title.length==0 || KindEditor.instances[0].isEmpty()) {
		alert("标题或内容为空！");
		return false;
	};
	return true;
};

function update_about(){
	KindEditor.instances[1].sync(); 
	if (KindEditor.instances[1].isEmpty()) {
		alert("内容为空！");
		return false;
	};
	return true;
};

KindEditor.ready(function(K) {
	var content = K.create('textarea[name="content"]', {
		uploadJson : '/admin/file/upload',
	});

	var about = K.create('textarea[name="about"]', {
		uploadJson : '/admin/file/upload',
	});
});

function search_check(){
	var key = $(div#key).val()
	if key.length == 0{
		alert("搜索内容为空");
		return false;
	}
	return true;
};