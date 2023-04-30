---- drop ----
DROP TABLE IF EXISTS drink_histories;

---- create ----
create table IF not exists drink_histories
(
 `id`               INT(20) AUTO_INCREMENT,
 `user_id`          INT(20) NOT NULL,
 `description`      TEXT NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into drink_histories values (1, 1, "週3日飲みます", current_timestamp, current_timestamp);
