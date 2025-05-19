IF OBJECT_ID('daily', 'U') IS NOT NULL DROP TABLE daily;
IF OBJECT_ID('recipe', 'U') IS NOT NULL DROP TABLE recipe;
IF OBJECT_ID('food_records', 'U') IS NOT NULL DROP TABLE food_records;
IF OBJECT_ID('nutrition_info', 'U') IS NOT NULL DROP TABLE nutrition_info;

CREATE TABLE nutrition_info
(
  id INT IDENTITY(1,1) PRIMARY KEY,
  calories INT NOT NULL,
  total_fat INT,
  saturated_fat INT,
  poly_fat INT,
  mono_fat INT,
  cholesterol INT,
  sodium INT,
  carbs INT,
  fiber INT,
  sugar INT,
  protein INT,
  potassium INT,
  phosphorus INT
);

CREATE TABLE daily
(
  id INT IDENTITY(1,1) PRIMARY KEY,
  food_string VARCHAR(255) NOT NULL,
  date DATE NOT NULL,
  nutrition_id INT NOT NULL,
  FOREIGN KEY (nutrition_id) REFERENCES nutrition_info(id)
);

CREATE TABLE custom_recipe
(
  id INT IDENTITY(1,1) PRIMARY KEY,
  recipe_name VARCHAR(255) NOT NULL,
  alt_recipe_names VARCHAR(512),
  food_string VARCHAR(255) NOT NULL,
  serving_size INT NOT NULL,
  active BIT NOT NULL,
  nutrition_id INT NOT NULL,
  last_modified DATE NOT NULL,
  FOREIGN KEY (nutrition_id) REFERENCES nutrition_info(id)
);

CREATE TABLE food_records
(
  id INT IDENTITY(1,1) PRIMARY KEY,
  food VARCHAR(64) NOT NULL,
  count INT NOT NULL
);
