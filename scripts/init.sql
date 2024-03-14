
SHOW DATABASES;

CREATE DATABASE IF NOT EXISTS `xyz_vote`;

USE `xyz_vote`;

SHOW TABLES;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id`    varchar(64) NOT NULL DEFAULT '',
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
VALUES ("ea7ee4ac-4b0b-4f78-846e-26e7ea70411d", "sjxiang1997", "123456789qwe", "1535484943@qq.com");

INSERT INTO `user` (`user_id`, `username`, `password`, `email`)
VALUES ("f2d274f5-bb8b-4175-a60b-a2d113df4818", "sjxiang2024", "123456789qwe", "cisco@qq.com");


CREATE TABLE `form` ( 
    `id`         bigint NOT NULL AUTO_INCREMENT, 
    `title`      varchar(255) DEFAULT NULL COMMENT '标题', 
    `type`       int DEFAULT NULL COMMENT '0单选 1多选', 
    `status`     int DEFAULT NULL COMMENT '0正常 1超时', 
    `duration`   bigint DEFAULT NULL COMMENT '有效时长、持续时间', 
    `user_id`    varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
    `created_at` datetime DEFAULT NULL, 
    `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP, 
    PRIMARY KEY (`id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT = '表单设置'; 

-- 晚饭吃什么
INSERT INTO `form` (`title`, `type`, `status`, `duration`, `user_id`, `created_at`, `updated_at`)
VALUES ("today eat food", 0, 0, 86400, "ea7ee4ac-4b0b-4f78-846e-26e7ea70411d", NOW(), NOW());


CREATE TABLE `form_opt` ( 
    `id` bigint NOT NULL AUTO_INCREMENT, 
    `name` varchar(255) DEFAULT NULL COMMENT '选项名称', 
    `vote_count` int DEFAULT NULL COMMENT '选项的投票数量', 
    `form_id` bigint DEFAULT NULL, 
    `created_at` datetime DEFAULT NULL, 
    `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP, 
    PRIMARY KEY (`id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '问卷选项'; 

-- 肯德基
INSERT INTO `form_opt` (`name`, `vote_count`, `form_id`, `created_at`, `updated_at`)
VALUES ("kfc", 0, 1, NOW(), NOW());

INSERT INTO `form_opt` (`name`, `vote_count`, `form_id`, `created_at`, `updated_at`)
VALUES ("shaxian", 0, 1, NOW(), NOW());


CREATE TABLE `form_opt_user` ( 
    `id` bigint NOT NULL AUTO_INCREMENT, 
    `user_id`   varchar(64) NOT NULL DEFAULT '' COMMENT '投票用户',
    `form_id`   bigint DEFAULT NULL COMMENT '投票项目', 
    `option_id` bigint DEFAULT NULL COMMENT '投票选项', 
    `created_at` datetime DEFAULT NULL COMMENT '投票时间',  
    PRIMARY KEY (`id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '用户投票详情表';


-- 表之间的关系
-- 投票和选项之间是 1:n，用户和投票之间是 n:n，必须三张表。