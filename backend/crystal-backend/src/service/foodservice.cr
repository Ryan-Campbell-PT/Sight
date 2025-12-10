module FoodService
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
  ) : Array(FoodItem)
    ret = Array(FoodItem).new

    recipeFoodList = Array(FoodItem).new
    recipeList.each do |r|
      recipeFoodItem = FoodService.convert_recipe_to_food_item(r)
      next unless recipeFoodItem
      recipeFoodList << recipeFoodItem
    end

    ret.concat(recipeFoodList)
    ret
  end

  # i think this will be the better function but we will see
  def convert_recipe_to_food_item(r : Recipe) : FoodItem?
    ret = FoodItem.new
    ret.serving_qty = r.serving_size
    ret.is_recipe = true

    nutInfo = NutritionInfoService.get(r.nutrition_id)
    unless nutInfo
      return nil
    end

    # automation for this is complicated and i dont understand it, so for now I will not use it
    # reflection functionality isnt realy built into crystal, so it may just not be possible
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Calories)] = nutInfo.calories.to_f32.round(2)
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::TotalFat)] = nutInfo.total_fat.to_f32.round(2)
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::SaturatedFat)] = nutInfo.saturated_fat.to_f32.round(2)
    # ret.full_nutrient_dict[get_nutrition_id(NutritionValues::TransFat)] = nutInfo.trans_fat.to_f32
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Protein)] = nutInfo.protein.to_f32.round(2)
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Sugar)] = nutInfo.sugar.to_f32.round(2)
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Sodium)] = nutInfo.sodium.to_f32.round(2)
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Fiber)] = nutInfo.fiber.to_f32.round(2)
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Cholesterol)] = nutInfo.cholesterol.to_f32.round(2)
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Potassium)] = nutInfo.potassium.to_f32.round(2)
    ret.full_nutrient_dict[get_nutrition_id(NutritionValues::Phosphorus)] = nutInfo.phosphorus.to_f32.round(2)

    ret
  end

  def analyse_user_food_query(foodQuery : String) : FoodQueryAnalysis
    llm = LLM.new(foodQuery)
    # if a recipe is included in the query, confirm it by the user first (could match a recipe they didnt mean/know)
    # if (llm.get_only_recipe_items.size > 0) # && !ignoreRecipe)
    #   llm.get_only_recipe_items.each do |qb|
    #     response.errors << AnalysisErrorObject.new(qb.full_string)
    #   end
    #   if response.errors.size > 0
    #     next response.to_json
    #   end
    # end

    nix_response = FoodService.natural_language_query(llm.original_query_string)
    error_list = llm.check_for_errors(nix_response)
    recipeIdList = llm.get_only_recipe_items.compact_map(&.recipe_id)
    recipe_list = RecipeService.get_many(recipeIdList)

    return FoodQueryAnalysis.new(llm, nix_response, recipe_list, error_list)

    return analysis
  end
end
