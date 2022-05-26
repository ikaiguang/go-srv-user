-- 用户
CREATE TABLE ikg_user
(
    id             BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID：雪花算法',
    user_nickname  VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户昵称',
    user_mobile    VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户手机号',
    user_email     VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户邮箱',
    user_avatar    VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户头像',
    user_sex       TINYINT(1)   NOT NULL DEFAULT '0' COMMENT '用户性别；MALE、FEMALE、SECRET',
    user_status    TINYINT(1)   NOT NULL DEFAULT '0' COMMENT '用户状态；INACTIVATED、ENABLE、DISABLE、BLACKLIST、DELETED',
    enable_time    DATETIME(3)  NULL COMMENT '激活时间',
    disable_time   DATETIME(3)  NULL COMMENT '禁用时间',
    blacklist_time DATETIME(3)  NULL COMMENT '黑名单时间',
    deleted_time   DATETIME(3)  NULL COMMENT '删除时间',
    created_time   DATETIME(3)  NOT NULL COMMENT '创建时间',
    updated_time   DATETIME(3)  NOT NULL COMMENT '最后修改时间',
    PRIMARY KEY (id)
) ENGINE InnoDB
  DEFAULT CHARSET utf8mb4
    COMMENT '用户';

-- 用户-注册-用户名
CREATE TABLE ikg_user_reg_username
(
    id           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID',
    user_id      BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
    username     VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '用户名',
    created_time DATETIME(3)     NOT NULL COMMENT '创建时间',
    updated_time DATETIME(3)     NOT NULL COMMENT '最后修改时间',
    PRIMARY KEY (id),
    UNIQUE KEY username (username),
    KEY user_id (user_id)
) ENGINE InnoDB
  DEFAULT CHARSET utf8mb4
    COMMENT '用户-注册-用户名';

-- 用户-注册-手机号
CREATE TABLE ikg_user_reg_mobile
(
    id           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID',
    user_id      BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
    user_mobile  VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '用户手机号码',
    created_time DATETIME(3)     NOT NULL COMMENT '创建时间',
    updated_time DATETIME(3)     NOT NULL COMMENT '最后修改时间',
    PRIMARY KEY (id),
    UNIQUE KEY user_mobile (user_mobile),
    KEY user_id (user_id)
) ENGINE InnoDB
  DEFAULT CHARSET utf8mb4
    COMMENT '用户-注册-手机号';

-- 用户-注册-邮箱
CREATE TABLE ikg_user_reg_email
(
    id           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID',
    user_id      BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
    user_email   VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '用户手机号码',
    created_time DATETIME(3)     NOT NULL COMMENT '创建时间',
    updated_time DATETIME(3)     NOT NULL COMMENT '最后修改时间',
    PRIMARY KEY (id),
    UNIQUE KEY user_email (user_email),
    KEY user_id (user_id)
) ENGINE InnoDB
  DEFAULT CHARSET utf8mb4
    COMMENT '用户-注册-邮箱';
