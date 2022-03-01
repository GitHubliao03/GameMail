--邮箱
create table `game_email`.`mailbox`
(
    `uid`         int(11) NOT NULL COMMENT '用户id',
    `eid`         int(11) NOT NULL COMMENT '邮件id',
    `status`      int(11) NOT NULL DEFAULT 0 COMMENT '0未读，1已读未领取，2已读已领取',
    `is_del`      int(11) NOT NULL DEFAULT 0 COMMENT '0未删除，1已删除',
    `create_time` datetime NOT NULL COMMENT '创建时间',
    `update_time` datetime NOT NULL COMMENT '更新时间',
    PRIMARY key (`uid`, `eid`),
    index(`status`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;