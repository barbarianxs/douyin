drop table if exists comment;
create table comment
(
    id          bigint primary key auto_increment,
    user_id     bigint        not null,
    content     varchar(1024) not null,
    video_id    bigint        not null,
    parent_id   bigint        not null default 0,
    is_valid    bool          not null default true,
    create_time datetime      not null default CURRENT_TIMESTAMP,
    index (video_id, parent_id)
) char set utf8mb4;