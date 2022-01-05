
-- +migrate Up
CREATE TABLE IF NOT EXISTS `echodock`.`boards` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '投稿ID',
  `title` varchar (255) NOT NULL COMMENT 'タイトル',
  `body` text NOT NULL COMMENT '本文',
  `name` varchar(255) NOT NULL COMMENT '投稿者',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='投稿';

-- +migrate Down
DROP TABLE IF EXISTS `echodock`.`boards`;