CREATE
DATABASE IF NOT EXISTS bingFood
	DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_unicode_ci;

CREATE TABLE `t_user`
(
    `user_id`         varchar(36) NOT NULL DEFAULT '' COMMENT 'ID',
    `user_real_name`  varchar(50)          DEFAULT NULL COMMENT '真实姓名',
    `user_status`     int         NOT NULL DEFAULT '1' COMMENT '状态 1 正常 0 无效',
    `user_mobile`     varchar(20)          DEFAULT NULL COMMENT '手机号码',
    `user_region`     varchar(36) NOT NULL DEFAULT '' COMMENT '用户所在地区',
    `user_wx_number`  varchar(36) NOT NULL DEFAULT '' COMMENT '微信号',
    `create_at`       datetime             DEFAULT CURRENT_TIMESTAMP COMMENT '订购时间',
    `update_at`       datetime             DEFAULT NULL COMMENT '订单更新时间',
    `user_sex`        tinyint(1) DEFAULT 0 COMMENT '性别1男 2 女',
    `user_birth_date` char(10)             DEFAULT NULL COMMENT '例如：2009-11-27',
    `user_mail`       varchar(100)         DEFAULT NULL COMMENT '用户邮箱',
    `login_password`  varchar(255)         DEFAULT NULL COMMENT '登录密码',
    `user_lasttime`   datetime             DEFAULT NULL COMMENT '最后登录时间',
    `user_regip`      varchar(50)          DEFAULT NULL COMMENT '注册IP',
    `user_lastip`     varchar(50)          DEFAULT NULL COMMENT '最后登录IP',
    `score`           int                  DEFAULT '0' COMMENT '用户积分',

    PRIMARY KEY (`user_id`) USING BTREE,
    UNIQUE KEY `uidx_user_mail` (`user_mail`) USING BTREE,
    UNIQUE KEY `uidx_user_unique_mobile` (`user_mobile`) USING BTREE,
    UNIQUE KEY `uidx_user_wx_number` (`user_wx_number`) USING BTREE,
    KEY               `idx_user_region` (`user_region`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户表';