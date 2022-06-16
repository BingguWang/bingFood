CREATE TABLE `t_shop`
(
    `shop_id`               bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '店铺id',

    `shop_name`             varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL COMMENT '店铺名称',
    `intro`                 varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '店铺简介',
    `shop_logo`             varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '店铺logo(可修改)',
    `mobile_background_pic` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '店铺移动端背景图',

    `shop_status`           tinyint(1) unsigned DEFAULT 1 COMMENT '店铺状态(-1:已删除 0: 停业中 1:营业中,因为我这没有写后台系统所以就是直接默认是营业的)',

    `business_license`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '营业执照',
    `identity_card_front`   varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '身份证正面',
    `identity_card_later`   varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '身份证反面',

    `type`                  tinyint(1) unsigned DEFAULT NULL COMMENT '店铺类型',
    `create_at`             datetime                                                      DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`             datetime                                                      DEFAULT NULL COMMENT '最近更新时间',
    PRIMARY KEY (`shop_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='店铺表';