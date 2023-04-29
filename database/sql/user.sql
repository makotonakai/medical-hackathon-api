---- drop ----
DROP TABLE IF EXISTS users;

---- create ----
create table IF not exists users
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

INSERT INTO users VALUES (1, "user", current_timestamp, current_timestamp);
