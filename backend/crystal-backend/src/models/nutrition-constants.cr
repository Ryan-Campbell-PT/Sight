# AUTO-GENERATED — DO NOT EDIT
module NutritionValues
  Calories = "Calories"
  TotalCarbohydrate = "TotalCarbohydrate"
  TotalFat = "TotalFat"
  SaturatedFat = "SaturatedFat"
  TransFat = "TransFat"
  Protein = "Protein"
  Sugar = "Sugar"
  Sodium = "Sodium"
  Fiber = "Fiber"
  Cholesterol = "Cholesterol"
  Potassium = "Potassium"
  Iron = "Iron"
  Caffeine = "Caffeine"
  Phosphorus = "Phosphorus"
end

NUTRITION_MAP = {
  NutritionValues::Calories => {id: 208, unit: "kcal", daily_value: 2000, dbName: "calories", display_name: "Calories"},
  NutritionValues::TotalCarbohydrate => {id: 205, unit: "g", daily_value: nil, dbName: "carbs", display_name: "Carbs"},
  NutritionValues::TotalFat => {id: 204, unit: "g", daily_value: 78, dbName: "total_fat", display_name: "Total Fat"},
  NutritionValues::SaturatedFat => {id: 606, unit: "g", daily_value: nil, dbName: "saturated_fat", display_name: "Saturated Fat"},
  NutritionValues::TransFat => {id: 605, unit: "g", daily_value: nil, dbName: nil, display_name: "Trans Fat"},
  NutritionValues::Protein => {id: 203, unit: "g", daily_value: nil, dbName: "protein", display_name: "Protein"},
  NutritionValues::Sugar => {id: 269, unit: "g", daily_value: 50, dbName: "sugar", display_name: "Sugar"},
  NutritionValues::Sodium => {id: 307, unit: "mg", daily_value: 2300, dbName: "sodium", display_name: "Sodium"},
  NutritionValues::Fiber => {id: 291, unit: "g", daily_value: 28, dbName: "fiber", display_name: "Fiber"},
  NutritionValues::Cholesterol => {id: 601, unit: "mg", daily_value: 300, dbName: "cholesterol", display_name: "Cholesterol"},
  NutritionValues::Potassium => {id: 306, unit: "mg", daily_value: nil, dbName: "potassium", display_name: "Potassium"},
  NutritionValues::Iron => {id: 303, unit: "mg", daily_value: nil, dbName: nil, display_name: "Iron"},
  NutritionValues::Caffeine => {id: 262, unit: "mg", daily_value: nil, dbName: nil, display_name: "Caffeine"},
  NutritionValues::Phosphorus => {id: 305, unit: "mg", daily_value: nil, dbName: "phosphorus", display_name: "Phosphorus"},
}

def get_nutrition_id(key)
  NUTRITION_MAP[key][:id]
end
