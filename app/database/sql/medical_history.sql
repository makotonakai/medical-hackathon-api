---- drop ----
DROP TABLE IF EXISTS medical_histories;

---- create ----
create table IF not exists medical_histories
(
 `id`               INT(20) AUTO_INCREMENT,
 `user_id`          INT(20) NOT NULL,
 `disease`          VARCHAR(20) NOT NULL,
 `description`      TEXT NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

insert into medical_histories values (1, 1, "その他", "6歳まで小児喘息でした", current_timestamp, current_timestamp);
