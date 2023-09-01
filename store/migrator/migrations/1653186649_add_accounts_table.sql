-- +migrate Up

-- create languages table
create table if not exists `accounts`
(
    `id`          bigint unsigned                                                 not null primary key auto_increment,
    `created_at`  timestamp default current_timestamp                             not null,
    `updated_at`  datetime  default current_timestamp on update current_timestamp not null,
    `app_version` text                                                            not null,
    `email`       text                                                            not null,
    `is_deleted`  tinyint(1)                                                      not null default 0
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_unicode_ci;

-- +migrate Down
drop table accounts;
