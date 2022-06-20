CREATE TABLE `t_basket`
(
    `basket_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `shop_id`   int DEFAULT NULL COMMENT '店铺id',
    `prod_id`   bigint unsigned NOT NULL COMMENT '商品ID',
    `sku_id`    bigint unsigned NOT NULL COMMENT 'skuID',
    `user_id`   bigint unsigned NOT NULL COMMENT 'userid',
    `prod_nums` int DEFAULT 0 COMMENT '商品数',

    `create_at`   datetime                                                DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`   datetime                                                DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近更新时间',

    PRIMARY KEY (`basket_id`) USING BTREE,
    UNIQUE KEY `uidx_user_shop_sku` (`sku_id`,`user_id`,`shop_id`) USING BTREE,
    KEY         `shop_id` (`shop_id`) USING BTREE,
    KEY         `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='购物车';