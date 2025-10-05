require "./models/nutritionixmodels"
require "./models/foodquerymodels"
require "./models/nutrition-constants.cr"
require "./models/databasemodels"

module Foods
  extend self

  # Takes a users foodList query, reaches out to the nutritionix api, returns the json from the api
  def natural_language_query(userFoodQuery : String)
    cfg = Util::EnvConfig.from_env
    body = {
      query: userFoodQuery,
    }.to_json
    url = cfg.nutritionix_domain + cfg.nutritionix_natural_language
    headers = HTTP::Headers{
      "Content-Type" => cfg.nutritionix_content_type,
      "x-app-id"     => cfg.nutritionix_appid,
      "x-app-key"    => cfg.nutritionix_appkey,
    }
    r = HTTP::Client.post(
      url,
      headers: headers,
      body: body
    )

    # turn into a custom object
    return NutritionixNaturalLangaugeResponse.from_json(r.body)
  end

  # this may need to be modified/broke up in the future,
  # as your doing the recipe work in here when it could be elsewhere (maybe)
  def combine_food_with_recipes(
    foodList : Array(FoodItem) = [] of FoodItem,
    recipeList : Array(Recipe) = [] of Recipe,
  ) : ListOfFoods
  end

  def convert_nutrition_info_to_food_item(nutInfo : NutritionInfo) : FoodItem
    # TODO this needs to be done using const values to make the fooditem dict
    ret = FoodItem.new
  end

  # i think this will be the better function but we will see
  def convert_recipe_to_food_item(r : Recipe) : FoodItem
    ret = FoodItem.new
    ret.serving_qty = r.serving_size
    # ret.serving_unit = r.serving_unit
    ret.is_recipe = true

    # TODO put this somewhere else, for now just prototype
    nutInfo = Database.db.query_one?("SELECT * FROM nutrition_info WHERE id = $1", r.nutrition_id)

    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Calories)] = nutInfo.calories
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::TotalFat)] = nutInfo.total_fat
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::SaturatedFat)] = nutInfo.saturated_fat
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::TransFar)] = nutInfo.trans_fat
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Protein)] = nutInfo.protein
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Sugar)] = nutInfo.sugar
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Sodium)] = nutInfo.sodium
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Fiber)] = nutInfo.fiber
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Cholesterol)] = nutInfo.cholesterol
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Potassium)] = nutInfo.potassium
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Phosphorus)] = nutInfo.phosphorus
  end
end
