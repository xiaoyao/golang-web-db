CREATE DATABASE qiniu_bank DEFAULT CHARSET 'utf8' COLLATE 'utf8_general_ci';

CREATE TABLE `account` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '储户账号',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '储户姓名',
  `gender` char(1) NOT NULL DEFAULT '' COMMENT '储户性别',
  `age` int(11) NOT NULL COMMENT '储户年龄',
  `balance` int(11) NOT NULL COMMENT '储户余额',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `transfer` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '交易流水号',
  `from_account_id` int(10) unsigned NOT NULL COMMENT '转账人id',
  `to_account_id` int(10) unsigned NOT NULL COMMENT '收款人id',
  `amount` int(11) NOT NULL COMMENT '转账金额',
  `created_at` datetime NOT NULL COMMENT '转账时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
