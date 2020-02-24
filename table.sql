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

create table question_answer_rel (
    id bigint(20) primary key auto_increment not null,
    question_id bigint(20) not null,
    answer_id bigint(20) not null,
    create_time timestamp null default current_timestamp,
    UNIQUE key `id_question_answer` (question_id, answer_id)
)

create table answer (
    id bigint(20) primary key auto_increment not null,
    answer_id bigint(20) not null,
    content text not null,
    comment_count int(10) unsigned not null default '0',
    voteup_count int(11) not null default '0',
    author_id bigint(20) not null,
    status tinyint(3) unsigned not null default '1',
    can_comment tinyint(3) unsigned not null default '1',
    create_time timestamp default current_timestamp,
    update_time timestamp default current_timestamp on update current_timestamp,
    UNIQUE key `idx_answer_id` (answer_id),
    key `idx_author_id` (author_id)
)

create table comment (
    id bigint(20) primary key auto_increment not null,
    comment_id bigint(20) not null,
    content text not null,
    answer_id bigint(20) not null COMMENT '评论所属的答案id',
    parent_id bigint(20) not null default '0' COMMENT '父comment_id, 为0时表示直接一级评论',
    to_user_id bigint(20) not null default '0' COMMENT '评论的用户的id, 为0时表示 一级评论，没有用户id',
    from_user_id bigint(20) not null COMMENT '发评论的用户的id',
    is_del tinyint(3) not null default '0' COMMENT '是否被删除，0未被删除，1被删除',
    create_time timestamp not null default current_timestamp,
    key `idx_answer_id` (answer_id)
)

create table favorite_dir (
    id bigint primary key auto_increment not null,
    dir_id bigint not null,
    dir_name varchar(128) not null,
    user_id bigint not null,
    key `idx_dir_name` (dir_name),
    key `idx_user_id` (user_id)
)

create table favorite (
    id bigint primary key auto_increment not null,
    answer_id bigint not null,
    dir_id bigint not null,
    user_id bigint not null,
    key `idx_dir_id` (dir_id),
    key `idx_user_id` (user_id)
)