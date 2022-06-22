DROP TABLE if EXISTS bingFood.t_order_pay;

CREATE TABLE `t_order_pay`
(
    `pay_id`        bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '支付表ID',
    `order_number`  varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单号',
    `shop_id`       int                                                             DEFAULT NULL COMMENT '店铺id',

    `pay_no`        varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '支付单号，用于传给第三方',
    `pay_type`      tinyint(1) unsigned DEFAULT NULL COMMENT '支付方式  1 微信支付 2 支付宝 ',
    `pay_type_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci          DEFAULT NULL COMMENT '支付方式名称',
    `pay_amount`    int                                                    NOT NULL DEFAULT 0 COMMENT '支付金额',
    `user_id`       bigint unsigned NOT NULL COMMENT 'userid',
    `user_mobile`   varchar(20)                                                     DEFAULT NULL COMMENT '用户手机号码',
    `pay_status`    tinyint(1) unsigned DEFAULT 0 COMMENT '支付状态 0 未支付 1支付',

    `create_at`     datetime                                                        DEFAULT CURRENT_TIMESTAMP COMMENT '创建订单时间',
    `update_at`     datetime                                                        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单最近更新时间',


    PRIMARY KEY (`pay_id`) USING BTREE,
    KEY             `idx_user_id` (`user_id`) USING BTREE,
    KEY             `idx_user_mobile` (`user_mobile`) USING BTREE,
    KEY             `idx_shop_id` (`shop_id`) USING BTREE,
    KEY             `idx_order_number` (`order_number`) USING BTREE,
    KEY             `idx_pay_no` (`pay_no`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='订单支付表';