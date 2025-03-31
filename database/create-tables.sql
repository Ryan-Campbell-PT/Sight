DROP TABLE IF EXISTS daily;
CREATE TABLE daily (
  id             INT AUTO_INCREMENT NOT NULL,
  foodString     VARCHAR(255) NOT NULL,
  date           VARCHAR(64) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO daily
  (foodString, date)
VALUES
    ("medium cheese pizza, 1 pound of green beans", "January 1, 2025"),
    ("3 cookies, banana, 2 skinless chicken thighs", "1/1/2025");