/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : localhost:3306
 Source Schema         : mozi

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 22/09/2018 12:03:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `password` varchar(128) NOT NULL DEFAULT '',
  `google_secret` varchar(36) NOT NULL DEFAULT '',
  `google_secret_status` tinyint(1) NOT NULL DEFAULT '0',
  `role` varchar(32) NOT NULL DEFAULT '0' COMMENT '组',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` varchar(14) NOT NULL DEFAULT '' COMMENT '创建时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of admin
-- ----------------------------
BEGIN;
INSERT INTO `admin` VALUES (1, 'boss', '$2a$10$1HdEWW.7Ob6aCZUkuhKl7O16gs0jYwe9/jsN4NYsQKZok8eQ/uxFS', '', 0, '1', 1, '');
COMMIT;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sort` int(11) DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  `open` int(1) DEFAULT NULL,
  `text` varchar(64) DEFAULT '',
  `icon` varchar(128) DEFAULT '',
  `url` varchar(128) DEFAULT '',
  `target_type` varchar(64) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=903 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` VALUES (1, 1, 0, 0, '活动管理', 'fa fa-gift', '', '');
INSERT INTO `admin_menu` VALUES (2, 2, 0, 0, '公告管理', 'fa fa-envelope', '', '');
INSERT INTO `admin_menu` VALUES (3, 3, 0, 0, '账户管理', 'fa fa-users', '', '');
INSERT INTO `admin_menu` VALUES (4, 4, 0, 0, '游戏管理', 'fa fa-gamepad', '', '');
INSERT INTO `admin_menu` VALUES (5, 5, 0, 0, '现金系统', 'fa fa-fw fa-money', '', '');
INSERT INTO `admin_menu` VALUES (6, 6, 0, 0, '退佣管理', 'fa fa-fw fa-calculator', '', '');
INSERT INTO `admin_menu` VALUES (7, 7, 0, 0, '报表管理', 'fa fa-fw fa-area-chart', '', '');
INSERT INTO `admin_menu` VALUES (8, 8, 0, 0, '系统管理', 'fa fa-fw fa-cogs', '', '');
INSERT INTO `admin_menu` VALUES (9, 9, 0, 0, '权限管理', 'fa fa-fw fa-cogs', '', '');
INSERT INTO `admin_menu` VALUES (110, 1, 1, 0, '活动记录', 'fa fa-list-alt', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (111, 2, 1, 0, '红包活动', 'fa fa-toggle-on', ' ', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (112, 3, 1, 0, '签到活动', 'fa fa-toggle-on', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (113, 4, 1, 0, '首存礼包', 'fa fa-toggle-on', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (114, 5, 1, 0, '每日存款', 'fa fa-toggle-on', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (130, 1, 2, 0, '通用公告', 'fa fa-edit', 'pages/notice/normal_notice.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (131, 2, 2, 0, '滚动公告', 'fa fa-edit', 'pages/notice/roll_notice.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (160, 1, 3, 0, '账户列表', 'fa fa-user', 'page/account/list', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (161, 2, 3, 0, '登入记录', 'fa fa-list-alt', 'pages/account/record_login.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (162, 3, 3, 0, '账号银行卡', 'fa fa-list-alt', 'pages/account/member_banks.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (163, 4, 3, 0, 'IP黑名单', 'fa fa-list-alt', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (164, 5, 3, 0, '管理列表', 'fa fa-fw fa-user-secret', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (180, 1, 4, 0, '游戏结果', 'fa fa-plus-square-o', '', '');
INSERT INTO `admin_menu` VALUES (181, 2, 4, 0, '玩家订单', 'fa fa-plus-square-o', '', '');
INSERT INTO `admin_menu` VALUES (183, 3, 4, 0, '彩票管理', 'fa fa-plus-square-o', '', '');
INSERT INTO `admin_menu` VALUES (200, 1, 180, 0, '彩票开奖', 'fa fa-life-bouy', 'pages/lotto/result.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (221, 1, 181, 0, '彩票订单', 'fa fa-life-bouy', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (222, 2, 181, 0, '电子订单', 'fa fa-life-bouy', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (223, 3, 181, 0, '体育订单', 'fa fa-life-bouy', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (224, 4, 181, 0, 'AG订单', 'fa fa-life-bouy', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (230, 1, 183, 0, '赔率房间', 'fa fa-list-alt', 'game/oddsRoom', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (231, 2, 183, 0, '赔率设置', 'fa fa-toggle-on', 'game/lottoOdds', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (232, 3, 183, 0, '实时操盘', 'fa fa-hand-pointer-o', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (233, 4, 183, 0, '实时统计', 'fa fa-bar-chart', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (234, 5, 183, 0, '彩种设置', 'fa fa-hand-pointer-o', 'pages/lotto/code_lotto.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (250, 0, 5, 0, '银行卡管理', 'fa fa-fw fa-bank', 'pages/bank/code_bank.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (251, 1, 5, 0, '收款账户', '', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (252, 2, 5, 0, '第三方收款账户', '', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (253, 3, 5, 0, '会员入款单', '', '', '');
INSERT INTO `admin_menu` VALUES (254, 4, 5, 0, '会员收款单', '', '', '');
INSERT INTO `admin_menu` VALUES (255, 5, 5, 0, '手动入款', 'fa fa-fw fa-hand-o-right', 'pages/finance/manual_deposit.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (256, 6, 5, 0, '手动出款', 'fa fa-fw fa-hand-o-right', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (257, 7, 5, 0, '手动转点', '', '', '');
INSERT INTO `admin_menu` VALUES (258, 8, 5, 0, '流水查询', 'fa fa-fw fa-file-text', 'pages/finance/record_amount_change.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (259, 9, 5, 0, '稽核分查询', 'fa fa-fw fa-file-text', 'pages/finance/record_audit_score.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (260, 10, 5, 0, '财务汇总', '', '', '');
INSERT INTO `admin_menu` VALUES (261, 0, 6, 0, '分红管理', '', '', '');
INSERT INTO `admin_menu` VALUES (262, 1, 6, 0, '工资管理', '', '', '');
INSERT INTO `admin_menu` VALUES (263, 2, 6, 0, '退佣管理', '', '', '');
INSERT INTO `admin_menu` VALUES (270, 1, 7, 0, '游戏报表', 'fa fa-bar-chart', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (271, 2, 7, 0, '个人报表', 'fa fa-bar-chart', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (272, 3, 7, 0, '平台报表', 'fa fa-bar-chart', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (273, 4, 7, 0, '运营趋势', 'fa fa-bar-chart', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (300, 1, 8, 0, '系统设置', 'fa fa-bar-chart', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (301, 2, 8, 0, '异动记录', 'fa fa-bar-chart', 'pages/sys/record_admin_operate.html', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (302, 3, 8, 0, '操作记录', 'fa fa-bar-chart', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (303, 4, 8, 0, '群组管理', 'fa fa-bar-chart', '', 'iframe-tab');
INSERT INTO `admin_menu` VALUES (900, 1, 9, 0, '管理员', '', '', '');
INSERT INTO `admin_menu` VALUES (901, 2, 9, 0, '群组', '', '', '');
INSERT INTO `admin_menu` VALUES (902, 3, 9, 0, '群组权限', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `menu` text NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `remark` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
BEGIN;
INSERT INTO `admin_role` VALUES (1, '管理员', '900,901,902,9,234,1,2,3,4,5,6,7,8,110,111,112,113,114,130,131,160,161,162,163,180,181,182,183,200,201,202,221,222,223,224,230,231,232,233,250,251,252,253,254,255,256,257,258,259,260,261,262,263,270,271,272,273,300,301,302,303', 1, '');
COMMIT;

-- ----------------------------
-- Table structure for code_lotto
-- ----------------------------
DROP TABLE IF EXISTS `code_lotto`;
CREATE TABLE `code_lotto` (
  `lotto_id` int(11) NOT NULL,
  `name` varchar(36) NOT NULL DEFAULT '' COMMENT '彩票名',
  `lotto_type` int(11) NOT NULL DEFAULT '0' COMMENT '彩票类型',
  `game_kind` int(11) NOT NULL DEFAULT '0' COMMENT '游戏类型',
  `game_type` int(11) NOT NULL DEFAULT '0' COMMENT '游戏种类',
  `sort_index` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0:停止,1:正常,2:停售',
  `introduction` varchar(64) NOT NULL DEFAULT '' COMMENT '简介',
  PRIMARY KEY (`lotto_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of code_lotto
