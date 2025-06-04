-- nutrition_info needs to go first due to foreign key contraints
INSERT INTO nutrition_info
    (calories, total_fat, saturated_fat, poly_fat, mono_fat, cholesterol, sodium, carbs, fiber, sugar, protein, potassium, phosphorus)
VALUES
    (250, 10, 3, 2, 3, 30, 300, 30, 4, 5, 8, 400, 150);

INSERT INTO nutrition_info
    (calories, total_fat, saturated_fat, poly_fat, mono_fat, cholesterol, sodium, carbs, fiber, sugar, protein, potassium, phosphorus)
VALUES
    (180, 7, 2, 1, 2, 15, 150, 20, 3, 7, 5, 320, 120);

INSERT INTO nutrition_info
    (calories, total_fat, saturated_fat, poly_fat, mono_fat, cholesterol, sodium, carbs, fiber, sugar, protein, potassium, phosphorus)
VALUES
    (90, 1, 0, 0, 0, 0, 20, 15, 1, 2, 3, 200, 50);

INSERT INTO nutrition_info
    (calories, total_fat, saturated_fat, poly_fat, mono_fat, cholesterol, sodium, carbs, fiber, sugar, protein, potassium, phosphorus)
VALUES
    (400, 18, 6, 3, 5, 50, 450, 35, 5, 10, 15, 700, 250);

INSERT INTO nutrition_info
    (calories, total_fat, saturated_fat, poly_fat, mono_fat, cholesterol, sodium, carbs, fiber, sugar, protein, potassium, phosphorus)
VALUES
    (120, 5, 1, 1, 2, 5, 100, 10, 2, 3, 6, 280, 90);

INSERT INTO nutrition_info
    (calories, total_fat, saturated_fat, poly_fat, mono_fat, cholesterol, sodium, carbs, fiber, sugar, protein, potassium, phosphorus)
VALUES
    (310, 12, 4, 2, 4, 35, 380, 28, 4, 6, 12, 540, 180);

INSERT INTO nutrition_info
    (calories, total_fat, saturated_fat, poly_fat, mono_fat, cholesterol, sodium, carbs, fiber, sugar, protein, potassium, phosphorus)
VALUES
    (160, 6, 2, 1, 2, 20, 250, 22, 3, 4, 9, 350, 130);

INSERT INTO nutrition_info
    (calories, total_fat, saturated_fat, poly_fat, mono_fat, cholesterol, sodium, carbs, fiber, sugar, protein, potassium, phosphorus)
VALUES
    (500, 22, 8, 4, 6, 60, 600, 40, 6, 12, 20, 800, 300);


-- daily
INSERT INTO daily
    (food_string, date, nutrition_id)
VALUES
    ('Grilled Chicken Salad', '2025-06-01', 1);

INSERT INTO daily
    (food_string, date, nutrition_id)
VALUES
    ('Banana Smoothie', '2025-06-01', 3);

INSERT INTO daily
    (food_string, date, nutrition_id)
VALUES
    ('Avocado Toast with Egg', '2025-06-02', 5);

-- recipes
INSERT INTO custom_recipe
    (recipe_name, alt_recipe_names, food_string, serving_size, active, nutrition_id)
VALUES
    ('Protein Power Bowl', 'High Protein Bowl|Muscle Meal', 'Quinoa, Chicken, Black Beans, Avocado', 1, TRUE, 4);

INSERT INTO custom_recipe
    (recipe_name, alt_recipe_names, food_string, serving_size, active, nutrition_id)
VALUES
    ('Vegan Green Smoothie', 'Green Detox|Kale Smoothie', 'Kale, Banana, Almond Milk, Chia Seeds', 1, TRUE, 2);

INSERT INTO custom_recipe
    (recipe_name, alt_recipe_names, food_string, serving_size, active, nutrition_id)
VALUES
    ('Breakfast Wrap', NULL, 'Egg, Spinach, Cheese, Whole Wheat Wrap', 1, FALSE, 6);
