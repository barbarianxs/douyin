CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
  `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
  `name`   varchar(128) NOT NULL DEFAULT 'xiaodouyin' COMMENT 'Name',
  
  `follow_count` bigint NOT NULL DEFAULT 0,
  `follower_count` bigint NOT NULL DEFAULT 0,
  `is_follow`      boolean  NOT NULL DEFAULT 1 COMMENT 'IsFollow',
  `avatar`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Avatar',
  `background_image`   varchar(128) NOT NULL DEFAULT '' COMMENT 'BackgroundImage',
  `signature`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Signature',
  `total_favorited`   varchar(128) NOT NULL DEFAULT '' COMMENT 'TotalFavorited',
  `work_count` bigint NOT NULL DEFAULT 0,
  `favorite_count` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `author_id` bigint NOT NULL,
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
<<<<<<< HEAD
<<<<<<< HEAD
  `play_url` varchar(128) NOT NULL,
  `cover_url` varchar(128) NOT NULL,
=======
  `file_path` varchar(128) NOT NULL,
  `cover_path` varchar(128) NOT NULL,
>>>>>>> origin/guo
=======
  `play_url` varchar(128) NOT NULL,
  `cover_url` varchar(128) NOT NULL,
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
  `favorite_count` bigint DEFAULT 0,
  `comment_count` bigint DEFAULT 0,
  `title` varchar(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'video create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'video update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'video delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_video_of_user_id` (`author_id`) COMMENT 'VideoOfUserId index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `message`
(
    `id`         bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_user_id` bigint NOT NULL COMMENT 'FromUserID',
    `to_user_id` bigint NOT NULL COMMENT 'ToUserID',
    `content`    TEXT NULL COMMENT 'Content',
<<<<<<< HEAD
<<<<<<< HEAD
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Message create_time',
=======
>>>>>>> origin/guo
=======
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Message create_time',
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Message create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Message update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Message delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`from_user_id`) COMMENT 'UserId index'
    
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Message table';

<<<<<<< HEAD
<<<<<<< HEAD
CREATE TABLE `relation` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `from_user_id` bigint NOT NULL,
  `to_user_id` bigint NOT NULL,
  -- `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Relation create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Relation update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Relation delete time',
  PRIMARY KEY (`id`),
  KEY          `idx_from_user_id` (`from_user_id`) COMMENT 'Relation index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Relation table';
=======
CREATE TABLE `follow` (
=======
CREATE TABLE `relation` (
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `from_user_id` bigint NOT NULL,
  `to_user_id` bigint NOT NULL,
  -- `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Relation create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Relation update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Relation delete time',
  PRIMARY KEY (`id`),
<<<<<<< HEAD
  KEY          `idx_from_user_id` (`from_user_id`) COMMENT 'Follower index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Follow table';
>>>>>>> origin/guo
=======
  KEY          `idx_from_user_id` (`from_user_id`) COMMENT 'Relation index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Relation table';
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a

CREATE TABLE `favorite` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'PK',
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'follow create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'follow update time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'follow delete time',
  PRIMARY KEY (`id`),
<<<<<<< HEAD
<<<<<<< HEAD
  KEY          `idx_from_user_id` (`user_id`) COMMENT 'Favorite index'
=======
  KEY          `idx_from_user_id` (`user_id`) COMMENT 'favorite index'
>>>>>>> origin/guo
=======
  KEY          `idx_from_user_id` (`user_id`) COMMENT 'Favorite index'
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favorite table';

ALTER TABLE `video` ADD FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);
ALTER TABLE `message` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `message` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);
<<<<<<< HEAD
<<<<<<< HEAD
ALTER TABLE `relation` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `relation` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `favorite` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `favorite` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);
=======
ALTER TABLE `follow` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `follow` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `favorite` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `favorite` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);
>>>>>>> origin/guo
=======
ALTER TABLE `relation` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `relation` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);
ALTER TABLE `favorite` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `favorite` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
