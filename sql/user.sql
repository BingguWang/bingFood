CREATE
DATABASE IF NOT EXISTS bingFood
	DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_unicode_ci;

CREATE TABLE `t_user`
(
    `user_id`              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'userid',
    `user_real_name`       varchar(50)          DEFAULT NULL COMMENT '真实姓名',
    `user_nick_name`       varchar(50)          DEFAULT NULL COMMENT '昵称',
    `user_status`          int         NOT NULL DEFAULT '1' COMMENT '状态 1 正常 0 无效',
    `user_mobile`          varchar(20)          DEFAULT NULL COMMENT '手机号码',
    `user_wx_number`       varchar(36) NOT NULL DEFAULT '' COMMENT '微信号',
    `create_at`            datetime             DEFAULT CURRENT_TIMESTAMP COMMENT '',
    `update_at`            datetime             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '',
    `user_sex`             tinyint(1) DEFAULT NULL COMMENT '性别1男 2 女',
    `user_birth_date`      char(10)             DEFAULT NULL COMMENT '例如：2009-11-27',
    `user_mail`            varchar(100)         DEFAULT NULL COMMENT '用户邮箱',
    `login_password`       varchar(255)         DEFAULT NULL COMMENT '登录密码',
    `user_reg_region`      varchar(36) NOT NULL DEFAULT '' COMMENT '用户注册所在地区',
    `user_reg_ip`          varchar(50)          DEFAULT NULL COMMENT '注册IP',
    `last_login_region`    varchar(50)          DEFAULT NULL COMMENT '最后登录所在地区',
    `last_login_ip`        varchar(36) NOT NULL DEFAULT '' COMMENT '最后登录IP',
    `user_last_login_time` datetime             DEFAULT NULL COMMENT '最后登录时间',
    `score`                int                  DEFAULT '0' COMMENT '用户积分',

    PRIMARY KEY (`user_id`) USING BTREE,
    UNIQUE KEY `uidx_user_mail` (`user_mail`) USING BTREE,
    UNIQUE KEY `uidx_user_unique_mobile` (`user_mobile`) USING BTREE,
    UNIQUE KEY `uidx_user_wx_number` (`user_wx_number`) USING BTREE,
    KEY                    `idx_user_region` (`user_reg_region`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户表';