
-- +migrate Up
CREATE TABLE IF NOT EXISTS `echodock`.`board_tag_relation` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT ' 投稿タグリレーションID',
  `board_id` int unsigned NOT NULL COMMENT '投稿ID',
  `tag_id` int unsigned NOT NULL COMMENT 'タグID',
  PRIMARY KEY(id),
  UNIQUE KEY `uid_board_tag_id` (`board_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='投稿タグリレーション';

-- +migrate Down
DROP TABLE IF EXISTS `echodock`.`board_tag_relation`;