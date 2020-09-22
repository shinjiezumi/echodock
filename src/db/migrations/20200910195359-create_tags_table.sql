
-- +migrate Up
CREATE TABLE IF NOT EXISTS `echodock`.`tags` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT ' タグID',
  `name` varchar(255) NOT NULL COMMENT 'タグ名',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='タグ';

-- +migrate Down
DROP TABLE IF EXISTS `echodock`.`tags`;