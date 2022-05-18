CREATE DATABASE IF NOT EXISTS douyin_db CHARACTER SET 'utf8';

USE douyin_db;

# 创建用户表
CREATE TABLE `user`(
`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id(主键)',
`username` VARCHAR(50) DEFAULT NULL COMMENT '用户名称',
`password` VARCHAR(80) NOT NULL COMMENT '用户密码',
`phone` VARCHAR(20) NOT NULL COMMENT '手机号',
`follow_count` BIGINT UNSIGNED COMMENT '关注数',
`follower_count` BIGINT UNSIGNED COMMENT '粉丝数',
`IsFollow` BIT  DEFAULT NULL COMMENT '是否关注',
`create_at` DATE DEFAULT NULL COMMENT '创建时间',
`update_at` DATE DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE
);

DROP TABLE IF EXISTS `user`;

# FollowCount   int64  `json:"follow_count,omitempty"`
#	FollowerCount int64  `json:"follower_count,omitempty"`
#	IsFollow      bool   `json:"is_follow,omitempty"`

DESC `user`;

INSERT INTO `user`(id,username,`password`,phone)
VALUES (1,'kkite','abc123','12345678910');

SELECT * FROM USER;