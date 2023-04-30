---- drop ----
DROP TABLE IF EXISTS medicines;

---- create ----
create table IF not exists medicines
(
 `id`               INT(20) AUTO_INCREMENT,
 `user_id`          INT(20) NOT NULL,
 `description`      TEXT NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into medicines values (1, 1, "ロキソニン飲んでます", current_timestamp, current_timestamp);
