create database zhihu;

create table user (
    id bigint(20) primary key auto_increment not null,
    user_id bigint(20) not null,
    username varchar(64) not null,
    nickname varchar(64) not null default '',
    password varchar(64) not null,
    email varchar(64),
    phone varchar(64),
    sex tinyint(4) not null default '0',
    create_time timestamp null default current_timestamp,
    update_time timestamp null default current_timestamp on update current_timestamp,
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
);

create table question (
    id bigint(20) primary key auto_increment not null,
    question_id bigint(20) not null,
    title varchar(128) not null,
    content varchar(8192) not null,
    status tinyint(4) not null default '0',
    author_id bigint(20) not null,
    category_id bigint(20) not null,
    create_time timestamp null default current_timestamp,
    update_time timestamp null default current_timestamp ON UPDATE current_timestamp,
    KEY `idx_author_id` (`author_id`) USING BTREE,
    KEY `idx_category_id` (`category_id`) USING BTREE,
    KEY `idx_question_id` (`question_id`) USING BTREE
)

create table category (
    id int(11) primary key auto_increment not null,
    category_id int(10) unsigned not null,
    category_name varchar(128) not null,
    create_time timestamp not null default current_timestamp,
    update_time timestamp null default current_timestamp on update current_timestamp,
    UNIQUE key `idx_category_id` (category_id),
    UNIQUE key  `idx_category_name` (category_name)
)
insert into category (id, category_id, category_name) values(1, 1, "科技")