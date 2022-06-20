CREATE TABLE `t_sku`
(
    `sku_id`      bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '单品ID',
    `sku_name`    varchar(120) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'sku名称',
    `prod_id`     bigint unsigned NOT NULL COMMENT '商品ID',
    `price`       int NOT NULL                                            DEFAULT 0 COMMENT '价格',
    `weight`      int NOT NULL COMMENT '份量，单位克',
    `sell_status` tinyint                                                 DEFAULT 0 COMMENT '是否  售罄0 未售罄 1 售罄',
    `stock`       int NOT NULL COMMENT 'sku库存',
    `create_at`   datetime                                                DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`   datetime                                                DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近更新时间',
    `delete_at`   datetime                                                DEFAULT NULL COMMENT '删除时间,软删除',

    PRIMARY KEY (`sku_id`) USING BTREE,
    KEY           `prod_id` (`prod_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='单品SKU表';