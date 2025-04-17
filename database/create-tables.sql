SET foreign_key_checks = 0;

DROP TABLE IF EXISTS nutrition_info;
-- i think you may want to add additional values to this, like unsaturated fats. but for now this is fine
CREATE TABLE nutrition_info (
  id              INT AUTO_INCREMENT NOT NULL,
  calories        INT NOT NULL,
  total_fat       INT,
  saturated_fat   INT,
  poly_fat        INT,
  mono_fat        INT,
  cholesterol     INT,
  sodium          INT,
  carbs           INT,
  fiber           INT,
  sugar           INT,
  protein         INT,
  potassium       INT,
  phosphorus      INT,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS daily;
CREATE TABLE daily (
  id             INT AUTO_INCREMENT NOT NULL,
  food_string    VARCHAR(255) NOT NULL,
  date           VARCHAR(64) NOT NULL,
  nutrition_id   INT NOT NULL,
  FOREIGN KEY (nutrition_id) REFERENCES nutrition_info(id),
  INDEX (nutrition_id),
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS recipe;
CREATE TABLE recipe (
  id              INT AUTO_INCREMENT NOT NULL,
  recipe_name     VARCHAR(255) NOT NULL,
  food_string     VARCHAR(255) NOT NULL,
  serving_size    INT NOT NULL,
  nutrition_id    INT NOT NULL,
  FOREIGN KEY (nutrition_id) REFERENCES nutrition_info(id),
  INDEX (nutrition_id),
  PRIMARY KEY (id)
);

SET foreign_key_checks = 1;