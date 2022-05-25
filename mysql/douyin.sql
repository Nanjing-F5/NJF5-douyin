CREATE DATABASE IF NOT EXISTS douyin_db CHARACTER SET 'utf8';

USE douyin_db;

# 1. 用户表
CREATE TABLE `user`(
`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id(主键)',
`name` VARCHAR(50) DEFAULT NULL COMMENT '用户名称',
`password` VARCHAR(80) NOT NULL COMMENT '用户密码',
`phone` VARCHAR(20) NOT NULL COMMENT '手机号',
`follow_count` BIGINT UNSIGNED COMMENT '关注数',
`follower_count` BIGINT UNSIGNED COMMENT '粉丝数',
`is_follow` BOOL  DEFAULT NULL COMMENT '是否关注',
`create_at` DATETIME DEFAULT NULL COMMENT '创建时间',
`update_at` DATETIME DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE
);

DROP TABLE IF EXISTS `user`;

# FollowCount   int64  `json:"follow_count,omitempty"`
#	FollowerCount int64  `json:"follower_count,omitempty"`
#	IsFollow      bool   `json:"is_follow,omitempty"`

DESC `user`;

INSERT INTO `user`(id,`name`,`password`,phone)
VALUES 
(1,'kkite','abc123','12345678910'),
(2,'zhangsan','abc123','12345678911');

INSERT INTO `user`(id,`name`,`password`,phone,follow_count,follower_count)
VALUES
(3,'daisy','123456',"232212312321",100,132303);

SELECT * FROM USER;

# 2. video表
CREATE TABLE `video`(
`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "视频id（主键）",
`author_id` BIGINT UNSIGNED NOT NULL COMMENT "作者id（外键）",
`play_url` VARCHAR(100) NOT NULL COMMENT "视频url",
`cover_url` VARCHAR(100) NOT NULL COMMENT "视频封面url",
`favorite_count` BIGINT UNSIGNED DEFAULT '0' COMMENT "点赞数",
`comment_count` BIGINT UNSIGNED DEFAULT '0' COMMENT "评论数",
`is_favorite` BOOL DEFAULT '0' COMMENT "是否点赞",
`create_at` DATETIME DEFAULT NULL COMMENT '创建时间',
PRIMARY KEY (`id`) USING BTREE,
CONSTRAINT `user_video_fk` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`)
);

DESC video;

INSERT INTO `video`(id,author_id,play_url,cover_url)
VALUES
(1,2,"http://10.136.240.231:8080/static/bear.mp4","http://10.136.240.231:8080/static/bear-1283347_1280.jpg")

SELECT * FROM video;