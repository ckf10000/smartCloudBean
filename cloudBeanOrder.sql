/*
 Navicat Premium Data Transfer

 Source Server         : 公司办公电脑
 Source Server Type    : MySQL
 Source Server Version : 50744
 Source Host           : 192.168.3.232:3306
 Source Schema         : cloudBeanOrder

 Target Server Type    : MySQL
 Target Server Version : 50744
 File Encoding         : 65001

 Date: 28/06/2024 12:42:33
*/

CREATE DATABASE IF NOT EXISTS cloudBeanOrder;
USE cloudBeanOrder;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_all_orders
-- ----------------------------
DROP TABLE IF EXISTS `t_all_orders`;
CREATE TABLE `t_all_orders`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `p_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '平台单号',
  `pre_order_id` int(11) NOT NULL COMMENT '预售单号',
  `pre_total_price` decimal(10, 2) DEFAULT NULL COMMENT '预售单总价',
  `pre_order_status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '待支付' COMMENT '预售单状态',
  `trip_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '行程类型，单程|多程|往返',
  `source_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '来源类型，国内|国际',
  `source_ota` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '来源供应商名，携程|就旅行|同程',
  `source_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '来源供应商编号，SKW|SZW|AZI',
  `policy_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '政策ID',
  `policy_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '政策编号',
  `policy_name` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '政策名',
  `policy_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '政策类型',
  `rule_constraints` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '规则约束',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  `deleted_time` bigint(13) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_pre_order_id_and_deleted_time`(`pre_order_id`, `deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 76 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_order_book_passenger
