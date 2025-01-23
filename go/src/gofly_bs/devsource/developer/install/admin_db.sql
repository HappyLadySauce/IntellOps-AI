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

 Date: 28/12/2024 14:34:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_account
-- ----------------------------
DROP TABLE IF EXISTS `admin_account`;
CREATE TABLE `admin_account`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '账号id/用于数据权限',
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
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updatetime` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `pwd_reset_time` datetime(0) NULL DEFAULT NULL COMMENT '修改密码时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户端-用户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_account
-- ----------------------------
INSERT INTO `admin_account` VALUES (1, 1, 3, 'gofly', '1323a478e990b04633af35175a06e1cd', '1709274330', '管理员', 'gofly后台', 'resource/uploads/20241219/e3cfbacb7dda28266f2b4165bbe07362.jpg', '88422345', '18988274072', '504500934@qq.com', '', 1735367101, 0, 0, '对的', '昆明', '开发测试账号', 'GoFLy科技', '', 'chaoyang', '2024-02-05 15:28:03', '2024-12-28 14:25:02', '2024-12-19 13:31:33');
INSERT INTO `admin_account` VALUES (3, 1, 3, 'test', 'f54a88feb9e617a379fa44fe4073a096', '1717830349', '测试账号2', '', 'http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/static/avatar.png', '', '', '', '', 1717855568, 0, 0, '', '', '', '试试', '', '', '2024-02-05 15:28:29', '2024-12-26 14:05:03', NULL);
INSERT INTO `admin_account` VALUES (4, 1, 1, 'user', 'ae43d869ba6dbf533cfc22e6e731f7c2', '1717810362', '销售员', '', 'http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/static/avatar.png', '', '', '', '', 1718092191, 0, 0, '', '', '', '', '', '', '2024-02-05 15:28:25', '2024-06-20 16:50:11', NULL);

-- ----------------------------
-- Table structure for admin_auth_dept
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth_dept`;
CREATE TABLE `admin_auth_dept`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '上级部门',
  `weigh` int(11) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理后台部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth_dept
-- ----------------------------
INSERT INTO `admin_auth_dept` VALUES (1, 1, '市场部门', 0, 1, 0, '营销', '2024-02-05 15:28:03');
INSERT INTO `admin_auth_dept` VALUES (2, 1, '第一组', 1, 2, 0, '', '2024-02-05 15:28:40');
INSERT INTO `admin_auth_dept` VALUES (3, 1, '研发部门', 1, 3, 0, '', '2024-02-05 15:28:43');
INSERT INTO `admin_auth_dept` VALUES (6, 1, '人事组', 1, 6, 0, '', '2024-02-05 15:28:46');

-- ----------------------------
-- Table structure for admin_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth_role`;
CREATE TABLE `admin_auth_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加用户id',
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
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限分组' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth_role
-- ----------------------------
INSERT INTO `admin_auth_role` VALUES (1, 1, 0, '超级管理组', '*', '*', '', 0, 2, '账号的总管理员', 1, '2024-02-05 15:28:51');
INSERT INTO `admin_auth_role` VALUES (5, 1, 1, '销售员', '11,13,1,2,3,7,4,5,6,8,9,10,12,14,107,108,21,22,23,24,25,26,28,29,30,31,32,33,34,35,36,37,38', '[11,13,1,2,3,7,4,5,6,8,9,10,12,14]', '[107,108,21,22,23,24,25,26,28,29,30,31,32,33,34,35,36,37,38]', 0, 0, '产品销售组', 2, '2024-02-05 15:28:53');
INSERT INTO `admin_auth_role` VALUES (6, 1, 1, '管理员', '11,13,1,2,3,4,5,6,7,12,107,108,21,22', '[11,13,1,2,3,4,5,6]', '[107,108,21,22]', 0, 1, '', 3, '2024-02-05 15:28:54');
INSERT INTO `admin_auth_role` VALUES (7, 1, 6, '编辑组', '11,7,13,12,1,2,3,4,5,6,107,108,21,22', '[11,7,13,12,1,2,3,4,5,6]', '[107,108,21,22]', 0, 0, '', 4, '2024-02-05 15:28:55');
INSERT INTO `admin_auth_role` VALUES (8, 1, 6, '兼职组', '11,12,34,7,33', '[11,12,34,7,33]', '', 0, 0, 'ceshi', 5, '2024-02-05 15:28:56');

