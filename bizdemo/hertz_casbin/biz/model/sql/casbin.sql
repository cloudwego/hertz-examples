/*
 Navicat Premium Data Transfer

 Source Server         : 9912
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:9912
 Source Schema         : casbin

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 03/02/2023 12:19:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  KEY `IDX_casbin_rule_v5` (`v5`),
  KEY `IDX_casbin_rule_p_type` (`p_type`),
  KEY `IDX_casbin_rule_v0` (`v0`),
  KEY `IDX_casbin_rule_v1` (`v1`),
  KEY `IDX_casbin_rule_v2` (`v2`),
  KEY `IDX_casbin_rule_v3` (`v3`),
  KEY `IDX_casbin_rule_v4` (`v4`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/*', '*', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'role', '/v1/role/create/', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'permission', '/v1/permission/create/', 'POST', NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for permission_roles
-- ----------------------------
DROP TABLE IF EXISTS `permission_roles`;
CREATE TABLE `permission_roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rid` bigint DEFAULT NULL,
  `pid` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of permission_roles
-- ----------------------------
BEGIN;
INSERT INTO `permission_roles` (`id`, `rid`, `pid`) VALUES (6, 1, 1);
INSERT INTO `permission_roles` (`id`, `rid`, `pid`) VALUES (7, 2, 2);
INSERT INTO `permission_roles` (`id`, `rid`, `pid`) VALUES (8, 3, 3);
COMMIT;

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `v1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `v2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of permissions
-- ----------------------------
BEGIN;
INSERT INTO `permissions` (`id`, `v1`, `v2`) VALUES (1, '/*', '*');
INSERT INTO `permissions` (`id`, `v1`, `v2`) VALUES (2, '/v1/role/create/', 'POST');
INSERT INTO `permissions` (`id`, `v1`, `v2`) VALUES (3, '/v1/permission/create/', 'POST');
COMMIT;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` (`id`, `name`) VALUES (1, 'admin');
INSERT INTO `roles` (`id`, `name`) VALUES (2, 'role');
INSERT INTO `roles` (`id`, `name`) VALUES (3, 'permission');
INSERT INTO `roles` (`id`, `name`) VALUES (7, 'admin123');
COMMIT;

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rid` bigint DEFAULT NULL,
  `uid` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
BEGIN;
INSERT INTO `user_roles` (`id`, `rid`, `uid`) VALUES (1, 1, 5);
INSERT INTO `user_roles` (`id`, `rid`, `uid`) VALUES (5, 2, 6);
INSERT INTO `user_roles` (`id`, `rid`, `uid`) VALUES (6, 3, 7);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `username`, `password`) VALUES (5, 'admin', '202cb962ac59075b964b07152d234b70');
INSERT INTO `users` (`id`, `username`, `password`) VALUES (6, 'role_user', '202cb962ac59075b964b07152d234b70');
INSERT INTO `users` (`id`, `username`, `password`) VALUES (7, 'permission_user', '202cb962ac59075b964b07152d234b70');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
