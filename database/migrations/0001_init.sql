-- +goose Up
CREATE TABLE IF NOT EXISTS `records` (
	`id`   	     varchar(255),
	`the_num`    INTEGER,
	`the_str`    varchar(255),
	`created_at` TIMESTAMP,
	`updated_at` TIMESTAMP
);

CREATE TABLE members (
    id 	  	   INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    name       varchar(255),
    birthday   TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY key (id)
);
-- +goose Down
DROP TABLE records;
DROP TABLE members;