-- ----------------------------
-- Table structure for admin_auth_role_access
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth_role_access`;
CREATE TABLE `admin_auth_role_access`  (
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '账号id',
  `role_id` int(11) NOT NULL DEFAULT 0 COMMENT '授权id'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'admin端菜单权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth_role_access
-- ----------------------------
INSERT INTO `admin_auth_role_access` VALUES (4, 6);
INSERT INTO `admin_auth_role_access` VALUES (1, 1);
INSERT INTO `admin_auth_role_access` VALUES (3, 7);

-- ----------------------------
-- Table structure for admin_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth_rule`;
CREATE TABLE `admin_auth_rule`  (
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
) ENGINE = InnoDB AUTO_INCREMENT = 85 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'A端后台菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth_rule
-- ----------------------------
INSERT INTO `admin_auth_rule` VALUES (1, 1, '概况', '', 'menu.home', 1, 1, 0, 'icon-dashboard', '/home', 'home', '/dashboard/workplace/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (2, 1, '个人中心', '', '', 2, 1, 0, 'icon-user', '/usersetting', 'usersetting', 'system/usersetting/index.vue', '', '', '', 0, 0, 0, 1, 1, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (3, 1, '业务端管理', '', 'business.rootname', 3, 0, 0, 'icon-book', '/business', 'business', 'LAYOUT', '/business/bizuser', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (4, 1, '菜单管理', '', 'menu.system.rule', 1, 1, 3, '', 'bizrule', 'bizrule', '/business/bizrule/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (5, 1, '账号管理', '', 'business.bizuser.title', 3, 1, 3, '', 'bizuser', 'bizuser', '/business/bizuser/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (6, 1, '业务角色', '', 'business.role', 2, 1, 3, '', 'bizrole', 'bizrole', '/business/bizrole/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (7, 1, '系统设置', '', '', 7, 0, 0, 'icon-settings', '/system', 'system', 'LAYOUT', '/system/account', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (8, 1, '菜单管理', '', '', 1, 1, 7, '', 'rule', 'rule', '/system/rule/index', '', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (9, 1, '部门管理', '', '', 2, 1, 7, '', 'dept', 'dept', '/system/dept/index', '', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (10, 1, '账户管理', '', '', 4, 1, 7, '', 'account', 'account', '/system/account/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (11, 1, '角色管理', '', '', 3, 1, 7, '', 'role', 'role', '/system/role/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (12, 1, '数据中心', '', '', 12, 0, 0, 'icon-storage', '/datacenter', 'datacenter', 'LAYOUT', '/datacenter/logininfo', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (13, 1, '字典数据', '', '', 13, 1, 12, '', 'dictionary', 'dictionary', 'datacenter/dictionary/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (14, 1, '配置管理', '', '', 14, 1, 12, '', 'configuration', 'configuration', '/datacenter/configuration/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (15, 1, '素材管理', '', '', 15, 0, 0, 'icon-folder', '/matter', 'matter', 'LAYOUT', '/matter/attachment', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (16, 1, '系统附件', '', '', 16, 1, 15, '', 'attachment', 'attachment', '/matter/attachment/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (17, 1, '公共图片库', '', '', 17, 1, 15, '', 'picture', 'picture', '/matter/picture/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05 15:30:25');
INSERT INTO `admin_auth_rule` VALUES (18, 1, '生成代码示例', '', '', 18, 0, 0, 'icon-common', '/createcode', 'createcode', 'LAYOUT', '/createcode/product', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-04-30 21:24:54');
INSERT INTO `admin_auth_rule` VALUES (20, 1, '测试代码产品', '', '', 20, 1, 18, '', 'product', 'product', 'createcode/product/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-04-30 21:30:41');
INSERT INTO `admin_auth_rule` VALUES (21, 1, '个人信息', '', '', 21, 2, 2, '', '', '', '', '', '/admin/user/setting/getUserinfo', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 12:21:47');
INSERT INTO `admin_auth_rule` VALUES (22, 1, '修改', '', '', 22, 2, 2, '', '', '', '', '', '/admin/user/setting/saveInfo', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 13:22:29');
INSERT INTO `admin_auth_rule` VALUES (23, 1, '查看', '', '', 23, 2, 9, '', '', '', '', '', '/admin/system/dept/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:30:15');
INSERT INTO `admin_auth_rule` VALUES (24, 1, '添加/编辑', '', '', 24, 2, 9, '', '', '', '', '', '/admin/system/dept/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:30:26');
INSERT INTO `admin_auth_rule` VALUES (25, 1, '删除', '', '', 25, 2, 9, '', '', '', '', '', '/admin/system/dept/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:30:35');
INSERT INTO `admin_auth_rule` VALUES (26, 1, '状态', '', '', 26, 2, 9, '', '', '', '', '', '/admin/system/dept/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:30:41');
INSERT INTO `admin_auth_rule` VALUES (28, 1, '查看', '', '', 28, 2, 10, '', '', '', '', '', '/admin/system/account/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:32:45');
INSERT INTO `admin_auth_rule` VALUES (29, 1, '添加/编辑', '', '', 29, 2, 10, '', '', '', '', '', '/admin/system/account/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:33:01');
INSERT INTO `admin_auth_rule` VALUES (30, 1, '删除', '', '', 30, 2, 10, '', '', '', '', '', '/admin/system/account/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:33:07');
INSERT INTO `admin_auth_rule` VALUES (31, 1, '状态', '', '', 31, 2, 10, '', '', '', '', '', '/admin/system/account/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:33:14');
INSERT INTO `admin_auth_rule` VALUES (32, 1, '查看', '', '', 32, 2, 8, '', '', '', '', '', '/admin/system/rule/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:36:23');
INSERT INTO `admin_auth_rule` VALUES (33, 1, '添加/编辑', '', '', 33, 2, 8, '', '', '', '', '', '/admin/system/rule/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:36:32');
INSERT INTO `admin_auth_rule` VALUES (34, 1, '删除', '', '', 34, 2, 8, '', '', '', '', '', '/admin/system/rule/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:36:44');
INSERT INTO `admin_auth_rule` VALUES (35, 1, '状态', '', '', 35, 2, 8, '', '', '', '', '', '/admin/system/rule/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:36:51');
INSERT INTO `admin_auth_rule` VALUES (36, 1, '查看', '', '', 36, 2, 11, '', '', '', '', '', '/admin/system/role/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:54:31');
INSERT INTO `admin_auth_rule` VALUES (37, 1, '添加/编辑', '', '', 37, 2, 11, '', '', '', '', '', '/admin/system/role/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:54:51');
INSERT INTO `admin_auth_rule` VALUES (38, 1, '删除', '', '', 38, 2, 11, '', '', '', '', '', '/admin/system/role/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 14:55:01');
INSERT INTO `admin_auth_rule` VALUES (39, 1, '查看', '', '', 39, 2, 13, '', '', '', '', '', '/admin/datacenter/dictionary/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:03:57');
INSERT INTO `admin_auth_rule` VALUES (40, 1, '添加/编辑', '', '', 40, 2, 13, '', '', '', '', '', '/admin/datacenter/dictionary/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:18:28');
INSERT INTO `admin_auth_rule` VALUES (41, 1, '删除', '', '', 41, 2, 13, '', '', '', '', '', '/admin/datacenter/dictionary/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:18:38');
INSERT INTO `admin_auth_rule` VALUES (42, 1, '状态', '', '', 42, 2, 13, '', '', '', '', '', '/admin/datacenter/dictionary/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:27:57');
INSERT INTO `admin_auth_rule` VALUES (43, 1, '添加分组', '', '', 43, 2, 13, '', '', '', '', '', '/admin/datacenter/tabledata/save', 'addcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:28:46');
INSERT INTO `admin_auth_rule` VALUES (44, 1, '删除分组', '', '', 44, 2, 13, '', '', '', '', '', '/admin/datacenter/tabledata/del', 'delcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:29:23');
INSERT INTO `admin_auth_rule` VALUES (45, 1, '系统配置', '', '', 45, 2, 14, '', '', '', '', '', '/admin/datacenter/common_config/getConfig', 'syscnf', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:34:21');
INSERT INTO `admin_auth_rule` VALUES (46, 1, '邮箱配置', '', '', 46, 2, 14, '', '', '', '', '', '/admin/datacenter/configuration/getEmail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:35:01');
INSERT INTO `admin_auth_rule` VALUES (47, 1, '动态配置', '', '', 47, 2, 14, '', '', '', '', '', '/admin/datacenter/configuration/getCodestoreConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:35:13');
INSERT INTO `admin_auth_rule` VALUES (48, 1, '配置状态', '', '', 48, 2, 14, '', '', '', '', '', '/admin/datacenter/configuration/upConfigStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:35:26');
INSERT INTO `admin_auth_rule` VALUES (49, 1, '修改邮箱', '', '', 49, 2, 14, '', '', '', '', '', '/admin/datacenter/configuration/saveEmail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:35:39');
INSERT INTO `admin_auth_rule` VALUES (50, 1, '修改动态配置', '', '', 50, 2, 14, '', '', '', '', '', '/admin/datacenter/configuration/saveCodeStoreConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:35:56');
INSERT INTO `admin_auth_rule` VALUES (51, 1, '修改系统配置', '', '', 51, 2, 14, '', '', '', '', '', '/admin/datacenter/common_config/saveConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:36:14');
INSERT INTO `admin_auth_rule` VALUES (52, 1, '查看', '', '', 52, 2, 16, '', '', '', '', '', '/admin/matter/attachment/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:40:08');
INSERT INTO `admin_auth_rule` VALUES (53, 1, '删除', '', '', 53, 2, 16, '', '', '', '', '', '/admin/matter/attachment/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:40:17');
INSERT INTO `admin_auth_rule` VALUES (54, 1, '查看', '', '', 54, 2, 17, '', '', '', '', '', '/admin/matter/picture/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:41:13');
INSERT INTO `admin_auth_rule` VALUES (55, 1, '添加/编辑', '', '', 55, 2, 17, '', '', '', '', '', '/admin/matter/picture/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:48:38');
INSERT INTO `admin_auth_rule` VALUES (56, 1, '删除', '', '', 56, 2, 17, '', '', '', '', '', '/admin/matter/picture/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:57:33');
INSERT INTO `admin_auth_rule` VALUES (57, 1, '状态', '', '', 57, 2, 17, '', '', '', '', '', '/admin/matter/picture/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 15:57:39');
INSERT INTO `admin_auth_rule` VALUES (58, 1, '查看', '', '', 58, 2, 20, '', '', '', '', '', '/admin/createcode/product/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 16:25:26');
INSERT INTO `admin_auth_rule` VALUES (59, 1, '添加/编辑', '', '', 59, 2, 20, '', '', '', '', '', '/admin/createcode/product/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 16:25:33');
INSERT INTO `admin_auth_rule` VALUES (60, 1, '删除', '', '', 60, 2, 20, '', '', '', '', '', '/admin/createcode/product/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 16:25:37');
INSERT INTO `admin_auth_rule` VALUES (61, 1, '状态', '', '', 61, 2, 20, '', '', '', '', '', '/admin/createcode/product/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 16:25:42');
INSERT INTO `admin_auth_rule` VALUES (62, 1, '详情', '', '', 62, 2, 20, '', '', '', '', '', '/admin/createcode/product/getContent', 'details', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 16:26:00');
INSERT INTO `admin_auth_rule` VALUES (63, 1, '查看', '', '', 63, 2, 6, '', '', '', '', '', '/admin/business/bizrole/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 16:35:28');
INSERT INTO `admin_auth_rule` VALUES (64, 1, '查看', '', '', 64, 2, 5, '', '', '', '', '', '/admin/business/bizuser/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 16:43:18');
INSERT INTO `admin_auth_rule` VALUES (65, 1, '查看', '', '', 65, 2, 4, '', '', '', '', '', '/admin/business/bizrule/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 16:51:15');
INSERT INTO `admin_auth_rule` VALUES (66, 1, '添加/编辑', '', '', 66, 2, 5, '', '', '', '', '', '/admin/business/bizuser/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:07:38');
INSERT INTO `admin_auth_rule` VALUES (67, 1, '删除', '', '', 67, 2, 5, '', '', '', '', '', '/admin/business/bizuser/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:07:43');
INSERT INTO `admin_auth_rule` VALUES (68, 1, '状态', '', '', 68, 2, 5, '', '', '', '', '', '/admin/business/bizuser/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:07:47');
INSERT INTO `admin_auth_rule` VALUES (69, 1, '添加/编辑', '', '', 69, 2, 6, '', '', '', '', '', '/admin/business/bizrole/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:10:33');
INSERT INTO `admin_auth_rule` VALUES (70, 1, '删除', '', '', 70, 2, 6, '', '', '', '', '', '/admin/business/bizrole/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:10:38');
INSERT INTO `admin_auth_rule` VALUES (71, 1, '状态', '', '', 71, 2, 6, '', '', '', '', '', '/admin/business/bizrole/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:10:47');
INSERT INTO `admin_auth_rule` VALUES (72, 1, '添加/编辑', '', '', 72, 2, 4, '', '', '', '', '', '/admin/business/bizrule/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:40:13');
INSERT INTO `admin_auth_rule` VALUES (73, 1, '删除', '', '', 73, 2, 4, '', '', '', '', '', '/admin/business/bizrule/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:40:19');
INSERT INTO `admin_auth_rule` VALUES (74, 1, '状态', '', '', 74, 2, 4, '', '', '', '', '', '/admin/business/bizrule/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19 17:40:25');
INSERT INTO `admin_auth_rule` VALUES (75, 1, '系统日志', '', '', 75, 1, 7, '', 'log', 'log', '/system/log/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 21:16:18');
INSERT INTO `admin_auth_rule` VALUES (76, 1, '登录日志', '', '', 76, 2, 75, '', '', '', '', '', '/admin/system/log/getLogin', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 21:18:34');
INSERT INTO `admin_auth_rule` VALUES (77, 1, '操作日志', '', '', 77, 2, 75, '', '', '', '', '', '/admin/system/log/getOperation', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 21:18:52');
INSERT INTO `admin_auth_rule` VALUES (78, 1, '操作日志详情', '', '', 78, 2, 75, '', '', '', '', '', '/admin/system/log/getOperationDetail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 21:19:04');
INSERT INTO `admin_auth_rule` VALUES (79, 1, '删除登录日志', '', '', 79, 2, 75, '', '', '', '', '', '/admin/system/log/delLastLogin', 'delLastLogin', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 21:19:28');
INSERT INTO `admin_auth_rule` VALUES (80, 1, '删除操作日志', '', '', 80, 2, 75, '', '', '', '', '', '/admin/system/log/delLastOperation', 'delLastOperation', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21 21:19:44');
INSERT INTO `admin_auth_rule` VALUES (81, 1, '状态', '', '', 81, 2, 11, '', '', '', '', '', '/admin/system/role/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-26 14:04:41');
INSERT INTO `admin_auth_rule` VALUES (84, 1, '分组列表', '', '', 84, 2, 13, '', '', '', '', '', '/admin/datacenter/tabledata/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-26 22:32:52');

-- ----------------------------
-- Table structure for common_picture
-- ----------------------------
DROP TABLE IF EXISTS `common_picture`;
CREATE TABLE `common_picture`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加账号',
  `cid` int(11) NOT NULL DEFAULT 0 COMMENT '分类id',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附件原来名称',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件名称',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型0=素材图1=插图,2=视频，3=音频',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问路径',
  `imagewidth` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '宽度',
  `imageheight` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '高度',
  `filesize` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
  `mimetype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'mime类型',
  `storage` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'local' COMMENT '存储位置',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频封面',
  `sha1` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件 sha1编码',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '图片库' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_picture
-- ----------------------------
INSERT INTO `common_picture` VALUES (9, 1, 19, 8, 'sw1.jpg', 'loginbanner1', 0, 'resource/uploads/20241219/5c684f1773c11baffb7e860bd3885cb2.png', '', '', 25384, 'image/jpeg', '/dataDB/project/go/gofly_singleresource\\uploads\\20230609\\c895e724853152e06b5915f046348808.jpg', '', '8a81b3c0d0f346d7a36a4573e7196408', 0, '2024-02-05 15:35:59');
INSERT INTO `common_picture` VALUES (10, 1, 24, 6, '信息.png', '查找', 1, 'resource/uploads/20240629/50d582edbe21fedcb7cd5f9e6ee7d871.png', '', '', 65892, 'image/png', '/dataDB/project/go/gofly_singleresource\\uploads\\20230609\\46e5cc40453791e1db8c0e25a1c8ff9c.png', '', 'd58b80c230362875af642143b6bd3a70', 0, '2024-02-05 15:35:59');

-- ----------------------------
-- Table structure for common_picture_cate
-- ----------------------------
DROP TABLE IF EXISTS `common_picture_cate`;
CREATE TABLE `common_picture_cate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '添加账号',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型0=素材图1=插图,2=两种共有',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类名称',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态1=禁用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `createtime` datetime(0) NULL DEFAULT NULL COMMENT '上传时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '分类名称' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of common_picture_cate
-- ----------------------------
INSERT INTO `common_picture_cate` VALUES (1, 1, 1, 0, '商务', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (2, 1, 2, 2, '科技', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (3, 1, 3, 0, '教育', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (4, 1, 4, 0, '风景', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (5, 1, 5, 0, '建筑', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (6, 1, 6, 2, '人物', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (7, 1, 7, 0, '金融', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (8, 1, 8, 0, '城市', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (9, 1, 9, 0, '运动', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (10, 1, 10, 2, '美食', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (11, 1, 11, 0, '交通', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (12, 1, 12, 0, '植物', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (13, 1, 13, 2, '动物', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (14, 1, 14, 0, '生活', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (15, 1, 15, 0, '创意', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (16, 1, 16, 0, '艺术', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (17, 1, 17, 0, '场景', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (18, 1, 18, 0, '生产', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (19, 1, 19, 0, '军事', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (20, 1, 20, 0, '背景', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (21, 1, 21, 1, '产品', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (22, 1, 22, 1, '浮漂', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (23, 1, 23, 1, '水墨', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (24, 1, 24, 1, '特效', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (25, 1, 25, 1, '动物', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (26, 1, 26, 1, '自然', 0, '', '2024-02-05 15:35:59');
INSERT INTO `common_picture_cate` VALUES (27, 1, 27, 1, '文字', 0, '', '2024-02-05 15:35:59');

SET FOREIGN_KEY_CHECKS = 1;
