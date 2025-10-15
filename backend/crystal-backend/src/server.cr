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

get "/coolStuff" do
  cfg = Util::EnvConfig.from_env
  newStr = ""
  r = Recipe.new
  r.recipe_name = "Test"
  r.food_string = "1 banana"
  r.serving_size = 1
  r.active = true
  r.nutrition_id = 1
  r.id = 0
  Database.db.exec("INSERT INTO recipe (recipe_name, food_string, serving_size, active, nutrition_id) VALUES ($1, $2, $3, $4, $5)",
    r.recipe_name,
    r.food_string,
    r.serving_size,
    r.active,
    r.nutrition_id)
  # Connection string format:
  # postgres://user:password@host:port/database
  # DB.open "postgres://#{cfg.postgres_user}:#{URI.encode_path(cfg.postgres_password)}@localhost:5432/#{cfg.postgres_database}" do |db|
  #   newStr = ""
  #   db.query "SELECT * from recipe" do |rs|
  #     rs.each do
  #       id = rs.read(Int32)
  #       recipe_name = rs.read(String)
  #       food_string = rs.read(String)
  #       serving_size = rs.read(Int32)
  #       active = rs.read(Bool)
  #       nutrition_id = rs.read(Int32)
  #       newStr = newStr + "#{id} #{recipe_name} #{food_string} #{serving_size} #{active} #{nutrition_id} \n"
  #     end
  #   end

  #   next newStr
  # end
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

  llm, nixResponse, errors = FoodService.analyse_user_food_query(foodQuery)
  if (errors.size > 0)
    # if there are any errors, dont go through with creating/updating a new recipe
    # ensure the errors are resolved before continuing
    response.error_list = errors
    next response
  end

  # transform the nix response into something to use
  foodList = ListOfFoods.new(nixResponse)
  # get all the recipeIds to pass into the db
  recipeIdList = llm.get_only_recipe_items.compact_map(&.recipe_id)
  if recipeIdList.size > 0
    recipeList = RecipeService.get_many(recipeIdList)

    # recipes contain the id of their nutrition_info
    # convert those ids to foodItems to use throughout the app
    # combine any recipes with nix foods to create a totals array
    listOfFoods = FoodService.combine_food_with_recipes(foodList.food_list, recipeList)
    foodList.food_list = listOfFoods
  end

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

post "/user_food_query" do |env|
  # grab the users food string
  foodQuery = env.params.json["user_food_query"].as(String)

  llm, nixResponse, errors = FoodService.analyse_user_food_query(foodQuery)
  response = UserFoodQueryResponse.new(nixResponse)
  if (errors.size > 0)
    response.error_list = errors
    response.success = false
    next response.to_json
  end

  # analyse the string provided and get whatever information it catches
  llm = LLM.new(foodQuery)

  # the food string may contain food, recipes, or errors. handle all three
  # send the food string to the nutritionix api, get back the raw nutrition info
  nixResponse = FoodService.natural_language_query(llm.original_query_string)

  # errorList = llm.check_for_errors

  # transform that raw data into something that will be manipulated throughout the app

  response.list_of_foods = ListOfFoods.new(nixResponse)
  response.success = true
  response.to_json
end

Kemal.run
