CREATE TABLE `common_user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
    `user_name` varchar (100)  NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
    `password` varchar (60) NOT NULL DEFAULT '' COMMENT '密码',
    `tel` varchar (60) NOT NULL DEFAULT '' COMMENT '电话',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='普通用户表';


CREATE TABLE `fans_user` (
                             `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                             `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
                             `nickname` varchar(100) NOT NULL DEFAULT '' COMMENT '昵称',
                             `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
                             `union_id` varchar(60) NOT NULL DEFAULT '' COMMENT '微信唯一ID',
                             `xcx_openid` varchar(60) NOT NULL DEFAULT '' COMMENT '小程序openid',
                             `gzh_openid` varchar(60) NOT NULL DEFAULT '' COMMENT '某一个公众号的openid',
                             `has_sub` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '已订阅数',
                             `sub_top` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '订阅上限',
                             `expire_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
                             `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `uk_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
