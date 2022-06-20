CREATE TABLE `t_prod`
(
    `prod_id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品id',
    `prod_name`        varchar(300) NOT NULL DEFAULT '' COMMENT '商品名称',
    `shop_id`          int                   DEFAULT NULL COMMENT '店铺id',
    `ori_price`        int          NOT NULL DEFAULT 0 COMMENT '原价',
    `price`            int          NOT NULL DEFAULT 0 COMMENT '现价',
    `packing_fee`      int          NOT NULL DEFAULT 0 COMMENT '打包费',
    `pic`              varchar(255)          DEFAULT NULL COMMENT '商品主图',
    `imgs`             varchar(1000)         DEFAULT NULL COMMENT '商品图片，以,分割',
    `description`      varchar(1000)         DEFAULT '' COMMENT '简要描述,卖点等',
    `content`          text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '套餐内容',
    `prod_status`      tinyint(1) unsigned DEFAULT 0 COMMENT '默认是1，表示正常状态, -1表示删除, 0下架',
    `is_auto_add`      tinyint(1) unsigned DEFAULT 0 COMMENT '默认是0，是否自动补充，即库存是否无限制 0不是 1 是',

    `sold_history_num` int                   DEFAULT 0 COMMENT '销量',
    `sold_month_num`   int                   DEFAULT 0 COMMENT '月销量',
    `total_stock`      int                   DEFAULT 0 COMMENT '总库存',
    `score`            int                   DEFAULT 0 COMMENT '可得积分',
    `buy_limit`        int                   DEFAULT '0' COMMENT '限购数量 0 不限制 大于0表示限制数量',
    `create_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`        datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近更新时间',
    `delete_at`        datetime              DEFAULT NULL COMMENT '删除时间,软删除',

    `category_id`      bigint unsigned DEFAULT NULL COMMENT '商品分类',

    PRIMARY KEY (`prod_id`) USING BTREE,
    KEY                `idx_shop_id` (`shop_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='商品表';