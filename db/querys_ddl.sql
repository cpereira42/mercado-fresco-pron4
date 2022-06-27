-- MySQL Workbench Synchronization
-- Generated: 2022-06-22 19:53
-- Model: New Model
-- Version: 1.0
-- Project: Name of the project
-- Author: Eneas Da Silva Sena

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE DATABASE `mercadofresco`;

ALTER SCHEMA `mercadofresco`  DEFAULT CHARACTER SET utf8  DEFAULT COLLATE utf8_general_ci ;

CREATE TABLE IF NOT EXISTS `mercadofresco`.`sellers` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `cid` INT(11) NOT NULL,
  `company_name` VARCHAR(100) NOT NULL,
  `address` VARCHAR(50) NOT NULL,
  `telephone` CHAR(8) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofresco`.`warehouses` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `address` VARCHAR(50) NOT NULL,
  `telephone` VARCHAR(20) NOT NULL,
  `warehouse_code` VARCHAR(100) NOT NULL,
  `minimum_capacity` INT(11) NOT NULL,
  `minimum_temperature` INT(11) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofresco`.`sections` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `section_number` INT(11) NOT NULL,
  `current_temperature` INT(11) NOT NULL,
  `minimum_temperature` INT(11) NOT NULL,
  `current_capacity` INT(11) NOT NULL,
  `minimum_capacity` INT(11) NOT NULL,
  `maximum_capacity` INT(11) NOT NULL,
  `warehouse_id` INT(11) NOT NULL,
  `product_type_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk2_warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  CONSTRAINT `fk2_warehouse_id`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouses` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofresco`.`products` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `product_code` VARCHAR(45) NOT NULL,
  `description` VARCHAR(45) NOT NULL,
  `width` DECIMAL(2,1) NOT NULL,
  `height` DECIMAL(2,1) NOT NULL,
  `length` DECIMAL(2,1) NOT NULL,
  `net_weight` DECIMAL(2,1) NOT NULL,
  `expiration_rate` INT(11) NOT NULL,
  `recommended_freezing_temperature` DECIMAL(2,1) NOT NULL,
  `freezing_rate` INT(11) NOT NULL,
  `product_type_id` INT(11) NOT NULL,
  `seller_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofresco`.`employees` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `card_number_id` INT(11) NOT NULL,
  `first_name` VARCHAR(50) NOT NULL,
  `last_name` VARCHAR(50) NOT NULL,
  `warehouse_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk1_warehouse_id_idx` (`warehouse_id` ASC) VISIBLE,
  CONSTRAINT `fk1_warehouse_id`
    FOREIGN KEY (`warehouse_id`)
    REFERENCES `mercadofresco`.`warehouses` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofresco`.`buyers` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `card_number_id` INT(11) NOT NULL,
  `first_name` VARCHAR(45) NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
