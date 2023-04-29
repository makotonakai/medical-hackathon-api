---- drop ----
DROP TABLE IF EXISTS clinics;

---- create ----
create table IF not exists clinics
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into clinics values (1, "hoge", current_timestamp, current_timestamp);
