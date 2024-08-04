create table tb_user(
    id bigint unsigned primary key comment 'uid',
    phone char(11) unique not null default '' comment '电话号码',
    `password` varchar(20) not null default '' comment '登录密码',
    nickname varchar(10) not null default '' comment '昵称',
    avatar_url char(50) not null default '' comment '头像url',
    sex char(2) not null default '保密' comment '性别：保密/男/女',
    email varchar(50) not null default '' comment '邮箱地址',
    birth_at bigint unsigned not null default 0 comment '出生时间',
    created_at bigint unsigned not null default 0 comment '创建时间',
    updated_at bigint unsigned not null default 0 comment '更新时间',
    deleted_at bigint unsigned not null default 0 comment '删除时间'
) comment '用户信息表';

select * from tb_user;