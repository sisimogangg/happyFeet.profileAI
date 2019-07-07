CREATE DATABASE IF NOT EXISTS `profile`
USE `profile`

CREATE TABLE `profile` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
    `sgb_member` int(1) DEFAULT '0',

    PRIMARY KEY (`id`)

)ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



LOCK TABLES `profile` WRITE;
INSERT INTO `profile` values(1001,'c016d1dd-4eb2-42fe-80ab-019aa52bca3b', 0);
INSERT INTO `profile` values(1005,'c3f0ad14-ceac-4cff-ba15-8e9f18e950bf', 1);
INSERT INTO `profile` values(1006,'cc365f23-9819-46fd-a24b-7eb44ccc57d1', 0);
INSERT INTO `profile` values(1007,'45920ba5-e85f-41c8-a2d1-f5704e23ed50', 0);
UNLOCK TABLES;

-----------------------------------------------------------------------------------------------

CREATE TABLE `personal_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `surname` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `date_of_birth` datetime DEFAULT NULL,
  `phone_number` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `emergency_number` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `profile_id` int(11) NOT NULL,

   PRIMARY KEY (`id`),

  FOREIGN KEY(`profile_id`) REFERENCES profile(`id`)

)ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


LOCK TABLES `personal_details` WRITE;
INSERT INTO `personal_details` values(1001,'Godfrey', 'Sisimogang', '1990-03-10 00:00:00', '0826955179', '0848861429', 1005);
INSERT INTO `personal_details` values(1002,'Antoinette', 'Mthembu', '1991-02-26 00:00:00', '0848861429', '0826955179' , 1001);
UNLOCK TABLES;

---------------------------------------------------------------------------------------------------
DROP TABLE IF EXISTS `address`;

CREATE TABLE `address` (
  `id` int(11) NOT NULL,
  `street_number` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `street_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `suburb` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `city` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `postal_code` int(11) DEFAULT '0',
  `personal_details_id` int(11) NOT NULL,

  PRIMARY KEY (`id`),

  FOREIGN KEY(`personal_details_id`) REFERENCES personal_details(`id`)

) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



LOCK TABLES `address` WRITE;
INSERT INTO `address` values(1001,'36', 'Lonsdale Drive', 'Durban North', 'Durban', 4051, 1001);
INSERT INTO `address` values(1002,'77', 'Armstrong Avenue', 'La Lucia', 'Durban', 4051, 1002);
INSERT INTO `address` values(1003,'215', 'Peter Mokaba Rd', 'Morningside', 'Durban', 4051, 1001);
UNLOCK TABLES;


-----------------------------------------------------------------------------------------------------------
CREATE TABLE `kid`(
    `id` int(11) NOT NULL,
    `name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
    `surname` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
    `date_of_birth` datetime DEFAULT NULL,
    `enrollment_date` datetime DEFAULT NULL,
    `profile_id` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY(`profile_id`) REFERENCES profile(`id`)

)ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


LOCK TABLES `kid` WRITE;
INSERT INTO `kid` values(1001,'Thubelihle', 'Mthembu','grade R', '2012-01-21 00:00:00', '2016-01-04 00:00:00', 1001);
INSERT INTO `kid` values(1002, 'Thabo', 'Sisimogang','grade R','2012-01-21 00:00:00', '2016-01-04 00:00:00', 1005);
INSERT INTO `kid` values(1003, 'Jaden', 'Sisimogang','toddler','2014-04-23 00:00:00', '2016-01-04 00:00:00',1005);
UNLOCK TABLES;
-----------------------------------------------------------------------------------------------------------
select * from `address`