
-- +migrate Up
CREATE TABLE IF NOT EXISTS `echodock`.`comments` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT ' コメントID',
  `board_id` int unsigned NOT NULL COMMENT '投稿ID',
  `comment` text NOT NULL COMMENT 'コメント',
  `name` varchar(255) NOT NULL COMMENT '投稿者',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='コメント';

-- +migrate Down
DROP TABLE IF EXISTS `echodock`.`comments`;