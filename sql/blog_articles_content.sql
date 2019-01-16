CREATE TABLE `blog_articles_content` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '文章ID',
  `content` text COMMENT '文章内容',
  `html_content` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;