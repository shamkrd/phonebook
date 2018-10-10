--
-- Скрипт сгенерирован Devart dbForge Studio for MySQL, Версия 7.2.63.0
-- Домашняя страница продукта: http://www.devart.com/ru/dbforge/mysql/studio
-- Дата скрипта: 09.10.2018 22:41:53
-- Версия сервера: 5.7.18-log
-- Версия клиента: 4.1
--


-- 
-- Отключение внешних ключей
-- 
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;

-- 
-- Установить режим SQL (SQL mode)
-- 
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 
-- Установка кодировки, с использованием которой клиент будет посылать запросы на сервер
--
SET NAMES 'utf8';

-- 
-- Установка базы данных по умолчанию
--
USE phonebook;

--
-- Описание для таблицы contacts
--
DROP TABLE IF EXISTS contacts;
CREATE TABLE contacts (
  id BIGINT(20) NOT NULL AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  PRIMARY KEY (id)
)
ENGINE = INNODB
AUTO_INCREMENT = 30
AVG_ROW_LENGTH = 5461
CHARACTER SET utf8
COLLATE utf8_general_ci
ROW_FORMAT = DYNAMIC;

--
-- Описание для таблицы phones
--
DROP TABLE IF EXISTS phones;
CREATE TABLE phones (
  id BIGINT(20) NOT NULL AUTO_INCREMENT,
  phone VARCHAR(11) NOT NULL,
  contact_id BIGINT(20) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT FK_phones_contact_id2 FOREIGN KEY (contact_id)
    REFERENCES contacts(id) ON DELETE CASCADE ON UPDATE CASCADE
)
ENGINE = INNODB
AUTO_INCREMENT = 15
AVG_ROW_LENGTH = 2340
CHARACTER SET utf8
COLLATE utf8_general_ci
ROW_FORMAT = DYNAMIC;

-- 
-- Вывод данных для таблицы contacts
--
INSERT INTO contacts VALUES
(2, 'Петров Олег'),
(23, 'Александр'),
(24, 'Александр 1');

-- 
-- Вывод данных для таблицы phones
--
INSERT INTO phones VALUES
(4, '324343', 2),
(9, '33333333', 2),
(10, '44444444', 2),
(11, '4545454', 23),
(12, '545445454', 23),
(13, '432234234', 24),
(14, '424323423', 24);

-- 
-- Восстановить предыдущий режим SQL (SQL mode)
-- 
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;

-- 
-- Включение внешних ключей
-- 
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;