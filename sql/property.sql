DROP TABLE if EXISTS bingFood.t_prod_prop;

CREATE TABLE `t_prod_prop`
(
    `prop_id`   bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '属性id',
    `prop_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '属性名称',
    `shop_id`   bigint unsigned DEFAULT NULL COMMENT '店铺id',
    `prod_id`   bigint unsigned DEFAULT NULL COMMENT '商品id',

    `create_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近更新时间',
    `delete_at`        datetime              DEFAULT NULL COMMENT '删除时间,软删除',
    PRIMARY KEY (`prop_id`) USING BTREE,
    KEY         `idx_shop_id_prod_id` (`shop_id`,`prod_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='商品属性, 口味, 和份量价格无关';