/*
 Navicat Premium Data Transfer

 Source Server         : win10 mysql
 Source Server Type    : MySQL
 Source Server Version : 50540
 Source Host           : localhost:3306
 Source Schema         : ginblog

 Target Server Type    : MySQL
 Target Server Version : 50540
 File Encoding         : 65001

 Date: 06/08/2022 13:09:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `titile` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `cid` int(11) NOT NULL,
  `desc` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `img` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tittle` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `comment_count` bigint(20) NOT NULL DEFAULT 0,
  `read_count` bigint(20) NOT NULL DEFAULT 0,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_article_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (2, '2022-08-03 22:32:28', '2022-08-03 22:32:28', '', 5, '12213123123', '<p>121231231231231</p>', 'http://rfd843vfl.hd-bkt.clouddn.com/Fgw0ey2B826nxRfZXlJnHfSiqMaA', '12', NULL, 0, 0, '');
INSERT INTO `article` VALUES (3, '2022-08-03 22:27:33', '2022-08-03 22:27:33', '', 4, '123132', '<p>12123123</p>', '12', '12', NULL, 0, 0, '');
INSERT INTO `article` VALUES (4, '2022-08-03 22:44:56', '2022-08-03 22:44:56', '', 2, 'wakk1', '<p>wakk1</p>', '', 'wakk1', NULL, 0, 0, '');
INSERT INTO `article` VALUES (5, '2022-08-03 22:45:07', '2022-08-03 22:45:07', '', 2, 'wakk2', '<p>wakk2</p>', '', 'wakk2', NULL, 0, 0, '');
INSERT INTO `article` VALUES (6, '2022-08-04 14:34:44', '2022-08-04 14:34:44', '', 4, 'test', '<div><br /><br />\n<div>\n<h2 data-first-child=\"\">用一句话来总结，就是，一个可以扣掉大招和第二天赋，且只能对单的，如此特化的角色，对单普攻DPS甚至也只能算比较优秀，算不上顶尖</h2>\n<p data-pid=\"cw0SkaLw\">。。。。。。。。。。。。。。。。。。。</p>\n<h2><strong>站在2.0版本的玩家角度（不排除未来可期），个人评价宵宫的设计比较失败。</strong></h2>\n<p data-pid=\"2ryV13Vw\">1.大招只能队友触发</p>\n<p data-pid=\"j9TXo18T\">2.第二天赋只能给队友增伤</p>\n<p data-pid=\"_0dSNa3r\">3.附魔8秒真空期</p>\n<h2><strong>从这里可以看出，官方设定宵宫是副c，但是。。。。。</strong></h2>\n<p data-pid=\"gMP8OIkF\">1.追忆套-15能量，大招要75能量才能开</p>\n<p data-pid=\"QlRZQOK5\">2.大招只标记一个人</p>\n<p data-pid=\"do0evh_D\">3.大招2s触发一次，且只持续10s，转移时不刷新时间</p>\n<p data-pid=\"v6uMZNPq\">4.倍率相当低230，被香菱大招全方位碾压</p>\n<p data-pid=\"IHQgZrR-\">5.给队友的buff仅有17%攻击力，比较鸡肋，还不如胡桃雪梅香，开e就能触发</p>\n<h2><strong>这是一个相当无力的大招（上一个这样食之无味，弃之不可惜的大招还是老版钟离），再加上专武和圣遗物套装属性来看，官方又并不推荐玩家释放大招。</strong></h2>\n<p data-pid=\"C14rHYrC\">从这里就能看出<a class=\"css-pgtd2j\" href=\"https://www.zhihu.com/search?q=%E5%AE%B5%E5%AE%AB&amp;search_source=Entity&amp;hybrid_search_source=Entity&amp;hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A2053907455%7D\" target=\"_blank\" rel=\"noopener\" data-za-not-track-link=\"true\">宵宫</a>设定的矛盾了，也是我认为失败的地方。</p>\n<p data-pid=\"bq04tzbZ\">1.不同于胡桃公子，虽然附魔有较长空档期，但是大招都是较高的伤害技</p>\n<p data-pid=\"kTmznbJh\">2.又不同于莫娜钟离，自己有一定伤害的同时，能给予队友较大的加成</p>\n<p data-pid=\"zeCdUyOT\">3.还不同于奥兹锅巴等一众副c，有非常优秀的脱手技能</p>\n<p data-pid=\"t9FMqLA-\">所以先不论宵宫的普攻伤害如何，<a class=\"css-pgtd2j\" href=\"https://www.zhihu.com/search?q=%E9%9C%84%E5%AE%AB&amp;search_source=Entity&amp;hybrid_search_source=Entity&amp;hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A2053907455%7D\" target=\"_blank\" rel=\"noopener\" data-za-not-track-link=\"true\">霄宫</a>就是这样一个角色:</p>\n<p data-pid=\"26Zt_RXA\">1.有胡桃公子一样的<a class=\"css-pgtd2j\" href=\"https://www.zhihu.com/search?q=%E7%9C%9F%E7%A9%BA%E6%9C%9F&amp;search_source=Entity&amp;hybrid_search_source=Entity&amp;hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A2053907455%7D\" target=\"_blank\" rel=\"noopener\" data-za-not-track-link=\"true\">真空期</a>，却没有同级别的爆发</p>\n<p data-pid=\"9lUz68pm\">2.有后台队友加成属性，但是必须大招触发，75能量条件比较苛刻，且加成效果也不理想</p>\n<p data-pid=\"ubKLljN-\">3.有脱手技能，但伤害完全不够看，挂火也不理想</p>\n<h2><strong>属于是缝合角色了，不敢给太强，当然绷带占模另说（）</strong></h2>\n<h2><strong>改进建议：大招标记持续25s，能量改为50，最多标记2名敌人，共享2s触发cd</strong></h2>\n<p data-pid=\"Tpw8GAP6\">。。。。。。。。。。。。。。。。。。。</p>\n<p data-pid=\"ce49iXGA\">再说非常关注的普通攻击：</p>\n</div>\n</div>', 'http://rfd843vfl.hd-bkt.clouddn.com/FhccXA9oZhT1RHHvh86yihmujbKb', 'test', NULL, 0, 0, '');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (2, '2022-07-20 22:25:27', '2022-07-20 22:25:27', 'wakk');
INSERT INTO `category` VALUES (4, '2022-07-30 23:03:20', '2022-07-30 23:03:20', 'java');
INSERT INTO `category` VALUES (5, '2022-07-30 22:30:26', '2022-08-03 22:32:28', '123');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `article_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `title` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `username` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `content` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` tinyint(4) NULL DEFAULT 2,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_comment_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, '2022-08-06 12:16:35', '2022-08-06 12:17:21', NULL, 0, 2, '', 'undefined', '123123', 1);
INSERT INTO `comment` VALUES (2, '2022-08-06 12:17:06', '2022-08-06 12:17:20', NULL, 0, 2, '', 'undefined', '12312312', 1);

-- ----------------------------
-- Table structure for profile
-- ----------------------------
DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `desc` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `qqchat` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `wechat` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `weibo` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `bili` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `img` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `avatar` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icp_record` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of profile
-- ----------------------------
INSERT INTO `profile` VALUES (1, 'WAKK', 'WAKK', 'WAKK', 'WAKK', 'WAKK', 'WAKK', 'WAKK', 'http://rfd843vfl.hd-bkt.clouddn.com/FmxjQjDUmYPlOwwFmvzEHoAixx6M', 'http://rfd843vfl.hd-bkt.clouddn.com/FkaOvIzcJ9--_A7WSeNrrmVsvuQj', 'WAKK');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `user_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `pass_word` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `role` bigint(20) NULL DEFAULT 2,
  `deleted_at` datetime NULL DEFAULT NULL,
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (4, '2022-07-30 22:20:07', '2022-07-30 22:20:07', 'wakk', 'nijn2gVHw9qobg==', 1, NULL, '', '');
INSERT INTO `user` VALUES (6, '2022-07-30 18:37:11', '2022-07-30 18:37:11', '1234', 'nijn2gVHw9qobg==', 1, NULL, '', '');
INSERT INTO `user` VALUES (7, '2022-07-26 16:50:25', '2022-07-26 16:50:25', 'wakk1', '6dzz49qDNRZJZQ==', 1, NULL, '', '');
INSERT INTO `user` VALUES (8, '2022-07-30 20:43:04', '2022-07-30 20:43:04', '11111', '6dzz49qDNRZJZQ==', 1, NULL, '', '');
INSERT INTO `user` VALUES (9, '2022-08-03 23:58:49', '2022-08-03 23:58:49', 'user1', '6dzz49qDNRZJZQ==', 1, NULL, '', '');
INSERT INTO `user` VALUES (10, '2022-07-30 20:44:32', '2022-07-30 20:44:32', 'user2', '6dzz49qDNRZJZQ==', 1, NULL, '', '');
INSERT INTO `user` VALUES (11, '2022-07-30 21:20:45', '2022-07-30 21:20:45', 'user3', '6dzz49qDNRZJZQ==', 1, NULL, '', '');
INSERT INTO `user` VALUES (12, '2022-07-30 21:21:06', '2022-07-30 21:21:06', 'user4', '6dzz49qDNRZJZQ==', 1, NULL, '', '');
INSERT INTO `user` VALUES (13, '2022-07-30 21:41:58', '2022-07-30 21:41:58', 'user5', '6dzz49qDNRZJZQ==', 2, NULL, '', '');
INSERT INTO `user` VALUES (14, '2022-07-30 21:51:33', '2022-07-30 21:51:33', 'user6', '6dzz49qDNRZJZQ==', 1, NULL, '', '');
INSERT INTO `user` VALUES (15, '2022-08-04 15:00:05', '2022-08-04 15:00:05', 'user_test', '6dzz49qDNRZJZQ==', 2, NULL, '', '');

SET FOREIGN_KEY_CHECKS = 1;
