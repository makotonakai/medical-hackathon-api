---- drop ----
DROP TABLE IF EXISTS profiles;

---- create ----
create table IF not exists profiles
(
 `id`               INT(20) AUTO_INCREMENT,
 `user_id`          INT(20) NOT NULL,
 `sex`      VARCHAR(20) NOT NULL, 
 `is_pregnant` BOOLEAN NOT NULL,
 `height` INT(20) NOT NULL,
 `weight` INT(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into profiles values (1, 1, "男性", 0, 170, 65, current_timestamp, current_timestamp);
