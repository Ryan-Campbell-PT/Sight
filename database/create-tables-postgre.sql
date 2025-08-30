DROP TABLE IF EXISTS food_records;
DROP TABLE IF EXISTS recipe;
DROP TABLE IF EXISTS daily;
DROP TABLE IF EXISTS nutrition_info;

CREATE TABLE nutrition_info
(
    id SERIAL PRIMARY KEY,
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
    id SERIAL PRIMARY KEY,
    food_string VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    nutrition_id INT NOT NULL REFERENCES nutrition_info(id)
);

-- Add index on nutrition_id in daily
CREATE INDEX daily_nutrition_idx ON daily(nutrition_id);

CREATE TABLE recipe
(
    id SERIAL PRIMARY KEY,
    recipe_name VARCHAR(255) NOT NULL,
    food_string VARCHAR(255) NOT NULL,
    serving_size INT NOT NULL,
    active BOOLEAN NOT NULL,
    nutrition_id INT NOT NULL REFERENCES nutrition_info(id)
);

-- Add index on nutrition_id in recipe
CREATE INDEX recipe_nutrition_idx ON recipe(nutrition_id);

CREATE TABLE food_records
(
    id SERIAL PRIMARY KEY,
    food VARCHAR(64) NOT NULL,
    count INT NOT NULL
);

INSERT INTO nutrition_info
    (
    calories, total_fat, saturated_fat, poly_fat, mono_fat,
    cholesterol, sodium, carbs, fiber, sugar, protein,
    potassium, phosphorus
    )
VALUES
    (550, 20, 8, 4, 6, 60, 800, 60, 5, 10, 30, 400, 300),
    (400, 15, 4, 3, 5, 50, 600, 25, 4, 5, 35, 350, 250),
    (300, 5, 1, 2, 1, 10, 400, 40, 6, 8, 10, 300, 200),
    (600, 22, 9, 5, 6, 70, 850, 65, 4, 12, 32, 420, 310),
    (420, 16, 5, 3, 5, 55, 620, 27, 5, 6, 36, 360, 270);

INSERT INTO recipe
    (
    recipe_name, food_string, serving_size, active, nutrition_id
    )
VALUES
    ('Spaghetti Bolognese', 'spaghetti, ground beef, tomato sauce, onion, garlic', 2, TRUE, 1),
    ('Grilled Chicken Salad', 'chicken breast, lettuce, tomato, cucumber, vinaigrette', 1, TRUE, 2),
    ('Vegetable Stir Fry', 'broccoli, bell pepper, carrot, soy sauce, ginger', 2, FALSE, 3);
