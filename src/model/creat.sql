CREATE DATABASE IF NOT EXISTS `grayRelease`;

USE `grayRelease`;

CREATE TABLE IF NOT EXISTS `rules` (
    id INT AUTO_INCREMENT primary key comment '自增id，与业务无关',
    aid INT COMMENT 'app的唯一标识，逻辑主键',
    platform VARCHAR(20) COMMENT '操作系统，Android | Apple',
    download_url VARCHAR(128) COMMENT '下载链接',
    update_version_code VARCHAR(128) COMMENT '版本号',
    md5 VARCHAR(128) COMMENT 'md5值',
    device_id_list TEXT COMMENT '白名单',
    max_update_version_code VARCHAR(128) COMMENT '最大版本号',
    min_update_version_code VARCHAR(128) COMMENT '最小版本号',
    max_os_api INT COMMENT '最大操作系统版本',
    min_os_api INT COMMENT '最小操作系统版本',
    cpu_arch VARCHAR(32) COMMENT 'cpu 架构',
    channel VARCHAR(128) COMMENT '渠道',
    title VARCHAR(256) COMMENT '弹窗标题',
    update_tips VARCHAR(1024) COMMENT '弹窗的更新文本',
    create_time DATETIME COMMENT '创建时间',
    delete_time DATETIME COMMENT '删除时间',
    is_delete INT COMMENT '规则是否启用（1：删除，0：未删除）',
    is_release INT COMMENT '是否发布上线（1：上线，0：下线）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
