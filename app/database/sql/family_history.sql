---- drop ----
DROP TABLE IF EXISTS family_histories;

---- create ----
create table IF not exists family_histories
(
 `id`               INT(20) AUTO_INCREMENT,
 `user_id`          INT(20) NOT NULL,
 `description`      TEXT NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into family_histories values (1, 1, "父が頭痛持ちです", current_timestamp, current_timestamp);
