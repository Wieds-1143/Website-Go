create table blog_posts (
	id		smallserial primary key,
	name	varchar(32) not null,
	post	text not null
);

create table visit_log(
	uuid	uuid primary key,
	path	varchar(256),
	ip		varchar(64),
	userAgent varchar(256),
	date	timestamp
);

create table images (
	name		varchar(32) primary key,
	image_path	varchar(256) not null,
	type		varchar(10) not null
);

create table css (
	name	varchar(32) primary key,
	doc		text not null
);