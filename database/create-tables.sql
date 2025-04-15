DROP TABLE IF EXISTS Daily;
CREATE TABLE Daily (
  id             INT AUTO_INCREMENT NOT NULL,
  foodString     VARCHAR(255) NOT NULL,
  date           VARCHAR(64) NOT NULL,
  PRIMARY KEY (`id`)
);