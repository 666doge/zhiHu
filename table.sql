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