CREATE DATABASE IF NOT EXISTS notifications_db;

USE notifications_db;

CREATE TABLE IF NOT EXISTS `templates`(
    `id`        BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `bodyEmail` VARCHAR(2048),
    `bodySMS`   VARCHAR(2048),
    `bodyPush`  VARCHAR(2048),
	`language`  VARCHAR(3)      NOT NULL,
	`type`      VARCHAR(8)      NOT NULL
);

CREATE TABLE IF NOT EXISTS `notifications`(
    `id`                   BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `appId`                VARCHAR(16)     NOT NULL,
    `templateId`           BIGINT UNSIGNED NOT NULL,
    `contactInfo`          VARCHAR(168),
    `title`                VARCHAR(128)    NOT NULL,
    `message`              VARCHAR(2048)   NOT NULL,
    `sentTime`             TIMESTAMP       NOT NULL DEFAULT(CURRENT_TIMESTAMP())
);

CREATE TABLE IF NOT EXISTS `clients`(
    `id`           VARCHAR(16)      PRIMARY KEY,
    `secret`       VARCHAR(128)     NOT NULL,
    `permissions`  INTEGER UNSIGNED NOT NULL,
    `creationTime` TIMESTAMP        NOT NULL DEFAULT(CURRENT_TIMESTAMP())
);
