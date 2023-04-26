CREATE DATABASE IF NOT EXISTS notifications_db;

USE notifications_db;

CREATE TABLE IF NOT EXISTS `templates`(
    `id`        INTEGER       PRIMARY KEY AUTO_INCREMENT,
    `bodyEmail` VARCHAR(2048),
    `bodySMS`   VARCHAR(2048),
    `bodyPush`  VARCHAR(2048),
	`language`  VARCHAR(3)    NOT NULL,
	`type`      VARCHAR(8)    NOT NULL
);

CREATE TABLE IF NOT EXISTS `notifications`(
    `id`                   INTEGER       PRIMARY KEY AUTO_INCREMENT,
    `appId`                VARCHAR(16)   NOT NULL,
    `templateId`           INTEGER       NOT NULL,
    `contactInfo`          VARCHAR(168),
    `title`                VARCHAR(128)  NOT NULL,
    `message`              VARCHAR(2048) NOT NULL,
    `sentTime`             INTEGER       NOT NULL DEFAULT(UNIX_TIMESTAMP())
);

CREATE TABLE IF NOT EXISTS `clients`(
    `id`           VARCHAR(16)  PRIMARY KEY,
    `secret`       VARCHAR(128) NOT NULL,
    `permissions`  INTEGER      NOT NULL,
    `creationTime` INTEGER      NOT NULL DEFAULT(UNIX_TIMESTAMP())
);
