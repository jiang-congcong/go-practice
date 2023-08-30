DROP DATABASE IF EXISTS go_db;
CREATE DATABASE IF NOT EXISTS go_db;
USE go_db;

CREATE TABLE IF NOT EXISTS user(
	id bigint not null auto_increment,
	username varchar(50) not null,
	`password` varchar(30) not null,
    age int,
    tel varchar(20),
    create_time timestamp null,
    creater_id bigint,
    modify_time timestamp null,
    modifier_id bigint,
	primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS books(
    id bigint auto_increment not null,
    book_name varchar(100) not null,
    author varchar(100),
    publish_time timestamp null,
    create_time timestamp null,
    update_time timestamp null,
    primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4