CREATE TABLE `t_order_item`
(
    `order_item_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单项ID',
    `order_number`  varchar(50)   NOT NULL COMMENT '订单号',

    `shop_id`       int           NOT NULL COMMENT '店铺id',
    `prod_id`       bigint unsigned NOT NULL COMMENT '产品ID',
    `prod_name`     varchar(120)  NOT NULL                                       DEFAULT '' COMMENT '产品名称',
    `product_nums`  int                                                          DEFAULT NULL COMMENT '商品个数',
    `pic`           varchar(255)  NOT NULL                                       DEFAULT '' COMMENT '产品主图片路径',
    `sku_id`        bigint unsigned NOT NULL COMMENT '产品SkuID',
    `sku_name`      varchar(120)                                                 DEFAULT NULL COMMENT 'sku名称',

    `prop_id`       bigint unsigned NOT NULL COMMENT '属性id',
    `prop_name`     varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '属性名称',

    `price`         int           NOT NULL COMMENT '产品单价',
    `prod_amount`   int           NOT NULL                                       DEFAULT 0 COMMENT '仅订单项商品总价',

    `user_id`       bigint unsigned NOT NULL COMMENT 'userid',
    `score`         int                                                          DEFAULT 0 COMMENT '此订单项可得积分',

    `is_commented`  tinyint(1) unsigned DEFAULT NULL COMMENT '是否评价 0 未评价 1 已评价 ',
    `is_good`       tinyint(1) unsigned DEFAULT NULL COMMENT ' 1 好评 2 差评 3 一般',
    `comment`       varchar(1000) NOT NULL                                       DEFAULT '' COMMENT '评语',

    `create_at`     datetime                                                     DEFAULT CURRENT_TIMESTAMP COMMENT '创建订单时间',
    `update_at`     datetime                                                     DEFAULT NULL COMMENT '订单最近更新时间',
    `comment_at`    datetime                                                     DEFAULT NULL COMMENT '评论时间',


    PRIMARY KEY (`order_item_id`) USING BTREE,
    KEY             `idx_order_number` (`order_number`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='订单项表';