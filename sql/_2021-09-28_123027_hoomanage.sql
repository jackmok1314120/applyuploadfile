/*!40101 SET NAMES utf8 */;
/*!40014 SET FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/ hoomanage /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE hoomanage;

DROP TABLE IF EXISTS apply_info;
CREATE TABLE `apply_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '申请人名字',
  `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `c_id` int(11) DEFAULT NULL COMMENT '公司Id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS apply_pending;
CREATE TABLE `apply_pending` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `name` varchar(255) DEFAULT NULL COMMENT '用户名称',
  `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `introduce` text COMMENT '公司介绍',
  `coin_name` varchar(255) DEFAULT NULL COMMENT '币种名称',
  `id_card_picture` varchar(255) DEFAULT NULL COMMENT '身份证件',
  `business_picture` varchar(255) DEFAULT NULL COMMENT '营业执照证',
  `pass` int(11) DEFAULT NULL COMMENT '是否通过',
  `create_time` datetime DEFAULT NULL COMMENT 'create time',
  `update_time` datetime DEFAULT NULL COMMENT 'update time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS company;
CREATE TABLE `company` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `coin_name` varchar(255) DEFAULT NULL COMMENT '币种名称',
  `introduce` text COMMENT '介绍',
  `id_card_picture` varchar(255) DEFAULT NULL COMMENT '身份证复印件',
  `business_picture` varchar(255) DEFAULT NULL COMMENT '营业执照',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;