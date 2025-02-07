-- 数据库导出时间: 2025-02-07 11:10:43
-- 数据库: gofly_bs

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;


-- 删除表 attachment
DROP TABLE IF EXISTS attachment CASCADE;

-- 创建表 attachment
CREATE TABLE attachment (
id bigint NOT NULL,
    account_id integer NOT NULL,
    business_id integer NOT NULL,
    uid integer NOT NULL,
    cid integer NOT NULL,
    url character varying(255) NOT NULL,
    imagewidth character varying(30) NOT NULL,
    imageheight character varying(30) NOT NULL,
    imagetype character varying(30) NOT NULL,
    imageframes bigint NOT NULL,
    filesize bigint NOT NULL,
    mimetype character varying(100) NOT NULL,
    extparam character varying(255) NOT NULL,
    storage character varying(100) NOT NULL,
    sha1 character varying(40) NOT NULL,
    title character varying(50) NOT NULL,
    name character varying(255) NOT NULL,
    cover_url character varying(255) NOT NULL,
    pid integer NOT NULL,
    type smallint NOT NULL,
    is_common smallint NOT NULL,
    weigh integer NOT NULL,
    createtime timestamp without time zone,
    updatetime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 attachment 数据
INSERT INTO attachment (id, account_id, business_id, uid, cid, url, imagewidth, imageheight, imagetype, imageframes, filesize, mimetype, extparam, storage, sha1, title, name, cover_url, pid, type, is_common, weigh, createtime, updatetime) VALUES (11, 0, 0, 1, 0, '', '', '', '', 0, 0, '', '', 'local', '', '新建文件夹', '', '', 0, 1, 0, 11, NULL, NULL);
INSERT INTO attachment (id, account_id, business_id, uid, cid, url, imagewidth, imageheight, imagetype, imageframes, filesize, mimetype, extparam, storage, sha1, title, name, cover_url, pid, type, is_common, weigh, createtime, updatetime) VALUES (28, 0, 0, 1, 0, '', '', '', '', 0, 0, '', '', 'local', '', '新建文件夹1', '', '', 0, 1, 0, 28, '2024-05-20T22:46:55', '2024-05-20T22:46:55');
INSERT INTO attachment (id, account_id, business_id, uid, cid, url, imagewidth, imageheight, imagetype, imageframes, filesize, mimetype, extparam, storage, sha1, title, name, cover_url, pid, type, is_common, weigh, createtime, updatetime) VALUES (42, 0, 0, 0, 0, 'resource/uploads/20241219/e3cfbacb7dda28266f2b4165bbe07362.jpg', '', '', '', 0, 52295, 'image/jpeg', '', 'local', 'd2a7fddd3a1349016e244168c0c85ed2', '微信图片_20230816001705', '微信图片_20230816001705.jpg', '', 0, 0, 0, 42, '2024-12-19T16:23:47', '2024-12-19T16:23:47');

-- 删除表 business_account
DROP TABLE IF EXISTS business_account CASCADE;

-- 创建表 business_account
CREATE TABLE business_account (
id bigint NOT NULL,
    account_id integer NOT NULL,
    business_id integer NOT NULL,
    main_account smallint NOT NULL,
    dept_id integer NOT NULL,
    username character varying(30) NOT NULL,
    password character varying(80) NOT NULL,
    salt character varying(30) NOT NULL,
    name character varying(50) NOT NULL,
    nickname character varying(50) NOT NULL,
    avatar character varying(145) NOT NULL,
    tel character varying(30) NOT NULL,
    mobile character varying(30) NOT NULL,
    email character varying(50) NOT NULL,
    last_login_ip character varying(30) NOT NULL,
    last_login_time integer NOT NULL,
    status smallint NOT NULL,
    validtime integer NOT NULL,
    address character varying(255) NOT NULL,
    city character varying(100) NOT NULL,
    remark character varying(255) NOT NULL,
    company character varying(100) NOT NULL,
    province character varying(30) NOT NULL,
    area character varying(50) NOT NULL,
    "fileSize" numeric NOT NULL,
    createtime timestamp without time zone,
    updatetime timestamp without time zone,
    deletetime timestamp without time zone,
    pwd_reset_time timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 business_account 数据
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (3, 1, 3, 1, 4, 'test', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '测试账号biz', '', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '试试', '', '', 2147483648, '2024-02-05T15:35:59', '2024-12-19T17:08:04', NULL, NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (12, 1, 1, 0, 1, 'user', 'cec9fcc9526b54902b1d046f4fe47a7e', '784851', '管理员账号', '测试子账号', 'http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20240128/f326244d70c1d1db48b5ef9b8095da8d.png', '', '', '', '', 1734512324, 0, 0, '五华区霖雨路江东耀龙康城27幢二单元502', '昆明市', '', '云律科技（云南）有限公司', '', '', 2147483647, '2024-02-05T15:35:59', '2024-12-22T22:40:42', NULL, NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (22, 12, 1, 0, 0, 'chuser', 'c18cd0d7e4c6a1112e7d91981c8837aa', '373413', '账号添加', '张三', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-06-09T00:18:35', '2024-12-21T20:55:50', NULL, NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (23, 1, 23, 1, 0, 'saas1', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', '主户', 'saas账号', 'resource/uploads/static/avatar.png', '', '', '', '', 1733644407, 0, 0, '', '', '', '', '', '', 2147483648, '2024-06-09T10:32:24', '2024-12-19T17:08:12', NULL, NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (24, 4, 24, 1, 0, 'usaas', '9bb610df8adde220720f23dabad486e0', '3305628230121721621', 'user开的saas', 'user开账户', 'resource/uploads/static/avatar.png', '', '', '', '', 1717910665, 0, 0, '', '', '', '', '', '', 2147483648, '2024-06-09T12:02:32', '2024-12-19T17:08:22', NULL, NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (25, 23, 23, 0, 0, 'saasusr', '5ef9181bafa87b4f30b4c7bd94bf79ee', '862911', '子账号saas1', 'saas1_user', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-06-09T12:48:35', '2024-06-09T12:48:35', NULL, NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (27, 1, 1, 0, 1, 'test2', 'b0cc47aa3cbc0a329944c1939fc778e0', '181683', '测试', 'test', 'resource/uploads/20241208/d737f7a009fbbf348630202550ee63ac.jpg', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-12-17T22:48:20', '2024-12-17T22:48:26', '2024-12-17T22:48:28', NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (28, 1, 1, 0, 2, 'test1', 'c2f989d40c53795c41bf37852458e798', '934942', '测试', 'ee', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-12-17T23:05:33', '2024-12-17T23:05:33', '2024-12-17T23:06:18', NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (29, 1, 1, 0, 3, 'goflyuser', '6d46963fc82ab957d9751678b0e67458', '291275', 'gofly用户', '', 'resource/uploads/static/avatar.png', '', '', '', '', 0, 0, 0, '', '', '', '', '', '', 3787456512, '2024-12-18T13:15:36', '2024-12-18T21:15:58', NULL, NULL);
INSERT INTO business_account (id, account_id, business_id, main_account, dept_id, username, password, salt, name, nickname, avatar, tel, mobile, email, last_login_ip, last_login_time, status, validtime, address, city, remark, company, province, area, "fileSize", createtime, updatetime, deletetime, pwd_reset_time) VALUES (1, 1, 1, 1, 3, 'gofly', 'caae530801e35a128c3ad835af9d9aef', '376718', '开发管理员', 'GoFly', 'http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20240809/ed10a94e2174afa9386ce57d566cc8f3.png', '88422345', '18988375982', '504500934@qq.com', '', 1738759240, 0, 0, '王府井', '昆明', '开发测试账号', 'GoFLy科技1', '', 'chaoyang', 3221225472, '2024-02-05T15:35:59', '2025-02-05T20:40:40.893025', NULL, '2024-12-19T20:36:24');

-- 删除表 business_attachment
DROP TABLE IF EXISTS business_attachment CASCADE;

-- 创建表 business_attachment
CREATE TABLE business_attachment (
id bigint NOT NULL,
    business_id integer NOT NULL,
    pid integer NOT NULL,
    name character varying(50) NOT NULL,
    title character varying(50) NOT NULL,
    type smallint NOT NULL,
    weigh integer NOT NULL,
    url character varying(255) NOT NULL,
    imagewidth character varying(30) NOT NULL,
    imageheight character varying(30) NOT NULL,
    filesize bigint NOT NULL,
    mimetype character varying(100) NOT NULL,
    extparam character varying(255) NOT NULL,
    storage character varying(100) NOT NULL,
    cover_url character varying(255) NOT NULL,
    sha1 character varying(40) NOT NULL,
    is_common smallint NOT NULL,
    createtime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 business_attachment 数据
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (1, 0, 0, '', '默认文件', 1, 1, '', '', '', 0, '', '', 'local', '', '', 1, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (25, 1, 0, '', '新建文件夹8', 1, 25, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (26, 1, 0, '', '新建文件夹9', 1, 26, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (27, 1, 0, '', '新建文件夹10', 1, 27, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (28, 1, 0, '', '新建文件夹11', 1, 28, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (29, 1, 0, '', '新建文件夹12', 1, 29, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (31, 1, 0, '', '新建文件夹13', 1, 31, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (32, 1, 0, '', '新建文件夹15', 1, 32, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (33, 1, 0, '', '新建文件夹16', 1, 33, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (34, 1, 0, '', '新建文件夹16', 1, 34, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (35, 1, 0, '', '新建文件夹18', 1, 35, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (36, 1, 0, '', '新建文件夹19', 1, 36, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (37, 1, 0, '', '新建文件夹20', 1, 37, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (85, 1, 0, '', '新建文件夹21', 1, 85, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (97, 1, 1, '', '新建文件夹1', 1, 97, '', '', '', 0, '', '', 'local', '', '', 0, '2024-02-05T15:35:59');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (155, 1, 0, '', '新建文件夹1', 1, 155, '', '', '', 0, '', '', 'local', '', '', 0, '2024-06-27T21:50:53');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (204, 1, 1, 'article.zip', 'article', 4, 0, 'resource/uploads/20240731/cc34f4b2874e77b707877dd1e5931fea.zip', '', '', 18460, 'application/octet-stream', '', 'local', '', 'f1c9696a7dbb9382ea62eb34f492cd93', 0, '2024-07-31T23:07:21');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (205, 1, 0, 'gofly软件科技头像.png', 'gofly软件科技头像', 0, 205, 'resource/uploads/20240809/ed10a94e2174afa9386ce57d566cc8f3.png', '', '', 118119, 'image/png', '', 'local', '', '26d9b35717e3787b4718d3e7d95f78d6', 0, '2024-08-09T16:46:18');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (206, 1, 0, 'a310c1cd82b93f4feabf883595d6fbef.jpg', 'a310c1cd82b93f4feabf883595d6fbef', 0, 206, 'resource/uploads/20240930/0f8374d9779353d0183005361c50aecb.jpg', '', '', 382803, 'image/jpeg', '', 'local', '', 'be37d080da82dd2fbb93d4860dead15f', 0, '2024-09-30T17:42:38');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (216, 1, 0, 'mov_bbb.mp4', 'mov_bbb', 2, 229, 'resource/uploads/20241129/b912ee35cf385da12a70842661fe6459.mp4', '', '', 788493, 'video/mp4', '', 'local', 'resource/uploads/20241129/b912ee35cf385da12a70842661fe6459.png', '0d306c5e986211c246fe4be4f696d67f', 0, '2024-11-29T10:51:20');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (217, 1, 0, 'test.docx', 'test', 4, 217, 'resource/uploads/20241129/69f9343ddd4d90c5a9b191da22d2b17a.docx', '', '', 433392, 'application/vnd.openxmlformats-officedocument.wordprocessingml.document', '', 'local', '', '072b4e146e8b0e539c89f89eb95839d8', 0, '2024-11-29T10:57:28');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (218, 1, 0, 'ebf7195d4f742c9acb6fce27760586fa.jpg', 'ebf7195d4f742c9acb6fce27760586fa', 0, 218, 'resource/uploads/20241208/d737f7a009fbbf348630202550ee63ac.jpg', '', '', 196457, 'image/jpeg', '', 'local', '', 'de205b6d875df848836b6cad06e51268', 0, '2024-12-08T20:26:09');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (219, 1, 0, '63c4b478438dbc7d0ff3e08caf0f6b4a.jpg', '63c4b478438dbc7d0ff3e08caf0f6b4a', 0, 219, 'resource/uploads/20241208/eb21c3b2fe2742e5c94230b2a7d5b21d.jpg', '', '', 335082, 'image/jpeg', '', 'local', '', 'e582a3f65a951e3752009f4b233a5981', 0, '2024-12-08T20:26:09');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (220, 1, 0, 'bf3086c71f2391436135c0fd4f99a21c.jpg', 'bf3086c71f2391436135c0fd4f99a21c', 0, 220, 'resource/uploads/20241208/23d1ffc7a3c9e9705e5e095ada383dde.jpg', '', '', 241216, 'image/jpeg', '', 'local', '', '7f8b67885543851f841747f09a75dce8', 0, '2024-12-08T20:26:09');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (221, 1, 0, 'f31bd9e4f226aecd51d07abd1b913e8a.jpg', 'f31bd9e4f226aecd51d07abd1b913e8a', 0, 221, 'resource/uploads/20241208/67f3ef0660e5cd5cec70cfe88cb8d011.jpg', '', '', 441091, 'image/jpeg', '', 'local', '', 'e1e43119ba644a25092263ddbc11e9c6', 0, '2024-12-08T20:26:09');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (222, 1, 0, '23936c87bd3a1e72de4044d420ce7624.jpg', '23936c87bd3a1e72de4044d420ce7624', 0, 222, 'resource/uploads/20241208/cf78412be9a1bbb950b4409b4f512d0d.jpg', '', '', 375459, 'image/jpeg', '', 'local', '', 'ea050552312f2e3fdf639b75b27f126e', 0, '2024-12-08T20:26:09');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (223, 1, 0, 'dff86306862826bd949617d8d7b09649.jpg', 'dff86306862826bd949617d8d7b09649', 0, 223, 'resource/uploads/20241208/2f9365aea081be00daf02b9771724426.jpg', '', '', 277534, 'image/jpeg', '', 'local', '', 'fc73e891232ecc9aa1f158de1286692b', 0, '2024-12-08T20:26:09');
INSERT INTO business_attachment (id, business_id, pid, name, title, type, weigh, url, imagewidth, imageheight, filesize, mimetype, extparam, storage, cover_url, sha1, is_common, createtime) VALUES (224, 1, 0, '曾经的你-许巍.mp3', '曾经的你-许巍', 3, 224, 'resource/uploads/20241208/96238ffc4d67315e9e1720f53d9046a7.mp3', '', '', 10463820, 'audio/mpeg', '', 'local', '', '35d37db79c38508e39d67b387e3c8592', 0, '2024-12-08T22:38:45');

-- 删除表 business_auth_dept
DROP TABLE IF EXISTS business_auth_dept CASCADE;

-- 创建表 business_auth_dept
CREATE TABLE business_auth_dept (
id bigint NOT NULL,
    business_id integer NOT NULL,
    account_id integer NOT NULL,
    name character varying(100) NOT NULL,
    pid integer NOT NULL,
    weigh integer NOT NULL,
    status smallint NOT NULL,
    remark character varying(255) NOT NULL,
    createtime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 business_auth_dept 数据
INSERT INTO business_auth_dept (id, business_id, account_id, name, pid, weigh, status, remark, createtime) VALUES (1, 1, 1, '市场部门', 0, 1, 0, '营销', '2024-02-05T15:35:59');
INSERT INTO business_auth_dept (id, business_id, account_id, name, pid, weigh, status, remark, createtime) VALUES (2, 1, 1, '第一组', 1, 2, 0, '', '2024-02-05T15:35:59');
INSERT INTO business_auth_dept (id, business_id, account_id, name, pid, weigh, status, remark, createtime) VALUES (3, 1, 1, '研发部门', 1, 3, 0, '', '2024-02-05T15:35:59');
INSERT INTO business_auth_dept (id, business_id, account_id, name, pid, weigh, status, remark, createtime) VALUES (9, 23, 23, '主账号部门', 0, 9, 0, '是主账号添加的部门', '2024-06-09T15:16:50');

-- 删除表 business_auth_role
DROP TABLE IF EXISTS business_auth_role CASCADE;

-- 创建表 business_auth_role
CREATE TABLE business_auth_role (
id bigint NOT NULL,
    business_id integer NOT NULL,
    account_id integer NOT NULL,
    pid integer NOT NULL,
    name character varying(50) NOT NULL,
    rules text NOT NULL,
    menu text NOT NULL,
    btns text NOT NULL,
    status smallint NOT NULL,
    data_access smallint NOT NULL,
    remark character varying(255) NOT NULL,
    weigh integer NOT NULL,
    createtime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 business_auth_role 数据
INSERT INTO business_auth_role (id, business_id, account_id, pid, name, rules, menu, btns, status, data_access, remark, weigh, createtime) VALUES (1, 0, 0, 0, '超级管理组', '*', '*', '', 0, 2, '账号的总管理员', 1, '2024-02-05T15:35:59');
INSERT INTO business_auth_role (id, business_id, account_id, pid, name, rules, menu, btns, status, data_access, remark, weigh, createtime) VALUES (29, 1, 1, 1, 'gofly角色', '1,2,3,8,46,5,4,7,6,9,10,47,11,12,13,14,15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,48,49,50,51,52,53,54,55', '[1,2,3,8,46,5,4,7,6,9,10,47,11,12,13,14]', '[15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,48,49,50,51,52,53,54,55]', 0, 0, '', 29, '2024-06-09T15:38:17');
INSERT INTO business_auth_role (id, business_id, account_id, pid, name, rules, menu, btns, status, data_access, remark, weigh, createtime) VALUES (30, 1, 1, 1, '领导组', '1,2,3,11,9,10,8,5,4,7,6,12,13,14,15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,56,57', '[1,2,3,11,9,10,8,5,4,7,6,12,13,14]', '[15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,56,57]', 0, 1, '管理层', 30, '2024-06-30T21:58:46');
INSERT INTO business_auth_role (id, business_id, account_id, pid, name, rules, menu, btns, status, data_access, remark, weigh, createtime) VALUES (33, 1, 12, 30, '12新增角色', '1,2,3,8,11,46,5,7,4,6,9,10,14,12,13,15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38', '[1,2,3,8,11,46,5,7,4,6,9,10,14,12,13]', '[15,16,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38]', 0, 0, '测', 33, '2024-12-18T14:07:56');

-- 删除表 business_auth_role_access
DROP TABLE IF EXISTS business_auth_role_access CASCADE;

-- 创建表 business_auth_role_access
CREATE TABLE business_auth_role_access (
uid integer NOT NULL,
    role_id integer NOT NULL
);

-- 插入表 business_auth_role_access 数据
INSERT INTO business_auth_role_access (uid, role_id) VALUES (1, 1);
INSERT INTO business_auth_role_access (uid, role_id) VALUES (29, 29);
INSERT INTO business_auth_role_access (uid, role_id) VALUES (3, 29);
INSERT INTO business_auth_role_access (uid, role_id) VALUES (23, 30);
INSERT INTO business_auth_role_access (uid, role_id) VALUES (24, 29);
INSERT INTO business_auth_role_access (uid, role_id) VALUES (22, 30);
INSERT INTO business_auth_role_access (uid, role_id) VALUES (12, 30);

-- 删除表 business_auth_rule
DROP TABLE IF EXISTS business_auth_rule CASCADE;

-- 创建表 business_auth_rule
CREATE TABLE business_auth_rule (
id bigint NOT NULL,
    uid integer NOT NULL,
    title character varying(200) NOT NULL,
    des character varying(255) NOT NULL,
    locale character varying(50) NOT NULL,
    weigh integer NOT NULL,
    type smallint NOT NULL,
    pid integer NOT NULL,
    icon character varying(50) NOT NULL,
    routepath character varying(100) NOT NULL,
    routename character varying(100) NOT NULL,
    component character varying(100) NOT NULL,
    redirect character varying(100) NOT NULL,
    path character varying(200) NOT NULL,
    permission character varying(80) NOT NULL,
    status smallint NOT NULL,
    isext smallint NOT NULL,
    keepalive smallint NOT NULL,
    requiresauth smallint NOT NULL,
    hideinmenu smallint NOT NULL,
    hidechildreninmenu smallint NOT NULL,
    activemenu smallint NOT NULL,
    noaffix smallint NOT NULL,
    onlypage smallint NOT NULL,
    createtime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 business_auth_rule 数据
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (1, 1, '概况', '', 'menu.home', 1, 1, 0, 'icon-dashboard', '/home', 'home', '/dashboard/workplace/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (2, 1, '个人中心', '', '', 2, 1, 0, 'icon-user', '/usersetting', 'usersetting', 'system/usersetting/index.vue', '', '', '', 0, 0, 0, 1, 1, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (3, 1, '系统设置', '', 'system', 3, 0, 0, 'icon-settings', '/system', 'system', 'LAYOUT', '/system/account', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (4, 1, '部门管理', '', 'system.dept.title', 2, 1, 3, '', 'dept', 'dept', '/system/dept/index', '', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (5, 1, '菜单管理', '', 'system.rule.title', 1, 1, 3, '', 'rule', 'rule', '/system/rule/index', '', '', '', 0, 0, 1, 1, 2, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (6, 1, '账户管理', '', 'system.account.title', 4, 1, 3, '', 'account', 'account', '/system/account/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (7, 1, '角色管理', '', 'system.role.title', 3, 1, 3, '', 'role', 'role', '/system/role/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (8, 1, '数据中心', '', 'datacenter', 4, 0, 0, 'icon-storage', '/datacenter', 'datacenter', 'LAYOUT', '/datacenter/dictionary', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (9, 1, '字典数据', '', 'datacenter.data.title', 9, 1, 8, '', 'data', 'data', '/datacenter/dictionary/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (10, 1, '配置管理', '', 'datacenter.configuration.title', 12, 1, 8, '', 'configuration', 'configuration', '/datacenter/configuration/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (11, 1, '开发者工具', '', '', 5, 0, 0, 'icon-code', '/developer', 'developer', 'LAYOUT', '/developer/generatecode', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (12, 1, '生成代码', '', '', 1, 1, 11, '', 'generatecode', 'generatecode', '/developer/generatecode/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (13, 1, '代码仓库', '', '', 2, 1, 11, '', 'codestore', 'codestore', '/developer/codestore/index', '', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (14, 1, '代码生成器', '', '', 3, 1, 11, '', 'codemaker', 'codemaker', '/developer/generatecode/codemaker.vue', '', '', '', 0, 0, 0, 1, 1, 0, 0, 1, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (15, 1, '删除', '', '', 2, 2, 6, '', '', '', '', '', '/business/system/account/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (16, 1, '新建', '创建部门权限', '', 2, 2, 4, '', '', '', '', '', '/business/system/dept/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-02-05T15:35:59');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (18, 1, '系统配置', '', '', 1, 2, 10, '', '', '', '', '', '/business/datacenter/common_config/getConfig', 'syscnf', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-06-27T16:08:27');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (19, 1, '添加/编辑', '创建账号权限', '', 3, 2, 6, '', '', '', '', '', '/business/system/account/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-07-03T19:09:02');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (20, 1, '查看', '', '', 1, 2, 6, '', '', '', '', '', '/business/system/account/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-07-04T12:33:02');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (21, 1, '查看', '', '', 1, 2, 5, '', '', '', '', '', '/business/system/rule/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T17:23:26');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (22, 1, '删除', '', '', 2, 2, 5, '', '', '', '', '', '/business/system/rule/del', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T17:24:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (23, 1, '添加/编辑', '', '', 3, 2, 5, '', '', '', '', '', '/business/system/rule/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T17:25:03');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (24, 1, '排序', '', '', 4, 2, 5, '', '', '', '', '', '/business/system/rule/tableWeigh', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T17:26:19');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (25, 1, '状态', '', '', 5, 2, 5, '', '', '', '', '', '/business/system/rule/upStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T17:28:11');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (26, 1, '查看', '查看列表数据', '', 1, 2, 4, '', '', '', '', '', '/business/system/dept/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T21:14:45');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (27, 1, '删除', '', '', 3, 2, 4, '', '', '', '', '', '/business/system/dept/del', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T22:07:22');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (28, 1, '状态', '', '', 4, 2, 4, '', '', '', '', '', '/business/system/dept/upStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T22:11:12');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (29, 1, '状态', '', '', 4, 2, 6, '', '', '', '', '', '/business/system/account/upStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T22:53:51');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (30, 1, '查看', '查看数据列表', '', 1, 2, 7, '', '', '', '', '', '/business/system/role/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T22:54:34');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (31, 1, '添加/编辑', '添加数据', '', 2, 2, 7, '', '', '', '', '', '/business/system/role/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T23:12:49');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (32, 1, '删除', '', '', 3, 2, 7, '', '', '', '', '', '/business/system/role/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-17T23:17:47');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (33, 1, '添加/编辑', '', '', 2, 2, 9, '', '', '', '', '', '/business/datacenter/dictionary/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T14:42:07');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (34, 1, '查看', '', '', 1, 2, 9, '', '', '', '', '', '/business/datacenter/dictionary/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T14:45:06');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (35, 1, '状态', '', '', 4, 2, 7, '', '', '', '', '', '/business/system/role/upStatus', 'upStatus', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T14:46:19');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (36, 1, '删除', '', '', 3, 2, 9, '', '', '', '', '', '/business/datacenter/dictionary/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:16:15');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (37, 1, '状态', '', '', 4, 2, 9, '', '', '', '', '', '/business/datacenter/dictionary/upStatus', 'upStatus', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:16:47');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (38, 1, '添加分组', '', '', 5, 2, 9, '', '', '', '', '', '/business/datacenter/tabledata/save', 'addcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:20:32');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (39, 1, '删除分组', '', '', 6, 2, 9, '', '', '', '', '', '/business/datacenter/tabledata/del', 'delcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:20:54');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (40, 1, '邮箱配置', '', '', 2, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/getEmail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:38:05');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (41, 1, '动态配置', '', '', 3, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/getCodestoreConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:41:07');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (42, 1, '配置状态', '', '', 4, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/upConfigStatus', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:42:00');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (43, 1, '修改邮箱', '', '', 5, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/saveEmail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:43:23');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (44, 1, '修改动态配置', '', '', 6, 2, 10, '', '', '', '', '', '/business/datacenter/configuration/saveCodeStoreConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:45:12');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (45, 1, '修改系统配置', '', '', 7, 2, 10, '', '', '', '', '', '/business/datacenter/common_config/saveConfig', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T15:57:48');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (46, 1, '生成代码示例', '', '', 6, 0, 0, 'icon-compass', '/createcode', 'createcode', 'LAYOUT', '/createcode/product', '', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-04-30T17:07:40');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (47, 1, '测试代码产品', '', '', 47, 1, 46, '', 'product', 'product', 'createcode/product/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (48, 1, '查看', '', '', 1, 2, 47, '', '', '', '', '', '/business/createcode/product/getList', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (49, 1, '添加/编辑', '', '', 2, 2, 47, '', '', '', '', '', '/business/createcode/product/save', 'add', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (50, 1, '删除', '', '', 3, 2, 47, '', '', '', '', '', '/business/createcode/product/del', 'del', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (51, 1, '状态', '', '', 4, 2, 47, '', '', '', '', '', '/business/createcode/product/upStatus', 'status', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (52, 1, '详情', '', '', 5, 2, 47, '', '', '', '', '', '/business/createcode/product/getContent', 'details', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (53, 1, '添加分类', '', '', 6, 2, 47, '', '', '', '', '', '/business/createcode/productcate/save', 'addcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (54, 1, '删除分类', '', '', 7, 2, 47, '', '', '', '', '', '/business/createcode/productcate/del', 'delcate', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (55, 1, '分类状态', '', '', 8, 2, 47, '', '', '', '', '', '/business/createcode/productcate/upStatus', 'catestatus', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-18T19:05:33');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (56, 1, '用户信息', '获取用户信息', '', 56, 2, 2, '', '', '', '', '', '/business/user/setting/getUserinfo', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19T18:06:21');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (57, 1, '修改', '修改密码、手机号等用户信息', '', 57, 2, 2, '', '', '', '', '', '/business/user/setting/saveInfo', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-19T18:07:40');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (60, 1, '系统日志', '', '', 60, 1, 3, '', 'log', 'log', '/system/log/index', '', '', '', 0, 0, 1, 1, 0, 0, 0, 0, 0, '2024-12-19T20:53:01');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (65, 1, '登录日志', '查看登录日志', '', 65, 2, 60, '', '', '', '', '', '/business/system/log/getLogin', 'view', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-20T19:15:12');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (66, 1, '操作日志', '', '', 66, 2, 60, '', '', '', '', '', '/business/system/log/getOperation', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-20T23:18:06');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (67, 1, '操作日志详情', '', '', 67, 2, 60, '', '', '', '', '', '/business/system/log/getOperationDetail', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21T18:12:22');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (68, 1, '删除登录日志', '', '', 68, 2, 60, '', '', '', '', '', '/business/system/log/delLastLogin', 'delLastLogin', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21T20:06:27');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (69, 1, '删除操作日志', '', '', 69, 2, 60, '', '', '', '', '', '/business/system/log/delLastOperation', 'delLastOperation', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-21T20:07:05');
INSERT INTO business_auth_rule (id, uid, title, des, locale, weigh, type, pid, icon, routepath, routename, component, redirect, path, permission, status, isext, keepalive, requiresauth, hideinmenu, hidechildreninmenu, activemenu, noaffix, onlypage, createtime) VALUES (71, 1, '分类列表', '', '', 71, 2, 9, '', '', '', '', '', '/business/datacenter/tabledata/getList', '', 0, 0, 0, 1, 0, 0, 0, 0, 0, '2024-12-26T22:30:16');

-- 删除表 business_home_quickop
DROP TABLE IF EXISTS business_home_quickop CASCADE;

-- 创建表 business_home_quickop
CREATE TABLE business_home_quickop (
id bigint NOT NULL,
    business_id integer NOT NULL,
    uid integer NOT NULL,
    is_common smallint NOT NULL,
    type smallint NOT NULL,
    name character varying(50) NOT NULL,
    path_url character varying(50) NOT NULL,
    icon character varying(50) NOT NULL,
    weigh integer NOT NULL,
    PRIMARY KEY (id)
);

-- 插入表 business_home_quickop 数据
INSERT INTO business_home_quickop (id, business_id, uid, is_common, type, name, path_url, icon, weigh) VALUES (1, 1, 1, 0, 0, '文档接口', 'devapi', 'svgfont-caozuo-banli', 1);
INSERT INTO business_home_quickop (id, business_id, uid, is_common, type, name, path_url, icon, weigh) VALUES (2, 1, 1, 0, 0, '生成代码', 'generatecode', 'icon-code-sandbox', 2);

-- 删除表 business_user
DROP TABLE IF EXISTS business_user CASCADE;

-- 创建表 business_user
CREATE TABLE business_user (
id bigint NOT NULL,
    business_id integer NOT NULL,
    username character varying(32),
    name character varying(50) NOT NULL,
    nickname character varying(50),
    remark character varying(255) NOT NULL,
    password character varying(32),
    salt character varying(30),
    email character varying(100),
    mobile character varying(11),
    avatar character varying(255),
    level smallint NOT NULL,
    sex smallint NOT NULL,
    birthday date,
    money numeric NOT NULL,
    score integer NOT NULL,
    successions bigint NOT NULL,
    maxsuccessions bigint NOT NULL,
    prevtime bigint,
    logintime bigint,
    loginip character varying(50),
    loginfailure smallint NOT NULL,
    status smallint NOT NULL,
    createtime timestamp without time zone,
    updatetime timestamp without time zone,
    deletetime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 business_user 数据
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (3, 1, 'gofly', '黄兄', 'gofly', '测试账号', '1323a478e990b04633af35175a06e1cd', '1709274330', '', '', '', 2, 2, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-05-21T20:16:29', '2024-07-09T18:33:44', NULL);
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (4, 1, '', 'GoFly技术', 'U_4', '', '', '', '', '', '', 1, 1, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-05-21T21:39:21', '2024-11-02T22:27:14', NULL);
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (5, 1, '', '', 'u5', '', '', '', '', '', '', 0, 1, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-07-13T22:57:26', '2024-07-13T22:57:27', NULL);
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (8, 1, '', '小程序', 'us8', '测试2', '', '', '', '', '', 0, 1, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-07-14T19:45:50', '2024-07-16T16:50:20', NULL);
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (9, 1, '', '小程', 'u9', '', '', '', '', '', '', 0, 1, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-07-16T10:54:34', '2024-07-16T16:35:30', NULL);
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (105, 1, '', 'huang', 'u105', '', '', '', '', '18988274055', '', 0, 2, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-07-18T22:34:50', '2024-08-16T22:51:58', NULL);
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (107, 1, 'goflys', '', 'GoFly账号', '', '12ad842c4b4c7ed622ee7a4b3e1fa98d', '381119', '', '', 'resource/uploads/20240129/ae5d85e8f6745a4c8e5a2c8d09b958c3.png', 0, 0, NULL, 0.00, 0, 1, 1, NULL, 1726493591, '', 0, 0, '2024-09-10T16:31:30', '2024-09-16T21:33:12', NULL);
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (108, 1, 'zhang', '', '', '', '2e91be15a783d18ac8589f87afd28b69', '389685', '', '', NULL, 0, 0, NULL, 0.00, 0, 1, 1, NULL, NULL, '', 0, 0, '2024-09-10T18:48:23', '2024-09-10T18:48:23', NULL);
INSERT INTO business_user (id, business_id, username, name, nickname, remark, password, salt, email, mobile, avatar, level, sex, birthday, money, score, successions, maxsuccessions, prevtime, logintime, loginip, loginfailure, status, createtime, updatetime, deletetime) VALUES (109, 1, 'test', '', '测试账号', '', '73cf613a5ed38322250c6311f5abcaa4', '857989', '', '', 'resource/uploads/20240809/ed10a94e2174afa9386ce57d566cc8f3.png', 0, 0, NULL, 0.00, 0, 1, 1, NULL, 1726411527, '', 0, 0, '2024-09-11T11:04:32', '2024-09-15T22:45:28', NULL);

-- 删除表 common_api
DROP TABLE IF EXISTS common_api CASCADE;

-- 创建表 common_api
CREATE TABLE common_api (
id bigint NOT NULL,
    title character varying(255) NOT NULL,
    tablename character varying(120) NOT NULL,
    fields text NOT NULL,
    istoken smallint NOT NULL,
    createtime timestamp without time zone NOT NULL,
    PRIMARY KEY (id)
);

-- 插入表 common_api 数据
INSERT INTO common_api (id, title, tablename, fields, istoken, createtime) VALUES (26, '低代码接口', 'business_user', 'avatar,birthday,business_id,createtime', 0, '2024-07-02T17:44:19');

-- 删除表 common_dictionary_data
DROP TABLE IF EXISTS common_dictionary_data CASCADE;

-- 创建表 common_dictionary_data
CREATE TABLE common_dictionary_data (
id bigint NOT NULL,
    data_from character varying(255) NOT NULL,
    group_id integer NOT NULL,
    keyname character varying(100) NOT NULL,
    keyvalue character varying(255) NOT NULL,
    tagcolor character varying(20) NOT NULL,
    des character varying(100) NOT NULL,
    status smallint NOT NULL,
    weigh integer NOT NULL,
    createtime timestamp without time zone,
    updatetime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 common_dictionary_data 数据
INSERT INTO common_dictionary_data (id, data_from, group_id, keyname, keyvalue, tagcolor, des, status, weigh, createtime, updatetime) VALUES (1, 'business', 2, '管理层', 'mteam', '#D91AD9', '公司领导', 0, 1, '2024-02-05T15:35:59', '2024-12-19T15:25:03');
INSERT INTO common_dictionary_data (id, data_from, group_id, keyname, keyvalue, tagcolor, des, status, weigh, createtime, updatetime) VALUES (2, 'business', 2, '业务员', 'salesman', 'orange', '员工', 0, 2, '2024-02-05T15:35:59', '2024-12-26T15:02:21');
INSERT INTO common_dictionary_data (id, data_from, group_id, keyname, keyvalue, tagcolor, des, status, weigh, createtime, updatetime) VALUES (5, 'business', 4, '汽车', 'car', '#00B42A', '', 0, 5, '2024-06-30T17:25:54', '2024-07-02T21:55:02');
INSERT INTO common_dictionary_data (id, data_from, group_id, keyname, keyvalue, tagcolor, des, status, weigh, createtime, updatetime) VALUES (6, 'business', 4, '飞机', 'air', '#3C7EFF', '', 0, 6, '2024-06-30T22:25:44', '2024-07-02T21:54:51');

-- 删除表 common_dictionary_group
DROP TABLE IF EXISTS common_dictionary_group CASCADE;

-- 创建表 common_dictionary_group
CREATE TABLE common_dictionary_group (
id bigint NOT NULL,
    business_id integer NOT NULL,
    data_from character varying(255) NOT NULL,
    title character varying(50) NOT NULL,
    remark character varying(200) NOT NULL,
    db_way character varying(10) NOT NULL,
    tablename character varying(50) NOT NULL,
    status smallint NOT NULL,
    weigh integer NOT NULL,
    createtime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 common_dictionary_group 数据
INSERT INTO common_dictionary_group (id, business_id, data_from, title, remark, db_way, tablename, status, weigh, createtime) VALUES (2, 1, 'business', '用户分组', '用户分组', 'sys', 'common_dictionary_data', 0, 2, '2024-02-05T15:35:59');
INSERT INTO common_dictionary_group (id, business_id, data_from, title, remark, db_way, tablename, status, weigh, createtime) VALUES (4, 1, 'business', '出行方式', '用来存储出行字段', 'sys', 'common_dictionary_data', 0, 4, '2024-06-30T16:54:36');

-- 删除表 common_email
DROP TABLE IF EXISTS common_email CASCADE;

-- 创建表 common_email
CREATE TABLE common_email (
id bigint NOT NULL,
    data_from character varying(255) NOT NULL,
    business_id integer NOT NULL,
    sender_email character varying(50) NOT NULL,
    auth_code character varying(50) NOT NULL,
    mail_title character varying(80) NOT NULL,
    mail_body text NOT NULL,
    service_host character varying(30) NOT NULL,
    service_port integer NOT NULL,
    PRIMARY KEY (id)
);

-- 插入表 common_email 数据
INSERT INTO common_email (id, data_from, business_id, sender_email, auth_code, mail_title, mail_body, service_host, service_port) VALUES (1, 'business', 1, '504500934@qq.com', 'amidmyjnnxy(youkey)', 'GoFly验证码', '你的验证码为：{code}', 'smtp.qq.com', 587);
INSERT INTO common_email (id, data_from, business_id, sender_email, auth_code, mail_title, mail_body, service_host, service_port) VALUES (2, 'common', 0, '504500934@qq.com', 'amidmyjnnxyvbgfb', 'GoFly验证码', '你的验证码为：{code}', 'smtp.qq.com', 587);

-- 删除表 common_generatecode
DROP TABLE IF EXISTS common_generatecode CASCADE;

-- 创建表 common_generatecode
CREATE TABLE common_generatecode (
id bigint NOT NULL,
    fromtype smallint NOT NULL,
    tablename character varying(50) NOT NULL,
    comment character varying(100) NOT NULL,
    engine character varying(50) NOT NULL,
    table_rows integer NOT NULL,
    "collation" character varying(50) NOT NULL,
    auto_increment integer NOT NULL,
    status smallint NOT NULL,
    pid integer NOT NULL,
    icon character varying(50) NOT NULL,
    routepath character varying(255) NOT NULL,
    routename character varying(100) NOT NULL,
    component character varying(100) NOT NULL,
    godir character varying(100) NOT NULL,
    api_path character varying(60) NOT NULL,
    api_filename character varying(50) NOT NULL,
    fields text NOT NULL,
    rule_id integer NOT NULL,
    rule_name character varying(30) NOT NULL,
    codelocation character varying(20) NOT NULL,
    is_install smallint NOT NULL,
    tpl_type character varying(20) NOT NULL,
    cate_tablename character varying(50) NOT NULL,
    createtime timestamp without time zone,
    updatetime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 common_generatecode 数据
INSERT INTO common_generatecode (id, fromtype, tablename, comment, engine, table_rows, "collation", auto_increment, status, pid, icon, routepath, routename, component, godir, api_path, api_filename, fields, rule_id, rule_name, codelocation, is_install, tpl_type, cate_tablename, createtime, updatetime) VALUES (706, 0, 'createcode_product', '测试代码产品', 'InnoDB', 3, 'utf8mb4_general_ci', 12, 0, 46, '', 'product', 'product', 'createcode/product/index', '', 'business/createcode', 'product.go', 'id,title,image,cid,userType,images,likeColor,record_audio,price,des,sex,workerway,status,updatetime', 47, '测试代码产品', 'busDirName', 1, 'contentcatelist', 'createcode_product_cate', '2024-04-16T15:26:35', '2024-12-28T11:33:15');
INSERT INTO common_generatecode (id, fromtype, tablename, comment, engine, table_rows, "collation", auto_increment, status, pid, icon, routepath, routename, component, godir, api_path, api_filename, fields, rule_id, rule_name, codelocation, is_install, tpl_type, cate_tablename, createtime, updatetime) VALUES (708, 0, 'createcode_product', '测试代码产品', 'InnoDB', 3, 'utf8mb4_general_ci', 12, 0, 92, '', 'product', 'product', 'createcode/product/index', '', 'admin/createcode', 'product.go', 'id,title,image,cid,num,price,sex,likeColor,userType,images,status,createtime,updatetime,des', 20, '测试代码产品', 'adminDirName', 1, 'sitecatelist', 'createcode_product_cate', '2024-04-30T21:12:25', '2024-12-19T00:02:07');

-- 删除表 common_generatecode_field
DROP TABLE IF EXISTS common_generatecode_field CASCADE;

-- 创建表 common_generatecode_field
CREATE TABLE common_generatecode_field (
id bigint NOT NULL,
    generatecode_id integer NOT NULL,
    islist smallint NOT NULL,
    name character varying(50) NOT NULL,
    field character varying(50) NOT NULL,
    isorder smallint NOT NULL,
    align character varying(10) NOT NULL,
    width integer NOT NULL,
    show_ui character varying(20) NOT NULL,
    isform smallint NOT NULL,
    required smallint NOT NULL,
    formtype character varying(15) NOT NULL,
    datatable character varying(30) NOT NULL,
    datatablename character varying(30) NOT NULL,
    dic_group_id integer NOT NULL,
    issearch smallint NOT NULL,
    searchway character varying(15) NOT NULL,
    searchtype character varying(30) NOT NULL,
    field_weigh integer NOT NULL,
    list_weigh integer NOT NULL,
    search_weigh integer NOT NULL,
    def_value character varying(255) NOT NULL,
    option_value character varying(255) NOT NULL,
    gridwidth integer NOT NULL,
    searchwidth integer NOT NULL,
    PRIMARY KEY (id)
);

-- 插入表 common_generatecode_field 数据
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (146, 706, 0, '球类', 'ballType', 0, 'left', 190, '', 0, 0, 'text', '', '', 0, 0, '=', 'text', 3, 7, 6, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (147, 706, 1, '分类', 'cid', 0, 'left', 100, '', 1, 0, 'belongto', 'createcode_product_cate', 'name', 0, 1, '=', 'belongto', 2, 4, 4, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (148, 706, 0, '内容详情', 'content', 0, 'left', 100, '', 1, 0, 'editor', '', '', 0, 0, '=', 'text', 19, 6, 7, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (149, 706, 0, '注意内容', 'contenttow', 0, 'left', 100, '', 1, 0, 'editor', '', '', 0, 0, '=', 'text', 20, 8, 8, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (150, 706, 0, '上传时间', 'createtime', 0, 'left', 80, '', 0, 0, 'time', '', '', 0, 1, 'between', 'daterange', 9, 9, 3, '', '', 12, 230);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (151, 706, 0, '附件', 'file', 0, 'left', 100, '', 1, 0, 'file', '', '', 0, 0, '=', 'text', 16, 10, 9, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (152, 706, 1, 'ID', 'id', 1, 'left', 60, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 18, 1, 10, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (153, 706, 1, '单张图', 'image', 0, 'left', 80, 'image', 1, 0, 'image', '', '', 0, 0, '=', 'text', 11, 3, 11, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (154, 706, 1, '多张图', 'images', 0, 'left', 100, 'images', 1, 0, 'images', '', '', 0, 0, '=', 'text', 13, 11, 12, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (155, 706, 1, '喜欢颜色', 'likeColor', 0, 'left', 100, 'color', 1, 0, 'colorpicker', '', '', 0, 0, '=', 'text', 12, 12, 13, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (156, 706, 0, '图组', 'moreimgs', 0, 'left', 280, '', 1, 0, 'images', '', '', 0, 0, '=', 'text', 14, 13, 14, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (157, 706, 0, '库存', 'num', 0, 'left', 100, '', 1, 0, 'number', '', '', 0, 0, '=', 'text', 4, 14, 15, '0', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (158, 706, 1, '价格', 'price', 0, 'left', 100, '', 1, 1, 'number', '', '', 0, 0, '=', 'text', 8, 16, 16, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (159, 706, 1, '性别', 'sex', 0, 'left', 100, 'gender', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 6, 18, 17, '0', '0=未知,1=男,2=女', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (160, 706, 1, '状态', 'status', 0, 'left', 100, 'cellstatus', 1, 0, 'radio', '', '', 0, 1, '=', 'select', 7, 20, 2, '0', '0=正常,1=隐藏', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (161, 706, 1, '标题', 'title', 0, 'left', 190, '', 1, 1, 'text', '', '', 0, 1, 'like', 'text', 1, 2, 1, '', '', 12, 150);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (162, 706, 1, '更新时间', 'updatetime', 0, 'left', 190, 'datetime', 0, 0, 'datetime', '', '', 0, 0, '=', 'text', 21, 21, 18, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (163, 706, 1, '用户类型', 'userType', 0, 'left', 110, 'dic', 1, 0, 'belongDic', '', '', 2, 1, '=', 'dic', 5, 5, 5, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (164, 706, 1, '出行方式', 'workerway', 0, 'left', 180, 'tags', 1, 0, 'checkbox', '', '', 0, 0, '=', 'text', 10, 19, 19, '', 'car=汽车,bus=公交,air=飞机', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (165, 708, 1, 'ID', 'id', 1, 'left', 60, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 1, 1, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (166, 708, 1, '分类', 'cid', 0, 'left', 100, '', 1, 1, 'belongto', 'createcode_product_cate', 'name', 0, 0, '=', 'text', 3, 4, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (167, 708, 1, '标题', 'title', 0, 'left', 190, '', 1, 1, 'text', '', '', 0, 1, '=', 'text', 2, 2, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (168, 708, 1, '库存', 'num', 0, 'left', 100, '', 1, 0, 'number', '', '', 0, 0, '=', 'text', 4, 5, 0, '0', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (169, 708, 1, '价格', 'price', 0, 'left', 100, '', 1, 1, 'number', '', '', 0, 0, '=', 'text', 6, 6, 0, '0', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (170, 708, 1, '性别', 'sex', 0, 'left', 100, 'gender', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 7, 7, 0, '0', '0=未知,1=男,2=女', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (171, 708, 0, '内容详情', 'content', 0, 'left', 100, '', 1, 0, 'editor', '', '', 0, 0, '=', 'text', 9, 8, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (172, 708, 1, '喜欢的颜色', 'likeColor', 0, 'left', 190, 'color', 1, 0, 'colorpicker', '', '', 0, 0, '=', 'text', 10, 9, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (173, 708, 0, '球类', 'ballType', 0, 'left', 190, '', 0, 0, 'text', '', '', 0, 0, '=', 'text', 11, 10, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (174, 708, 1, '用户类型', 'userType', 0, 'left', 190, 'dic', 1, 0, 'belongDic', 'common_dictionary_data', '', 2, 1, '=', 'dic', 12, 11, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (175, 708, 1, '单张图', 'image', 0, 'left', 80, 'image', 1, 0, 'image', '', '', 0, 0, '=', 'text', 13, 3, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (176, 708, 1, '多张图', 'images', 0, 'left', 100, 'images', 1, 0, 'images', '', '', 0, 0, '=', 'text', 14, 12, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (177, 708, 0, '图组', 'moreimgs', 0, 'left', 280, 'images', 1, 0, 'images', '', '', 0, 0, '=', 'text', 15, 13, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (178, 708, 0, '注意内容', 'contenttow', 0, 'left', 100, '', 0, 0, 'editor', '', '', 0, 0, '=', 'text', 16, 14, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (179, 708, 0, '附件', 'file', 0, 'left', 100, '', 1, 0, 'file', '', '', 0, 0, '=', 'text', 22, 15, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (180, 708, 0, '出行方式', 'workerway', 0, 'left', 250, '', 1, 0, 'checkbox', '', '', 0, 0, '=', 'text', 8, 16, 0, '', 'car=汽车,bus=公交,air=飞机', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (181, 708, 1, '状态', 'status', 0, 'left', 100, 'dotstatus', 1, 0, 'radio', '', '', 0, 1, '=', 'text', 5, 17, 0, '0', '0=正常,1=隐藏', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (182, 708, 1, '上传时间', 'createtime', 0, 'left', 160, 'datetime', 0, 0, 'time', '', '', 0, 1, 'between', 'daterange', 17, 18, 0, '', '', 12, 230);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (183, 708, 1, '更新时间', 'updatetime', 0, 'left', 160, 'datetime', 0, 0, 'datetime', '', '', 0, 0, '=', 'text', 18, 19, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (184, 710, 1, 'ID', 'id', 1, 'left', 60, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 1, 1, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (185, 710, 1, '任务名称', 'name', 0, 'left', 190, '', 1, 1, 'text', '', '', 0, 1, 'like', 'text', 2, 2, 0, '', '', 24, 180);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (186, 710, 1, '任务分组', 'cid', 0, 'left', 100, '', 1, 1, 'belongto', 'business_crontask_cate', 'name', 0, 0, '=', 'text', 3, 3, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (187, 710, 1, '任务表达式', 'cron_expression', 0, 'left', 280, '', 1, 1, 'text', '', '', 0, 0, '=', 'text', 6, 4, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (188, 710, 0, '调用目标', 'invoke_target', 0, 'left', 280, '', 1, 0, 'text', '', '', 0, 0, '=', 'text', 5, 5, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (189, 710, 1, '状态', 'status', 0, 'left', 100, '', 1, 0, 'radio', '', '', 0, 1, '=', 'select', 10, 6, 0, '1', '0=启用,1=停用', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (190, 710, 0, '目标参数', 'args', 0, 'left', 280, '', 1, 0, 'text', '', '', 0, 0, '=', 'text', 4, 7, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (191, 710, 1, '是否并发', 'concurrent', 0, 'left', 100, '', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 7, 8, 0, '1', '0=允许,1=禁止', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (192, 710, 0, '调用类型', 'job_type', 0, 'left', 100, '', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 8, 9, 0, '1', '1=接口,2=函数', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (193, 710, 0, '执行策略', 'exe_policy', 0, 'left', 100, '', 1, 0, 'radio', '', '', 0, 0, '=', 'text', 9, 11, 0, '1', '1=立即执行,2=执行一次,3=放弃执行', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (194, 710, 1, '创建时间', 'createtime', 0, 'left', 80, '', 0, 0, 'time', '', '', 0, 1, '=', 'text', 11, 12, 0, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (195, 710, 0, '修改时间', 'updatetime', 1, 'left', 150, '', 0, 0, 'datetime', '', '', 0, 0, '=', 'text', 12, 13, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (196, 710, 1, '上次执行时间', 'addtime', 0, 'left', 120, '', 0, 0, 'time', '', '', 0, 0, '=', 'text', 13, 10, 0, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (209, 708, 0, '删除时间', 'deletetime', 0, 'left', 80, '', 0, 0, 'time', '', '', 0, 0, '=', 'text', 19, 183, 183, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (210, 706, 0, '删除时间', 'deletetime', 0, 'left', 80, '', 0, 0, 'time', '', '', 0, 0, '=', 'text', 22, 22, 20, '', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (211, 706, 1, '描述', 'des', 0, 'left', 280, 'des', 1, 0, 'textarea', '', '', 0, 0, '=', 'text', 17, 17, 21, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (241, 708, 1, '描述', 'des', 0, 'left', 280, 'des', 1, 0, 'textarea', '', '', 0, 0, '=', 'text', 20, 209, 209, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (254, 706, 1, '记录录音', 'record_audio', 0, 'left', 190, 'cellcopy', 1, 0, 'audio', '', '', 0, 0, '=', 'text', 15, 15, 22, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (265, 708, 0, '记录录音', 'record_audio', 0, 'left', 100, '', 1, 0, 'audio', '', '', 0, 0, '=', 'text', 21, 241, 241, '', '', 24, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (388, 706, 0, '账号id/记录那个账号添加', 'account_id', 0, 'left', 100, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 254, 23, 254, '0', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (389, 706, 0, '业务主账号id', 'business_id', 0, 'left', 100, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 254, 24, 254, '0', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (503, 708, 0, '账号id/记录那个账号添加', 'account_id', 0, 'left', 100, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 265, 265, 265, '0', '', 12, 120);
INSERT INTO common_generatecode_field (id, generatecode_id, islist, name, field, isorder, align, width, show_ui, isform, required, formtype, datatable, datatablename, dic_group_id, issearch, searchway, searchtype, field_weigh, list_weigh, search_weigh, def_value, option_value, gridwidth, searchwidth) VALUES (504, 708, 0, '业务主账号id', 'business_id', 0, 'left', 100, '', 0, 0, 'number', '', '', 0, 0, '=', 'text', 265, 265, 265, '0', '', 12, 120);

-- 删除表 common_message
DROP TABLE IF EXISTS common_message CASCADE;

-- 创建表 common_message
CREATE TABLE common_message (
id bigint NOT NULL,
    usertype smallint NOT NULL,
    account_id integer NOT NULL,
    adduid integer NOT NULL,
    touid integer NOT NULL,
    type smallint NOT NULL,
    title character varying(255) NOT NULL,
    path character varying(255) NOT NULL,
    content text NOT NULL,
    isread smallint NOT NULL,
    createtime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 common_message 数据

-- 删除表 common_sys_login_log
DROP TABLE IF EXISTS common_sys_login_log CASCADE;

-- 创建表 common_sys_login_log
CREATE TABLE common_sys_login_log (
id bigint NOT NULL,
    account_id integer NOT NULL,
    business_id integer NOT NULL,
    type character varying(20) NOT NULL,
    uid integer NOT NULL,
    ip character varying(50) NOT NULL,
    address character varying(255) NOT NULL,
    des character varying(100) NOT NULL,
    user_agent character varying(255) NOT NULL,
    error_msg character varying(255) NOT NULL,
    status smallint NOT NULL,
    createtime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 common_sys_login_log 数据
INSERT INTO common_sys_login_log (id, account_id, business_id, type, uid, ip, address, des, user_agent, error_msg, status, createtime) VALUES (6, 1, 1, 'business', 1, '127.0.0.1', '0|0|0|内网IP|内网IP', '账号登录', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:134.0) Gecko/20100101 Firefox/134.0', '', 0, '2025-01-20T09:39:07');

-- 删除表 common_sys_operation_log
DROP TABLE IF EXISTS common_sys_operation_log CASCADE;

-- 创建表 common_sys_operation_log
CREATE TABLE common_sys_operation_log (
id bigint NOT NULL,
    account_id integer NOT NULL,
    business_id integer NOT NULL,
    type character varying(20) NOT NULL,
    uid integer NOT NULL,
    request_method character varying(10) NOT NULL,
    url character varying(255) NOT NULL,
    ip character varying(50) NOT NULL,
    address character varying(255) NOT NULL,
    des character varying(100) NOT NULL,
    req_headers text NOT NULL,
    req_body text NOT NULL,
    resp_headers text NOT NULL,
    resp_body text NOT NULL,
    latency numeric NOT NULL,
    status integer NOT NULL,
    createtime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 common_sys_operation_log 数据

-- 删除表 createcode_product
DROP TABLE IF EXISTS createcode_product CASCADE;

-- 创建表 createcode_product
CREATE TABLE createcode_product (
id bigint NOT NULL,
    account_id integer NOT NULL,
    business_id integer NOT NULL,
    cid integer NOT NULL,
    title character varying(50) NOT NULL,
    num integer NOT NULL,
    price numeric NOT NULL,
    sex smallint NOT NULL,
    des character varying(255) NOT NULL,
    content text NOT NULL,
    "likeColor" character varying(30) NOT NULL,
    "ballType" character varying(20) NOT NULL,
    "userType" character varying(20) NOT NULL,
    image character varying(145) NOT NULL,
    images text NOT NULL,
    moreimgs character varying(255) NOT NULL,
    contenttow text NOT NULL,
    record_audio character varying(145) NOT NULL,
    file character varying(145) NOT NULL,
    workerway character varying(100) NOT NULL,
    status smallint NOT NULL,
    createtime timestamp without time zone,
    updatetime timestamp without time zone,
    deletetime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 createcode_product 数据
INSERT INTO createcode_product (id, account_id, business_id, cid, title, num, price, sex, des, content, "likeColor", "ballType", "userType", image, images, moreimgs, contenttow, record_audio, file, workerway, status, createtime, updatetime, deletetime) VALUES (5, 1, 1, 1, '无图数据', 23, 18.00, 0, '', '<div data-element="root" class="am-engine"><p data-id="p838747a-JRN0dh9q"><br></p></div>', '#C396ED', 'pq', '', '', '', '', '<div data-element="root" class="am-engine"><p data-id="p838747a-foomPXK5"><br></p></div>', '', '', '', 0, '2024-02-05T16:07:45', '2024-12-18T19:02:08', NULL);
INSERT INTO createcode_product (id, account_id, business_id, cid, title, num, price, sex, des, content, "likeColor", "ballType", "userType", image, images, moreimgs, contenttow, record_audio, file, workerway, status, createtime, updatetime, deletetime) VALUES (8, 1, 1, 1, '单图数据', 22, 12.86, 1, '', '<div data-element="root" class="am-engine"><p data-id="p838747a-mSXI1S4K"><span style="font-size: 15px;"><span style="font-family: &quot;Microsoft YaHei&quot;, 微软雅黑, &quot;PingFang SC&quot;, SimHei, STHeiti, sans-serif;">GoFly快速开发版本满足中小企业的项目，golang语言的稳定和性能优异下，部署再配置不错的云服务器上，您不必担心项目性能、稳定、安全等情况，Golang作为一种<strong>高效</strong>、<strong>安全</strong>的编程语言，完全可以帮助开发者快速构建</span><span style="font-family: &quot;Microsoft YaHei&quot;, 微软雅黑, &quot;PingFang SC&quot;, SimHei, STHeiti, sans-serif;"><strong>高效、可靠、安全的应用。GoFly基于Go特性完全满足了您快速开发（像PHP感觉一样开发项目），并</strong></span></span></p></div>', '#D91AD9', 'zq', 'salesman', 'resource/uploads/20241208/cf78412be9a1bbb950b4409b4f512d0d.jpg', 'resource/uploads/20241208/67f3ef0660e5cd5cec70cfe88cb8d011.jpg,resource/uploads/20241208/d737f7a009fbbf348630202550ee63ac.jpg,resource/uploads/20241208/2f9365aea081be00daf02b9771724426.jpg', '', '<div data-v-71ecfed1="" data-element="root" class="am-engine"><div data-type="video" data-value="data:%7B%22status%22%3A%22done%22%2C%22name%22%3A%22mov_bbb.mp4%22%2C%22size%22%3A788493%2C%22id%22%3A%22afLuW%22%2C%22type%22%3A%22block%22%2C%22percent%22%3A100%2C%22video_id%22%3A216%2C%22url%22%3A%22resource%2Fuploads%2F20241129%2Fb912ee35cf385da12a70842661fe6459.mp4%22%2C%22cover%22%3A%22resource%2Fuploads%2F20241129%2Fb912ee35cf385da12a70842661fe6459.png%22%2C%22naturalWidth%22%3A320%2C%22naturalHeight%22%3A176%2C%22width%22%3A413%2C%22height%22%3A227%7D"><video controls="" src="http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20241129/b912ee35cf385da12a70842661fe6459.mp4" poster="http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20241129/b912ee35cf385da12a70842661fe6459.png" webkit-playsinline="webkit-playsinline" playsinline="playsinline" style="outline:none;"></video></div></div>', '', '', '', 0, '2024-03-12T11:11:52', '2024-12-26T12:11:14', NULL);
INSERT INTO createcode_product (id, account_id, business_id, cid, title, num, price, sex, des, content, "likeColor", "ballType", "userType", image, images, moreimgs, contenttow, record_audio, file, workerway, status, createtime, updatetime, deletetime) VALUES (9, 1, 1, 1, '多图数据', 0, 0.00, 2, '如果公众号基于安全等考虑，需要获知微信服务器的IP地址列表，以便进行相关限制，可以通过该接口获得微信服务器IP地址列表或者IP网段信息。', '<div data-element="root" class="am-engine"><p data-id="p838747a-ViU0CHhf">你的验证码为：{code}2</p><p data-id="p838747a-WlJ31ZcZ"><br></p></div>', '#3C7EFF', 'zq', 'mteam', 'resource/uploads/20240930/0f8374d9779353d0183005361c50aecb.jpg', 'resource/uploads/20241208/23d1ffc7a3c9e9705e5e095ada383dde.jpg,resource/uploads/20241208/eb21c3b2fe2742e5c94230b2a7d5b21d.jpg', '', '<div data-v-71ecfed1="" data-element="root" class="am-engine"><p data-id="p838747a-APqF60Yo">dddd222</p><p data-id="p838747a-jWeNV0da"><img src="http://localhost:8200/common/uploadfile/getfile?url=resource/uploads/20241208/cf78412be9a1bbb950b4409b4f512d0d.jpg" style="width: 802px; height: 802px; visibility: visible;" data-type="inline"></p></div>', 'resource/uploads/20240725/a27486f670a9f08ae6c73174f42b16a6.mp3', 'resource/uploads/20240725/e2f5778a15ecb45c17a856513a819c09.txt', 'car,   bus', 1, '2024-03-12T12:16:15', '2024-12-26T22:45:07', NULL);
INSERT INTO createcode_product (id, account_id, business_id, cid, title, num, price, sex, des, content, "likeColor", "ballType", "userType", image, images, moreimgs, contenttow, record_audio, file, workerway, status, createtime, updatetime, deletetime) VALUES (10, 0, 1, 1, '小娃', 0, 0.00, 0, '', '', '', '', 'salesman', '', '', '', '', '', '', '', 0, '2024-12-18T17:20:51', '2024-12-18T17:20:51', '2024-12-18T17:25:30');
INSERT INTO createcode_product (id, account_id, business_id, cid, title, num, price, sex, des, content, "likeColor", "ballType", "userType", image, images, moreimgs, contenttow, record_audio, file, workerway, status, createtime, updatetime, deletetime) VALUES (11, 0, 1, 1, '测试', 0, 12.00, 0, '', '', '', '', 'salesman', '', '', '', '', '', '', '', 0, '2024-12-18T17:25:27', '2024-12-18T17:25:27', '2024-12-18T17:25:31');

-- 删除表 createcode_product_cate
DROP TABLE IF EXISTS createcode_product_cate CASCADE;

-- 创建表 createcode_product_cate
CREATE TABLE createcode_product_cate (
id bigint NOT NULL,
    account_id integer NOT NULL,
    business_id integer NOT NULL,
    pid integer NOT NULL,
    name character varying(50) NOT NULL,
    remark character varying(255) NOT NULL,
    weigh integer NOT NULL,
    status smallint NOT NULL,
    createtime timestamp without time zone,
    deletetime timestamp without time zone,
    PRIMARY KEY (id)
);

-- 插入表 createcode_product_cate 数据
INSERT INTO createcode_product_cate (id, account_id, business_id, pid, name, remark, weigh, status, createtime, deletetime) VALUES (1, 1, 1, 0, '测试分类', '', 0, 0, '2024-02-05T15:35:59', NULL);
INSERT INTO createcode_product_cate (id, account_id, business_id, pid, name, remark, weigh, status, createtime, deletetime) VALUES (4, 1, 1, 0, '水果', '', 4, 0, '2024-03-22T19:20:25', NULL);
INSERT INTO createcode_product_cate (id, account_id, business_id, pid, name, remark, weigh, status, createtime, deletetime) VALUES (5, 1, 1, 0, '素菜', '', 5, 0, '2024-03-22T19:21:09', NULL);
INSERT INTO createcode_product_cate (id, account_id, business_id, pid, name, remark, weigh, status, createtime, deletetime) VALUES (6, 1, 1, 4, '苹果12', '', 6, 0, '2024-03-22T19:23:29', NULL);
INSERT INTO createcode_product_cate (id, account_id, business_id, pid, name, remark, weigh, status, createtime, deletetime) VALUES (7, 1, 1, 4, '芒果', '', 7, 0, '2024-05-19T14:56:25', NULL);
