require "../models/databasemodels"
require "../models/foodquerymodels"

require "../database"

module NutritionInfoService
  extend self

  def get(id : Int32) : NutritionInfo?
    Database.db.query_one?("SELECT * FROM nutrition_info WHERE id = $1", id, as: NutritionInfo)
  end

  def create_from_fooditem(f : FoodItem) : Int32?
    n = NutritionInfo.new(
      -1, # id
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Calories)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::TotalFat)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::SaturatedFat)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::PolyFat)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::MonoFat)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Cholesterol)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Sodium)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::TotalCarbohydrate)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Fiber)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Sugar)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Protein)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Potassium)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Phosphorus)].to_i32,
    )

    NutritionInfoService.create(n)
  end

  def create(n : NutritionInfo) : Int32?
    # exec does not support returning the inserted id in postgres
    # so you have to query for the id
    Database.db.query_one("INSERT INTO nutrition_info
        (
            calories,
            total_fat,
            saturated_fat,
            poly_fat,
            mono_fat,
            cholesterol,
            sodium,
            carbs,
            fiber,
            sugar,
            protein,
            potassium,
            phosphorus
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7,
            $8, $9, $10, $11, $12, $13
        )
        RETURNING id", # required for postgres, as it doesnt support last_insert_id natively
      n.calories,
      n.total_fat,
      n.saturated_fat,
      n.poly_fat,
      n.mono_fat,
      n.cholesterol,
      n.sodium,
      n.carbs,
      n.fiber,
      n.sugar,
      n.protein,
      n.potassium,
      n.phosphorus,
      as: Int32)
  end

  def update(n : NutritionInfo) : Bool
    statement = Database.db.exec("UPDATE nutrition_info
        SET calories = $1,
            total_fat = $2,
            saturated_fat = $3,
            poly_fat = $4,
            mono_fat = $5,
            cholesterol = $6,
            sodium = $7,
            carbs = $8,
            fiber = $9,
            sugar = $10,
            protein = $11,
            potassium = $12,
            phosphorus = $13
        WHERE id = $14",
      n.calories,
      n.total_fat,
      n.saturated_fat,
      n.poly_fat,
      n.mono_fat,
      n.cholesterol,
      n.sodium,
      n.carbs,
      n.fiber,
      n.sugar,
      n.protein,
      n.potassium,
      n.phosphorus,
      n.id)
    statement.rows_affected > 0 ? true : false
  end

  def update_from_fooditem(id : Int32, f : FoodItem) : Bool
    n = NutritionInfo.new(
      id,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Calories)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::TotalFat)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::SaturatedFat)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::PolyFat)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::MonoFat)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Cholesterol)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Sodium)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::TotalCarbohydrate)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Fiber)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Sugar)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Protein)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Potassium)].to_i32,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Phosphorus)].to_i32,
    )

    NutritionInfoService.update(n)
  end
end
