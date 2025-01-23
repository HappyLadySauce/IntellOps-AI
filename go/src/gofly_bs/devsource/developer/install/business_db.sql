/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : gofly_bs_db

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 28/12/2024 14:33:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for attachment
-- ----------------------------
DROP TABLE IF EXISTS `attachment`;
CREATE TABLE `attachment`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '商户账号',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '上传用户',
  `cid` int(11) NOT NULL DEFAULT 0 COMMENT '分类',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问路径',
  `imagewidth` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '宽度',
  `imageheight` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '高度',
  `imagetype` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片类型',
  `imageframes` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '图片帧数',
  `filesize` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
  `mimetype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'mime类型',
  `extparam` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '透传数据',
  `storage` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'local' COMMENT '存储位置',
  `sha1` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件 sha1编码',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件名称',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附件名称',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频封面',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '附件',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '文件类型0=图片，1=文件夹,2=视频，3=音频，4=文件类',
  `is_common` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否公共1=是',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '上传时间',
  `updatetime` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '附件管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of attachment
-- ----------------------------
INSERT INTO `attachment` VALUES (11, 0, 0, 1, 0, '', '', '', '', 0, 0, '', '', 'local', '', '新建文件夹', '', '', 0, 1, 0, 11, NULL, NULL);
INSERT INTO `attachment` VALUES (28, 0, 0, 1, 0, '', '', '', '', 0, 0, '', '', 'local', '', '新建文件夹1', '', '', 0, 1, 0, 28, '2024-05-20 22:46:55', '2024-05-20 22:46:55');
INSERT INTO `attachment` VALUES (42, 0, 0, 0, 0, 'resource/uploads/20241219/e3cfbacb7dda28266f2b4165bbe07362.jpg', '', '', '', 0, 52295, 'image/jpeg', '', 'local', 'd2a7fddd3a1349016e244168c0c85ed2', '微信图片_20230816001705', '微信图片_20230816001705.jpg', '', 0, 0, 0, 42, '2024-12-19 16:23:47', '2024-12-19 16:23:47');

-- ----------------------------
-- Table structure for business_account
-- ----------------------------
DROP TABLE IF EXISTS `business_account`;
CREATE TABLE `business_account`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '账号id/记录那个账号添加',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `main_account` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为主账号:0=否,1=是',
  `dept_id` int(11) NOT NULL COMMENT '部门id',
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户账号',
  `password` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码盐',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '姓名',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(145) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `tel` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备用电话用户自己填写',
  `mobile` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `last_login_ip` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '最后登录IP',
  `last_login_time` int(11) NOT NULL DEFAULT 0 COMMENT '最后登录时间',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态0=正常，1=禁用',
  `validtime` int(11) NOT NULL DEFAULT 0 COMMENT '账号有效时间0=无限',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地址',
  `city` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '城市',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `company` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '公司名称',
  `province` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '省份',
  `area` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地区',
  `fileSize` bigint(11) UNSIGNED NOT NULL DEFAULT 3787456512 COMMENT '附件存储空间',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updatetime` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `deletetime` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  `pwd_reset_time` datetime(0) NULL DEFAULT NULL COMMENT '修改密码时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户端-用户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_account
