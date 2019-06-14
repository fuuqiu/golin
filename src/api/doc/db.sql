CREATE TABLE `t_room` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '群标题',
  `level` int(11) DEFAULT NULL COMMENT '群组等级',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
)COMMENT='群';