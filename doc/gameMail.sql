create database game_mail;

use game_mail;

create table Title(
      tid int primary key auto_increment comment '标题ID',
      title varchar(30) not null default '' comment '标题文本'
)

create table Content(
        cid int primary key auto_increment comment '文本ID',
        content varchar(30) not null default '' comment '内容文本'
)

create table AllMails(
     mail_id int primary key auto_increment comment '全体邮件自增长主键',
     title_id int not null comment '标题ID',
     content_id int not null comment '文本ID',
     prop_id int comment '道具',
     send_date datetime  not null default '1000-01-01 00:00:00' comment '邮件发出日期',
     expiry_date datetime not null default '1000-01-01 00:00:00' comment '邮件过期日期'
);

alter table AllMails add constraint title_fk foreign key(title_id) references Title (tid);
alter table AllMails add constraint content_fk foreign key(content_id) references Content (cid);