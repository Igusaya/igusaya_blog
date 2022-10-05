CREATE TABLE `d_article` (
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '記事識別子',
    `subject`  VARCHAR(100)    NOT NULL                COMMENT '表題',
    `body`     TEXT            NOT NULL                COMMENT '本文',
    `modified` DATETIME(6)     NOT NULL                COMMENT '更新日時',
    `created`  DATETIME(6)     NOT NULL                COMMENT '作成日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='記事   ';