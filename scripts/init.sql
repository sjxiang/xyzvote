
SHOW DATABASES;

CREATE DATABASE IF NOT EXISTS `xyz_vote` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `xyz_vote`;

SHOW TABLES;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id`    bigint(20) NOT NULL DEFAULT '0',
    `username`   varchar(64) NOT NULL DEFAULT '',
    `password`   varchar(64) NOT NULL DEFAULT '',
    `email`      varchar(64) NOT NULL DEFAULT '',
    `gender`     tinyint(4) NOT NULL DEFAULT '0',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT = '用户表';


INSERT INTO `user` (`user_id`, `username`, `password`, `email`)
VALUES (10001, "sjxiang", "shgqmrf19", "1535484943@qq.com");


CREATE TABLE `vote` ( 
    `id`         bigint NOT NULL AUTO_INCREMENT, 
    `title`      varchar(255) DEFAULT NULL COMMENT '标题', 
    `type`       int DEFAULT NULL COMMENT '0单选 1多选', 
    `status`     int DEFAULT NULL COMMENT '0正常 1超时', 
    `time`       bigint DEFAULT NULL COMMENT '有效时长', 
    `user_id`    bigint DEFAULT NULL COMMENT '创建人', 
    `created_at` datetime DEFAULT NULL, 
    `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP, 
    PRIMARY KEY (`id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT = '投票活动表'; 

-- 晚饭吃什么
INSERT INTO `vote` (`title`, `type`, `status`, `time`, `user_id`, `created_at`, `updated_at`)
VALUES ("today eat food", 0, 0, 86400, 10001, NOW(), NOW());

-- 周末 city walk
INSERT INTO `vote` (`title`, `type`, `status`, `time`, `user_id`, `created_at`, `updated_at`)
VALUES ("today city walk", 0, 0, 86400, 10001, NOW(), NOW());


CREATE TABLE `vote_opt` ( 
    `id` bigint NOT NULL AUTO_INCREMENT, 
    `name` varchar(255) DEFAULT NULL COMMENT '选项名称', 
    `count` int DEFAULT NULL COMMENT '得票数', 
    `vote_id` bigint DEFAULT NULL COMMENT '投票活动', 
    `created_at` datetime DEFAULT NULL, 
    `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP, 
    PRIMARY KEY (`id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '投票选项表'; 

-- 肯德基
INSERT INTO `vote_opt` (`name`, `count`, `vote_id`, `created_at`, `updated_at`)
VALUES ("kfc", 0, 1, NOW(), NOW());

-- 麦当劳
INSERT INTO `vote_opt` (`name`, `count`, `vote_id`, `created_at`, `updated_at`)
VALUES ("m", 0, 1, NOW(), NOW());

-- 沙县小吃
INSERT INTO `vote_opt` (`name`, `count`, `vote_id`, `created_at`, `updated_at`)
VALUES ("shaxian", 0, 1, NOW(), NOW());


CREATE TABLE `vote_opt_user` ( 
    `id` bigint NOT NULL AUTO_INCREMENT, 
    `user_id` bigint DEFAULT NULL COMMENT '用户', 
    `vote_id` bigint DEFAULT NULL COMMENT '投票项目', 
    `vote_option_id` bigint DEFAULT NULL COMMENT '投票选项', 
    `created_at` datetime DEFAULT NULL COMMENT '投票时间',  
    PRIMARY KEY (`id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '用户投票详情表';


-- 表之间的关系
-- 投票和选项之间是 1:n，用户和投票之间是 n:n，必须三张表。