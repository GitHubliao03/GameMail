--服务器存放的所有邮件
create table `game_email`.`mail`
(
    `id`          int           NOT NULL AUTO_INCREMENT COMMENT '邮件id',
    `content`     varchar(1024) NOT NULL COMMENT '邮件的内容',
    `send_time`   datetime      NOT NULL COMMENT '邮件发送的时间',
    `clear_time`  datetime      NOT NULL COMMENT '被清理的时间',
    `filter`      varchar(40)   NOT NULL COMMENT '过滤条件',
    `is_del`      int(11) NOT NULL DEFAULT 0 COMMENT '0未删除，1已删除',
    `create_time` datetime      NOT NULL COMMENT '创建的时间',
    `update_time` datetime      NOT NULL COMMENT '更新的时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;