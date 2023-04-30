---- drop ----
DROP TABLE IF EXISTS clinics;

---- create ----
create table IF not exists clinics
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(100) NOT NULL,
 `postal_code`      VARCHAR(20) NOT NULL,
 `address`          TEXT NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into clinics values (1, "足立慶友整形外科", "120-0015", "東京都足立区足立1-12-12", current_timestamp, current_timestamp);
