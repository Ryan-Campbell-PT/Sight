// AUTO-GENERATED — DO NOT EDIT
export enum NutritionValues {
  Calories = "Calories",
  TotalCarbohydrate = "TotalCarbohydrate",
  TotalFat = "TotalFat",
  SaturatedFat = "SaturatedFat",
  PolyFat = "PolyFat",
  MonoFat = "MonoFat",
  TransFat = "TransFat",
  Protein = "Protein",
  Sugar = "Sugar",
  Sodium = "Sodium",
  Fiber = "Fiber",
  Cholesterol = "Cholesterol",
  Potassium = "Potassium",
  Iron = "Iron",
  Caffeine = "Caffeine",
  Phosphorus = "Phosphorus",
}

export const NutritionMap = {
  [NutritionValues.Calories]: { id: 208, unit: "kcal", dailyValue: 2000, dbName: "calories", displayName:  "Calories"},
  [NutritionValues.TotalCarbohydrate]: { id: 205, unit: "g", dailyValue: null, dbName: "carbs", displayName:  "Carbs"},
  [NutritionValues.TotalFat]: { id: 204, unit: "g", dailyValue: 78, dbName: "total_fat", displayName:  "Total Fat"},
  [NutritionValues.SaturatedFat]: { id: 606, unit: "g", dailyValue: null, dbName: "saturated_fat", displayName:  "Saturated Fat"},
  [NutritionValues.PolyFat]: { id: 646, unit: "g", dailyValue: null, dbName: "poly_fat", displayName:  "Polyunsaturated Fat"},
  [NutritionValues.MonoFat]: { id: 645, unit: "g", dailyValue: null, dbName: "mono_fat", displayName:  "Monounsaturated Fat"},
  [NutritionValues.TransFat]: { id: 605, unit: "g", dailyValue: null, dbName: null, displayName:  "Trans Fat"},
  [NutritionValues.Protein]: { id: 203, unit: "g", dailyValue: null, dbName: "protein", displayName:  "Protein"},
  [NutritionValues.Sugar]: { id: 269, unit: "g", dailyValue: 50, dbName: "sugar", displayName:  "Sugar"},
  [NutritionValues.Sodium]: { id: 307, unit: "mg", dailyValue: 2300, dbName: "sodium", displayName:  "Sodium"},
  [NutritionValues.Fiber]: { id: 291, unit: "g", dailyValue: 28, dbName: "fiber", displayName:  "Fiber"},
  [NutritionValues.Cholesterol]: { id: 601, unit: "mg", dailyValue: 300, dbName: "cholesterol", displayName:  "Cholesterol"},
  [NutritionValues.Potassium]: { id: 306, unit: "mg", dailyValue: null, dbName: "potassium", displayName:  "Potassium"},
  [NutritionValues.Iron]: { id: 303, unit: "mg", dailyValue: null, dbName: null, displayName:  "Iron"},
  [NutritionValues.Caffeine]: { id: 262, unit: "mg", dailyValue: null, dbName: null, displayName:  "Caffeine"},
  [NutritionValues.Phosphorus]: { id: 305, unit: "mg", dailyValue: null, dbName: "phosphorus", displayName:  "Phosphorus"},
};

export function getNutritionId(key: NutritionValues): number {
  return NutritionMap[key].id;
}
