CREATE TABLE `blog_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(60) NOT NULL DEFAULT '' COMMENT '文章名称',
  `describe` varchar(255) NOT NULL DEFAULT '一句话描述',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别id群',
  `author_id` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '作者ID',
  `remove` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0/否1/是2/待发布',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;