-- ----------------------------
DROP TABLE IF EXISTS `t_order_book_passenger`;
CREATE TABLE `t_order_book_passenger`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pre_order_id` int(11) NOT NULL COMMENT '预售单号',
  `ch_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '出单平台订单id',
  `card_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '身份证' COMMENT '证件类型',
  `card_id` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '证件id',
  `internal_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '我公司退票，改签负责人电话',
  `itinerary_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '票号',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  `deleted_time` bigint(13) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_pre_order_id_and_card_id_and_deleted_time`(`pre_order_id`, `card_id`, `deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_order_book_payment
-- ----------------------------
DROP TABLE IF EXISTS `t_order_book_payment`;
CREATE TABLE `t_order_book_payment`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pre_order_id` int(11) NOT NULL COMMENT '预售单号',
  `ch_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '出单渠道商订单id',
  `payment_amount` decimal(10, 2) DEFAULT NULL COMMENT '出单时的实际支付金额',
  `payment_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '支付方式',
  `payment_account` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '支付账号，当支付方式是钱包，余额时，为登录账号，当为银行卡时，为银行卡卡号',
  `payment_time` datetime(0) DEFAULT NULL COMMENT '支付时间',
  `payment_status` int(2) NOT NULL COMMENT '支付状态，0 支付成功 -1待支付 1 支付待确认 2 支付失败',
  `bill_no` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '支付流水',
  `payment_context` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '支付成功上下文',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  `deleted_time` bigint(13) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_pre_order_id_and_ch_order_id_and_deleted_time`(`pre_order_id`, `ch_order_id`, `deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_order_book_record
-- ----------------------------
DROP TABLE IF EXISTS `t_order_book_record`;
CREATE TABLE `t_order_book_record`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pre_order_id` int(11) NOT NULL COMMENT '预售单号',
  `ch_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '出单渠道商订单id',
  `ticket_channel` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '出单渠道',
  `channel_user` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '渠道账号',
  `out_ticket_amount` decimal(10, 2) DEFAULT NULL COMMENT '出单金额',
  `book_context` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '预订成功上下文',
  `book_combine` int(1) DEFAULT 1 COMMENT '订单是分批预订，还是组合预订，1：组合，0：分批',
  `ch_order_status` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '渠道商订单状态',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  `deleted_time` bigint(13) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_pre_order_id_and_ch_order_id_and_deleted_time`(`pre_order_id`, `ch_order_id`, `deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_order_book_report
-- ----------------------------
DROP TABLE IF EXISTS `t_order_book_report`;
CREATE TABLE `t_order_book_report`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pre_order_id` int(11) NOT NULL COMMENT '预售单号',
  `ch_order_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '出单渠道商订单id',
  `out_pf` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应采购回填信息的出票平台',
  `out_ticket_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应采购回填信息的出票账号',
  `pay_account_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应采购回填信息的采购账号类型',
  `pay_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应采购回填信息的采购账号',
  `payment_time` datetime(0) DEFAULT NULL COMMENT '支付时间',
  `payment_amount` decimal(10, 2) DEFAULT NULL COMMENT '出单时的实际支付金额',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  `deleted_time` bigint(13) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_pre_order_id_and_ch_order_id_and_deleted_time`(`pre_order_id`, `ch_order_id`, `deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_order_failed_records
-- ----------------------------
DROP TABLE IF EXISTS `t_order_failed_records`;
CREATE TABLE `t_order_failed_records`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pre_order_id` int(11) NOT NULL COMMENT '劲旅平台订单号',
  `platform` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '平台',
  `platform_order_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '平台订单号',
  `describe` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '失败描述',
  `endpoint` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '失败端点',
  `code` int(6) NOT NULL COMMENT '失败编号',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15970 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_order_flight
-- ----------------------------
DROP TABLE IF EXISTS `t_order_flight`;
CREATE TABLE `t_order_flight`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pre_order_id` int(11) NOT NULL COMMENT '预售单号',
  `flight_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '航班编号',
  `cabin` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '舱位类型',
  `child_cabin` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '子舱位类型',
  `air_co_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '航司编号',
  `air_co_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '航司名称',
  `segment` int(1) DEFAULT 1 COMMENT '航段',
  `departure_airport_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '起飞机场编号',
  `arrive_airport_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '抵达机场编号',
  `departure_city_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '起飞城市编号',
  `arrive_city_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '抵达城市编号',
  `departure_city_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '起飞城市名',
  `arrive_city_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '抵达城市名',
  `departure_time` datetime(0) DEFAULT NULL COMMENT '起飞时间',
  `arrive_time` datetime(0) DEFAULT NULL COMMENT '抵达时间',
  `transfer` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '是否中转',
  `price` decimal(10, 2) DEFAULT NULL COMMENT '预售单票面价格',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  `deleted_time` bigint(13) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_pre_order_id_and_flight_no_and_deleted_time`(`pre_order_id`, `flight_no`, `deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 76 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_order_passenger
-- ----------------------------
DROP TABLE IF EXISTS `t_order_passenger`;
CREATE TABLE `t_order_passenger`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pre_order_id` int(11) NOT NULL COMMENT '预售单号',
  `passenger` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '乘客姓名',
  `gender` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '性别',
  `age_stage` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '乘客年龄阶段(成人，儿童，婴儿)',
  `card_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '身份证' COMMENT '证件类型',
  `card_id` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '证件id',
  `passenger_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '乘客电话',
  `birth_day` datetime(0) DEFAULT NULL COMMENT '乘客生日',
  `ticket_price` decimal(10, 2) DEFAULT NULL COMMENT '票面价',
  `sale_price` decimal(10, 2) DEFAULT NULL COMMENT '销售价',
  `airport_fee` decimal(10, 2) DEFAULT NULL COMMENT '机建费',
  `fuel_fee` decimal(10, 2) DEFAULT NULL COMMENT '燃油费',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  `deleted_time` bigint(13) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_pre_order_id_and_card_id_and_deleted_time`(`pre_order_id`, `card_id`, `deleted_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 96 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_orders
-- ----------------------------
DROP TABLE IF EXISTS `t_orders`;
CREATE TABLE `t_orders`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `pre_order_id` int(11) NOT NULL COMMENT '劲旅平台订单号',
  `departure_city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '起飞城市编号',
  `arrive_city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '抵达城市编号',
  `departure_time` datetime(0) DEFAULT NULL COMMENT '起飞时间',
  `out_total_price` decimal(10, 2) DEFAULT NULL COMMENT '劲旅平台出单总价',
  `pre_sale_amount` decimal(10, 2) DEFAULT NULL COMMENT '劲旅平台出单价格',
  `flight` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '航班编号',
  `passenger` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '乘客姓名',
  `age_stage` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '乘客年龄阶段(成人，儿童)',
  `card_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '身份证' COMMENT '证件类型',
  `card_id` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '证件id',
  `internal_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '我公司退票，改签负责人电话',
  `passenger_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '乘客电话',
  `ctrip_order_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '携程订单id',
  `payment_amount` decimal(10, 2) DEFAULT NULL COMMENT '携程订单金额',
  `payment_method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '支付方式',
  `itinerary_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '携程行程单号',
  `departure_city_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '起飞城市名称',
  `arrive_city_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '抵达城市名称',
  `arrive_time` datetime(0) DEFAULT NULL COMMENT '抵达时间',
  `ctrip_username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '携程用户名',
  `user_pass` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '携程用户密码',
  `out_pf` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应劲旅系统出票采购回填信息栏的出票平台',
  `out_ticket_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应劲旅系统出票采购回填信息栏的出票账号',
  `pay_account_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应劲旅系统出票采购回填信息栏的采购账号类型',
  `pay_account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '对应劲旅系统出票采购回填信息栏的采购账号',
  `oper` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '操作人，可以是自定义的机器人名',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `order_status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '待支付' COMMENT '订单状态',
  `payment_time` datetime(0) DEFAULT NULL COMMENT '订单生成时间',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_ctrip_order_id_and_card_id`(`ctrip_order_id`, `card_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32144 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_sms
-- ----------------------------
DROP TABLE IF EXISTS `t_sms`;
CREATE TABLE `t_sms`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `phone_num` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '手机号',
  `context` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '短信原文',
  `digital_verification_code` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '短信中的纯数字验证码',
  `is_deleted` int(1) DEFAULT 0 COMMENT '是否已删除',
  `create_time` datetime(0) DEFAULT CURRENT_TIMESTAMP COMMENT '数据插入时的时间',
  `update_time` datetime(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '数据更新的时候',
  `deleted_time` bigint(13) DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 153 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
