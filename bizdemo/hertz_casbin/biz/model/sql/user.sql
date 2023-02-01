/*
 Navicat Premium Data Transfer

 Source Server         : 9911
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:9911
 Source Schema         : gorm

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 01/02/2023 10:58:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `ID` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `role` varchar(30) DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`ID`, `username`, `password`, `email`, `role`, `deleted_at`, `updated_at`, `created_at`) VALUES (1, 'admin', '202cb962ac59075b964b07152d234b70', 'darren94me@gmail.com', 'admin', NULL, '2023-01-31 15:46:06', '2023-01-31 15:46:06');
INSERT INTO `users` (`ID`, `username`, `password`, `email`, `role`, `deleted_at`, `updated_at`, `created_at`) VALUES (2, 'darren', '202cb962ac59075b964b07152d234b70', 'darren@11.com', 'member', NULL, '2023-01-31 17:39:39', '2023-01-31 17:39:39');
INSERT INTO `users` (`ID`, `username`, `password`, `email`, `role`, `deleted_at`, `updated_at`, `created_at`) VALUES (3, 'zhangsan', '202cb962ac59075b964b07152d234b70', 'zhangsan@163.com', 'member', NULL, '2023-01-31 18:05:39', '2023-01-31 18:05:39');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
