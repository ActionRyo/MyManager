define("fishstrap/page/loginPage.js",function(t){var i=t("fishstrap/core/global.js"),o=t("fishstrap/ui/dialog.js");return{use:function(t){t=t||{};var n={title:"后台管理系统",login:function(){},init:function(){},button:[]};for(var a in t)n[a]=t[a];for(var d="",a=0;a!=n.button.length;++a)n.button[a].id=i.uniqueNum(),d+='<a class="rightbutton" href="#" id="'+n.button[a].id+'">'+n.button[a].name+"</a>";var e=i('<div class="container"><form class="form-signin" method="post" id="login-form"><div class="title">烘焙帮微信后台管理系统</div><div class="name"><input class="input-small" type="text" name="name" placeholder="账号"></div><div class="password"><input class="input-small" type="password" name="password" placeholder="密码"></div><div class="button"><button type="button" class="btn submit">登录</button>'+d+"</div></form></div>");i("body").append(e),i.addCssToHead('body {	padding-top: 40px;	padding-bottom: 40px;	background-color: #f5f5f5;}.title{	font-size:16px;	margin-bottom:20px;	text-align:center;	margin-top:10px;}.form-signin {	max-width: 420px;	padding: 5px 10px;	margin: 0 auto 20px;	background-color: #fff;	border: 1px solid #e5e5e5;	-webkit-border-radius: 5px;	-moz-border-radius: 5px;	border-radius: 5px;	-webkit-box-shadow: 0 1px 2px rgba(0, 0, 0, .05);	-moz-box-shadow: 0 1px 2px rgba(0, 0, 0, .05);	box-shadow: 0 1px 2px rgba(0, 0, 0, .05);	overflow:auto;}.name,.password{	width:80%;	margin:0 auto;}.form-signin input[type="text"],.form-signin input[type="password"] {	width:100%;	margin-bottom: 15px;	height: auto;	padding: 7px 9px;}.button{	overflow:auto;	zoom:1;	padding-left:10%;	padding-right:10%;	margin-bottom:15px;}.rightbutton{	margin-top:10px;	float:right;}.submit{	float:left;	padding-left:20px;	padding-right:20px;}'),e.find(".title").text(t.title),e.find("input").keydown(function(o){var o=o||event;if(13==o.keyCode){var n={name:i("input[name=name]").val(),password:i("input[name=password]").val()};t.login(n)}}),e.find(".submit").click(function(){var n={name:i.trim(i("input[name=name]").val()),password:i.trim(i("input[name=password]").val())};return""==n.name?void o.message("请输入帐号"):""==n.password?void o.message("请输入密码"):void t.login(n)});for(var a=0;a!=n.button.length;++a)i("#"+n.button[a].id).click(n.button[a].click);t.init()}}});