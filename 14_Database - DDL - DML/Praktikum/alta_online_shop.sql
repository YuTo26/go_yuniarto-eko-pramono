-- MySQL Script generated by MySQL Workbench
-- Sun Mar 19 10:01:03 2023
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema alterra_online_shop
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema alterra_online_shop
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `alterra_online_shop` DEFAULT CHARACTER SET utf8 ;
USE `alterra_online_shop` ;

-- -----------------------------------------------------
-- Table `alterra_online_shop`.`alamat`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`alamat` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`alamat` (
  `alamat_id` INT NOT NULL AUTO_INCREMENT,
  `customer_id` INT NOT NULL,
  `nama_penerima` VARCHAR(100) NOT NULL,
  `alamat` TEXT NOT NULL,
  `telp` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`alamat_id`),
  CONSTRAINT `fk_alamat_customers`
    FOREIGN KEY (`customer_id`)
    REFERENCES `alterra_online_shop`.`customers` (`customer_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_alamat_customers1_idx` ON `alterra_online_shop`.`alamat` (`customer_id` ASC) VISIBLE;

CREATE INDEX `fk_alamat_customers` ON `alterra_online_shop`.`alamat` (`customer_id` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`customers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`customers` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`customers` (
  `customer_id` INT NOT NULL,
  `username` VARCHAR(255) NOT NULL,
  `full_neme` VARCHAR(255) NOT NULL,
  `email_address` VARCHAR(100) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `address` TEXT NOT NULL,
  `phone_number` VARCHAR(20) NOT NULL,
  `birthdate` DATE NULL DEFAULT NULL,
  `gender` VARCHAR(45) NULL DEFAULT NULL,
  `status_user` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`customer_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`customers_payment_method_detail`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`customers_payment_method_detail` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`customers_payment_method_detail` (
  `customers_payment_method_detail_id` INT NOT NULL AUTO_INCREMENT,
  `customer_id` INT NOT NULL,
  `payment_method_id` INT NOT NULL,
  PRIMARY KEY (`customers_payment_method_detail_id`),
  CONSTRAINT `fk_customers_payment_method_detail_customers`
    FOREIGN KEY (`customer_id`)
    REFERENCES `alterra_online_shop`.`customers` (`customer_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_customers_payment_method_detail_payment_method`
    FOREIGN KEY (`payment_method_id`)
    REFERENCES `alterra_online_shop`.`payment_method` (`payment_method_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_customers_payment_method_detail_customers_idx` ON `alterra_online_shop`.`customers_payment_method_detail` (`customer_id` ASC) VISIBLE;

CREATE INDEX `fk_customers_payment_method_detail_payment_method1_idx` ON `alterra_online_shop`.`customers_payment_method_detail` (`payment_method_id` ASC) VISIBLE;

CREATE INDEX `fk_customers_payment_method_detail_customers` ON `alterra_online_shop`.`customers_payment_method_detail` (`customer_id` ASC) VISIBLE;

CREATE INDEX `fk_customers_payment_method_detail_payment_method` ON `alterra_online_shop`.`customers_payment_method_detail` (`payment_method_id` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`kurir`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`kurir` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`kurir` (
  `kurir_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`kurir_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`payment_method`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`payment_method` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`payment_method` (
  `payment_method_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`payment_method_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`payment_method_description`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`payment_method_description` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`payment_method_description` (
  `payment_method_description_id` INT NOT NULL AUTO_INCREMENT,
  `payment_method_id` INT NOT NULL,
  `description` TEXT NOT NULL,
  PRIMARY KEY (`payment_method_description_id`),
  CONSTRAINT `fk_payment_method_description_payment_method`
    FOREIGN KEY (`payment_method_id`)
    REFERENCES `alterra_online_shop`.`payment_method` (`payment_method_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_payment_method_description_payment_method1_idx` ON `alterra_online_shop`.`payment_method_description` (`payment_method_id` ASC) VISIBLE;

CREATE INDEX `fk_payment_method_description_payment_method` ON `alterra_online_shop`.`payment_method_description` (`payment_method_id` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`payment_method_detail`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`payment_method_detail` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`payment_method_detail` (
  `payment_method_detail_id` INT NOT NULL AUTO_INCREMENT,
  `payment_method_id` INT NOT NULL,
  `credit_card_number` VARCHAR(255) NOT NULL,
  `exp_date` DATE NOT NULL,
  `cardholder_name` VARCHAR(255) NOT NULL,
  `billing_address` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`payment_method_detail_id`),
  CONSTRAINT `fk_payment_method_detail_payment_method`
    FOREIGN KEY (`payment_method_id`)
    REFERENCES `alterra_online_shop`.`payment_method` (`payment_method_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_payment_method_detail_payment_method1_idx` ON `alterra_online_shop`.`payment_method_detail` (`payment_method_id` ASC) VISIBLE;

CREATE INDEX `fk_payment_method_detail_payment_method` ON `alterra_online_shop`.`payment_method_detail` (`payment_method_id` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`product`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`product` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`product` (
  `product_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `product_description` INT NOT NULL,
  `product_type` INT NOT NULL,
  `operator` INT NOT NULL,
  `status` SMALLINT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`product_id`),
  CONSTRAINT `fk_product_product_type`
    FOREIGN KEY (`product_type`)
    REFERENCES `alterra_online_shop`.`product_type` (`product_type_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_product_product_operator`
    FOREIGN KEY (`operator`)
    REFERENCES `alterra_online_shop`.`product_operator` (`product_operator_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_product_product_description`
    FOREIGN KEY (`product_description`)
    REFERENCES `alterra_online_shop`.`product_description` (`product_description_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE INDEX `fk_product_product_type_idx` ON `alterra_online_shop`.`product` (`product_type` ASC) VISIBLE;

CREATE INDEX `fk_product_product_operator_idx` ON `alterra_online_shop`.`product` (`operator` ASC) VISIBLE;

CREATE INDEX `fk_product_product_description_idx` ON `alterra_online_shop`.`product` (`product_description` ASC) VISIBLE;

CREATE INDEX `fk_product_product_type` ON `alterra_online_shop`.`product` (`product_type` ASC) VISIBLE;

CREATE INDEX `fk_product_product_operator` ON `alterra_online_shop`.`product` (`operator` ASC) VISIBLE;

CREATE INDEX `fk_product_product_description` ON `alterra_online_shop`.`product` (`product_description` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`product_description`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`product_description` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`product_description` (
  `product_description_id` INT NOT NULL AUTO_INCREMENT,
  `description` TEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`product_description_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`product_operator`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`product_operator` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`product_operator` (
  `product_operator_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`product_operator_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`product_type`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`product_type` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`product_type` (
  `product_type_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`product_type_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`transaction_details`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`transaction_details` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`transaction_details` (
  `transaction_detail_id` INT NOT NULL AUTO_INCREMENT,
  `transaction_id` INT NOT NULL,
  `product_id` INT NOT NULL,
  `quantity` INT NOT NULL,
  `price` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaction_detail_id`),
  CONSTRAINT `fk_transaction_details_product`
    FOREIGN KEY (`product_id`)
    REFERENCES `alterra_online_shop`.`product` (`product_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_transaction_details_transactions`
    FOREIGN KEY (`transaction_id`)
    REFERENCES `alterra_online_shop`.`transactions` (`transaction_id`)
    ON DELETE NO ACTION
    ON UPDATE CASCADE)
ENGINE = InnoDB;

CREATE INDEX `fk_transaction_details_transactions1_idx` ON `alterra_online_shop`.`transaction_details` (`transaction_id` ASC) VISIBLE;

CREATE INDEX `fk_transaction_details_product_idx` ON `alterra_online_shop`.`transaction_details` (`product_id` ASC) VISIBLE;

CREATE INDEX `fk_transaction_details_product` ON `alterra_online_shop`.`transaction_details` (`product_id` ASC) VISIBLE;

CREATE INDEX `fk_transaction_details_transactions` ON `alterra_online_shop`.`transaction_details` (`transaction_id` ASC) VISIBLE;


-- -----------------------------------------------------
-- Table `alterra_online_shop`.`transactions`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alterra_online_shop`.`transactions` ;

CREATE TABLE IF NOT EXISTS `alterra_online_shop`.`transactions` (
  `transaction_id` INT NOT NULL AUTO_INCREMENT,
  `customer_id` INT NOT NULL,
  `payment_method_id` INT NOT NULL,
  `total_price` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaction_id`),
  CONSTRAINT `fk_transactions_customers`
    FOREIGN KEY (`customer_id`)
    REFERENCES `alterra_online_shop`.`customers` (`customer_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_transactions_payment_method`
    FOREIGN KEY (`payment_method_id`)
    REFERENCES `alterra_online_shop`.`payment_method` (`payment_method_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_transactions_customers1_idx` ON `alterra_online_shop`.`transactions` (`customer_id` ASC) VISIBLE;

CREATE INDEX `fk_transactions_payment_method_idx` ON `alterra_online_shop`.`transactions` (`payment_method_id` ASC) VISIBLE;

CREATE INDEX `fk_transactions_customers` ON `alterra_online_shop`.`transactions` (`customer_id` ASC) VISIBLE;

CREATE INDEX `fk_transactions_payment_method` ON `alterra_online_shop`.`transactions` (`payment_method_id` ASC) VISIBLE;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
