DROP DATABASE IF EXISTS michishirube;
CREATE DATABASE michishirube;
USE michishirube;

---- drop ----
DROP TABLE IF EXISTS `locate`;
DROP TABLE IF EXISTS `emotional`;
DROP TABLE IF EXISTS `log`;

---- create ----
create table IF not exists `locate`
(
 `id`           INT(20) AUTO_INCREMENT,
 `locate`       VARCHAR(50) NOT NULL,
 `latitude`     DOUBLE NOT NULL,
 `longitude`    DOUBLE NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `emotional`
(
 `id`       INT(20) AUTO_INCREMENT,
 `glad`     INT(20) NOT NULL,
 `anger`    INT(20) NOT NULL,
 `sad`      INT(20) NOT NULL,
 `pleasant` INT(20) NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

create table IF not exists `log`
(
 `id`           INT(20) AUTO_INCREMENT,
 `count`        INT(20) NOT NULL,
 `timestamp`    VARCHAR(50) NOT NULL,
 `emotion`      INT(20) NOT NULL,
 `value`        INT(20) NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

---- insert ----
insert into `locate` values (1, '函館駅', 41.773746, 140.726399);
insert into `locate` values (2, '五稜郭駅', 41.803418, 140.733884);
insert into `locate` values (3, '桔梗駅', 41.846735, 140.723046);
insert into `locate` values (4, '大中山駅', 41.864738, 140.713482);

insert into `emotional` values (1, 50, 50, 50, 50);
insert into `emotional` values (2, 50, 50, 50, 50);
insert into `emotional` values (3, 50, 50, 50, 50);
insert into `emotional` values (4, 50, 50, 50, 50);