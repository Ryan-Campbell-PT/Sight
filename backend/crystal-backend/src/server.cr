require "pg"
require "kemal"
require "./crystal-backend"
require "./nutritionix"
require "./util"
require "./database"
require "./models/databasemodels"
require "./models/foodquerymodels"
require "./foods"
require "./llm"
require "./service/recipeservice"
require "./models/responsemodels"
require "./service/foodservice"

Kemal.config.port = 8080

before_all do |env|
  env.response.headers["Access-Control-Allow-Origin"] = "*"
  env.response.headers["Access-Control-Allow-Methods"] = "GET, POST, PUT, DELETE, OPTIONS"
  env.response.headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
  env.response.content_type = "application/json"
end

# Handle preflight OPTIONS requests
options "/*" do |env|
  env.response.status_code = 204
  ""
end

# this should account for both create and edit, passing the recipeId in if its an edit, or empty/-1 if not
# also passes in an 'ignore_recipe' property that allows you to add recipes to your recipe
post "/save_recipe" do |env|
  response = SaveRecipeResponse.new
  response.recipe_id = -1

  recipeId = env.params.json["recipe_id"].as(Int64).to_i32
  recipeName = env.params.json["recipe_name"].as(String)
  recipeServings = env.params.json["recipe_servings"].as(Int64).to_i32
  foodQuery = env.params.json["user_food_query"].as(String)

  r = Recipe.new(
    recipeId,
    recipeName,
    foodQuery,
    recipeServings,
    true,
    -1
  )
  if recipeId > 0
    r.nutrition_id = RecipeService.get_nutrition_id(recipeId)
  end

  analysis = FoodService.analyse_user_food_query(foodQuery)
  if (analysis.error_list.size > 0)
    # if there are any errors, dont go through with creating/updating a new recipe
    # ensure the errors are resolved before continuing
    response.error_list = analysis.error_list
    next response
  end

  # transform the nix response into something to use
  foodList = ListOfFoods.new(analysis.nix_response)

  # recipes contain the id of their nutrition_info
  # convert those ids to foodItems to use throughout the app
  # combine any recipes with nix foods to create a totals array
  listOfFoods = FoodService.combine_food_with_recipes(foodList.food_list, analysis.recipe_list)
  foodList.food_list = listOfFoods

  totalNutrition = foodList.get_total_nutrition_data
  if recipeId > 0
    RecipeService.update(r)
    NutritionInfoService.update_from_fooditem(r.nutrition_id, totalNutrition)
  else
    # recipes require nutInfo upon creation, so make that first
    nutId = NutritionInfoService.create_from_fooditem(totalNutrition)
    # ding = NutritionInfoService.get(11)
    # nutId = ding ? ding.id : nil
    if nutId
      r.nutrition_id = nutId
      recipeId = RecipeService.create(r)
      if recipeId
        response.recipe_id = recipeId
      end
    end
  end

  response.to_json
end

get "/get_active_recipes" do |env|
  rList = RecipeService.all_active(true)
  response = GetActiveRecipesResponse.new(rList)
  response.to_json
end

# TODO so right now what is happening
# if any errors whatsoever happen, success is set to false
# and essentially all data is thrown away and the nutrition label doesnt display any info
# (nutrition breakdown does which is interesting tho)

# so my thought process is that strict functionality should only be present for recipe saving
# for just the natrual language response, it should display any information thats present
post "/user_food_query" do |env|
  # grab the users food string
  foodQuery = env.params.json["user_food_query"].as(String)

  analysis = FoodService.analyse_user_food_query(foodQuery)
  response = NaturalLanguageResponse.new(analysis.nix_response)
  if (analysis.error_list.size > 0)
    response.error_list = analysis.error_list
    # response.success = false
    # next response.to_json
  end

  # analyse the string provided and get whatever information it catches
  # llm = LLM.new(foodQuery)

  # the food string may contain food, recipes, or errors. handle all three
  # send the food string to the nutritionix api, get back the raw nutrition info
  # nixResponse = FoodService.natural_language_query(llm.original_query_string)

  # errorList = llm.check_for_errors

  # transform that raw data into something that will be manipulated throughout the app

  nixFoodList = ListOfFoods.new(analysis.nix_response)
  response.list_of_foods = ListOfFoods.new(FoodService.combine_food_with_recipes(nixFoodList.food_list, analysis.recipe_list))
  response.total_nutrition_data = response.list_of_foods.get_total_nutrition_data
  response.success = true
  response.to_json
end

Kemal.run
