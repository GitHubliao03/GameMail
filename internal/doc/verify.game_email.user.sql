-- 用户信息表
create table `game_email`.`user`
(
    `id`               int         NOT NULL auto_increment,
    `level`            int(11) NOT NULL COMMENT '游戏等级',
    `name`             varchar(20) NOT NULL COMMENT '游戏昵称',
    `last_active_time` datatime    NOT NULL COMMENT '上次活跃时间',
    `status`           int(11) NOT NULL DEFAULT 0 COMMENT '0上线，1下线',
    `identity`         int(11) NOT NULL DEFAULT 0 COMMENT '0普通玩家，1白名单，2黑名单',
    `is_del`           int(11) NOT NULL DEFAULT 0 COMMENT '0未删除，1已删除',
    `create_time`      datetime    NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    index(`level`),
    index (`is_del`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