-- ----------------------------
BEGIN;
INSERT INTO `code_lotto` VALUES (1, '重庆时时彩', 1, 1, 0, 0, 1, '每天120期');
INSERT INTO `code_lotto` VALUES (2, '新疆时时彩', 1, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (3, '天津时时彩', 1, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (4, '分分彩', 1, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (5, '两分彩', 1, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (6, '三分彩', 1, 1, 0, 0, 2, '');
INSERT INTO `code_lotto` VALUES (7, '五分彩', 1, 1, 0, 0, 0, '');
INSERT INTO `code_lotto` VALUES (8, '山东11选5', 2, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (9, '江西11选5', 2, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (10, '广东11选5', 2, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (11, '江苏11选5', 2, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (12, '安徽11选5', 2, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (13, '山西11选5', 2, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (14, '上海11选5', 2, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (15, '分分11选5', 2, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (16, '江苏快3', 3, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (17, '安徽快3', 3, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (18, '湖北快3', 3, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (19, '河南快3', 3, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (20, '江苏骰宝', 3, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (21, '分分快3', 3, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (22, '北京PK10', 4, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (23, '幸运飞艇', 4, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (24, '分分PK10', 4, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (25, '福彩3D', 5, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (26, '排列3', 5, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (27, '排列5', 6, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (28, '广东快乐十分', 7, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (29, '重庆快乐十分', 7, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (30, '天津快乐十分', 7, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (31, '北京快乐8', 8, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (32, '加拿大基诺', 8, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (33, '分分快乐彩', 8, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (34, '台湾宾果', 8, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (35, '北京幸运28', 9, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (36, '加拿大幸运28', 9, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (37, '台湾幸运28', 9, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (38, '香港六合彩', 10, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (39, '五分六合彩', 10, 1, 0, 0, 1, '');
INSERT INTO `code_lotto` VALUES (40, '十分六合彩', 10, 1, 0, 0, 1, '');
COMMIT;

-- ----------------------------
-- Table structure for code_method
-- ----------------------------
DROP TABLE IF EXISTS `code_method`;
CREATE TABLE `code_method` (
  `lotto_type` int(11) NOT NULL,
  `method_code` varchar(32) NOT NULL,
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '名称',
  `odds_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1: 独立赔率, 2:多号码单赔率,3:多号码多赔率',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`lotto_type`,`method_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of code_method
-- ----------------------------
BEGIN;
INSERT INTO `code_method` VALUES (1, '10001', '第一球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10002', '第二球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10003', '第三球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10004', '第四球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10005', '第五球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10006', '第一球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10007', '第二球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10008', '第三球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10009', '第四球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10010', '第五球', 1, 1);
INSERT INTO `code_method` VALUES (1, '10011', '总和', 1, 1);
INSERT INTO `code_method` VALUES (1, '10012', '前三', 1, 1);
INSERT INTO `code_method` VALUES (1, '10013', '中三', 1, 1);
INSERT INTO `code_method` VALUES (1, '10014', '后三', 1, 1);
INSERT INTO `code_method` VALUES (1, '11001', '第一球', 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for issue_factory
-- ----------------------------
DROP TABLE IF EXISTS `issue_factory`;
CREATE TABLE `issue_factory` (
  `lotto_id` int(11) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `count` int(11) NOT NULL DEFAULT '0',
  `issue_interval` int(11) NOT NULL DEFAULT '0',
  `iss_bit` tinyint(1) NOT NULL DEFAULT '0',
  `block_sec` int(11) NOT NULL DEFAULT '0',
  `start_time` varchar(6) NOT NULL DEFAULT '',
  `end_time` varchar(6) NOT NULL DEFAULT '',
  `issue_type` tinyint(1) NOT NULL DEFAULT '1',
  `offset` int(3) NOT NULL DEFAULT '0',
  `extra_info` text NOT NULL,
  PRIMARY KEY (`lotto_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of issue_factory
-- ----------------------------
BEGIN;
INSERT INTO `issue_factory` VALUES (1, 1, 120, 600, 3, 0, '000000', '235959', 6, 0, '000000,015500,300|095000,220000,600|220000,235959,300');
INSERT INTO `issue_factory` VALUES (2, 1, 96, 600, 2, 90, '100000', '020000', 2, 0, ' ');
INSERT INTO `issue_factory` VALUES (3, 1, 84, 600, 3, 90, '090000', '225800', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (4, 1, 1440, 60, 4, 5, '000000', '235959', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (5, 1, 720, 120, 3, 5, '000000', '235959', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (6, 1, 480, 180, 3, 5, '000000', '235959', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (7, 1, 288, 300, 3, 5, '000000', '000000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (8, 1, 87, 600, 2, 90, '082500', '225500', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (9, 1, 84, 600, 2, 90, '090000', '230000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (10, 1, 84, 600, 2, 90, '090000', '230000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (11, 1, 82, 600, 2, 90, '082550', '220550', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (12, 1, 81, 600, 2, 90, '080000', '220000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (13, 2, 94, 600, 2, 90, '082000', '235500', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (14, 1, 90, 600, 2, 90, '085000', '235000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (15, 1, 1440, 60, 4, 5, '000000', '235959', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (16, 1, 82, 600, 3, 90, '083000', '221000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (17, 1, 80, 600, 3, 90, '084000', '220000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (18, 1, 78, 600, 3, 90, '090000', '220000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (19, 1, 84, 600, 3, 90, '083500', '223500', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (20, 1, 82, 600, 3, 90, '082000', '221000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (21, 1, 1440, 60, 4, 5, '000000', '235959', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (22, 1, 179, 300, 0, 45, '090230', '235730', 4, 0, '673722|20180329235700');
INSERT INTO `issue_factory` VALUES (23, 1, 180, 500, 3, 45, '120400', '040400', 2, 0, ' ');
INSERT INTO `issue_factory` VALUES (24, 1, 1440, 60, 4, 5, '000000', '235959', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (25, 1, 1, 85500, 3, 900, '203000', '203000', 5, 0, ' ');
INSERT INTO `issue_factory` VALUES (26, 1, 1, 85500, 3, 900, '211500', '211500', 5, 0, ' ');
INSERT INTO `issue_factory` VALUES (27, 1, 1, 85500, 3, 900, '211500', '211500', 5, 0, ' ');
INSERT INTO `issue_factory` VALUES (28, 1, 84, 600, 2, 90, '090000', '230000', 1, 0, ' ');
INSERT INTO `issue_factory` VALUES (29, 2, 97, 600, 2, 90, '', '', 6, 0, ' ');
INSERT INTO `issue_factory` VALUES (30, 1, 84, 600, 3, 90, '085500', '225500', 1, 0, ' ');
COMMIT;

-- ----------------------------
-- Table structure for lotto_odds
-- ----------------------------
DROP TABLE IF EXISTS `lotto_odds`;
CREATE TABLE `lotto_odds` (
  `lotto_id` int(11) NOT NULL COMMENT '彩票编号',
  `method_code` varchar(64) NOT NULL COMMENT '玩法编码',
  `play_code` varchar(64) NOT NULL COMMENT '细编码',
  `method_name` varchar(32) NOT NULL DEFAULT '' COMMENT '玩法名称',
  `play_name` varchar(32) NOT NULL DEFAULT '' COMMENT '细名称',
  `base_odds` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '基础赔率',
  `odds` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '当前赔率',
  `odds_min` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '系统最低赔率',
  `odds_max` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '系统最高赔率',
  `bet_min` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '最低下注',
  `bet_max` decimal(12,3) NOT NULL DEFAULT '50000000.000' COMMENT '最高下注',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示(0:不显示,1:显示)',
  PRIMARY KEY (`lotto_id`,`method_code`,`play_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for lotto_odds_template
-- ----------------------------
DROP TABLE IF EXISTS `lotto_odds_template`;
CREATE TABLE `lotto_odds_template` (
  `lotto_type` int(8) NOT NULL,
  `method_code` varchar(64) NOT NULL COMMENT '玩法编码',
  `play_code` varchar(64) NOT NULL COMMENT '细编码',
  `method_name` varchar(32) NOT NULL DEFAULT '' COMMENT '玩法名称',
  `play_name` varchar(32) NOT NULL DEFAULT '' COMMENT '细名称',
  `base_odds` decimal(12,5) NOT NULL DEFAULT '1.00000' COMMENT '基础赔率',
  `odds` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '当前赔率',
  `odds_min` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '系统最低赔率',
  `odds_max` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '系统最高赔率',
  `bet_min` decimal(12,3) NOT NULL DEFAULT '1.000' COMMENT '最低下注',
  `bet_max` decimal(12,3) NOT NULL DEFAULT '50000000.000' COMMENT '最高下注',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示(0:不显示,1:显示)',
  PRIMARY KEY (`lotto_type`,`method_code`,`play_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of lotto_odds_template
-- ----------------------------
BEGIN;
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '0', '第一球', '0', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '1', '第一球', '1', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '2', '第一球', '2', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '3', '第一球', '3', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '4', '第一球', '4', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '5', '第一球', '5', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '6', '第一球', '6', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '7', '第一球', '7', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '8', '第一球', '8', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10001', '9', '第一球', '9', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '0', '第二球', '0', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '1', '第二球', '1', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '2', '第二球', '2', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '3', '第二球', '3', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '4', '第二球', '4', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '5', '第二球', '5', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '6', '第二球', '6', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '7', '第二球', '7', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '8', '第二球', '8', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10002', '9', '第二球', '9', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '0', '第三球', '0', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '1', '第三球', '1', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '2', '第三球', '2', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '3', '第三球', '3', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '4', '第三球', '4', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '5', '第三球', '5', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '6', '第三球', '6', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '7', '第三球', '7', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '8', '第三球', '8', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10003', '9', '第三球', '9', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '0', '第四球', '0', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '1', '第四球', '1', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '2', '第四球', '2', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '3', '第四球', '3', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '4', '第四球', '4', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '5', '第四球', '5', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '6', '第四球', '6', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '7', '第四球', '7', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '8', '第四球', '8', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10004', '9', '第四球', '9', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '0', '第五球', '0', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '1', '第五球', '1', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '2', '第五球', '2', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '3', '第五球', '3', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '4', '第五球', '4', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '5', '第五球', '5', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '6', '第五球', '6', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '7', '第五球', '7', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '8', '第五球', '8', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10005', '9', '第五球', '9', 9.00000, 9.800, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10006', '单', '第一球', '单', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10006', '双', '第一球', '双', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10006', '大', '第一球', '大', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10006', '小', '第一球', '小', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10007', '单', '第二球', '单', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10007', '双', '第二球', '双', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10007', '大', '第二球', '大', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10007', '小', '第二球', '小', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10008', '单', '第三球', '单', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10008', '双', '第三球', '双', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10008', '大', '第三球', '大', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10008', '小', '第三球', '小', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10009', '单', '第四球', '单', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10009', '双', '第四球', '双', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10009', '大', '第四球', '大', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10009', '小', '第四球', '小', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10010', '单', '第五球', '单', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10010', '双', '第五球', '双', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10010', '大', '第五球', '大', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10010', '小', '第五球', '小', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10011', '单', '总和', '单', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10011', '双', '总和', '双', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10011', '和', '总和', '和', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10011', '大', '总和', '大', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10011', '小', '总和', '小', 1.80000, 1.900, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10011', '虎', '总和', '虎', 2.00000, 2.111, 1.000, 2.222, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10011', '龙', '总和', '龙', 2.00000, 2.111, 1.000, 2.222, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10012', '半顺', '前三', '半顺', 2.45000, 2.586, 1.000, 2.722, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10012', '对子', '前三', '对子', 3.33300, 3.518, 1.000, 3.704, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10012', '杂六', '前三', '杂六', 3.00000, 3.167, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10012', '豹子', '前三', '豹子', 90.00000, 95.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10012', '顺子', '前三', '顺子', 15.00000, 15.833, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10013', '半顺', '中三', '半顺', 2.45000, 2.586, 1.000, 2.722, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10013', '对子', '中三', '对子', 3.33300, 3.518, 1.000, 3.704, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10013', '杂六', '中三', '杂六', 3.00000, 3.167, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10013', '豹子', '中三', '豹子', 90.00000, 95.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10013', '顺子', '中三', '顺子', 15.00000, 15.833, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10014', '半顺', '后三', '半顺', 2.45000, 2.586, 1.000, 2.722, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10014', '对子', '后三', '对子', 3.33300, 3.518, 1.000, 3.704, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10014', '杂六', '后三', '杂六', 3.00000, 3.167, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10014', '豹子', '后三', '豹子', 90.00000, 95.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (1, '10014', '顺子', '后三', '顺子', 15.00000, 15.833, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '1', '第一球', '1', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '10', '第一球', '10', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '11', '第一球', '11', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '2', '第一球', '2', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '3', '第一球', '3', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '4', '第一球', '4', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '5', '第一球', '5', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '6', '第一球', '6', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '7', '第一球', '7', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '8', '第一球', '8', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11001', '9', '第一球', '9', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '1', '第二球', '1', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '10', '第二球', '10', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '11', '第二球', '11', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '2', '第二球', '2', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '3', '第二球', '3', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '4', '第二球', '4', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '5', '第二球', '5', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '6', '第二球', '6', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '7', '第二球', '7', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '8', '第二球', '8', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11002', '9', '第二球', '9', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '1', '第三球', '1', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '10', '第三球', '10', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '11', '第三球', '11', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '2', '第三球', '2', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '3', '第三球', '3', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '4', '第三球', '4', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '5', '第三球', '5', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '6', '第三球', '6', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '7', '第三球', '7', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '8', '第三球', '8', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11003', '9', '第三球', '9', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '1', '第四球', '1', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '10', '第四球', '10', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '11', '第四球', '11', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '2', '第四球', '2', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '3', '第四球', '3', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '4', '第四球', '4', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '5', '第四球', '5', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '6', '第四球', '6', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '7', '第四球', '7', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '8', '第四球', '8', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11004', '9', '第四球', '9', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '1', '第五球', '1', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '10', '第五球', '10', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '11', '第五球', '11', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '2', '第五球', '2', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '3', '第五球', '3', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '4', '第五球', '4', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '5', '第五球', '5', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '6', '第五球', '6', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '7', '第五球', '7', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '8', '第五球', '8', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11005', '9', '第五球', '9', 9.90000, 10.450, 1.000, 11.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11006', '单', '第一球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11006', '双', '第一球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11006', '大', '第一球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11006', '小', '第一球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11007', '单', '第二球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11007', '双', '第二球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11007', '大', '第二球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11007', '小', '第二球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11008', '单', '第三球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11008', '双', '第三球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11008', '大', '第三球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11008', '小', '第三球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11009', '单', '第四球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11009', '双', '第四球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11009', '大', '第四球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11009', '小', '第四球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11010', '单', '第五球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11010', '双', '第五球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11010', '大', '第五球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11010', '小', '第五球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11011', '单', '总合', '单', 1.76100, 1.860, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11011', '双', '总合', '双', 1.84000, 1.942, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11011', '大', '总合', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11011', '小', '总合', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11011', '尾大', '总合', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11011', '尾小', '总合', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11011', '虎', '总合', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11011', '龙', '总合', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11012', '任选1中1', '任选1中1', '任选1中1', 1.98000, 2.090, 1.000, 2.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11013', '任选2中2', '任选2中2', '任选2中2', 4.95000, 5.225, 1.000, 5.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11014', '任选3中3', '任选3中3', '任选3中3', 14.85000, 15.675, 1.000, 16.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11015', '任选4中4', '任选4中4', '任选4中4', 59.40000, 62.700, 1.000, 66.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11016', '任选5中5', '任选5中5', '任选5中5', 415.80000, 438.900, 1.000, 462.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11017', '任选6中5', '任选6中5', '任选6中5', 69.30000, 73.150, 1.000, 77.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11018', '任选7中5', '任选7中5', '任选7中5', 19.80000, 20.900, 1.000, 22.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11019', '任选8中5', '任选8中5', '任选8中5', 7.42500, 7.838, 1.000, 8.250, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11020', '前二组选', '前二组选', '前二组选', 49.50000, 52.250, 1.000, 55.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11021', '前三组选', '前三组选', '前三组选', 267.30000, 282.150, 1.000, 297.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11022', '前二直选', '前二直选', '前二直选', 99.00000, 104.500, 1.000, 110.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (2, '11023', '前三直选', '前三直选', '前三直选', 891.00000, 940.500, 1.000, 990.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '10', '点数', '10', 7.20000, 7.600, 1.000, 8.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '11', '点数', '11', 7.20000, 7.600, 1.000, 8.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '12', '点数', '12', 7.77600, 8.208, 1.000, 8.640, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '13', '点数', '13', 9.25700, 9.771, 1.000, 10.286, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '14', '点数', '14', 12.96000, 13.680, 1.000, 14.400, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '15', '点数', '15', 19.44000, 20.520, 1.000, 21.600, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '16', '点数', '16', 32.40000, 34.200, 1.000, 36.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '17', '点数', '17', 64.80000, 68.400, 1.000, 72.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '4', '点数', '4', 64.80000, 68.400, 1.000, 72.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '5', '点数', '5', 32.40000, 34.200, 1.000, 36.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '6', '点数', '6', 19.44000, 20.520, 1.000, 21.600, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '7', '点数', '7', 12.96000, 13.680, 1.000, 14.400, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '8', '点数', '8', 9.25700, 9.771, 1.000, 10.286, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12001', '9', '点数', '9', 7.77600, 8.208, 1.000, 8.640, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12002', '1', '三军', '1', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12002', '2', '三军', '2', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12002', '3', '三军', '3', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12002', '4', '三军', '4', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12002', '5', '三军', '5', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12002', '6', '三军', '6', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12003', '大', '大小', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12003', '小', '大小', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12004', '111', '围骰', '111', 194.40000, 205.200, 1.000, 216.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12004', '222', '围骰', '222', 194.40000, 205.200, 1.000, 216.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12004', '333', '围骰', '333', 194.40000, 205.200, 1.000, 216.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12004', '444', '围骰', '444', 194.40000, 205.200, 1.000, 216.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12004', '555', '围骰', '555', 194.40000, 205.200, 1.000, 216.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12004', '666', '围骰', '666', 194.40000, 205.200, 1.000, 216.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12005', '1', '全骰', '1', 32.40000, 34.200, 1.000, 36.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '12', '二不同', '12', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '13', '二不同', '13', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '14', '二不同', '14', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '15', '二不同', '15', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '16', '二不同', '16', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '23', '二不同', '23', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '24', '二不同', '24', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '25', '二不同', '25', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '26', '二不同', '26', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '34', '二不同', '34', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '35', '二不同', '35', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '36', '二不同', '36', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '45', '二不同', '45', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '46', '二不同', '46', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12006', '56', '二不同', '56', 6.48000, 6.840, 1.000, 7.200, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12007', '11', '二同', '11', 12.96000, 13.680, 1.000, 14.400, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12007', '22', '二同', '22', 12.96000, 13.680, 1.000, 14.400, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12007', '33', '二同', '33', 12.96000, 13.680, 1.000, 14.400, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12007', '44', '二同', '44', 12.96000, 13.680, 1.000, 14.400, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12007', '55', '二同', '55', 12.96000, 13.680, 1.000, 14.400, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (3, '12007', '66', '二同', '66', 12.96000, 13.680, 1.000, 14.400, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '1', '冠军', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '10', '冠军', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '2', '冠军', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '3', '冠军', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '4', '冠军', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '5', '冠军', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '6', '冠军', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '7', '冠军', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '8', '冠军', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13001', '9', '冠军', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '1', '亚军', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '10', '亚军', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '2', '亚军', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '3', '亚军', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '4', '亚军', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '5', '亚军', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '6', '亚军', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '7', '亚军', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '8', '亚军', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13002', '9', '亚军', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '1', '季军', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '10', '季军', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '2', '季军', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '3', '季军', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '4', '季军', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '5', '季军', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '6', '季军', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '7', '季军', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '8', '季军', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13003', '9', '季军', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '1', '第四名', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '10', '第四名', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '2', '第四名', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '3', '第四名', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '4', '第四名', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '5', '第四名', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '6', '第四名', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '7', '第四名', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '8', '第四名', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13004', '9', '第四名', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '1', '第五名', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '10', '第五名', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '2', '第五名', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '3', '第五名', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '4', '第五名', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '5', '第五名', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '6', '第五名', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '7', '第五名', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '8', '第五名', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13005', '9', '第五名', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '1', '第六名', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '10', '第六名', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '2', '第六名', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '3', '第六名', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '4', '第六名', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '5', '第六名', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '6', '第六名', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '7', '第六名', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '8', '第六名', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13006', '9', '第六名', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '1', '第七名', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '10', '第七名', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '2', '第七名', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '3', '第七名', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '4', '第七名', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '5', '第七名', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '6', '第七名', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '7', '第七名', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '8', '第七名', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13007', '9', '第七名', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '1', '第八名', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '10', '第八名', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '2', '第八名', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '3', '第八名', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '4', '第八名', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '5', '第八名', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '6', '第八名', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '7', '第八名', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '8', '第八名', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13008', '9', '第八名', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '1', '第九名', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '10', '第九名', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '2', '第九名', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '3', '第九名', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '4', '第九名', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '5', '第九名', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '6', '第九名', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '7', '第九名', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '8', '第九名', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13009', '9', '第九名', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '1', '第十名', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '10', '第十名', '10', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '2', '第十名', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '3', '第十名', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '4', '第十名', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '5', '第十名', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '6', '第十名', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '7', '第十名', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '8', '第十名', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13010', '9', '第十名', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13011', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13011', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13011', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13011', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13011', '虎', '双面', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13011', '龙', '双面', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13012', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13012', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13012', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13012', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13012', '虎', '双面', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13012', '龙', '双面', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13013', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13013', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13013', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13013', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13013', '虎', '双面', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13013', '龙', '双面', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13014', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13014', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13014', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13014', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13014', '虎', '双面', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13014', '龙', '双面', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13015', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13015', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13015', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13015', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13015', '虎', '双面', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13015', '龙', '双面', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13016', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13016', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13016', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13016', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13017', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13017', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13017', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13017', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13018', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13018', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13018', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13018', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13019', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13019', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13019', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13019', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13020', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13020', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13020', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13020', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13021', '单', '冠亚合双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13021', '双', '冠亚合双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13021', '大', '冠亚合双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13021', '小', '冠亚合双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '10', '冠亚', '10', 10.12500, 10.688, 1.000, 11.250, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '11', '冠亚', '11', 8.10000, 8.550, 1.000, 9.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '12', '冠亚', '12', 10.12500, 10.688, 1.000, 11.250, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '13', '冠亚', '13', 10.12500, 10.688, 1.000, 11.250, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '14', '冠亚', '14', 13.50000, 14.250, 1.000, 15.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '15', '冠亚', '15', 13.50000, 14.250, 1.000, 15.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '16', '冠亚', '16', 20.25000, 21.375, 1.000, 22.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '17', '冠亚', '17', 20.25000, 21.375, 1.000, 22.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '18', '冠亚', '18', 40.50000, 42.750, 1.000, 45.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '19', '冠亚', '19', 40.50000, 42.750, 1.000, 45.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '3', '冠亚', '3', 40.50000, 42.750, 1.000, 45.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '4', '冠亚', '4', 40.50000, 42.750, 1.000, 45.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '5', '冠亚', '5', 20.25000, 21.375, 1.000, 22.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '6', '冠亚', '6', 20.25000, 21.375, 1.000, 22.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '7', '冠亚', '7', 13.50000, 14.250, 1.000, 15.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '8', '冠亚', '8', 13.50000, 14.250, 1.000, 15.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (4, '13023', '9', '冠亚', '9', 10.12500, 10.688, 1.000, 11.250, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '0', '一字组合', '0', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '1', '一字组合', '1', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '2', '一字组合', '2', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '3', '一字组合', '3', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '4', '一字组合', '4', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '5', '一字组合', '5', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '6', '一字组合', '6', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '7', '一字组合', '7', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '8', '一字组合', '8', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14002', '9', '一字组合', '9', 3.00000, 3.170, 1.000, 3.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '00', '二字组合', '00', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '01', '二字组合', '01', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '02', '二字组合', '02', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '03', '二字组合', '03', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '04', '二字组合', '04', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '05', '二字组合', '05', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '06', '二字组合', '06', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '07', '二字组合', '07', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '08', '二字组合', '08', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '09', '二字组合', '09', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '11', '二字组合', '11', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '12', '二字组合', '12', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '13', '二字组合', '13', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '14', '二字组合', '14', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '15', '二字组合', '15', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '16', '二字组合', '16', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '17', '二字组合', '17', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '18', '二字组合', '18', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '19', '二字组合', '19', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '22', '二字组合', '22', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '23', '二字组合', '23', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '24', '二字组合', '24', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '25', '二字组合', '25', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '26', '二字组合', '26', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '27', '二字组合', '27', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '28', '二字组合', '28', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '29', '二字组合', '29', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '33', '二字组合', '33', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '34', '二字组合', '34', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '35', '二字组合', '35', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '36', '二字组合', '36', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '37', '二字组合', '37', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '38', '二字组合', '38', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '39', '二字组合', '39', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '44', '二字组合', '44', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '45', '二字组合', '45', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '46', '二字组合', '46', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '47', '二字组合', '47', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '48', '二字组合', '48', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '49', '二字组合', '49', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '55', '二字组合', '55', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '56', '二字组合', '56', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '57', '二字组合', '57', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '58', '二字组合', '58', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '59', '二字组合', '59', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '66', '二字组合', '66', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '67', '二字组合', '67', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '68', '二字组合', '68', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '69', '二字组合', '69', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '77', '二字组合', '77', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '78', '二字组合', '78', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '79', '二字组合', '79', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '88', '二字组合', '88', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '89', '二字组合', '89', 16.66700, 17.593, 1.000, 18.519, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14003', '99', '二字组合', '99', 32.14300, 33.929, 1.000, 335.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '000', '三字组合', '000', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '001', '三字组合', '001', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '002', '三字组合', '002', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '003', '三字组合', '003', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '004', '三字组合', '004', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '005', '三字组合', '005', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '006', '三字组合', '006', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '007', '三字组合', '007', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '008', '三字组合', '008', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '009', '三字组合', '009', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '011', '三字组合', '011', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '012', '三字组合', '012', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '013', '三字组合', '013', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '014', '三字组合', '014', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '015', '三字组合', '015', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '016', '三字组合', '016', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '017', '三字组合', '017', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '018', '三字组合', '018', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '019', '三字组合', '019', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '022', '三字组合', '022', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '023', '三字组合', '023', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '024', '三字组合', '024', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '025', '三字组合', '025', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '026', '三字组合', '026', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '027', '三字组合', '027', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '028', '三字组合', '028', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '029', '三字组合', '029', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '033', '三字组合', '033', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '034', '三字组合', '034', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '035', '三字组合', '035', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '036', '三字组合', '036', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '037', '三字组合', '037', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '038', '三字组合', '038', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '039', '三字组合', '039', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '044', '三字组合', '044', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '045', '三字组合', '045', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '046', '三字组合', '046', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '047', '三字组合', '047', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '048', '三字组合', '048', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '049', '三字组合', '049', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '055', '三字组合', '055', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '056', '三字组合', '056', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '057', '三字组合', '057', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '058', '三字组合', '058', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '059', '三字组合', '059', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '066', '三字组合', '066', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '067', '三字组合', '067', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '068', '三字组合', '068', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '069', '三字组合', '069', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '077', '三字组合', '077', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '078', '三字组合', '078', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '079', '三字组合', '079', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '088', '三字组合', '088', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '089', '三字组合', '089', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '099', '三字组合', '099', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '111', '三字组合', '111', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '112', '三字组合', '112', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '113', '三字组合', '113', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '114', '三字组合', '114', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '115', '三字组合', '115', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '116', '三字组合', '116', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '117', '三字组合', '117', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '118', '三字组合', '118', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '119', '三字组合', '119', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '122', '三字组合', '122', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '123', '三字组合', '123', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '124', '三字组合', '124', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '125', '三字组合', '125', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '126', '三字组合', '126', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '127', '三字组合', '127', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '128', '三字组合', '128', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '129', '三字组合', '129', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '133', '三字组合', '133', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '134', '三字组合', '134', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '135', '三字组合', '135', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '136', '三字组合', '136', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '137', '三字组合', '137', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '138', '三字组合', '138', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '139', '三字组合', '139', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '144', '三字组合', '144', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '145', '三字组合', '145', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '146', '三字组合', '146', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '147', '三字组合', '147', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '148', '三字组合', '148', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '149', '三字组合', '149', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '155', '三字组合', '155', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '156', '三字组合', '156', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '157', '三字组合', '157', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '158', '三字组合', '158', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '159', '三字组合', '159', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '166', '三字组合', '166', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '167', '三字组合', '167', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '168', '三字组合', '168', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '169', '三字组合', '169', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '177', '三字组合', '177', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '178', '三字组合', '178', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '179', '三字组合', '179', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '188', '三字组合', '188', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '189', '三字组合', '189', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '199', '三字组合', '199', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '222', '三字组合', '222', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '223', '三字组合', '223', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '224', '三字组合', '224', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '225', '三字组合', '225', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '226', '三字组合', '226', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '227', '三字组合', '227', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '228', '三字组合', '228', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '229', '三字组合', '229', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '233', '三字组合', '233', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '234', '三字组合', '234', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '235', '三字组合', '235', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '236', '三字组合', '236', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '237', '三字组合', '237', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '238', '三字组合', '238', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '239', '三字组合', '239', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '244', '三字组合', '244', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '245', '三字组合', '245', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '246', '三字组合', '246', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '247', '三字组合', '247', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '248', '三字组合', '248', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '249', '三字组合', '249', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '255', '三字组合', '255', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '256', '三字组合', '256', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '257', '三字组合', '257', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '258', '三字组合', '258', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '259', '三字组合', '259', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '266', '三字组合', '266', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '267', '三字组合', '267', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '268', '三字组合', '268', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '269', '三字组合', '269', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '277', '三字组合', '277', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '278', '三字组合', '278', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '279', '三字组合', '279', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '288', '三字组合', '288', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '289', '三字组合', '289', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '299', '三字组合', '299', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '333', '三字组合', '333', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '334', '三字组合', '334', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '335', '三字组合', '335', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '336', '三字组合', '336', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '337', '三字组合', '337', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '338', '三字组合', '338', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '339', '三字组合', '339', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '344', '三字组合', '344', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '345', '三字组合', '345', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '346', '三字组合', '346', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '347', '三字组合', '347', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '348', '三字组合', '348', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '349', '三字组合', '349', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '355', '三字组合', '355', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '356', '三字组合', '356', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '357', '三字组合', '357', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '358', '三字组合', '358', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '359', '三字组合', '359', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '366', '三字组合', '366', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '367', '三字组合', '367', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '368', '三字组合', '368', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '369', '三字组合', '369', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '377', '三字组合', '377', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '378', '三字组合', '378', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '379', '三字组合', '379', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '388', '三字组合', '388', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '389', '三字组合', '389', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '399', '三字组合', '399', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '444', '三字组合', '444', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '445', '三字组合', '445', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '446', '三字组合', '446', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '447', '三字组合', '447', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '448', '三字组合', '448', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '449', '三字组合', '449', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '455', '三字组合', '455', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '456', '三字组合', '456', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '457', '三字组合', '457', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '458', '三字组合', '458', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '459', '三字组合', '459', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '466', '三字组合', '466', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '467', '三字组合', '467', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '468', '三字组合', '468', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '469', '三字组合', '469', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '477', '三字组合', '477', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '478', '三字组合', '478', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '479', '三字组合', '479', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '488', '三字组合', '488', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '489', '三字组合', '489', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '499', '三字组合', '499', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '555', '三字组合', '555', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '556', '三字组合', '556', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '557', '三字组合', '557', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '558', '三字组合', '558', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '559', '三字组合', '559', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '566', '三字组合', '566', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '567', '三字组合', '567', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '568', '三字组合', '568', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '569', '三字组合', '569', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '577', '三字组合', '577', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '578', '三字组合', '578', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '579', '三字组合', '579', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '588', '三字组合', '588', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '589', '三字组合', '589', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '599', '三字组合', '599', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '666', '三字组合', '666', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '667', '三字组合', '667', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '668', '三字组合', '668', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '669', '三字组合', '669', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '677', '三字组合', '677', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '678', '三字组合', '678', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '679', '三字组合', '679', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '688', '三字组合', '688', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '689', '三字组合', '689', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '699', '三字组合', '699', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '777', '三字组合', '777', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '778', '三字组合', '778', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '779', '三字组合', '779', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '788', '三字组合', '788', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '789', '三字组合', '789', 150.00000, 149.500, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '799', '三字组合', '799', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '888', '三字组合', '888', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '889', '三字组合', '889', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '899', '三字组合', '899', 300.00000, 280.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14004', '999', '三字组合', '999', 900.00000, 920.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '0', '一字定位百', '0', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '1', '一字定位百', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '2', '一字定位百', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '3', '一字定位百', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '4', '一字定位百', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '5', '一字定位百', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '6', '一字定位百', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '7', '一字定位百', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '8', '一字定位百', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14005', '9', '一字定位百', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '0', '一字定位十', '0', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '1', '一字定位十', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '2', '一字定位十', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '3', '一字定位十', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '4', '一字定位十', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '5', '一字定位十', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '6', '一字定位十', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '7', '一字定位十', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '8', '一字定位十', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14006', '9', '一字定位十', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '0', '一字定位个', '0', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '1', '一字定位个', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '2', '一字定位个', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '3', '一字定位个', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '4', '一字定位个', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '5', '一字定位个', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '6', '一字定位个', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '7', '一字定位个', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '8', '一字定位个', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14007', '9', '一字定位个', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14008', '1', '二字定位-百十', '1', 90.00000, 95.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14009', '1', '二字定位-百个', '1', 90.00000, 95.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14010', '1', '二字定位-十个', '1', 90.00000, 95.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14011', '1', '三字定位', '1', 900.00000, 950.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 0', '二字和数-百十', ' 0', 90.00000, 87.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 1', '二字和数-百十', ' 1', 45.00000, 43.500, 1.000, 50.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 2', '二字和数-百十', ' 2', 30.00000, 29.000, 1.000, 33.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 3', '二字和数-百十', ' 3', 22.50000, 21.750, 1.000, 25.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 4', '二字和数-百十', ' 4', 18.00000, 17.400, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 5', '二字和数-百十', ' 5', 15.00000, 14.500, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 6', '二字和数-百十', ' 6', 12.85700, 12.430, 1.000, 14.286, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 7', '二字和数-百十', ' 7', 11.25000, 10.880, 1.000, 12.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 8', '二字和数-百十', ' 8', 10.00000, 9.670, 1.000, 11.111, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', ' 9', '二字和数-百十', ' 9', 9.00000, 8.700, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '10', '二字和数-百十', '10', 10.00000, 9.670, 1.000, 11.111, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '11', '二字和数-百十', '11', 11.25000, 10.880, 1.000, 12.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '12', '二字和数-百十', '12', 12.85700, 12.430, 1.000, 14.286, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '13', '二字和数-百十', '13', 15.00000, 14.500, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '14', '二字和数-百十', '14', 18.00000, 17.400, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '15', '二字和数-百十', '15', 22.50000, 21.750, 1.000, 25.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '16', '二字和数-百十', '16', 30.00000, 29.000, 1.000, 33.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '17', '二字和数-百十', '17', 45.00000, 43.500, 1.000, 50.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14012', '18', '二字和数-百十', '18', 90.00000, 87.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 0', '二字和数-百个', ' 0', 90.00000, 87.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 1', '二字和数-百个', ' 1', 45.00000, 43.500, 1.000, 50.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 2', '二字和数-百个', ' 2', 30.00000, 29.000, 1.000, 33.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 3', '二字和数-百个', ' 3', 22.50000, 21.750, 1.000, 25.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 4', '二字和数-百个', ' 4', 18.00000, 17.400, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 5', '二字和数-百个', ' 5', 15.00000, 14.500, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 6', '二字和数-百个', ' 6', 12.85700, 12.430, 1.000, 14.286, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 7', '二字和数-百个', ' 7', 11.25000, 10.880, 1.000, 12.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 8', '二字和数-百个', ' 8', 10.00000, 9.670, 1.000, 11.111, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', ' 9', '二字和数-百个', ' 9', 9.00000, 8.700, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '10', '二字和数-百个', '10', 10.00000, 9.670, 1.000, 11.111, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '11', '二字和数-百个', '11', 11.25000, 10.880, 1.000, 12.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '12', '二字和数-百个', '12', 12.85700, 12.430, 1.000, 14.286, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '13', '二字和数-百个', '13', 15.00000, 14.500, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '14', '二字和数-百个', '14', 18.00000, 17.400, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '15', '二字和数-百个', '15', 22.50000, 21.750, 1.000, 25.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '16', '二字和数-百个', '16', 30.00000, 29.000, 1.000, 33.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '17', '二字和数-百个', '17', 45.00000, 43.500, 1.000, 50.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14013', '18', '二字和数-百个', '18', 90.00000, 87.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 0', '二字和数-十个', ' 0', 90.00000, 87.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 1', '二字和数-十个', ' 1', 45.00000, 43.500, 1.000, 50.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 2', '二字和数-十个', ' 2', 30.00000, 29.000, 1.000, 33.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 3', '二字和数-十个', ' 3', 22.50000, 21.750, 1.000, 25.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 4', '二字和数-十个', ' 4', 18.00000, 17.400, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 5', '二字和数-十个', ' 5', 15.00000, 14.500, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 6', '二字和数-十个', ' 6', 12.85700, 12.430, 1.000, 14.286, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 7', '二字和数-十个', ' 7', 11.25000, 10.880, 1.000, 12.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 8', '二字和数-十个', ' 8', 10.00000, 9.670, 1.000, 11.111, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', ' 9', '二字和数-十个', ' 9', 9.00000, 8.700, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '10', '二字和数-十个', '10', 10.00000, 9.670, 1.000, 11.111, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '11', '二字和数-十个', '11', 11.25000, 10.880, 1.000, 12.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '12', '二字和数-十个', '12', 12.85700, 12.430, 1.000, 14.286, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '13', '二字和数-十个', '13', 15.00000, 14.500, 1.000, 16.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '14', '二字和数-十个', '14', 18.00000, 17.400, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '15', '二字和数-十个', '15', 22.50000, 21.750, 1.000, 25.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '16', '二字和数-十个', '16', 30.00000, 29.000, 1.000, 33.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '17', '二字和数-十个', '17', 45.00000, 43.500, 1.000, 50.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14014', '18', '二字和数-十个', '18', 90.00000, 87.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '0', ' 三字和数', '0', 900.00000, 720.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '1', ' 三字和数', '1', 300.00000, 240.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '10', '三字和数', '10', 14.28600, 13.900, 1.000, 15.873, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '11', '三字和数', '11', 13.04300, 12.800, 1.000, 14.492, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '12', '三字和数', '12', 12.32900, 12.300, 1.000, 13.699, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '13', '三字和数', '13', 12.00000, 12.000, 1.000, 13.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '14', '三字和数', '14', 12.00000, 1.980, 1.000, 13.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '15', '三字和数', '15', 12.32900, 12.300, 1.000, 13.699, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '16', '三字和数', '16', 13.04300, 12.800, 1.000, 14.492, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '17', '三字和数', '17', 14.28600, 13.900, 1.000, 15.873, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '18', '三字和数', '18', 16.36400, 15.800, 1.000, 18.182, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '19', '三字和数', '19', 20.00000, 19.300, 1.000, 22.222, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '2', ' 三字和数', '2', 150.00000, 136.000, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '20', '三字和数', '20', 25.00000, 24.300, 1.000, 27.778, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '21', '三字和数', '21', 32.14300, 31.000, 1.000, 35.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '22', '三字和数', '22', 42.85700, 39.000, 1.000, 47.619, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '23', '三字和数', '23', 60.00000, 54.500, 1.000, 66.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '24', '三字和数', '24', 90.00000, 82.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '25', '三字和数', '25', 150.00000, 136.000, 1.000, 166.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '26', '三字和数', '26', 300.00000, 240.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '27', '三字和数', '27', 900.00000, 720.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '3', ' 三字和数', '3', 90.00000, 82.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '4', ' 三字和数', '4', 60.00000, 54.000, 1.000, 66.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '5', ' 三字和数', '5', 42.85700, 39.000, 1.000, 47.619, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '6', ' 三字和数', '6', 32.14300, 31.000, 1.000, 35.714, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '7', ' 三字和数', '7', 25.00000, 24.300, 1.000, 27.778, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '8', ' 三字和数', '8', 20.00000, 19.300, 1.000, 22.222, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14015', '9', ' 三字和数', '9', 16.36400, 15.800, 1.000, 18.182, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '0', ' 二字和数尾-百十', '0', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '1', ' 二字和数尾-百十', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '2', ' 二字和数尾-百十', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '3', ' 二字和数尾-百十', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '4', ' 二字和数尾-百十', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '5', ' 二字和数尾-百十', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '6', ' 二字和数尾-百十', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '7', ' 二字和数尾-百十', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '8', ' 二字和数尾-百十', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14016', '9', ' 二字和数尾-百十', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '0', ' 二字和数尾-百个', '0', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '1', ' 二字和数尾-百个', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '2', ' 二字和数尾-百个', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '3', ' 二字和数尾-百个', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '4', ' 二字和数尾-百个', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '5', ' 二字和数尾-百个', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '6', ' 二字和数尾-百个', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '7', ' 二字和数尾-百个', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '8', ' 二字和数尾-百个', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14017', '9', ' 二字和数尾-百个', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '0', ' 二字和数尾-十个', '0', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '1', ' 二字和数尾-十个', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '2', ' 二字和数尾-十个', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '3', ' 二字和数尾-十个', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '4', ' 二字和数尾-十个', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '5', ' 二字和数尾-十个', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '6', ' 二字和数尾-十个', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '7', ' 二字和数尾-十个', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '8', ' 二字和数尾-十个', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14018', '9', ' 二字和数尾-十个', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '0', ' 二字和数尾-百十个', '0', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '1', ' 二字和数尾-百十个', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '2', ' 二字和数尾-百十个', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '3', ' 二字和数尾-百十个', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '4', ' 二字和数尾-百十个', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '5', ' 二字和数尾-百十个', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '6', ' 二字和数尾-百十个', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '7', ' 二字和数尾-百十个', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '8', ' 二字和数尾-百十个', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14019', '9', ' 二字和数尾-百十个', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '0', ' 跨度', '0', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '1', ' 跨度', '1', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '2', ' 跨度', '2', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '3', ' 跨度', '3', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '4', ' 跨度', '4', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '5', ' 跨度', '5', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '6', ' 跨度', '6', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '7', ' 跨度', '7', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '8', ' 跨度', '8', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14020', '9', ' 跨度', '9', 9.00000, 9.500, 1.000, 10.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14021', '单', ' 双面-百', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14021', '双', ' 双面-百', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14021', '合', ' 双面-百', '合', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14021', '大', ' 双面-百', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14021', '小', ' 双面-百', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14021', '质', ' 双面-百', '质', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14022', '单', ' 双面-十', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14022', '双', ' 双面-十', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14022', '合', ' 双面-十', '合', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14022', '大', ' 双面-十', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14022', '小', ' 双面-十', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14022', '质', ' 双面-十', '质', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14023', '单', ' 双面-个', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14023', '双', ' 双面-个', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14023', '合', ' 双面-个', '合', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14023', '大', ' 双面-个', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14023', '小', ' 双面-个', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14023', '质', ' 双面-个', '质', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14024', '单', ' 双面合数-百十', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14024', '双', ' 双面合数-百十', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14025', '单', ' 双面合数-百个', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14025', '双', ' 双面合数-百个', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14026', '单', ' 双面合数-十个', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14026', '双', ' 双面合数-十个', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14027', '单', ' 双面合数-百十个', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14027', '双', ' 双面合数-百十个', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14027', '大', ' 双面合数-百十个', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (5, '14027', '小', ' 双面合数-百十个', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '1', ' 第1球', '1', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '10', ' 第1球', '10', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '11', ' 第1球', '11', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '12', ' 第1球', '12', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '13', ' 第1球', '13', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '14', ' 第1球', '14', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '15', ' 第1球', '15', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '16', ' 第1球', '16', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '17', ' 第1球', '17', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '18', ' 第1球', '18', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '19', ' 第1球', '19', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '2', ' 第1球', '2', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '20', ' 第1球', '20', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '3', ' 第1球', '3', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '4', ' 第1球', '4', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '5', ' 第1球', '5', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '6', ' 第1球', '6', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '7', ' 第1球', '7', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '8', ' 第1球', '8', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15001', '9', ' 第1球', '9', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '1', '第2球', '1', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '10', '第2球', '10', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '11', '第2球', '11', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '12', '第2球', '12', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '13', '第2球', '13', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '14', '第2球', '14', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '15', '第2球', '15', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '16', '第2球', '16', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '17', '第2球', '17', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '18', '第2球', '18', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '19', '第2球', '19', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '2', '第2球', '2', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '20', '第2球', '20', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '3', '第2球', '3', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '4', '第2球', '4', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '5', '第2球', '5', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '6', '第2球', '6', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '7', '第2球', '7', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '8', '第2球', '8', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15002', '9', '第2球', '9', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '1', '第3球', '1', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '10', '第3球', '10', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '11', '第3球', '11', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '12', '第3球', '12', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '13', '第3球', '13', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '14', '第3球', '14', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '15', '第3球', '15', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '16', '第3球', '16', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '17', '第3球', '17', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '18', '第3球', '18', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '19', '第3球', '19', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '2', '第3球', '2', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '20', '第3球', '20', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '3', '第3球', '3', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '4', '第3球', '4', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '5', '第3球', '5', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '6', '第3球', '6', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '7', '第3球', '7', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '8', '第3球', '8', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15003', '9', '第3球', '9', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '1', '第4球', '1', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '10', '第4球', '10', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '11', '第4球', '11', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '12', '第4球', '12', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '13', '第4球', '13', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '14', '第4球', '14', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '15', '第4球', '15', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '16', '第4球', '16', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '17', '第4球', '17', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '18', '第4球', '18', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '19', '第4球', '19', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '2', '第4球', '2', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '20', '第4球', '20', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '3', '第4球', '3', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '4', '第4球', '4', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '5', '第4球', '5', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '6', '第4球', '6', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '7', '第4球', '7', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '8', '第4球', '8', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15004', '9', '第4球', '9', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '1', '第5球', '1', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '10', '第5球', '10', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '11', '第5球', '11', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '12', '第5球', '12', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '13', '第5球', '13', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '14', '第5球', '14', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '15', '第5球', '15', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '16', '第5球', '16', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '17', '第5球', '17', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '18', '第5球', '18', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '19', '第5球', '19', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '2', '第5球', '2', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '20', '第5球', '20', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '3', '第5球', '3', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '4', '第5球', '4', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '5', '第5球', '5', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '6', '第5球', '6', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '7', '第5球', '7', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '8', '第5球', '8', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15005', '9', '第5球', '9', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '1', '第6球', '1', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '10', '第6球', '10', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '11', '第6球', '11', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '12', '第6球', '12', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '13', '第6球', '13', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '14', '第6球', '14', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '15', '第6球', '15', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '16', '第6球', '16', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '17', '第6球', '17', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '18', '第6球', '18', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '19', '第6球', '19', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '2', '第6球', '2', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '20', '第6球', '20', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '3', '第6球', '3', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '4', '第6球', '4', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '5', '第6球', '5', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '6', '第6球', '6', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '7', '第6球', '7', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '8', '第6球', '8', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15006', '9', '第6球', '9', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '1', '第7球', '1', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '10', '第7球', '10', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '11', '第7球', '11', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '12', '第7球', '12', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '13', '第7球', '13', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '14', '第7球', '14', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '15', '第7球', '15', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '16', '第7球', '16', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '17', '第7球', '17', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '18', '第7球', '18', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '19', '第7球', '19', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '2', '第7球', '2', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '20', '第7球', '20', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '3', '第7球', '3', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '4', '第7球', '4', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '5', '第7球', '5', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '6', '第7球', '6', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '7', '第7球', '7', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '8', '第7球', '8', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15007', '9', '第7球', '9', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '1', '第8球', '1', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '10', '第8球', '10', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '11', '第8球', '11', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '12', '第8球', '12', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '13', '第8球', '13', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '14', '第8球', '14', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '15', '第8球', '15', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '16', '第8球', '16', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '17', '第8球', '17', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '18', '第8球', '18', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '19', '第8球', '19', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '2', '第8球', '2', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '20', '第8球', '20', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '3', '第8球', '3', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '4', '第8球', '4', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '5', '第8球', '5', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '6', '第8球', '6', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '7', '第8球', '7', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '8', '第8球', '8', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15008', '9', '第8球', '9', 18.00000, 19.000, 1.000, 20.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '东', '第1球', '东', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '北', '第1球', '北', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '单', '第1球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '南', '第1球', '南', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '双', '第1球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '合数单', '第1球', '合数单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '合数双', '第1球', '合数双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '大', '第1球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '小', '第1球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '尾大', '第1球', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '尾小', '第1球', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '虎', '第1球', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '西', '第1球', '西', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15009', '龙', '第1球', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '东', '第2球', '东', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '北', '第2球', '北', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '单', '第2球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '南', '第2球', '南', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '双', '第2球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '合数单', '第2球', '合数单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '合数双', '第2球', '合数双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '大', '第2球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '小', '第2球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '尾大', '第2球', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '尾小', '第2球', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '虎', '第2球', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '西', '第2球', '西', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15010', '龙', '第2球', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '东', '第3球', '东', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '北', '第3球', '北', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '单', '第3球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '南', '第3球', '南', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '双', '第3球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '合数单', '第3球', '合数单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '合数双', '第3球', '合数双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '大', '第3球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '小', '第3球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '尾大', '第3球', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '尾小', '第3球', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '虎', '第3球', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '西', '第3球', '西', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15011', '龙', '第3球', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '东', '第4球', '东', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '北', '第4球', '北', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '单', '第4球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '南', '第4球', '南', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '双', '第4球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '合数单', '第4球', '合数单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '合数双', '第4球', '合数双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '大', '第4球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '小', '第4球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '尾大', '第4球', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '尾小', '第4球', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '虎', '第4球', '虎', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '西', '第4球', '西', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15012', '龙', '第4球', '龙', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '东', '第5球', '东', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '北', '第5球', '北', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '单', '第5球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '南', '第5球', '南', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '双', '第5球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '合数单', '第5球', '合数单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '合数双', '第5球', '合数双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '大', '第5球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '小', '第5球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '尾大', '第5球', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '尾小', '第5球', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15013', '西', '第5球', '西', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '东', '第6球', '东', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '北', '第6球', '北', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '单', '第6球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '南', '第6球', '南', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '双', '第6球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '合数单', '第6球', '合数单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '合数双', '第6球', '合数双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '大', '第6球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '小', '第6球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '尾大', '第6球', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '尾小', '第6球', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15014', '西', '第6球', '西', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '东', '第7球', '东', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '北', '第7球', '北', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '单', '第7球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '南', '第7球', '南', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '双', '第7球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '合数单', '第7球', '合数单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '合数双', '第7球', '合数双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '大', '第7球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '小', '第7球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '尾大', '第7球', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '尾小', '第7球', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15015', '西', '第7球', '西', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '东', '第8球', '东', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '北', '第8球', '北', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '单', '第8球', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '南', '第8球', '南', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '双', '第8球', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '合数单', '第8球', '合数单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '合数双', '第8球', '合数双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '大', '第8球', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '小', '第8球', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '尾大', '第8球', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '尾小', '第8球', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15016', '西', '第8球', '西', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15017', '单', '总和', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15017', '双', '总和', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15017', '大', '总和', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15017', '小', '总和', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15017', '尾大', '总和', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15017', '尾小', '总和', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '1', '正码', '1', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '10', '正码', '10', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '11', '正码', '11', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '12', '正码', '12', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '13', '正码', '13', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '14', '正码', '14', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '15', '正码', '15', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '16', '正码', '16', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '17', '正码', '17', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '18', '正码', '18', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '19', '正码', '19', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '2', '正码', '2', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '20', '正码', '20', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '3', '正码', '3', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '4', '正码', '4', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '5', '正码', '5', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '6', '正码', '6', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '7', '正码', '7', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '8', '正码', '8', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (7, '15018', '9', '正码', '9', 2.25000, 2.375, 1.000, 2.500, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '810', '基诺', '810', 106.00000, 108.000, 1.000, 118.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '上', '基诺', '上', 2.25900, 2.385, 1.000, 2.510, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '下', '基诺', '下', 2.25900, 2.385, 1.000, 2.510, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '中', '基诺', '中', 4.42800, 4.674, 1.000, 4.920, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '偶', '基诺', '偶', 2.25900, 2.385, 1.000, 2.510, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '单', '基诺', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '双', '基诺', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '和', '基诺', '和', 4.42800, 4.674, 1.000, 4.920, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '土', '基诺', '土', 9.03700, 9.539, 1.000, 10.041, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '大', '基诺', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '大单', '基诺', '大单', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '大双', '基诺', '大双', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '奇', '基诺', '奇', 2.25900, 2.385, 1.000, 2.510, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '小', '基诺', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '小单', '基诺', '小单', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '小双', '基诺', '小双', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '木', '基诺', '木', 4.51800, 4.769, 1.000, 5.020, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '水', '基诺', '水', 2.35700, 2.488, 1.000, 2.619, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '火', '基诺', '火', 4.51800, 4.769, 1.000, 5.020, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (8, '16001', '金', '基诺', '金', 9.03700, 9.539, 1.000, 10.041, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '0', '数字', '0', 900.00000, 30.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '1', '数字', '1', 300.00000, 26.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '10', '数字', '10', 14.28300, 11.000, 1.000, 15.870, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '11', '数字', '11', 13.04600, 11.000, 1.000, 14.496, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '12', '数字', '12', 12.33000, 11.000, 1.000, 13.700, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '13', '数字', '13', 12.00000, 11.000, 1.000, 13.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '14', '数字', '14', 12.00000, 11.000, 1.000, 13.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '15', '数字', '15', 12.33000, 11.000, 1.000, 13.700, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '16', '数字', '16', 13.04600, 11.000, 1.000, 14.496, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '17', '数字', '17', 14.28300, 11.000, 1.000, 15.870, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '18', '数字', '18', 16.36600, 11.000, 1.000, 18.184, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '19', '数字', '19', 20.00000, 11.000, 1.000, 22.222, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '2', '数字', '2', 150.30000, 23.000, 1.000, 167.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '20', '数字', '20', 25.02000, 11.000, 1.000, 27.800, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '21', '数字', '21', 32.17500, 11.000, 1.000, 35.750, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '22', '数字', '22', 42.89400, 17.000, 1.000, 47.660, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '23', '数字', '23', 60.00000, 19.000, 1.000, 66.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '24', '数字', '24', 100.00000, 21.000, 1.000, 111.111, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '25', '数字', '25', 150.30000, 23.000, 1.000, 167.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '26', '数字', '26', 300.00000, 26.000, 1.000, 333.333, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '27', '数字', '27', 900.00000, 30.000, 1.000, 1000.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '3', '数字', '3', 100.00000, 21.000, 1.000, 111.111, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '4', '数字', '4', 60.00000, 19.000, 1.000, 66.667, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '5', '数字', '5', 42.89400, 17.000, 1.000, 47.660, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '6', '数字', '6', 32.17500, 11.000, 1.000, 35.750, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '7', '数字', '7', 25.02000, 11.000, 1.000, 27.800, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '8', '数字', '8', 20.00000, 11.000, 1.000, 22.222, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15009', '9', '数字', '9', 16.36600, 11.000, 1.000, 18.184, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15010', '红', '色波', '红', 2.70900, 3.000, 1.000, 3.010, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15010', '绿', '色波', '绿', 3.48700, 3.000, 1.000, 3.874, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15010', '蓝', '色波', '蓝', 3.48700, 3.000, 1.000, 3.874, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15011', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15011', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15011', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15011', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15012', '大单', '组合', '大单', 3.60000, 3.900, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15012', '大双', '组合', '大双', 3.60000, 3.900, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15012', '小单', '组合', '小单', 3.60000, 3.900, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15012', '小双', '组合', '小双', 3.60000, 3.900, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15013', '极大', '趣味', '极大', 16.20000, 10.000, 1.000, 18.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15013', '极小', '趣味', '极小', 16.20000, 10.000, 1.000, 18.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (9, '15013', '豹子', '趣味', '豹子', 90.00000, 50.000, 1.000, 100.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '1', '特码A', '1', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '10', '特码A', '10', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '11', '特码A', '11', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '12', '特码A', '12', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '13', '特码A', '13', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '14', '特码A', '14', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '15', '特码A', '15', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '16', '特码A', '16', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '17', '特码A', '17', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '18', '特码A', '18', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '19', '特码A', '19', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '2', '特码A', '2', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '20', '特码A', '20', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '21', '特码A', '21', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '22', '特码A', '22', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '23', '特码A', '23', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '24', '特码A', '24', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '25', '特码A', '25', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '26', '特码A', '26', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '27', '特码A', '27', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '28', '特码A', '28', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '29', '特码A', '29', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '3', '特码A', '3', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '30', '特码A', '30', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '31', '特码A', '31', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '32', '特码A', '32', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '33', '特码A', '33', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '34', '特码A', '34', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '35', '特码A', '35', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '36', '特码A', '36', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '37', '特码A', '37', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '38', '特码A', '38', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '39', '特码A', '39', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '4', '特码A', '4', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '40', '特码A', '40', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '41', '特码A', '41', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '42', '特码A', '42', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '43', '特码A', '43', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '44', '特码A', '44', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '45', '特码A', '45', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '46', '特码A', '46', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '47', '特码A', '47', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '48', '特码A', '48', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '49', '特码A', '49', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '5', '特码A', '5', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '6', '特码A', '6', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '7', '特码A', '7', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '8', '特码A', '8', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18001', '9', '特码A', '9', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '1', '特码B', '1', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '10', '特码B', '10', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '11', '特码B', '11', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '12', '特码B', '12', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '13', '特码B', '13', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '14', '特码B', '14', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '15', '特码B', '15', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '16', '特码B', '16', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '17', '特码B', '17', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '18', '特码B', '18', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '19', '特码B', '19', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '2', '特码B', '2', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '20', '特码B', '20', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '21', '特码B', '21', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '22', '特码B', '22', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '23', '特码B', '23', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '24', '特码B', '24', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '25', '特码B', '25', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '26', '特码B', '26', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '27', '特码B', '27', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '28', '特码B', '28', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '29', '特码B', '29', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '3', '特码B', '3', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '30', '特码B', '30', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '31', '特码B', '31', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '32', '特码B', '32', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '33', '特码B', '33', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '34', '特码B', '34', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '35', '特码B', '35', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '36', '特码B', '36', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '37', '特码B', '37', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '38', '特码B', '38', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '39', '特码B', '39', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '4', '特码B', '4', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '40', '特码B', '40', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '41', '特码B', '41', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '42', '特码B', '42', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '43', '特码B', '43', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '44', '特码B', '44', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '45', '特码B', '45', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '46', '特码B', '46', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '47', '特码B', '47', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '48', '特码B', '48', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '49', '特码B', '49', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '5', '特码B', '5', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '6', '特码B', '6', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '7', '特码B', '7', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '8', '特码B', '8', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18002', '9', '特码B', '9', 44.10000, 46.550, 1.000, 49.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '单', '双面', '单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '双', '双面', '双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '合单', '双面', '合单', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '合双', '双面', '合双', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '合大', '双面', '合大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '合小', '双面', '合小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '大', '双面', '大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '大单', '双面', '大单', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '大双', '双面', '大双', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '小', '双面', '小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '小单', '双面', '小单', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '小双', '双面', '小双', 3.60000, 3.800, 1.000, 4.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '尾大', '双面', '尾大', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
INSERT INTO `lotto_odds_template` VALUES (10, '18003', '尾小', '双面', '尾小', 1.80000, 1.950, 1.000, 2.000, 1.000, 50000000.000, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for lotto_result
-- ----------------------------
DROP TABLE IF EXISTS `lotto_result`;
CREATE TABLE `lotto_result` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `lotto_id` int(11) NOT NULL DEFAULT '0',
  `issue` varchar(32) NOT NULL DEFAULT '',
  `draw_number` varchar(128) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:未开奖, 1:已开奖, 2:停止购买',
  `start_time` varchar(14) NOT NULL DEFAULT '' COMMENT '开售时间',
  `close_time` varchar(14) NOT NULL DEFAULT '' COMMENT '封单时间',
  `end_time` varchar(14) NOT NULL DEFAULT '' COMMENT '开奖时间',
  `draw_time` varchar(14) NOT NULL DEFAULT '' COMMENT '实际开奖时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `lotto_id` (`lotto_id`,`issue`),
  KEY `lotto_id_2` (`lotto_id`,`issue`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for record_log_user_action
-- ----------------------------
DROP TABLE IF EXISTS `record_log_user_action`;
CREATE TABLE `record_log_user_action` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '操作编号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户编号',
  `action_module` smallint(8) NOT NULL DEFAULT '0' COMMENT '操作模块',
  `action_id` smallint(8) NOT NULL DEFAULT '0',
  `content` text COLLATE utf8_unicode_ci NOT NULL COMMENT '操作内容',
  `ip` varchar(16) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作IP',
  `record_at` varchar(14) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作时间',
  `success` tinyint(1) NOT NULL DEFAULT '0' COMMENT '操作结果',
  `message` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作消息',
  `operator_id` int(11) NOT NULL DEFAULT '0' COMMENT '操作人',
  `operator_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '操作用户类型,对应后台用户类型(0:用户自身,1:非用户自身)',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`,`action_module`,`action_id`,`record_at`),
  KEY `user_id_2` (`user_id`,`record_at`),
  KEY `action_module` (`action_module`,`action_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=COMPACT COMMENT='操作日志';

-- ----------------------------
-- Table structure for record_lotto_order
-- ----------------------------
DROP TABLE IF EXISTS `record_lotto_order`;
CREATE TABLE `record_lotto_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_no` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户编号',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `game_kind` int(11) NOT NULL DEFAULT '0',
  `game_type` int(11) NOT NULL DEFAULT '0',
  `room_id` int(11) NOT NULL DEFAULT '0' COMMENT '房间编号',
  `lotto_id` int(11) NOT NULL DEFAULT '0',
  `issue` varchar(32) NOT NULL DEFAULT '',
  `method_code` varchar(16) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT '0' COMMENT '下注注数',
  `content` varchar(128) NOT NULL DEFAULT '' COMMENT '下注内容',
  `win_content` varchar(128) NOT NULL DEFAULT '',
  `win_count` int(11) NOT NULL DEFAULT '0',
  `draw_number` varchar(64) NOT NULL DEFAULT '' COMMENT '开奖号码',
  `odds` decimal(11,3) NOT NULL DEFAULT '0.000' COMMENT '赔率',
  `amount` decimal(12,3) NOT NULL DEFAULT '0.000' COMMENT '下注金额',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '订单状态 0:未结算,1:中奖,2:未中奖,3:撤销,4:异常',
  `payout` decimal(12,3) NOT NULL DEFAULT '0.000' COMMENT '派奖',
  `profit` decimal(12,3) NOT NULL DEFAULT '0.000' COMMENT '盈亏',
  `flag` tinyint(1) NOT NULL DEFAULT '1' COMMENT '订单标识 0:无效投注,1:有效投注(当Status 是1/2)',
  `bet_date` varchar(8) NOT NULL DEFAULT '' COMMENT '下注日期',
  `calc_date` varchar(8) NOT NULL DEFAULT '' COMMENT '结算日期',
  `bet_time` varchar(14) NOT NULL DEFAULT '' COMMENT '下注时间',
  `update_time` varchar(14) NOT NULL DEFAULT '' COMMENT '结算时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for record_money_change
-- ----------------------------
DROP TABLE IF EXISTS `record_money_change`;
CREATE TABLE `record_money_change` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `link` varchar(64) NOT NULL DEFAULT '' COMMENT '关联编号',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户编号',
  `username` varchar(36) NOT NULL DEFAULT '',
  `game_kind` smallint(8) NOT NULL DEFAULT '0' COMMENT '游戏类型',
  `game_type` smallint(8) NOT NULL DEFAULT '0' COMMENT '游戏类型',
  `change_kind` smallint(8) NOT NULL DEFAULT '0' COMMENT '变化类型',
  `change_type` smallint(8) NOT NULL DEFAULT '0' COMMENT '变化种类',
  `before_balance` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '变化钱余额',
  `amount` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '金额',
  `after_balance` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '变化后余额',
  `record_date` varchar(8) NOT NULL DEFAULT '' COMMENT '纪录日期',
  `record_at` varchar(14) NOT NULL DEFAULT '' COMMENT '纪录时间',
  `operate_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '操作类型(0:用户;1:系统)',
  `operator_id` int(11) NOT NULL DEFAULT '0' COMMENT '操作者编号',
  `remark` varchar(128) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `user_id_r` (`user_id`,`record_at`),
  KEY `user_id_g_m` (`user_id`,`game_kind`,`change_kind`),
  KEY `game_id_m_r` (`game_kind`,`change_kind`,`record_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for record_user_login
-- ----------------------------
DROP TABLE IF EXISTS `record_user_login`;
CREATE TABLE `record_user_login` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `name` varchar(36) NOT NULL DEFAULT '',
  `device_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:Other,1:Web, 2:Mobile, 3:Api',
  `ip` varchar(16) NOT NULL DEFAULT '0.0.0.0' COMMENT '登入地址',
  `user_agent` text NOT NULL COMMENT '浏览器详情',
  `url` varchar(64) NOT NULL DEFAULT '' COMMENT '登入地址',
  `record_at` varchar(14) NOT NULL DEFAULT '' COMMENT '记录时间',
  `remark` varchar(128) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`,`record_at`),
  KEY `user_name` (`record_at`),
  KEY `ip` (`ip`,`record_at`),
  KEY `url` (`url`,`record_at`),
  KEY `record_time` (`record_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for sys_settings
-- ----------------------------
DROP TABLE IF EXISTS `sys_settings`;
CREATE TABLE `sys_settings` (
  `sys_key` varchar(64) NOT NULL,
  `sys_value` text,
  PRIMARY KEY (`sys_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_settings
-- ----------------------------
BEGIN;
INSERT INTO `sys_settings` VALUES ('default_parent_id', '1');
COMMIT;

-- ----------------------------
-- Table structure for user_power
-- ----------------------------
DROP TABLE IF EXISTS `user_power`;
CREATE TABLE `user_power` (
  `user_id` int(11) NOT NULL DEFAULT '1' COMMENT '0:禁用,1:启用',
  `power_bet` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0:禁用,1:启用',
  `power_login` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0:禁用,1:启用',
  `power_deposit` tinyint(1) NOT NULL,
  `power_withdraw` tinyint(1) NOT NULL DEFAULT '1' COMMENT '0:禁用,1:启用',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_power
-- ----------------------------
BEGIN;
INSERT INTO `user_power` VALUES (1, 0, 0, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for user_profile
-- ----------------------------
DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `user_profile` (
  `user_id` int(11) NOT NULL,
  `real_name` varchar(64) NOT NULL DEFAULT '',
  `nickname` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(64) NOT NULL DEFAULT '',
  `is_email_verified` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:未验证，1:已验证',
  `mobile` varchar(64) NOT NULL DEFAULT '',
  `is_mobile_verified` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0:未验证，1:已验证',
  `qq` varchar(64) NOT NULL DEFAULT '',
  `wechat` varchar(64) NOT NULL DEFAULT '',
  `register_ip` varchar(64) NOT NULL DEFAULT '',
  `registered` varchar(14) NOT NULL DEFAULT '',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user_profile
-- ----------------------------
BEGIN;
INSERT INTO `user_profile` VALUES (1, '', '', '', 0, '', 0, '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for user_relation
-- ----------------------------
DROP TABLE IF EXISTS `user_relation`;
CREATE TABLE `user_relation` (
  `user_id` int(11) NOT NULL,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `parents` text NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user_relation
-- ----------------------------
BEGIN;
INSERT INTO `user_relation` VALUES (1, 0, '1');
COMMIT;

-- ----------------------------
-- Table structure for user_score
-- ----------------------------
DROP TABLE IF EXISTS `user_score`;
CREATE TABLE `user_score` (
  `user_id` int(11) NOT NULL,
  `game_score` decimal(20,3) NOT NULL DEFAULT '0.000' COMMENT '余额(单位分)',
  `audit_score` decimal(20,3) NOT NULL DEFAULT '0.000',
  `up_score` decimal(20,3) NOT NULL DEFAULT '0.000',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user_score
-- ----------------------------
BEGIN;
INSERT INTO `user_score` VALUES (1, 0.000, 0.000, 0.000);
COMMIT;

-- ----------------------------
-- Table structure for user_wallet
-- ----------------------------
DROP TABLE IF EXISTS `user_wallet`;
CREATE TABLE `user_wallet` (
  `user_id` int(11) NOT NULL,
  `balance` decimal(12,3) NOT NULL DEFAULT '0.000',
  `password` varchar(128) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:未设置初始化密码，1:正常，2:禁用， 3:需要重置密码',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user_wallet
-- ----------------------------
BEGIN;
INSERT INTO `user_wallet` VALUES (1, 3.000, '', 0);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `name` varchar(32) NOT NULL COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态(0:冻结,1:正常, 2:开户成功但部分功能未开通)',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'admin', '', 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
