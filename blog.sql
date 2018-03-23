/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2018-03-23 08:47:06
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for qq_article
-- ----------------------------
DROP TABLE IF EXISTS `qq_article`;
CREATE TABLE `qq_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tags` varchar(100) DEFAULT NULL,
  `title` varchar(50) DEFAULT NULL,
  `summary` varchar(100) DEFAULT NULL,
  `thumb` varchar(100) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '1',
  `p_sort` smallint(5) DEFAULT NULL,
  `admin_id` int(11) unsigned DEFAULT NULL,
  `author_name` varchar(50) DEFAULT NULL,
  `scan_count` int(11) DEFAULT NULL,
  `comment_count` int(11) DEFAULT NULL,
  `is_recommend` tinyint(1) DEFAULT '1' COMMENT '推荐 1-默认  2-推荐',
  `is_top` tinyint(1) DEFAULT '1' COMMENT '推荐 1-默认  2-置顶',
  `is_hot` tinyint(1) DEFAULT '1' COMMENT '推荐 1-默认  2-热门',
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of qq_article
-- ----------------------------
INSERT INTO `qq_article` VALUES ('1', '1,10,11,2', '测试文章', '简介', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '2', '2', '2', '1520759229', '1520759229');
INSERT INTO `qq_article` VALUES ('2', '2,12,13,15', '测试2', '测试2', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '1', '2', '1', '1520763226', '1520763226');
INSERT INTO `qq_article` VALUES ('3', '2,1', '13', '13', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '3', '0', '1', '超级管理员', '0', '0', '1', '2', '1', '1520764344', '1520764344');
INSERT INTO `qq_article` VALUES ('4', '1,12', '42', '234', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '1', '2', '1', '1520764677', '1520764677');
INSERT INTO `qq_article` VALUES ('5', '1', '13', '132', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '1', '2', '1', '1520764710', '1520764710');
INSERT INTO `qq_article` VALUES ('6', '1,12', '132', '312', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '1', '2', '1', '1520764745', '1520764745');
INSERT INTO `qq_article` VALUES ('7', '1,12', '123', '13', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '2', '2', '2', '1520765341', '1520765341');
INSERT INTO `qq_article` VALUES ('8', '13,14', '123', '123', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '2', '2', '2', '1520825247', '1520825247');
INSERT INTO `qq_article` VALUES ('9', '13,14', '123', '123', 'http://localhost:7474/./upload/img/2018-3/12/15208252436960634006129484611666145821_l.jpg', '0', '0', '1', '超级管理员', '0', '0', '2', '2', '1', '1520825584', '1520825584');
INSERT INTO `qq_article` VALUES ('10', '1,13,15,2', '测试', '测试', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '1', '2', '1', '1520826533', '1520826533');
INSERT INTO `qq_article` VALUES ('11', '14,2', '123', '132', 'http://localhost:7474/./static/upload/img/2018-3/12/15208266371820137005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '1', '2', '1', '1520826638', '1520826638');
INSERT INTO `qq_article` VALUES ('12', '1,10,11,2', '123', '132', 'http://localhost:7474/./upload/img/2018-3/12/15208265324742175005577006791947779410_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '0', '0', '0', '1520909145', '1520909145');
INSERT INTO `qq_article` VALUES ('13', '1,10,11,2', '123', '132', 'http://localhost:7474/./static/upload/img/2018-3/13/15209185419624387008674665223082153551_l.png', '1', '0', '1', '超级管理员', '0', '0', '0', '0', '0', '1520918543', '1520918543');
INSERT INTO `qq_article` VALUES ('14', '1,10,11,2', '123', '132', 'http://localhost:7474/./static/upload/img/2018-3/13/15209185588945888006129484611666145821_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '0', '0', '0', '1520918559', '1520918559');
INSERT INTO `qq_article` VALUES ('15', '', '123', '132', 'http://localhost:7474/./static/upload/img/2018-3/13/15209185588945888006129484611666145821_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '0', '0', '0', '1521468513', '1521468513');
INSERT INTO `qq_article` VALUES ('16', '', '123', '132', 'http://localhost:7474/./static/upload/img/2018-3/13/15209185588945888006129484611666145821_l.jpg', '1', '0', '1', '超级管理员', '0', '0', '0', '0', '0', '1521469042', '1521469042');

-- ----------------------------
-- Table structure for qq_article_content
-- ----------------------------
DROP TABLE IF EXISTS `qq_article_content`;
CREATE TABLE `qq_article_content` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) unsigned DEFAULT NULL,
  `content` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of qq_article_content
-- ----------------------------
INSERT INTO `qq_article_content` VALUES ('1', '1', '测试');
INSERT INTO `qq_article_content` VALUES ('2', '2', '测试2');
INSERT INTO `qq_article_content` VALUES ('3', '3', '13');
INSERT INTO `qq_article_content` VALUES ('4', '4', '324');
INSERT INTO `qq_article_content` VALUES ('5', '5', '13');
INSERT INTO `qq_article_content` VALUES ('6', '6', '13');
INSERT INTO `qq_article_content` VALUES ('7', '7', '13');
INSERT INTO `qq_article_content` VALUES ('8', '8', '123');
INSERT INTO `qq_article_content` VALUES ('9', '9', '123');
INSERT INTO `qq_article_content` VALUES ('10', '10', '安德森');
INSERT INTO `qq_article_content` VALUES ('11', '11', '213');
INSERT INTO `qq_article_content` VALUES ('12', '12', '#### 请求头：\n\n|参数名|是否必须|类型|说明|\n|:----    |:---|:----- |-----   |\n|Content-Type |是  |string |请求类型： application/json   |\n|Content-MD5 |是  |string | 请求内容签名    |\n\n\n#### 请求参数:\n\n|参数名|是否必须|类型|说明|示例值\n|:----    |:---|:----- |-----   |-----   |\n|username |是  |string |用户名   | george518\n|password |是  |string | 密码    | george518\n\n#### 返回参数:\n\n|参数名|类型|说明|\n|:-----  |:-----|-----                           |\n|group_level |int   |用户组id，1：超级管理员；2：普通用户  |\n\n#### 返回示例:\n\n**正确时返回:**\n\n```\n{\n\"status\": 200,\n\"message\": \"Success\",\n\"data\": {\n    \"uid\": \"1\",\n    \"account\": \"admin\",\n    \"nickname\": \"Minho\",\n    \"group_level\": 0 ,\n    \"create_time\": \"1436864169\",\n    \"last_login_time\": \"0\",\n}\n}\n```\n\n**错误时返回:**\n\n\n```\n{\n\"status\": 300,\n\"message\": \"invalid username\"\n\"data\":{\n    \n}\n}\n```\n\n#### 调用示例:\n\n```\n\n<?php\n    echo \"Hello world\";\n```\n#### 备注:\n\n- 更多返回错误代码请看首页的错误代码描述');
INSERT INTO `qq_article_content` VALUES ('13', '13', '#### 请求头：\n\n|参数名|是否必须|类型|说明|\n|:----    |:---|:----- |-----   |\n|Content-Type |是  |string |请求类型： application/json   |\n|Content-MD5 |是  |string | 请求内容签名    |\n\n\n#### 请求参数:\n\n|参数名|是否必须|类型|说明|示例值\n|:----    |:---|:----- |-----   |-----   |\n|username |是  |string |用户名   | george518\n|password |是  |string | 密码    | george518\n\n#### 返回参数:\n\n|参数名|类型|说明|\n|:-----  |:-----|-----                           |\n|group_level |int   |用户组id，1：超级管理员；2：普通用户  |\n\n#### 返回示例:\n\n**正确时返回:**\n\n```\n{\n\"status\": 200,\n\"message\": \"Success\",\n\"data\": {\n    \"uid\": \"1\",\n    \"account\": \"admin\",\n    \"nickname\": \"Minho\",\n    \"group_level\": 0 ,\n    \"create_time\": \"1436864169\",\n    \"last_login_time\": \"0\",\n}\n}\n```\n\n**错误时返回:**\n\n\n```\n{\n\"status\": 300,\n\"message\": \"invalid username\"\n\"data\":{\n    \n}\n}\n```\n\n#### 调用示例:\n\n```\n\n<?php\n    echo \"Hello world\";\n```\n#### 备注:\n\n- 更多返回错误代码请看首页的错误代码描述');
INSERT INTO `qq_article_content` VALUES ('14', '14', '#### 请求头：\n\n|参数名|是否必须|类型|说明|\n|:----    |:---|:----- |-----   |\n|Content-Type |是  |string |请求类型： application/json   |\n|Content-MD5 |是  |string | 请求内容签名    |\n\n\n#### 请求参数:\n\n|参数名|是否必须|类型|说明|示例值\n|:----    |:---|:----- |-----   |-----   |\n|username |是  |string |用户名   | george518\n|password |是  |string | 密码    | george518\n\n#### 返回参数:\n\n|参数名|类型|说明|\n|:-----  |:-----|-----                           |\n|group_level |int   |用户组id，1：超级管理员；2：普通用户  |\n\n#### 返回示例:\n\n**正确时返回:**\n\n```\n{\n\"status\": 200,\n\"message\": \"Success\",\n\"data\": {\n    \"uid\": \"1\",\n    \"account\": \"admin\",\n    \"nickname\": \"Minho\",\n    \"group_level\": 0 ,\n    \"create_time\": \"1436864169\",\n    \"last_login_time\": \"0\",\n}\n}\n```\n\n**错误时返回:**\n\n\n```\n{\n\"status\": 300,\n\"message\": \"invalid username\"\n\"data\":{\n    \n}\n}\n```\n\n#### 调用示例:\n\n```\n\n<?php\n    echo \"Hello world\";\n```\n#### 备注:\n\n- 更多返回错误代码请看首页的错误代码描述');
INSERT INTO `qq_article_content` VALUES ('15', '15', '#### 请求头：\n\n|参数名|是否必须|类型|说明|\n|:----    |:---|:----- |-----   |\n|Content-Type |是  |string |请求类型： application/json   |\n|Content-MD5 |是  |string | 请求内容签名    |\n\n\n#### 请求参数:\n\n|参数名|是否必须|类型|说明|示例值\n|:----    |:---|:----- |-----   |-----   |\n|username |是  |string |用户名   | george518\n|password |是  |string | 密码    | george518\n\n#### 返回参数:\n\n|参数名|类型|说明|\n|:-----  |:-----|-----                           |\n|group_level |int   |用户组id，1：超级管理员；2：普通用户  |\n\n#### 返回示例:\n\n**正确时返回:**\n\n```\n{\n\"status\": 200,\n\"message\": \"Success\",\n\"data\": {\n    \"uid\": \"1\",\n    \"account\": \"admin\",\n    \"nickname\": \"Minho\",\n    \"group_level\": 0 ,\n    \"create_time\": \"1436864169\",\n    \"last_login_time\": \"0\",\n}\n}\n```\n\n**错误时返回:**\n\n\n```\n{\n\"status\": 300,\n\"message\": \"invalid username\"\n\"data\":{\n    \n}\n}\n```\n\n#### 调用示例:\n\n```\n\n<?php\n    echo \"Hello world\";\n```\n#### 备注:\n\n- 更多返回错误代码请看首页的错误代码描述');
INSERT INTO `qq_article_content` VALUES ('16', '16', '#### 请求头：\n\n|参数名|是否必须|类型|说明|\n|:----    |:---|:----- |-----   |\n|Content-Type |是  |string |请求类型： application/json   |\n|Content-MD5 |是  |string | 请求内容签名    |\n\n\n#### 请求参数:\n\n|参数名|是否必须|类型|说明|示例值\n|:----    |:---|:----- |-----   |-----   |\n|username |是  |string |用户名   | george518\n|password |是  |string | 密码    | george518\n\n#### 返回参数:\n\n|参数名|类型|说明|\n|:-----  |:-----|-----                           |\n|group_level |int   |用户组id，1：超级管理员；2：普通用户  |\n\n#### 返回示例:\n\n**正确时返回:**\n\n```\n{\n\"status\": 200,\n\"message\": \"Success\",\n\"data\": {\n    \"uid\": \"1\",\n    \"account\": \"admin\",\n    \"nickname\": \"Minho\",\n    \"group_level\": 0 ,\n    \"create_time\": \"1436864169\",\n    \"last_login_time\": \"0\",\n}\n}\n```\n\n**错误时返回:**\n\n\n```\n{\n\"status\": 300,\n\"message\": \"invalid username\"\n\"data\":{\n    \n}\n}\n```\n\n#### 调用示例:\n\n```\n\n<?php\n    echo \"Hello world\";\n```\n#### 备注:\n\n- 更多返回错误代码请看首页的错误代码描述');

-- ----------------------------
-- Table structure for qq_comment
-- ----------------------------
DROP TABLE IF EXISTS `qq_comment`;
CREATE TABLE `qq_comment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) unsigned DEFAULT NULL,
  `user_id` int(11) unsigned DEFAULT NULL,
  `reply_to` int(11) unsigned DEFAULT '0',
  `nickname` varchar(50) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `like_count` int(11) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0' COMMENT '点赞',
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of qq_comment
-- ----------------------------
INSERT INTO `qq_comment` VALUES ('1', '0', '0', '0', '', '', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('2', '10', '0', '0', '', '', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('3', '10', '0', '2', '', '1', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('4', '10', '0', '2', '', '2', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('5', '10', '3', '2', '', '3', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('6', '10', '3', '2', '', '4', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('7', '10', '3', '2', '', 'qwe5', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('8', '10', '3', '2', '', 'qwe6', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('9', '10', '3', '2', '', 'qwe7', '0', '0', '0', '0');
INSERT INTO `qq_comment` VALUES ('10', '10', '3', '2', '', 'qwe8', '0', '0', '1521702312', '0');
INSERT INTO `qq_comment` VALUES ('11', '9', '0', '0', '', '123', '0', '0', '1521702900', '0');
INSERT INTO `qq_comment` VALUES ('12', '9', '0', '0', '', '123', '0', '0', '1521702926', '0');
INSERT INTO `qq_comment` VALUES ('13', '9', '0', '0', '', '123', '0', '0', '1521702940', '0');
INSERT INTO `qq_comment` VALUES ('14', '10', '0', '2', '', '123123123213qweqe', '0', '0', '1521707319', '0');
INSERT INTO `qq_comment` VALUES ('15', '2', '0', '0', '', '13', '0', '0', '1521707346', '0');
INSERT INTO `qq_comment` VALUES ('16', '2', '0', '15', '', '123', '0', '0', '1521707511', '0');
INSERT INTO `qq_comment` VALUES ('17', '2', '0', '0', '', '123', '0', '0', '1521707524', '0');
INSERT INTO `qq_comment` VALUES ('18', '16', '0', '0', '', '123', '0', '0', '1521708717', '0');
INSERT INTO `qq_comment` VALUES ('19', '16', '0', '18', '', '123', '0', '0', '1521708722', '0');
INSERT INTO `qq_comment` VALUES ('20', '16', '0', '18', '', '123', '0', '0', '1521708724', '0');
INSERT INTO `qq_comment` VALUES ('21', '16', '0', '0', '', '1', '0', '0', '1521708971', '0');
INSERT INTO `qq_comment` VALUES ('22', '16', '0', '0', '', '123', '0', '0', '1521708975', '0');
INSERT INTO `qq_comment` VALUES ('23', '16', '0', '0', '', '213123', '0', '0', '1521708978', '0');
INSERT INTO `qq_comment` VALUES ('24', '16', '0', '0', '', '123123', '0', '0', '1521708982', '0');

-- ----------------------------
-- Table structure for qq_tag
-- ----------------------------
DROP TABLE IF EXISTS `qq_tag`;
CREATE TABLE `qq_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of qq_tag
-- ----------------------------
INSERT INTO `qq_tag` VALUES ('1', 'php');
INSERT INTO `qq_tag` VALUES ('2', 'go');
INSERT INTO `qq_tag` VALUES ('3', 'mysql');
INSERT INTO `qq_tag` VALUES ('4', 'linux');
INSERT INTO `qq_tag` VALUES ('5', 'mac');
INSERT INTO `qq_tag` VALUES ('6', '心情');
INSERT INTO `qq_tag` VALUES ('7', '其他');
INSERT INTO `qq_tag` VALUES ('8', '测试');
INSERT INTO `qq_tag` VALUES ('9', '攻防');
INSERT INTO `qq_tag` VALUES ('10', '黑客');
INSERT INTO `qq_tag` VALUES ('11', '服务器');
INSERT INTO `qq_tag` VALUES ('12', 'web');
INSERT INTO `qq_tag` VALUES ('13', '前端');
INSERT INTO `qq_tag` VALUES ('14', 'js');
INSERT INTO `qq_tag` VALUES ('15', 'jquery');

-- ----------------------------
-- Table structure for qq_uc_admin
-- ----------------------------
DROP TABLE IF EXISTS `qq_uc_admin`;
CREATE TABLE `qq_uc_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `phone` varchar(20) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Records of qq_uc_admin
-- ----------------------------
INSERT INTO `qq_uc_admin` VALUES ('1', 'admin', '超级管理员', 'bd730f3fa72d868a843a832e9f4bdead', '0', '13888888889', 'haodaquan2008@163.com', 'AmQX', '1521527041', '[', '1', '0', '1', '0', '1520610229');
INSERT INTO `qq_uc_admin` VALUES ('2', 'george518', 'georgeHao', 'e5d77ebaffd5e4fe7164b31c6d7f1921', '1,2', '13811558899', '12@11.com', 'ONNy', '1506125048', '127.0.0.1', '0', '0', '0', '0', '1515116737');
INSERT INTO `qq_uc_admin` VALUES ('3', 'haodaquan', '郝大全', 'e9fa9187e7497892c237433aee966cc1', '2,1', '13811559988', 'hao@123.com', '6fWE', '1505960085', '127.0.0.1', '1', '1', '0', '1505919245', '1513388240');
INSERT INTO `qq_uc_admin` VALUES ('4', 'ceshizhanghao', '测试姓名', 'fa3fb5825c2e64bc764f29245dd1ec7a', '2', '13988009988', '232@124.com', 'i8Nf', '0', '', '1', '1', '0', '1506047337', '1513388111');

-- ----------------------------
-- Table structure for qq_uc_auth
-- ----------------------------
DROP TABLE IF EXISTS `qq_uc_auth`;
CREATE TABLE `qq_uc_auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) unsigned NOT NULL DEFAULT '999' COMMENT '排序，越小越前',
  `icon` varchar(255) NOT NULL,
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示，0-隐藏，1-显示',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '操作者ID',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COMMENT='权限因子';

-- ----------------------------
-- Records of qq_uc_auth
-- ----------------------------
INSERT INTO `qq_uc_auth` VALUES ('1', '0', '所有权限', '/', '1', '', '0', '1', '1', '1', '1', '1505620970', '1505620970');
INSERT INTO `qq_uc_auth` VALUES ('2', '1', '权限管理', '/', '999', 'fa-id-card', '1', '1', '0', '1', '1', '0', '1505622360');
INSERT INTO `qq_uc_auth` VALUES ('3', '2', '管理员', '/admin/list', '1', 'fa-user-o', '1', '1', '1', '1', '1', '1505621186', '1505621186');
INSERT INTO `qq_uc_auth` VALUES ('4', '2', '角色管理', '/role/list', '2', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1505621852');
INSERT INTO `qq_uc_auth` VALUES ('5', '3', '新增', '/admin/add', '1', '', '0', '1', '0', '1', '1', '0', '1505621685');
INSERT INTO `qq_uc_auth` VALUES ('6', '3', '修改', '/admin/edit', '2', '', '0', '1', '0', '1', '1', '0', '1505621697');
INSERT INTO `qq_uc_auth` VALUES ('7', '3', '删除', '/admin/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621756', '1505621756');
INSERT INTO `qq_uc_auth` VALUES ('8', '4', '新增', '/role/add', '1', '', '1', '1', '0', '1', '1', '0', '1505698716');
INSERT INTO `qq_uc_auth` VALUES ('9', '4', '修改', '/role/edit', '2', '', '0', '1', '1', '1', '1', '1505621912', '1505621912');
INSERT INTO `qq_uc_auth` VALUES ('10', '4', '删除', '/role/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621951', '1505621951');
INSERT INTO `qq_uc_auth` VALUES ('11', '2', '权限因子', '/auth/list', '3', 'fa-list', '1', '1', '1', '1', '1', '1505621986', '1505621986');
INSERT INTO `qq_uc_auth` VALUES ('12', '11', '新增', '/auth/add', '1', '', '0', '1', '1', '1', '1', '1505622009', '1505622009');
INSERT INTO `qq_uc_auth` VALUES ('13', '11', '修改', '/auth/edit', '2', '', '0', '1', '1', '1', '1', '1505622047', '1505622047');
INSERT INTO `qq_uc_auth` VALUES ('14', '11', '删除', '/auth/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505622111', '1505622111');
INSERT INTO `qq_uc_auth` VALUES ('15', '1', '个人中心', 'profile/edit', '1001', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1506001114');
INSERT INTO `qq_uc_auth` VALUES ('16', '1', '文章管理', '/', '1', 'fa-cubes', '1', '1', '0', '1', '1', '0', '1520758866');
INSERT INTO `qq_uc_auth` VALUES ('17', '16', '文章', '/article/list', '1', 'fa-link', '1', '1', '0', '1', '1', '0', '1520758888');
INSERT INTO `qq_uc_auth` VALUES ('19', '20', '公共文档', '/apipublic/list', '3', 'fa-files-o', '1', '1', '0', '1', '0', '0', '1516155852');
INSERT INTO `qq_uc_auth` VALUES ('20', '1', '基础设置', '/', '2', 'fa-cogs', '1', '1', '1', '1', '0', '1505622601', '1505622601');
INSERT INTO `qq_uc_auth` VALUES ('21', '20', '分组设置', '/group/list', '1', 'fa-object-ungroup', '1', '1', '1', '1', '0', '1505622645', '1505622645');
INSERT INTO `qq_uc_auth` VALUES ('22', '20', '环境设置', '/env/list', '2', 'fa-tree', '1', '1', '1', '1', '0', '1505622681', '1505622681');
INSERT INTO `qq_uc_auth` VALUES ('23', '20', '状态码设置', '/code/list', '3', 'fa-code', '1', '1', '1', '1', '0', '1505622728', '1505622728');
INSERT INTO `qq_uc_auth` VALUES ('24', '15', '资料修改', '/user/edit', '1', 'fa-edit', '1', '1', '0', '1', '1', '0', '1506057468');
INSERT INTO `qq_uc_auth` VALUES ('25', '21', '新增', '/group/add', '1', 'n', '0', '1', '0', '1', '0', '0', '1513324170');
INSERT INTO `qq_uc_auth` VALUES ('26', '21', '修改', '/group/edit', '2', 'fa', '0', '0', '0', '0', '0', '1506237920', '1506237920');
INSERT INTO `qq_uc_auth` VALUES ('27', '21', '删除', '/group/ajaxdel', '3', 'fa', '0', '0', '0', '0', '0', '1506237948', '1506237948');
INSERT INTO `qq_uc_auth` VALUES ('28', '22', '新增', '/env/add', '1', 'fa', '0', '0', '0', '0', '0', '1506316506', '1506316506');
INSERT INTO `qq_uc_auth` VALUES ('29', '22', '修改', '/env/edit', '2', 'fa', '0', '0', '0', '0', '0', '1506316532', '1506316532');
INSERT INTO `qq_uc_auth` VALUES ('30', '22', '删除', '/env/ajaxdel', '3', 'fa', '0', '0', '0', '0', '0', '1506316567', '1506316567');
INSERT INTO `qq_uc_auth` VALUES ('31', '23', '新增', '/code/add', '1', 'fa', '0', '0', '0', '0', '0', '1506327812', '1506327812');
INSERT INTO `qq_uc_auth` VALUES ('32', '23', '修改', '/code/edit', '2', 'fa', '0', '0', '0', '0', '0', '1506327831', '1506327831');
INSERT INTO `qq_uc_auth` VALUES ('33', '23', '删除', '/code/ajadel', '3', 'fa', '0', '0', '0', '0', '0', '1506327857', '1506327857');
INSERT INTO `qq_uc_auth` VALUES ('34', '17', '新增文章', '/article/add', '1', 'fa-link', '1', '1', '0', '1', '1', '0', '1520758903');
INSERT INTO `qq_uc_auth` VALUES ('35', '17', '修改文章', '/article/edit', '2', 'fa-link', '1', '1', '0', '1', '1', '0', '1520758914');
INSERT INTO `qq_uc_auth` VALUES ('36', '17', '删除文章', '/article/ajaxdel', '3', 'fa-link', '1', '1', '0', '1', '1', '0', '1520758934');
INSERT INTO `qq_uc_auth` VALUES ('37', '17', '查看文章', '/article/detail', '4', '', '0', '1', '0', '1', '1', '0', '1520758947');
INSERT INTO `qq_uc_auth` VALUES ('38', '17', '批量审核文章', '/article/ajaxchangestatus', '5', '', '0', '1', '0', '1', '1', '0', '1520758959');
INSERT INTO `qq_uc_auth` VALUES ('39', '16', 'API资源', '/apisource/list', '1', 'fa-skyatlas', '1', '1', '0', '1', '0', '0', '1515984925');
INSERT INTO `qq_uc_auth` VALUES ('40', '39', '新增', '/apisource/add', '1', '', '0', '1', '0', '1', '0', '0', '1515905034');
INSERT INTO `qq_uc_auth` VALUES ('41', '39', '修改', '/apisource/edit', '2', '', '0', '1', '0', '1', '0', '0', '1515905044');
INSERT INTO `qq_uc_auth` VALUES ('42', '39', '删除', '/apisource/ajaxdel', '3', '', '0', '1', '0', '1', '0', '0', '1515905055');
INSERT INTO `qq_uc_auth` VALUES ('43', '17', '审核页面', '/api/audit', '6', '', '0', '1', '1', '1', '0', '1516000189', '1516000189');
INSERT INTO `qq_uc_auth` VALUES ('44', '19', '新增', '/apipublic/add', '1', '', '0', '1', '1', '1', '0', '1516067809', '1516067809');
INSERT INTO `qq_uc_auth` VALUES ('45', '19', '修改', '/apipublic/edit', '2', '', '0', '1', '1', '1', '0', '1516067841', '1516067841');
INSERT INTO `qq_uc_auth` VALUES ('46', '19', '删除', '/apipublic/ajaxdel', '3', '', '0', '1', '1', '1', '0', '1516067881', '1516067881');
INSERT INTO `qq_uc_auth` VALUES ('47', '20', '模板设置', '/template/list', '4', 'fa-file-text', '1', '1', '1', '1', '0', '1516083233', '1516083233');
INSERT INTO `qq_uc_auth` VALUES ('48', '47', '新增', '/template/add', '1', '', '0', '1', '1', '1', '0', '1516083262', '1516083262');
INSERT INTO `qq_uc_auth` VALUES ('49', '47', '修改', '/template/edit', '2', '', '0', '1', '1', '1', '0', '1516083296', '1516083296');
INSERT INTO `qq_uc_auth` VALUES ('50', '47', '删除', '/template/ajaxdel', '3', '', '0', '1', '1', '1', '0', '1516083335', '1516083335');
INSERT INTO `qq_uc_auth` VALUES ('51', '22', '新增', '/env/add', '1', 'fa', '0', '1', '1', '1', '0', '1520327210', '1520327210');

-- ----------------------------
-- Table structure for qq_uc_role
-- ----------------------------
DROP TABLE IF EXISTS `qq_uc_role`;
CREATE TABLE `qq_uc_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改这ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Records of qq_uc_role
-- ----------------------------
INSERT INTO `qq_uc_role` VALUES ('1', '文章管理员', '拥有文章所有权限', '0', '1', '1', '1520759152', '1520759152');
INSERT INTO `qq_uc_role` VALUES ('2', '系统管理员', '系统管理员', '0', '1', '1', '1520759118', '1520759118');
INSERT INTO `qq_uc_role` VALUES ('3', '测试', '测试', '0', '1', '0', '1520296554', '1520296564');

-- ----------------------------
-- Table structure for qq_uc_role_auth
-- ----------------------------
DROP TABLE IF EXISTS `qq_uc_role_auth`;
CREATE TABLE `qq_uc_role_auth` (
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `auth_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`role_id`,`auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限和角色关系表';

-- ----------------------------
-- Records of qq_uc_role_auth
-- ----------------------------
INSERT INTO `qq_uc_role_auth` VALUES ('1', '0');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '1');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '15');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '16');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '17');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '24');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '34');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '35');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '36');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '37');
INSERT INTO `qq_uc_role_auth` VALUES ('1', '38');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '1');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '2');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '3');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '4');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '5');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '6');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '7');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '8');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '9');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '10');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '11');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '12');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '13');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '14');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '15');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '16');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '17');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '24');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '34');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '35');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '36');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '37');
INSERT INTO `qq_uc_role_auth` VALUES ('2', '38');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '1');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '16');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '17');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '34');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '35');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '36');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '37');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '39');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '40');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '41');
INSERT INTO `qq_uc_role_auth` VALUES ('3', '42');
