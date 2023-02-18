CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
  `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
  `avatar`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Avatar',
  `follow_count` bigint NOT NULL DEFAULT 0,
  `follower_count` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT COMMENT 'PK',
  `author_id` bigint NOT NULL,
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `file_path` varchar(128) NOT NULL,
  `cover_path` varchar(128) NOT NULL,
  `favorite_count` bigint DEFAULT 0,
  `comment_count` bigint DEFAULT 0,
  `title` varchar(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'video update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'video delete time'
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='video table';

CREATE TABLE `message`
(
    `id`         bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_user_id` bigint NOT NULL COMMENT 'FromUserID',
    `to_user_id` bigint NOT NULL COMMENT 'ToUserID',
    `content`    TEXT NULL COMMENT 'Content',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Message create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Message update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Message delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`from_user_id`) COMMENT 'UserId index'
    
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Message table';

ALTER TABLE `message` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `video` ADD FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);
ALTER TABLE `message` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);