-- ----------------------------
INSERT INTO `business_account` VALUES (1, 1, 1, 1, 3, 'gofly', '1323a478e990b04633af35175a06e1cd', '1709274330', '开发管理员', 'GoFly', 'http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20240809/ed10a94e2174afa9386ce57d566cc8f3.png', '88422345', '18988375982', '504500934@qq.com', '', 1735350259, 0, 0, '王府井', '昆明', '开发测试账号', 'GoFLy科技1', '', 'chaoyang', 3221225472, '2024-02-05 15:35:59', '2024-12-28 09:44:19', NULL, '2024-12-19 20:36:24');
INSERT INTO `business_account` VALUES (3, 1, 3, 1, 4, 'test', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '测试账号biz', '', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '试试', '', '', 2147483648, '2024-02-05 15:35:59', '2024-12-19 17:08:04', NULL, NULL);
INSERT INTO `business_account` VALUES (12, 1, 1, 0, 1, 'user', 'cec9fcc9526b54902b1d046f4fe47a7e', '784851', '管理员账号', '测试子账号', 'http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20240128/f326244d70c1d1db48b5ef9b8095da8d.png', '', '', '', '', 1734512324, 0, 0, '五华区霖雨路江东耀龙康城27幢二单元502', '昆明市', '', '云律科技（云南）有限公司', '', '', 2147483647, '2024-02-05 15:35:59', '2024-12-22 22:40:42', NULL, NULL);
INSERT INTO `business_account` VALUES (22, 12, 1, 0, 0, 'chuser', 'c18cd0d7e4c6a1112e7d91981c8837aa', '373413', '账号添加', '张三', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-06-09 00:18:35', '2024-12-21 20:55:50', NULL, NULL);
INSERT INTO `business_account` VALUES (23, 1, 23, 1, 0, 'saas1', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '主户', 'saas账号', 'resource/uploads/static/avatar.png', '', '', '', '', 1733644407, 0, 0, '', '', '', '', '', '', 2147483648, '2024-06-09 10:32:24', '2024-12-19 17:08:12', NULL, NULL);
INSERT INTO `business_account` VALUES (24, 4, 24, 1, 0, 'usaas', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', 'user开的saas', 'user开账户', 'resource/uploads/static/avatar.png', '', '', '', '', 1717910665, 0, 0, '', '', '', '', '', '', 2147483648, '2024-06-09 12:02:32', '2024-12-19 17:08:22', NULL, NULL);
INSERT INTO `business_account` VALUES (25, 23, 23, 0, 0, 'saasusr', '5ef9181bafa87b4f30b4c7bd94bf79ee', '862911', '子账号saas1', 'saas1_user', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-06-09 12:48:35', '2024-06-09 12:48:35', NULL, NULL);
INSERT INTO `business_account` VALUES (27, 1, 1, 0, 1, 'test2', 'b0cc47aa3cbc0a329944c1939fc778e0', '181683', '测试', 'test', 'resource/uploads/20241208/d737f7a009fbbf348630202550ee63ac.jpg', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-12-17 22:48:20', '2024-12-17 22:48:26', '2024-12-17 22:48:28', NULL);
INSERT INTO `business_account` VALUES (28, 1, 1, 0, 2, 'test1', 'c2f989d40c53795c41bf37852458e798', '934942', '测试', 'ee', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-12-17 23:05:33', '2024-12-17 23:05:33', '2024-12-17 23:06:18', NULL);
INSERT INTO `business_account` VALUES (29, 1, 1, 0, 3, 'goflyuser', '6d46963fc82ab957d9751678b0e67458', '291275', 'gofly用户', '', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-12-18 13:15:36', '2024-12-18 21:15:58', NULL, NULL);

-- ----------------------------
-- Table structure for business_attachment
-- ----------------------------
DROP TABLE IF EXISTS `business_attachment`;
CREATE TABLE `business_attachment`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件类型0=图片，1=文件夹,2=视频，3=音频,4=文档',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '附件',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附件原来名称',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件名称',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '文件类型0=图片，1=文件夹,2=视频，3=音频,4=文档,5=其他',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问路径',
  `imagewidth` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '宽度',
  `imageheight` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '高度',
  `filesize` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
  `mimetype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'mime类型',
  `extparam` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '透传数据',
  `storage` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'local' COMMENT '存储位置',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频封面',
  `sha1` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件 sha1编码',
  `is_common` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否公共1=是',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 228 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户端附件' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_attachment
-- ----------------------------
INSERT INTO `business_attachment` VALUES (1, 0, 0, '', '默认文件', 1, 1, '', '', '', 0, '', '', 'local', '', '', 1, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (25, 1, 0, '', '新建文件夹8', 1, 25, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (26, 1, 0, '', '新建文件夹9', 1, 26, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (27, 1, 0, '', '新建文件夹10', 1, 27, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (28, 1, 0, '', '新建文件夹11', 1, 28, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (29, 1, 0, '', '新建文件夹12', 1, 29, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (31, 1, 0, '', '新建文件夹13', 1, 31, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (32, 1, 0, '', '新建文件夹15', 1, 32, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (33, 1, 0, '', '新建文件夹16', 1, 33, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (34, 1, 0, '', '新建文件夹16', 1, 34, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (35, 1, 0, '', '新建文件夹18', 1, 35, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (36, 1, 0, '', '新建文件夹19', 1, 36, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (37, 1, 0, '', '新建文件夹20', 1, 37, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (85, 1, 0, '', '新建文件夹21', 1, 85, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (97, 1, 1, '', '新建文件夹1', 1, 97, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05 15:35:59');
INSERT INTO `business_attachment` VALUES (155, 1, 0, '', '新建文件夹1', 1, 155, '', '', '', 0, '', '', 'local', '', '', 0, '2024-06-27 21:50:53');
INSERT INTO `business_attachment` VALUES (204, 1, 1, 'article.zip', 'article', 4, 0, 'resource/uploads/20240731/cc34f4b2874e77b707877dd1e5931fea.zip', '', '', 18460, 'application/octet-stream', '', 'local', '', 'f1c9696a7dbb9382ea62eb34f492cd93', 0, '2024-07-31 23:07:21');
INSERT INTO `business_attachment` VALUES (205, 1, 0, 'gofly软件科技头像.png', 'gofly软件科技头像', 0, 205, 'resource/uploads/20240809/ed10a94e2174afa9386ce57d566cc8f3.png', '', '', 118119, 'image/png', '', 'local', '', '26d9b35717e3787b4718d3e7d95f78d6', 0, '2024-08-09 16:46:18');
INSERT INTO `business_attachment` VALUES (206, 1, 0, 'a310c1cd82b93f4feabf883595d6fbef.jpg', 'a310c1cd82b93f4feabf883595d6fbef', 0, 206, 'resource/uploads/20240930/0f8374d9779353d0183005361c50aecb.jpg', '', '', 382803, 'image/jpeg', '', 'local', '', 'be37d080da82dd2fbb93d4860dead15f', 0, '2024-09-30 17:42:38');
INSERT INTO `business_attachment` VALUES (216, 1, 0, 'mov_bbb.mp4', 'mov_bbb', 2, 229, 'resource/uploads/20241129/b912ee35cf385da12a70842661fe6459.mp4', '', '', 788493, 'video/mp4', '', 'local', 'resource/uploads/20241129/b912ee35cf385da12a70842661fe6459.png', '0d306c5e986211c246fe4be4f696d67f', 0, '2024-11-29 10:51:20');
INSERT INTO `business_attachment` VALUES (217, 1, 0, 'test.docx', 'test', 4, 217, 'resource/uploads/20241129/69f9343ddd4d90c5a9b191da22d2b17a.docx', '', '', 433392, 'application/vnd.openxmlformats-officedocument.wordprocessingml.document', '', 'local', '', '072b4e146e8b0e539c89f89eb95839d8', 0, '2024-11-29 10:57:28');
INSERT INTO `business_attachment` VALUES (218, 1, 0, 'ebf7195d4f742c9acb6fce27760586fa.jpg', 'ebf7195d4f742c9acb6fce27760586fa', 0, 218, 'resource/uploads/20241208/d737f7a009fbbf348630202550ee63ac.jpg', '', '', 196457, 'image/jpeg', '', 'local', '', 'de205b6d875df848836b6cad06e51268', 0, '2024-12-08 20:26:09');
INSERT INTO `business_attachment` VALUES (219, 1, 0, '63c4b478438dbc7d0ff3e08caf0f6b4a.jpg', '63c4b478438dbc7d0ff3e08caf0f6b4a', 0, 219, 'resource/uploads/20241208/eb21c3b2fe2742e5c94230b2a7d5b21d.jpg', '', '', 335082, 'image/jpeg', '', 'local', '', 'e582a3f65a951e3752009f4b233a5981', 0, '2024-12-08 20:26:09');
INSERT INTO `business_attachment` VALUES (220, 1, 0, 'bf3086c71f2391436135c0fd4f99a21c.jpg', 'bf3086c71f2391436135c0fd4f99a21c', 0, 220, 'resource/uploads/20241208/23d1ffc7a3c9e9705e5e095ada383dde.jpg', '', '', 241216, 'image/jpeg', '', 'local', '', '7f8b67885543851f841747f09a75dce8', 0, '2024-12-08 20:26:09');
INSERT INTO `business_attachment` VALUES (221, 1, 0, 'f31bd9e4f226aecd51d07abd1b913e8a.jpg', 'f31bd9e4f226aecd51d07abd1b913e8a', 0, 221, 'resource/uploads/20241208/67f3ef0660e5cd5cec70cfe88cb8d011.jpg', '', '', 441091, 'image/jpeg', '', 'local', '', 'e1e43119ba644a25092263ddbc11e9c6', 0, '2024-12-08 20:26:09');
INSERT INTO `business_attachment` VALUES (222, 1, 0, '23936c87bd3a1e72de4044d420ce7624.jpg', '23936c87bd3a1e72de4044d420ce7624', 0, 222, 'resource/uploads/20241208/cf78412be9a1bbb950b4409b4f512d0d.jpg', '', '', 375459, 'image/jpeg', '', 'local', '', 'ea050552312f2e3fdf639b75b27f126e', 0, '2024-12-08 20:26:09');
INSERT INTO `business_attachment` VALUES (223, 1, 0, 'dff86306862826bd949617d8d7b09649.jpg', 'dff86306862826bd949617d8d7b09649', 0, 223, 'resource/uploads/20241208/2f9365aea081be00daf02b9771724426.jpg', '', '', 277534, 'image/jpeg', '', 'local', '', 'fc73e891232ecc9aa1f158de1286692b', 0, '2024-12-08 20:26:09');
INSERT INTO `business_attachment` VALUES (224, 1, 0, '曾经的你-许巍.mp3', '曾经的你-许巍', 3, 224, 'resource/uploads/20241208/96238ffc4d67315e9e1720f53d9046a7.mp3', '', '', 10463820, 'audio/mpeg', '', 'local', '', '35d37db79c38508e39d67b387e3c8592', 0, '2024-12-08 22:38:45');

-- ----------------------------
-- Table structure for business_auth_dept
-- ----------------------------
DROP TABLE IF EXISTS `business_auth_dept`;
CREATE TABLE `business_auth_dept`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '添加账号',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '上级部门',
  `weigh` int(11) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理后台部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_dept
-- ----------------------------
INSERT INTO `business_auth_dept` VALUES (1, 1, 1, '市场部门', 0, 1, 0, '营销', '2024-02-05 15:35:59');
INSERT INTO `business_auth_dept` VALUES (2, 1, 1, '第一组', 1, 2, 0, '', '2024-02-05 15:35:59');
INSERT INTO `business_auth_dept` VALUES (3, 1, 1, '研发部门', 1, 3, 0, '', '2024-02-05 15:35:59');
INSERT INTO `business_auth_dept` VALUES (9, 23, 23, '主账号部门', 0, 9, 0, '是主账号添加的部门', '2024-06-09 15:16:50');

-- ----------------------------
-- Table structure for business_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `business_auth_role`;
CREATE TABLE `business_auth_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户id',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '父级',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `rules` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '规则ID 所拥有的权限包扣父级',
  `menu` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '选择的id，用于编辑赋值',
  `btns` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '按钮id，用于编辑赋值',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `data_access` tinyint(1) NOT NULL DEFAULT 0 COMMENT '数据权限0=自己1=自己及子权限，2=全部',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `weigh` int(11) NOT NULL COMMENT '排序',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_role
-- ----------------------------
INSERT INTO `business_auth_role` VALUES (1, 0, 0, 0, '超级管理组', '*', '*', '', 0, 2, '账号的总管理员', 1, '2024-02-05 15:35:59');
INSERT INTO `business_auth_role` VALUES (29, 1, 1, 1, 'gofly角色', '1,2,3,8,46,5,4,7,6,9,10,47,11,12,13,14,15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,48,49,50,51,52,53,54,55', '[1,2,3,8,46,5,4,7,6,9,10,47,11,12,13,14]', '[15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,48,49,50,51,52,53,54,55]', 0, 0, '', 29, '2024-06-09 15:38:17');
INSERT INTO `business_auth_role` VALUES (30, 1, 1, 1, '领导组', '1,2,3,11,9,10,8,5,4,7,6,12,13,14,15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,56,57', '[1,2,3,11,9,10,8,5,4,7,6,12,13,14]', '[15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,56,57]', 0, 1, '管理层', 30, '2024-06-30 21:58:46');
INSERT INTO `business_auth_role` VALUES (33, 1, 12, 30, '12新增角色', '1,2,3,8,11,46,5,7,4,6,9,10,14,12,13,15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38', '[1,2,3,8,11,46,5,7,4,6,9,10,14,12,13]', '[15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38]', 0, 0, '测', 33, '2024-12-18 14:07:56');

-- ----------------------------
-- Table structure for business_auth_role_access
-- ----------------------------
DROP TABLE IF EXISTS `business_auth_role_access`;
CREATE TABLE `business_auth_role_access`  (
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `role_id` int(11) NOT NULL DEFAULT 0 COMMENT '授权id'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商务端菜单授权' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_role_access
-- ----------------------------
INSERT INTO `business_auth_role_access` VALUES (1, 1);
INSERT INTO `business_auth_role_access` VALUES (29, 29);
INSERT INTO `business_auth_role_access` VALUES (3, 29);
INSERT INTO `business_auth_role_access` VALUES (23, 30);
INSERT INTO `business_auth_role_access` VALUES (24, 29);
INSERT INTO `business_auth_role_access` VALUES (22, 30);
INSERT INTO `business_auth_role_access` VALUES (12, 30);

-- ----------------------------
-- Table structure for business_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `business_auth_rule`;
CREATE TABLE `business_auth_rule`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `locale` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中英文标题key',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型 0=目录，1=菜单，2=按钮',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '上一级',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标',
  `routepath` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由地址',
  `routename` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由名称',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '组件路径',
  `redirect` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '重定向地址',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '接口路径',
  `permission` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限标识',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态 0=启用1=禁用',
  `isext` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否外链 0=否1=是',
  `keepalive` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否缓存 0=否1=是',
  `requiresauth` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否需要登录鉴权 0=否1=是',
  `hideinmenu` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否在左侧菜单中隐藏该项 0=否1=是',
  `hidechildreninmenu` tinyint(1) NOT NULL DEFAULT 0 COMMENT '强制在左侧菜单中显示单项 0=否1=是',
  `activemenu` tinyint(1) NOT NULL DEFAULT 0 COMMENT '高亮设置的菜单项 0=否1=是',
  `noaffix` tinyint(1) NOT NULL DEFAULT 0 COMMENT '如果设置为true，标签将不会添加到tab-bar中 0=否1=是',
  `onlypage` tinyint(1) NOT NULL DEFAULT 0 COMMENT '独立页面不需layout和登录，如登录页、数据大屏',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 72 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'B端后台菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_auth_rule
-- ----------------------------
INSERT INTO `business_auth_rule` VALUES (1, 1, '概况', '', 'menu.home', 1, 1, 0, 'icon-dashboard', '/home', 'home', '/dashboard/workplace/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (2, 1, '个人中心', '', '', 2, 1, 0, 'icon-user', '/usersetting', 'usersetting', 'system/usersetting/index.vue', '', '', '', 0, 0, 0, 1, 1, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (3, 1, '系统设置', '', 'system', 3, 0, 0, 'icon-settings', '/system', 'system', 'LAYOUT', '/system/account', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (4, 1, '部门管理', '', 'system.dept.title', 2, 1, 3, '', 'dept', 'dept', '/system/dept/index', '', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (5, 1, '菜单管理', '', 'system.rule.title', 1, 1, 3, '', 'rule', 'rule', '/system/rule/index', '', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (6, 1, '账户管理', '', 'system.account.title', 4, 1, 3, '', 'account', 'account', '/system/account/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (7, 1, '角色管理', '', 'system.role.title', 3, 1, 3, '', 'role', 'role', '/system/role/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (8, 1, '数据中心', '', 'datacenter', 4, 0, 0, 'icon-storage', '/datacenter', 'datacenter', 'LAYOUT', '/datacenter/dictionary', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (9, 1, '字典数据', '', 'datacenter.data.title', 9, 1, 8, '', 'data', 'data', '/datacenter/dictionary/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (10, 1, '配置管理', '', 'datacenter.configuration.title', 12, 1, 8, '', 'configuration', 'configuration', '/datacenter/configuration/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (11, 1, '开发者工具', '', '', 5, 0, 0, 'icon-code', '/developer', 'developer', 'LAYOUT', '/developer/generatecode', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (12, 1, '生成代码', '', '', 1, 1, 11, '', 'generatecode', 'generatecode', '/developer/generatecode/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (13, 1, '代码仓库', '', '', 2, 1, 11, '', 'codestore', 'codestore', '/developer/codestore/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (14, 1, '代码生成器', '', '', 3, 1, 11, '', 'codemaker', 'codemaker', '/developer/generatecode/codemaker.vue', '', '', '', 0, 0, 0, 1, 1, 0, 0, 1, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (15, 1, '删除', '', '', 2, 2, 6, '', '', '', '', '', '/business/system/account/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (16, 1, '新建', '创建部门权限', '', 2, 2, 4, '', '', '', '', '', '/business/system/dept/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:35:59');
INSERT INTO `business_auth_rule` VALUES (18, 1, '系统配置', '', '', 1, 2, 10, '', '', '', '', '', '/business/datacenter/common_config/getConfig', 'syscnf', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-06-27 16:08:27');
INSERT INTO `business_auth_rule` VALUES (19, 1, '添加/编辑', '创建账号权限', '', 3, 2, 6, '', '', '', '', '', '/business/system/account/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-07-03 19:09:02');
INSERT INTO `business_auth_rule` VALUES (20, 1, '查看', '', '', 1, 2, 6, '', '', '', '', '', '/business/system/account/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-07-04 12:33:02');
INSERT INTO `business_auth_rule` VALUES (21, 1, '查看', '', '', 1, 2, 5, '', '', '', '', '', '/business/system/rule/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 17:23:26');
INSERT INTO `business_auth_rule` VALUES (22, 1, '删除', '', '', 2, 2, 5, '', '', '', '', '', '/business/system/rule/del', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 17:24:33');
INSERT INTO `business_auth_rule` VALUES (23, 1, '添加/编辑', '', '', 3, 2, 5, '', '', '', '', '', '/business/system/rule/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 17:25:03');
INSERT INTO `business_auth_rule` VALUES (24, 1, '排序', '', '', 4, 2, 5, '', '', '', '', '', '/business/system/rule/tableWeigh', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 17:26:19');
INSERT INTO `business_auth_rule` VALUES (25, 1, '状态', '', '', 5, 2, 5, '', '', '', '', '', '/business/system/rule/upStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 17:28:11');
INSERT INTO `business_auth_rule` VALUES (26, 1, '查看', '查看列表数据', '', 1, 2, 4, '', '', '', '', '', '/business/system/dept/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 21:14:45');
INSERT INTO `business_auth_rule` VALUES (27, 1, '删除', '', '', 3, 2, 4, '', '', '', '', '', '/business/system/dept/del', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 22:07:22');
INSERT INTO `business_auth_rule` VALUES (28, 1, '状态', '', '', 4, 2, 4, '', '', '', '', '', '/business/system/dept/upStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 22:11:12');
INSERT INTO `business_auth_rule` VALUES (29, 1, '状态', '', '', 4, 2, 6, '', '', '', '', '', '/business/system/account/upStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 22:53:51');
INSERT INTO `business_auth_rule` VALUES (30, 1, '查看', '查看数据列表', '', 1, 2, 7, '', '', '', '', '', '/business/system/role/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 22:54:34');
INSERT INTO `business_auth_rule` VALUES (31, 1, '添加/编辑', '添加数据', '', 2, 2, 7, '', '', '', '', '', '/business/system/role/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 23:12:49');
INSERT INTO `business_auth_rule` VALUES (32, 1, '删除', '', '', 3, 2, 7, '', '', '', '', '', '/business/system/role/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17 23:17:47');
INSERT INTO `business_auth_rule` VALUES (33, 1, '添加/编辑', '', '', 2, 2, 9, '', '', '', '', '', '/business/datacenter/dictionary/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 14:42:07');
INSERT INTO `business_auth_rule` VALUES (34, 1, '查看', '', '', 1, 2, 9, '', '', '', '', '', '/business/datacenter/dictionary/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 14:45:06');
INSERT INTO `business_auth_rule` VALUES (35, 1, '状态', '', '', 4, 2, 7, '', '', '', '', '', '/business/system/role/upStatus', 'upStatus', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 14:46:19');
INSERT INTO `business_auth_rule` VALUES (36, 1, '删除', '', '', 3, 2, 9, '', '', '', '', '', '/business/datacenter/dictionary/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:16:15');
INSERT INTO `business_auth_rule` VALUES (37, 1, '状态', '', '', 4, 2, 9, '', '', '', '', '', '/business/datacenter/dictionary/upStatus', 'upStatus', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:16:47');
INSERT INTO `business_auth_rule` VALUES (38, 1, '添加分组', '', '', 5, 2, 9, '', '', '', '', '', '/business/datacenter/tabledata/save', 'addcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:20:32');
INSERT INTO `business_auth_rule` VALUES (39, 1, '删除分组', '', '', 6, 2, 9, '', '', '', '', '', '/business/datacenter/tabledata/del', 'delcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:20:54');
INSERT INTO `business_auth_rule` VALUES (40, 1, '邮箱配置', '', '', 2, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/getEmail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:38:05');
INSERT INTO `business_auth_rule` VALUES (41, 1, '动态配置', '', '', 3, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/getCodestoreConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:41:07');
INSERT INTO `business_auth_rule` VALUES (42, 1, '配置状态', '', '', 4, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/upConfigStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:42:00');
INSERT INTO `business_auth_rule` VALUES (43, 1, '修改邮箱', '', '', 5, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/saveEmail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:43:23');
INSERT INTO `business_auth_rule` VALUES (44, 1, '修改动态配置', '', '', 6, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/saveCodeStoreConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:45:12');
INSERT INTO `business_auth_rule` VALUES (45, 1, '修改系统配置', '', '', 7, 2, 10, '', '', '', '', '', '/business/datacenter/common_config/saveConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 15:57:48');
INSERT INTO `business_auth_rule` VALUES (46, 1, '生成代码示例', '', '', 6, 0, 0, 'icon-compass', '/createcode', 'createcode', 'LAYOUT', '/createcode/product', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-04-30 17:07:40');
INSERT INTO `business_auth_rule` VALUES (47, 1, '测试代码产品', '', '', 47, 1, 46, '', 'product', 'product', 'createcode/product/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (48, 1, '查看', '', '', 1, 2, 47, '', '', '', '', '', '/business/createcode/product/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (49, 1, '添加/编辑', '', '', 2, 2, 47, '', '', '', '', '', '/business/createcode/product/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (50, 1, '删除', '', '', 3, 2, 47, '', '', '', '', '', '/business/createcode/product/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (51, 1, '状态', '', '', 4, 2, 47, '', '', '', '', '', '/business/createcode/product/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (52, 1, '详情', '', '', 5, 2, 47, '', '', '', '', '', '/business/createcode/product/getContent', 'details', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (53, 1, '添加分类', '', '', 6, 2, 47, '', '', '', '', '', '/business/createcode/productcate/save', 'addcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (54, 1, '删除分类', '', '', 7, 2, 47, '', '', '', '', '', '/business/createcode/productcate/del', 'delcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (55, 1, '分类状态', '', '', 8, 2, 47, '', '', '', '', '', '/business/createcode/productcate/upStatus', 'catestatus', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18 19:05:33');
INSERT INTO `business_auth_rule` VALUES (56, 1, '用户信息', '获取用户信息', '', 56, 2, 2, '', '', '', '', '', '/business/user/setting/getUserinfo', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 18:06:21');
INSERT INTO `business_auth_rule` VALUES (57, 1, '修改', '修改密码、手机号等用户信息', '', 57, 2, 2, '', '', '', '', '', '/business/user/setting/saveInfo', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 18:07:40');
INSERT INTO `business_auth_rule` VALUES (60, 1, '系统日志', '', '', 60, 1, 3, '', 'log', 'log', '/system/log/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-12-19 20:53:01');
INSERT INTO `business_auth_rule` VALUES (65, 1, '登录日志', '查看登录日志', '', 65, 2, 60, '', '', '', '', '', '/business/system/log/getLogin', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-20 19:15:12');
INSERT INTO `business_auth_rule` VALUES (66, 1, '操作日志', '', '', 66, 2, 60, '', '', '', '', '', '/business/system/log/getOperation', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-20 23:18:06');
INSERT INTO `business_auth_rule` VALUES (67, 1, '操作日志详情', '', '', 67, 2, 60, '', '', '', '', '', '/business/system/log/getOperationDetail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 18:12:22');
INSERT INTO `business_auth_rule` VALUES (68, 1, '删除登录日志', '', '', 68, 2, 60, '', '', '', '', '', '/business/system/log/delLastLogin', 'delLastLogin', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 20:06:27');
INSERT INTO `business_auth_rule` VALUES (69, 1, '删除操作日志', '', '', 69, 2, 60, '', '', '', '', '', '/business/system/log/delLastOperation', 'delLastOperation', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 20:07:05');
INSERT INTO `business_auth_rule` VALUES (71, 1, '分类列表', '', '', 71, 2, 9, '', '', '', '', '', '/business/datacenter/tabledata/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-26 22:30:16');

-- ----------------------------
-- Table structure for business_home_quickop
-- ----------------------------
DROP TABLE IF EXISTS `business_home_quickop`;
CREATE TABLE `business_home_quickop`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加人',
  `is_common` tinyint(1) NOT NULL DEFAULT 0 COMMENT '公共1=是',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型1=外部',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '快捷名称',
  `path_url` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '跳转路径',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '权重',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '首页快捷操作' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_home_quickop
-- ----------------------------
INSERT INTO `business_home_quickop` VALUES (1, 1, 1, 0, 0, '文档接口', 'devapi', 'svgfont-caozuo-banli', 1);
INSERT INTO `business_home_quickop` VALUES (2, 1, 1, 0, 0, '生成代码', 'generatecode', 'icon-code-sandbox', 2);

-- ----------------------------
-- Table structure for business_user
-- ----------------------------
DROP TABLE IF EXISTS `business_user`;
CREATE TABLE `business_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '用户名',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '昵称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '密码盐',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '电子邮箱',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '头像',
  `level` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '等级',
  `sex` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别:1=男性,2=女性,0=未知',
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `money` decimal(10, 2) NOT NULL COMMENT '余额',
  `score` int(11) NOT NULL DEFAULT 0 COMMENT '积分',
  `successions` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '连续登录天数',
  `maxsuccessions` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '最大连续登录天数',
  `prevtime` bigint(20) NULL DEFAULT NULL COMMENT '上次登录时间',
  `logintime` bigint(20) NULL DEFAULT NULL COMMENT '登录时间',
  `loginip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP',
  `loginfailure` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '失败次数',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updatetime` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deletetime` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `username`(`username`) USING BTREE,
  INDEX `email`(`email`) USING BTREE,
  INDEX `mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 110 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '会员表(用户主表)' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of business_user
-- ----------------------------
INSERT INTO `business_user` VALUES (3, 1, 'gofly', '黄兄', 'gofly', '测试账号', '1323a478e990b04633af35175a06e1cd', '1709274330', '', '', '', 2, 2, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-05-21 20:16:29', '2024-07-09 18:33:44', NULL);
INSERT INTO `business_user` VALUES (4, 1, '', 'GoFly技术', 'U_4', '', '', '', '', '', '', 1, 1, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-05-21 21:39:21', '2024-11-02 22:27:14', NULL);
INSERT INTO `business_user` VALUES (5, 1, '', '', 'u5', '', '', '', '', '', '', 0, 1, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-07-13 22:57:26', '2024-07-13 22:57:27', NULL);
INSERT INTO `business_user` VALUES (8, 1, '', '小程序', 'us8', '测试2', '', '', '', '', '', 0, 1, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-07-14 19:45:50', '2024-07-16 16:50:20', NULL);
INSERT INTO `business_user` VALUES (9, 1, '', '小程', 'u9', '', '', '', '', '', '', 0, 1, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-07-16 10:54:34', '2024-07-16 16:35:30', NULL);
INSERT INTO `business_user` VALUES (105, 1, '', 'huang', 'u105', '', '', '', '', '18988274055', '', 0, 2, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-07-18 22:34:50', '2024-08-16 22:51:58', NULL);
INSERT INTO `business_user` VALUES (107, 1, 'goflys', '', 'GoFly账号', '', '12ad842c4b4c7ed622ee7a4b3e1fa98d', '381119', '', '', 'resource/uploads/20240129/ae5d85e8f6745a4c8e5a2c8d09b958c3.png', 0, 0, NULL, 0.00, 0, 1, 1, NULL, 1726493591, '', 0, 0, '2024-09-10 16:31:30', '2024-09-16 21:33:12', NULL);
INSERT INTO `business_user` VALUES (108, 1, 'zhang', '', '', '', '2e91be15a783d18ac8589f87afd28b69', '389685', '', '', NULL, 0, 0, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-09-10 18:48:23', '2024-09-10 18:48:23', NULL);
INSERT INTO `business_user` VALUES (109, 1, 'test', '', '测试账号', '', '73cf613a5ed38322250c6311f5abcaa4', '857989', '', '', 'resource/uploads/20240809/ed10a94e2174afa9386ce57d566cc8f3.png', 0, 0, NULL, 0.00, 0, 1, 1, NULL, 1726411527, '', 0, 0, '2024-09-11 11:04:32', '2024-09-15 22:45:28', NULL);

-- ----------------------------
-- Table structure for common_api
-- ----------------------------
DROP TABLE IF EXISTS `common_api`;
CREATE TABLE `common_api`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '接口名称',
  `tablename` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '数据表名',
  `fields` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '查询字段',
  `istoken` tinyint(1) NOT NULL DEFAULT 1 COMMENT '需要token 1=需要',
  `createtime` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '低代码接口-处理数据过程' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_api
-- ----------------------------
INSERT INTO `common_api` VALUES (26, '低代码接口', 'business_user', 'avatar,birthday,business_id,createtime', 0, '2024-07-02 17:44:19');

-- ----------------------------
-- Table structure for common_dictionary_data
-- ----------------------------
DROP TABLE IF EXISTS `common_dictionary_data`;
CREATE TABLE `common_dictionary_data`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `data_from` enum('common','business') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'common' COMMENT '数据来源common=公共，business=商业端',
  `group_id` int(10) NOT NULL DEFAULT 0 COMMENT '数据分组id',
  `keyname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典名称',
  `keyvalue` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典项值',
  `tagcolor` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签颜色',
  `des` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典描述',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updatetime` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据-公共表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_dictionary_data
-- ----------------------------
INSERT INTO `common_dictionary_data` VALUES (1, 'business', 2, '管理层', 'mteam', '#D91AD9', '公司领导', 0, 1, '2024-02-05 15:35:59', '2024-12-19 15:25:03');
INSERT INTO `common_dictionary_data` VALUES (2, 'business', 2, '业务员', 'salesman', 'orange', '员工', 0, 2, '2024-02-05 15:35:59', '2024-12-26 15:02:21');
INSERT INTO `common_dictionary_data` VALUES (5, 'business', 4, '汽车', 'car', '#00B42A', '', 0, 5, '2024-06-30 17:25:54', '2024-07-02 21:55:02');
INSERT INTO `common_dictionary_data` VALUES (6, 'business', 4, '飞机', 'air', '#3C7EFF', '', 0, 6, '2024-06-30 22:25:44', '2024-07-02 21:54:51');

-- ----------------------------
-- Table structure for common_dictionary_group
-- ----------------------------
DROP TABLE IF EXISTS `common_dictionary_group`;
CREATE TABLE `common_dictionary_group`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `data_from` enum('common','business') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'common' COMMENT '数据来源common=公共，business=商业端',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典分组名称',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `db_way` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'sys' COMMENT '数据存储位置:sys=公共表,alone=单独建表',
  `tablename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '数据表名称',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `weigh` int(11) NOT NULL DEFAULT 1 COMMENT '排序',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_dictionary_group
-- ----------------------------
INSERT INTO `common_dictionary_group` VALUES (2, 1, 'business', '用户分组', '用户分组', 'sys', 'common_dictionary_data', 0, 2, '2024-02-05 15:35:59');
INSERT INTO `common_dictionary_group` VALUES (4, 1, 'business', '出行方式', '用来存储出行字段', 'sys', 'common_dictionary_data', 0, 4, '2024-06-30 16:54:36');

-- ----------------------------
-- Table structure for common_email
-- ----------------------------
DROP TABLE IF EXISTS `common_email`;
CREATE TABLE `common_email`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `data_from` enum('common','business') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'common' COMMENT '数据来源common=公共，business=商业端',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `sender_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '发送者邮箱',
  `auth_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱授权码',
  `mail_title` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件标题',
  `mail_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件内容,可以是html',
  `service_host` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮件服务器',
  `service_port` int(11) NOT NULL DEFAULT 0 COMMENT '邮件服务器端口',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '业务端邮箱' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_email
-- ----------------------------
INSERT INTO `common_email` VALUES (1, 'business', 1, '504500934@qq.com', 'amidmyjnnxy(youkey)', 'GoFly验证码', '你的验证码为：{code}', 'smtp.qq.com', 587);
INSERT INTO `common_email` VALUES (2, 'common', 0, '504500934@qq.com', 'amidmyjnnxyvbgfb', 'GoFly验证码', '你的验证码为：{code}', 'smtp.qq.com', 587);

-- ----------------------------
-- Table structure for common_generatecode
-- ----------------------------
DROP TABLE IF EXISTS `common_generatecode`;
CREATE TABLE `common_generatecode`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `fromtype` tinyint(1) NOT NULL DEFAULT 0 COMMENT '数据类型0=数据表，1=代码工具',
  `tablename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表名称',
  `comment` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表备注',
  `engine` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '引擎',
  `table_rows` int(11) NOT NULL DEFAULT 0 COMMENT '记录数',
  `collation` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '编码',
  `auto_increment` int(11) NOT NULL DEFAULT 1 COMMENT '自增索引',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '菜单上级',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标',
  `routepath` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由地址',
  `routename` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由名称',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '组件路径',
  `godir` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '后端代码位置',
  `api_path` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '后端业务接口',
  `api_filename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '后端文件名',
  `fields` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '查询字段',
  `rule_id` int(11) NOT NULL DEFAULT 0 COMMENT '生成菜单id',
  `rule_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `codelocation` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'busDirName' COMMENT '生成代码位置',
  `is_install` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否安装0=未安装，1=已安装，2=已卸载',
  `tpl_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'list' COMMENT '模板类型list=仅一个数据，cate=数据加分类',
  `cate_tablename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类表名称',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '上传时间',
  `updatetime` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 800 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_generatecode
-- ----------------------------
INSERT INTO `common_generatecode` VALUES (706, 0, 'createcode_product', '测试代码产品', 'InnoDB', 3, 'utf8mb4_general_ci', 12, 0, 46, '', 'product', 'product', 'createcode/product/index', '', 'business/createcode', 'product.go', 'id,title,image,cid,userType,images,likeColor,record_audio,price,des,sex,workerway,status,updatetime', 47, '测试代码产品', 'busDirName', 1, 'contentcatelist', 'createcode_product_cate', '2024-04-16 15:26:35', '2024-12-28 11:33:15');
INSERT INTO `common_generatecode` VALUES (708, 0, 'createcode_product', '测试代码产品', 'InnoDB', 3, 'utf8mb4_general_ci', 12, 0, 92, '', 'product', 'product', 'createcode/product/index', '', 'admin/createcode', 'product.go', 'id,title,image,cid,num,price,sex,likeColor,userType,images,status,createtime,updatetime,des', 20, '测试代码产品', 'adminDirName', 1, 'sitecatelist', 'createcode_product_cate', '2024-04-30 21:12:25', '2024-12-19 00:02:07');

-- ----------------------------
-- Table structure for common_generatecode_field
-- ----------------------------
DROP TABLE IF EXISTS `common_generatecode_field`;
CREATE TABLE `common_generatecode_field`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `generatecode_id` int(10) NOT NULL COMMENT '关联列表',
  `islist` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是列表1=是',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段名称',
  `field` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段',
  `isorder` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否参与排序',
  `align` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'left' COMMENT '对齐方向',
  `width` int(10) NOT NULL DEFAULT 0 COMMENT '宽度',
  `show_ui` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '显示UI',
  `isform` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为表单字段',
  `required` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为必填项',
  `formtype` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表单类型',
  `datatable` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联数据表',
  `datatablename` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '关联显示字段',
  `dic_group_id` int(10) NOT NULL DEFAULT 0 COMMENT '关联字典分组id',
  `issearch` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否查询',
  `searchway` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '=' COMMENT '查询方式',
  `searchtype` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '查询文本类型',
  `field_weigh` int(10) NOT NULL COMMENT '表单排序',
  `list_weigh` int(10) NOT NULL COMMENT '列表排序',
  `search_weigh` int(10) NOT NULL DEFAULT 0 COMMENT '搜索排序',
  `def_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '默认值',
  `option_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '选项值',
  `gridwidth` int(10) NOT NULL DEFAULT 12 COMMENT '布局栅格',
  `searchwidth` int(10) NOT NULL DEFAULT 120 COMMENT '搜索表单宽',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 505 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '生成代码字段管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_generatecode_field
-- ----------------------------
INSERT INTO `common_generatecode_field` VALUES (146, 706, 0, '球类', 'ballType', 0, 'left', 190, '', 0, 0, 'text', '', '', 0, 0, '=', 'text', 3, 7, 6, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (147, 706, 1, '分类', 'cid', 0, 'left', 100, '', 1, 0, 'belongto', 'createcode_product_cate', 'name', 0, 1, '=', 'belongto', 2, 4, 4, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (148, 706, 0, '内容详情', 'content', 0, 'left', 100, '', 1, 0, 'editor', '', '', 0, 0, '=', 'text', 19, 6, 7, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (149, 706, 0, '注意内容', 'contenttow', 0, 'left', 100, '', 1, 0, 'editor', '', '', 0, 0, '=', 'text', 20, 8, 8, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (150, 706, 0, '上传时间', 'createtime', 0, 'left', 80, '', 0, 0, 'time', '', '', 0, 1, 'between', 'daterange', 9, 9, 3, '', '', 12, 230);
INSERT INTO `common_generatecode_field` VALUES (151, 706, 0, '附件', 'file', 0, 'left', 100, '', 1, 0, 'file', '', '', 0, 0, '=', 'text', 16, 10, 9, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (152, 706, 1, 'ID', 'id', 1, 'left', 60, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 18, 1, 10, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (153, 706, 1, '单张图', 'image', 0, 'left', 80, 'image', 1, 0, 'image', '', '', 0, 0, '=', 'text', 11, 3, 11, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (154, 706, 1, '多张图', 'images', 0, 'left', 100, 'images', 1, 0, 'images', '', '', 0, 0, '=', 'text', 13, 11, 12, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (155, 706, 1, '喜欢颜色', 'likeColor', 0, 'left', 100, 'color', 1, 0, 'colorpicker', '', '', 0, 0, '=', 'text', 12, 12, 13, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (156, 706, 0, '图组', 'moreimgs', 0, 'left', 280, '', 1, 0, 'images', '', '', 0, 0, '=', 'text', 14, 13, 14, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (157, 706, 0, '库存', 'num', 0, 'left', 100, '', 1, 0, 'number', '', '', 0, 0, '=', 'text', 4, 14, 15, '0', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (158, 706, 1, '价格', 'price', 0, 'left', 100, '', 1, 1, 'number', '', '', 0, 0, '=', 'text', 8, 16, 16, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (159, 706, 1, '性别', 'sex', 0, 'left', 100, 'gender', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 6, 18, 17, '0', '0=未知,1=男,2=女', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (160, 706, 1, '状态', 'status', 0, 'left', 100, 'cellstatus', 1, 0, 'radio', '', '', 0, 1, '=', 'select', 7, 20, 2, '0', '0=正常,1=隐藏', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (161, 706, 1, '标题', 'title', 0, 'left', 190, '', 1, 1, 'text', '', '', 0, 1, 'like', 'text', 1, 2, 1, '', '', 12, 150);
INSERT INTO `common_generatecode_field` VALUES (162, 706, 1, '更新时间', 'updatetime', 0, 'left', 190, 'datetime', 0, 0, 'datetime', '', '', 0, 0, '=', 'text', 21, 21, 18, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (163, 706, 1, '用户类型', 'userType', 0, 'left', 110, 'dic', 1, 0, 'belongDic', '', '', 2, 1, '=', 'dic', 5, 5, 5, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (164, 706, 1, '出行方式', 'workerway', 0, 'left', 180, 'tags', 1, 0, 'checkbox', '', '', 0, 0, '=', 'text', 10, 19, 19, '', 'car=汽车,bus=公交,air=飞机', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (165, 708, 1, 'ID', 'id', 1, 'left', 60, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 1, 1, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (166, 708, 1, '分类', 'cid', 0, 'left', 100, '', 1, 1, 'belongto', 'createcode_product_cate', 'name', 0, 0, '=', 'text', 3, 4, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (167, 708, 1, '标题', 'title', 0, 'left', 190, '', 1, 1, 'text', '', '', 0, 1, '=', 'text', 2, 2, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (168, 708, 1, '库存', 'num', 0, 'left', 100, '', 1, 0, 'number', '', '', 0, 0, '=', 'text', 4, 5, 0, '0', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (169, 708, 1, '价格', 'price', 0, 'left', 100, '', 1, 1, 'number', '', '', 0, 0, '=', 'text', 6, 6, 0, '0', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (170, 708, 1, '性别', 'sex', 0, 'left', 100, 'gender', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 7, 7, 0, '0', '0=未知,1=男,2=女', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (171, 708, 0, '内容详情', 'content', 0, 'left', 100, '', 1, 0, 'editor', '', '', 0, 0, '=', 'text', 9, 8, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (172, 708, 1, '喜欢的颜色', 'likeColor', 0, 'left', 190, 'color', 1, 0, 'colorpicker', '', '', 0, 0, '=', 'text', 10, 9, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (173, 708, 0, '球类', 'ballType', 0, 'left', 190, '', 0, 0, 'text', '', '', 0, 0, '=', 'text', 11, 10, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (174, 708, 1, '用户类型', 'userType', 0, 'left', 190, 'dic', 1, 0, 'belongDic', 'common_dictionary_data', '', 2, 1, '=', 'dic', 12, 11, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (175, 708, 1, '单张图', 'image', 0, 'left', 80, 'image', 1, 0, 'image', '', '', 0, 0, '=', 'text', 13, 3, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (176, 708, 1, '多张图', 'images', 0, 'left', 100, 'images', 1, 0, 'images', '', '', 0, 0, '=', 'text', 14, 12, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (177, 708, 0, '图组', 'moreimgs', 0, 'left', 280, 'images', 1, 0, 'images', '', '', 0, 0, '=', 'text', 15, 13, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (178, 708, 0, '注意内容', 'contenttow', 0, 'left', 100, '', 0, 0, 'editor', '', '', 0, 0, '=', 'text', 16, 14, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (179, 708, 0, '附件', 'file', 0, 'left', 100, '', 1, 0, 'file', '', '', 0, 0, '=', 'text', 22, 15, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (180, 708, 0, '出行方式', 'workerway', 0, 'left', 250, '', 1, 0, 'checkbox', '', '', 0, 0, '=', 'text', 8, 16, 0, '', 'car=汽车,bus=公交,air=飞机', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (181, 708, 1, '状态', 'status', 0, 'left', 100, 'dotstatus', 1, 0, 'radio', '', '', 0, 1, '=', 'text', 5, 17, 0, '0', '0=正常,1=隐藏', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (182, 708, 1, '上传时间', 'createtime', 0, 'left', 160, 'datetime', 0, 0, 'time', '', '', 0, 1, 'between', 'daterange', 17, 18, 0, '', '', 12, 230);
INSERT INTO `common_generatecode_field` VALUES (183, 708, 1, '更新时间', 'updatetime', 0, 'left', 160, 'datetime', 0, 0, 'datetime', '', '', 0, 0, '=', 'text', 18, 19, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (184, 710, 1, 'ID', 'id', 1, 'left', 60, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 1, 1, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (185, 710, 1, '任务名称', 'name', 0, 'left', 190, '', 1, 1, 'text', '', '', 0, 1, 'like', 'text', 2, 2, 0, '', '', 24, 180);
INSERT INTO `common_generatecode_field` VALUES (186, 710, 1, '任务分组', 'cid', 0, 'left', 100, '', 1, 1, 'belongto', 'business_crontask_cate', 'name', 0, 0, '=', 'text', 3, 3, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (187, 710, 1, '任务表达式', 'cron_expression', 0, 'left', 280, '', 1, 1, 'text', '', '', 0, 0, '=', 'text', 6, 4, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (188, 710, 0, '调用目标', 'invoke_target', 0, 'left', 280, '', 1, 0, 'text', '', '', 0, 0, '=', 'text', 5, 5, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (189, 710, 1, '状态', 'status', 0, 'left', 100, '', 1, 0, 'radio', '', '', 0, 1, '=', 'select', 10, 6, 0, '1', '0=启用,1=停用', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (190, 710, 0, '目标参数', 'args', 0, 'left', 280, '', 1, 0, 'text', '', '', 0, 0, '=', 'text', 4, 7, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (191, 710, 1, '是否并发', 'concurrent', 0, 'left', 100, '', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 7, 8, 0, '1', '0=允许,1=禁止', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (192, 710, 0, '调用类型', 'job_type', 0, 'left', 100, '', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 8, 9, 0, '1', '1=接口,2=函数', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (193, 710, 0, '执行策略', 'exe_policy', 0, 'left', 100, '', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 9, 11, 0, '1', '1=立即执行,2=执行一次,3=放弃执行', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (194, 710, 1, '创建时间', 'createtime', 0, 'left', 80, '', 0, 0, 'time', '', '', 0, 1, '=', 'text', 11, 12, 0, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (195, 710, 0, '修改时间', 'updatetime', 1, 'left', 150, '', 0, 0, 'datetime', '', '', 0, 0, '=', 'text', 12, 13, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (196, 710, 1, '上次执行时间', 'addtime', 0, 'left', 120, '', 0, 0, 'time', '', '', 0, 0, '=', 'text', 13, 10, 0, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (209, 708, 0, '删除时间', 'deletetime', 0, 'left', 80, '', 0, 0, 'time', '', '', 0, 0, '=', 'text', 19, 183, 183, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (210, 706, 0, '删除时间', 'deletetime', 0, 'left', 80, '', 0, 0, 'time', '', '', 0, 0, '=', 'text', 22, 22, 20, '', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (211, 706, 1, '描述', 'des', 0, 'left', 280, 'des', 1, 0, 'textarea', '', '', 0, 0, '=', 'text', 17, 17, 21, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (241, 708, 1, '描述', 'des', 0, 'left', 280, 'des', 1, 0, 'textarea', '', '', 0, 0, '=', 'text', 20, 209, 209, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (254, 706, 1, '记录录音', 'record_audio', 0, 'left', 190, 'cellcopy', 1, 0, 'audio', '', '', 0, 0, '=', 'text', 15, 15, 22, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (265, 708, 0, '记录录音', 'record_audio', 0, 'left', 100, '', 1, 0, 'audio', '', '', 0, 0, '=', 'text', 21, 241, 241, '', '', 24, 120);
INSERT INTO `common_generatecode_field` VALUES (388, 706, 0, '账号id/记录那个账号添加', 'account_id', 0, 'left', 100, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 254, 23, 254, '0', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (389, 706, 0, '业务主账号id', 'business_id', 0, 'left', 100, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 254, 24, 254, '0', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (503, 708, 0, '账号id/记录那个账号添加', 'account_id', 0, 'left', 100, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 265, 265, 265, '0', '', 12, 120);
INSERT INTO `common_generatecode_field` VALUES (504, 708, 0, '业务主账号id', 'business_id', 0, 'left', 100, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 265, 265, 265, '0', '', 12, 120);

-- ----------------------------
-- Table structure for common_message
-- ----------------------------
DROP TABLE IF EXISTS `common_message`;
CREATE TABLE `common_message`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `usertype` tinyint(1) NOT NULL DEFAULT 0 COMMENT '用户类型0=全部，1=系统-2=boss端，3=C端',
  `account_id` int(11) NOT NULL COMMENT '账号id',
  `adduid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `touid` int(11) NOT NULL DEFAULT 0 COMMENT '接收用户',
  `type` tinyint(1) NOT NULL DEFAULT 2 COMMENT '类型1=通知，2=消息，3=代办',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息标题',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '跳转路由',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息内容',
  `isread` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否已读1=已读',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '发送时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统通用消息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_message
-- ----------------------------

-- ----------------------------
-- Table structure for common_sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `common_sys_login_log`;
CREATE TABLE `common_sys_login_log`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '账号id/记录那个账号添加',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型:admin=admin后台,business=business后台',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
  `ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录IP',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地点',
  `des` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录行为',
  `user_agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '浏览器及浏览器版本信息',
  `error_msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录失败原因',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态:0=成功,1=失败',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统登录日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_sys_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for common_sys_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `common_sys_operation_log`;
CREATE TABLE `common_sys_operation_log`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '账号id/记录那个账号添加',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型:admin=admin后台,business=business后台',
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方法',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求地址',
  `ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录IP',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地点',
  `des` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录行为',
  `req_headers` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求头',
  `req_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求体',
  `resp_headers` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '响应头',
  `resp_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '响应体',
  `latency` decimal(10, 4) NOT NULL COMMENT '耗时',
  `status` int(2) NOT NULL DEFAULT 0 COMMENT '状态:0=成功,1=失败',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1304 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统操作日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_sys_operation_log
-- ----------------------------

-- ----------------------------
-- Table structure for createcode_product
-- ----------------------------
DROP TABLE IF EXISTS `createcode_product`;
CREATE TABLE `createcode_product`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '账号id/记录那个账号添加',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `cid` int(10) NOT NULL COMMENT '分类',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `num` int(10) NOT NULL DEFAULT 0 COMMENT '库存',
  `price` decimal(10, 2) NOT NULL COMMENT '价格',
  `sex` tinyint(1) NOT NULL DEFAULT 0 COMMENT '性别:0=未知,1=男,2=女',
  `des` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容详情',
  `likeColor` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '喜欢的颜色',
  `ballType` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '球类',
  `userType` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户类型',
  `image` varchar(145) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '单张图',
  `images` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '多张图',
  `moreimgs` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图组',
  `contenttow` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '注意内容',
  `record_audio` varchar(145) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '记录录音',
  `file` varchar(145) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附件',
  `workerway` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '出行方式:car=汽车,bus=公交,air=飞机',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态:0=正常,1=隐藏',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updatetime` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deletetime` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试代码产品' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of createcode_product
-- ----------------------------
INSERT INTO `createcode_product` VALUES (5, 1, 1, 1, '无图数据', 23, 18.00, 0, '', '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p838747a-JRN0dh9q\"><br></p></div>', '#C396ED', 'pq', '', '', '', '', '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p838747a-foomPXK5\"><br></p></div>', '', '', '', 0, '2024-02-05 16:07:45', '2024-12-18 19:02:08', NULL);
INSERT INTO `createcode_product` VALUES (8, 1, 1, 1, '单图数据', 22, 12.86, 1, '', '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p838747a-mSXI1S4K\"><span style=\"font-size: 15px;\"><span style=\"font-family: &quot;Microsoft YaHei&quot;, 微软雅黑, &quot;PingFang SC&quot;, SimHei, STHeiti, sans-serif;\">GoFly快速开发版本满足中小企业的项目，golang语言的稳定和性能优异下，部署再配置不错的云服务器上，您不必担心项目性能、稳定、安全等情况，Golang作为一种<strong>高效</strong>、<strong>安全</strong>的编程语言，完全可以帮助开发者快速构建</span><span style=\"font-family: &quot;Microsoft YaHei&quot;, 微软雅黑, &quot;PingFang SC&quot;, SimHei, STHeiti, sans-serif;\"><strong>高效、可靠、安全的应用。GoFly基于Go特性完全满足了您快速开发（像PHP感觉一样开发项目），并</strong></span></span></p></div>', '#D91AD9', 'zq', 'salesman', 'resource/uploads/20241208/cf78412be9a1bbb950b4409b4f512d0d.jpg', 'resource/uploads/20241208/67f3ef0660e5cd5cec70cfe88cb8d011.jpg,resource/uploads/20241208/d737f7a009fbbf348630202550ee63ac.jpg,resource/uploads/20241208/2f9365aea081be00daf02b9771724426.jpg', '', '<div data-v-71ecfed1=\"\" data-element=\"root\" class=\"am-engine\"><div data-type=\"video\" data-value=\"data:%7B%22status%22%3A%22done%22%2C%22name%22%3A%22mov_bbb.mp4%22%2C%22size%22%3A788493%2C%22id%22%3A%22afLuW%22%2C%22type%22%3A%22block%22%2C%22percent%22%3A100%2C%22video_id%22%3A216%2C%22url%22%3A%22resource%2Fuploads%2F20241129%2Fb912ee35cf385da12a70842661fe6459.mp4%22%2C%22cover%22%3A%22resource%2Fuploads%2F20241129%2Fb912ee35cf385da12a70842661fe6459.png%22%2C%22naturalWidth%22%3A320%2C%22naturalHeight%22%3A176%2C%22width%22%3A413%2C%22height%22%3A227%7D\"><video controls=\"\" src=\"http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20241129/b912ee35cf385da12a70842661fe6459.mp4\" poster=\"http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20241129/b912ee35cf385da12a70842661fe6459.png\" webkit-playsinline=\"webkit-playsinline\" playsinline=\"playsinline\" style=\"outline:none;\"></video></div></div>', '', '', '', 0, '2024-03-12 11:11:52', '2024-12-26 12:11:14', NULL);
INSERT INTO `createcode_product` VALUES (9, 1, 1, 1, '多图数据', 0, 0.00, 2, '如果公众号基于安全等考虑，需要获知微信服务器的IP地址列表，以便进行相关限制，可以通过该接口获得微信服务器IP地址列表或者IP网段信息。', '<div data-element=\"root\" class=\"am-engine\"><p data-id=\"p838747a-ViU0CHhf\">你的验证码为：{code}2</p><p data-id=\"p838747a-WlJ31ZcZ\"><br></p></div>', '#3C7EFF', 'zq', 'mteam', 'resource/uploads/20240930/0f8374d9779353d0183005361c50aecb.jpg', 'resource/uploads/20241208/23d1ffc7a3c9e9705e5e095ada383dde.jpg,resource/uploads/20241208/eb21c3b2fe2742e5c94230b2a7d5b21d.jpg', '', '<div data-v-71ecfed1=\"\" data-element=\"root\" class=\"am-engine\"><p data-id=\"p838747a-APqF60Yo\">dddd222</p><p data-id=\"p838747a-jWeNV0da\"><img src=\"http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20241208/cf78412be9a1bbb950b4409b4f512d0d.jpg\" style=\"width: 802px; height: 802px; visibility: visible;\" data-type=\"inline\"></p></div>', 'resource/uploads/20240725/a27486f670a9f08ae6c73174f42b16a6.mp3', 'resource/uploads/20240725/e2f5778a15ecb45c17a856513a819c09.txt', 'car,   bus', 1, '2024-03-12 12:16:15', '2024-12-26 22:45:07', NULL);
INSERT INTO `createcode_product` VALUES (10, 0, 1, 1, '小娃', 0, 0.00, 0, '', '', '', '', 'salesman', '', '', '', '', '', '', '', 0, '2024-12-18 17:20:51', '2024-12-18 17:20:51', '2024-12-18 17:25:30');
INSERT INTO `createcode_product` VALUES (11, 0, 1, 1, '测试', 0, 12.00, 0, '', '', '', '', 'salesman', '', '', '', '', '', '', '', 0, '2024-12-18 17:25:27', '2024-12-18 17:25:27', '2024-12-18 17:25:31');

-- ----------------------------
-- Table structure for createcode_product_cate
-- ----------------------------
DROP TABLE IF EXISTS `createcode_product_cate`;
CREATE TABLE `createcode_product_cate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '账号id/记录那个账号添加',
  `business_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务主账号id',
  `pid` int(10) NOT NULL DEFAULT 0 COMMENT '父级',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `deletetime` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试代码产品分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of createcode_product_cate
-- ----------------------------
INSERT INTO `createcode_product_cate` VALUES (1, 1, 1, 0, '测试分类', '', 0, 0, '2024-02-05 15:35:59', NULL);
INSERT INTO `createcode_product_cate` VALUES (4, 1, 1, 0, '水果', '', 4, 0, '2024-03-22 19:20:25', NULL);
INSERT INTO `createcode_product_cate` VALUES (5, 1, 1, 0, '素菜', '', 5, 0, '2024-03-22 19:21:09', NULL);
INSERT INTO `createcode_product_cate` VALUES (6, 1, 1, 4, '苹果12', '', 6, 0, '2024-03-22 19:23:29', NULL);
INSERT INTO `createcode_product_cate` VALUES (7, 1, 1, 4, '芒果', '', 7, 0, '2024-05-19 14:56:25', NULL);

SET FOREIGN_KEY_CHECKS = 1;
