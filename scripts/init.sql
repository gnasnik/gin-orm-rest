DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` longtext,
  `username` longtext,
  `pass_hash` longtext,
  `user_email` longtext,
  `address` longtext,
  `role` tinyint(4) NULL DEFAULT 0,
  `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime(3) DEFAULT ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;                                                                                                                   )


DROP TABLE IF EXISTS `login_log`;
CREATE TABLE `login_log`  (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`login_username` varchar(50) NULL DEFAULT '',
`ipaddr` varchar(50)  NULL DEFAULT '',
`login_location` varchar(255)  NULL DEFAULT '',
`browser` varchar(50)  NULL DEFAULT '',
`os` varchar(50)  NULL DEFAULT '',
`status` tinyint(4) NULL DEFAULT 0,
`msg` varchar(255)  NULL DEFAULT '',
`created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log`  (
 `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
 `title` varchar(50)  NULL DEFAULT '',
 `business_type` int(2) NULL DEFAULT 0 ,
 `method` varchar(100)  NULL DEFAULT '',
 `request_method` varchar(10)  NULL DEFAULT '',
 `operator_type` int(1) NULL DEFAULT 0,
 `operator_username` varchar(50)  NULL DEFAULT '',
 `operator_url` varchar(500)  NULL DEFAULT '',
 `operator_ip` varchar(50)  NULL DEFAULT '',
 `operator_location` varchar(255)  NULL DEFAULT '',
 `operator_param` text  NULL,
 `json_result` text  NULL,
 `status` int(1) NULL DEFAULT 0,
 `error_msg` varchar(2000)  NULL DEFAULT '',
 `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP,
 `updated_at` datetime(3) DEFAULT ON UPDATE CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4;


DROP TABLE IF EXISTS `schedulers`;
CREATE TABLE `schedulers` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`name` longtext,
`group` longtext,
`address` longtext,
`status` int(1) NULL DEFAULT 0,
`created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime(3) DEFAULT ON UPDATE CURRENT_TIMESTAMP,
`deleted_at` datetime(3) DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;