
-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `username` varchar(50) DEFAULT '' COMMENT '账号',
                             `password` varchar(50) DEFAULT '' COMMENT '密码',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES ('1', 'test', 'test123');

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
                            `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(100) DEFAULT '' COMMENT '标签名称',
                            `created_at` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
                            `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                            `modified_at` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                            `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                            `deleted_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
                            `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
