insert into images (name, image_path, type)
values ('picture.png', '
/Users/andrew/Programming/website/go/goWebsite/pages/pic.png',
 'png');

insert into css (name,doc)
values ('styles.css',
'.header{
	display: grid;
	background-color: darkmagenta;
	grid-template-columns: 100px 1fr 200px;
	padding-left: 15px;
	padding-right: 15px;
}
body{
	margin: 0;
	padding: 0;

}


nav a,
.mainlink a{
	color: white;
	text-decoration: none;
	font-size: 2em;
}

nav{
	grid-column: 3;
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-gap: 10px;

}

html{
	background-color: rgb(73, 23, 73);
	color: rgb(235, 235, 235);
}

.content{
	max-width: 50rem;
	margin-top: 50px;
	margin-left: auto;
	margin-right: auto;
	font-size: large;
	
}
.post{
	grid-column: 2;
}

h1{
	text-align: center;
}');

insert into blog_posts (id, name, post)
values (
	DEFAULT,
	'home.html',
	'<!DOCTYPE html>
<html>

<head>{{template "head" .}}



<body>
{{template "content" .}}

</body>

</html>'
),
(DEFAULT,
'404.html',
'{{define "content"}}page not found{{end}}');

insert into blog_posts (id, name, post)
values (
	DEFAULT,
	'body.html',
	'{{define "content"}}
<div class="content">
	<div class="post">
<h1>SOME AWESOME TITLE GOES HERE!!</h1>
	<p> Et quaerat fugiat consequatur voluptas reiciendis quia deserunt. Illo quo quaerat nulla dolorem odit. Blanditiis
		et accusamus laudantium sunt. Perferendis omnis eaque dolorum. Ratione vero quidem consequuntur omnis quos
		reprehenderit. Sunt in accusantium laudantium nam quibusdam est illum non.
	</p>
	<p>
		Nobis eaque ut sed nihil alias. Accusantium animi laboriosam placeat assumenda porro ut possimus sunt. Excepturi
		quidem sint eos.
	</p>
	<p>
		Eius non illo accusantium eveniet quaerat. Quas adipisci qui sed qui voluptatibus sit. Ipsa voluptatum
		temporibus perspiciatis enim similique eveniet dicta officiis. Perspiciatis officia numquam quia et. Praesentium
		sed fuga nam. Molestiae quod corporis enim rem nostrum.
	</p>
	<img src="pic.png">
	<p>
		Doloribus aut sapiente voluptatem. Illum ab doloremque sunt ut ducimus aperiam. Aut voluptatem explicabo sed
		porro accusamus pariatur sed dicta. Minus id aperiam aut et nam. Quibusdam quae blanditiis rerum et nemo ut
		voluptatum vero.
	</p>
	<p>
		Excepturi magnam qui error unde maiores aut. Velit qui incidunt qui sed totam perspiciatis accusamus magni. Id
		ullam aut est eveniet nesciunt ut qui. Ipsa quo illum perferendis mollitia. Facilis at culpa voluptates est et.
		Ad magni in accusantium veniam cupiditate.
	</p>
	</div>
</div>{{end}}'
),
(DEFAULT,
'head.html',
'{{define "head"}}
<title> title</title>
<meta name="viewport" content="width=device-width">
<!-- remove this next line once going live -->
<!--<script type="text/javascript" src="http://livejs.com/live.js"></script>-->
<link rel="stylesheet" href="styles.css">
<link rel="icon" href="favicon.svg">
</head>
<div class="header">
	<div class="mainlink">
		<a href="https://wieds.ca">Wieds</a>
	</div>
	<nav>
		<a href="http://localhost:8080/about">About</a>
		<a href="https://wieds.ca/projects">Projects</a>
	</nav>
</div>{{end}}');