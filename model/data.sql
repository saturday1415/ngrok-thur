-- 
create database if not exists `ngrok-thur`;

use ngrok-thur;

-- 用户表
create table if not exists `member`(
    `member_id` int unsigned auto_increment primary key,
    `member_number` varchar(32) unique not null comment '账号',
    `member_password` varchar(255) not null comment '密码',
    `member_phone` char(11) unique not null comment '手机号',
    `member_nickname` varchar(30) COLLATE utf8_unicode_ci comment '昵称',
    `member_avatar` varchar(50) comment '用户头像',
    `state` tinyint(1) default 0 comment '0-禁用 1-激活 3-锁定',
    `token` varchar(255) comment 'session_id or token',
    `created_at` INT comment '添加时间',
    `updated_at` INT comment '修改时间'
)engine=innodb default charset=utf8 comment '用户表';

-- 用户隧道权限表
create table if not exists `member_tunnel_auth`(

)engine=innodb default charset=utf8 comment '用户隧道权限表'