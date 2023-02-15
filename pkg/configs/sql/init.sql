CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `message`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_user_id` bigint NOT NULL, COMMENT 'FromUserID',
    `to_user_id` bigint NOT NULL, 0 COMMENT 'ToUserID',
    `content`    TEXT NULL COMMENT 'Content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Message create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Message update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Message delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`from_user_id`) COMMENT 'UserId index'
    FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Message table';