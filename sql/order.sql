DROP TABLE if EXISTS bingFood.t_order;

CREATE TABLE `t_order`
(
    `order_id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
    `order_number`       varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci   NOT NULL COMMENT '订单号',
    `shop_id`            int                                                               DEFAULT NULL COMMENT '店铺id',
    `user_id`            bigint unsigned NOT NULL COMMENT 'userid',
    `user_mobile`        varchar(20)                                                       DEFAULT NULL COMMENT '用户手机号码',
    `receiver_mobile`    varchar(20)                                                       DEFAULT NULL COMMENT '接收人手机号码',
    `receiver_addr`      JSON                                                              DEFAULT NULL COMMENT '接收人地址信息',

    `deliver_number`     varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci   NOT NULL COMMENT '配送单号',
    `prod_name`          varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '产品名称,多个产品将会以逗号隔开',
    `prod_nums`          int                                                               DEFAULT NULL COMMENT '订单商品总数',

    `order_status`       tinyint(1) unsigned DEFAULT NULL COMMENT '订单状态 0未支付 1已支付 2商家已接单 3骑手已接单 4已取消 5已完成 ',
    `delete_status`      tinyint(1) unsigned DEFAULT NULL COMMENT '订单删除状态',
    `pay_status`         tinyint(1) unsigned DEFAULT NULL COMMENT '支付状态，1：已经支付过，0：没有支付过',
    `refund_status`      tinyint(1) unsigned DEFAULT NULL COMMENT '0:默认  没退款,1:退款中,2:退款成功 3:退款失败 ',
    `deliver_status`     tinyint(1) unsigned DEFAULT NULL COMMENT ' ',

    `packing_amount`     int                                                      NOT NULL DEFAULT 0 COMMENT '打包费',
    `deliver_amount`     int                                                      NOT NULL DEFAULT 0 COMMENT '配送费',
    `prod_amount`        int                                                      NOT NULL DEFAULT 0 COMMENT '仅商品总价',
    `discount_amount`    int                                                      NOT NULL DEFAULT 0 COMMENT '优惠金额',
    `final_amount`       int                                                      NOT NULL DEFAULT 0 COMMENT '用户最终支付金额',
    `score`              int                                                               DEFAULT 0 COMMENT '本单可得积分',

    `pay_type`           tinyint(1) unsigned DEFAULT NULL COMMENT '支付方式',
    `deliver_type`       tinyint(1) unsigned DEFAULT NULL COMMENT '配送方式',

    `remark`             varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注',

    `create_at`          datetime                                                          DEFAULT CURRENT_TIMESTAMP COMMENT '创建订单时间',
    `update_at`          datetime                                                          DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '订单最近更新时间',
    `pay_at`             datetime                                                          DEFAULT NULL COMMENT '订单支付时间',
    `finish_at`          datetime                                                          DEFAULT NULL COMMENT '订单完成时间，送达则完成',
    `cancel_at`          datetime                                                          DEFAULT NULL COMMENT '订单取消时间',
    `delete_at`          datetime                                                          DEFAULT NULL COMMENT '删除时间,软删除',
    `cancel_apply_at`    datetime                                                          DEFAULT NULL COMMENT '订单取消申请时间',
    `cancel_reason_type` tinyint(1) unsigned DEFAULT NULL COMMENT '订单取消原因类型',


    PRIMARY KEY (`order_id`) USING BTREE,
    UNIQUE KEY `uidx_order_number` (`order_number`) USING BTREE,
    UNIQUE KEY `uidx_user_mobile_order_number` (`user_id`,`order_number`) USING BTREE,
    KEY                  `idx_shop_id` (`shop_id`) USING BTREE,
    KEY                  `idx_create_at` (`create_at`) USING BTREE,
    KEY                  `idx_finish_at` (`finish_at`) USING BTREE,

    KEY                  `idx_user_id` (`user_id`) USING BTREE,
    KEY                  `idx_user_mobile` (`user_mobile`) USING BTREE,
    KEY                  `idx_receiver_mobile` (`user_mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='订单表';