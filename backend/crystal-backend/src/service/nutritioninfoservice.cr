require "../models/databasemodels"
require "../database"

module NutritionInfoService
  extend self

  def get(id : Int32) : NutritionInfo?
    Database.db.query_one?("SELECT * FROM nutrition_info WHERE id = $1", id, as: NutritionInfo)
  end

  def create_from_foodinfo(f : FoodInfo) : Int32?
    n = NutritionInfo.new(
      -1, # id
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Calories)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::TotalFat)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::SaturatedFat)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::PolyFat)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::MonoFat)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Cholesterol)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Sodium)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Carbs)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Fiber)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Sugar)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Protein)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Potassium)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Phosphorus)],
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

  def update_from_foodinfo(id : Int32, f : FoodInfo) : Int32?
    n = NutritionInfo.new(
      id,
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Calories)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::TotalFat)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::SaturatedFat)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::PolyFat)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::MonoFat)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Cholesterol)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Sodium)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Carbs)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Fiber)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Sugar)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Protein)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Potassium)],
      f.full_nutrient_dict[get_nutrition_id(NutritionValues::Phosphorus)],
    )

    NutritionInfoService.update(n)
  end
end
