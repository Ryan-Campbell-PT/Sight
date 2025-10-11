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
post "/post_recipe" do |env|
  response = PostRecipeResponse.new
  response.recipe_id = -1

  # TODO more information needs to be passed in if this is an edit of a recipe
  # recipeId is really only used to determine which recipe needs to be updated
  # foodquery, serving size, name
  # and then create a recipe from it
  recipeId = env.params.json["recipe_id"].as(Int64).to_i32
  recipeName = env.params.json["recipe_name"].as(String)
  recipeServings = env.params.json["recipe_servings"].as(Int64).to_i32
  foodQuery = env.params.json["user_food_query"].as(String)
  ignoreRecipe = env.params.json["ignore_recipe"].as(Bool)

  r = Recipe.new(
    recipeId,
    recipeName,
    foodQuery,
    recipeServings,
    true,
    -1 # TODO every recipe should have their own distinct nutId, so youll need to grab it from the db to use
  )

  llm = LLM.new(foodQuery)
  # if a recipe is included in the query, confirm it by the user first (could match a recipe they didnt mean/know)
  if (llm.get_only_recipe_items.size > 0 && !ignoreRecipe)
    llm.get_only_recipe_items.each do |qb|
      response.errors << AnalysisErrorObject.new(qb.full_string)
    end
    next response.to_json
  end

  nixResponse = Foods.natural_language_query(llm.original_query_string)
  errorList = llm.check_for_errors(nixResponse)
  # dont go forward with creating/updating a recipe if there are issues with the string typed in
  if (errorList.size > 0)
    response.errors.concat(errorList)
    next response.to_json
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
    listOfFoods = Foods.combine_food_with_recipes(foodList.food_list, recipeList)
    foodList.food_list = listOfFoods
  end

  totalNutrition = foodList.get_total_nutrition_data

  if recipeId > 0
    RecipeService.update(r)
    NutritionInfoService.update_from_fooditem(r.nutrition_id, totalNutrition)
  else
    # recipes require nutInfo upon creation, so make that first
    nutId = NutritionInfoService.create_from_fooditem(totalNutrition)
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

# post "/userFoodQuery" do |env|
#   response = UserFoodQueryResponse.new

#   # grab the users food string
#   foodQuery = env.params.json["userFoodQuery"].as(String)

#   # analyse the string provided and get whatever information it catches
#   llm = LLM.new(foodQuery)

#   # the food string may contain food, recipes, or errors. handle all three
#   # send the food string to the nutritionix api, get back the raw nutrition info
#   nixResponse = Foods.natural_language_query(llm.original_query_string)

#   # errorList = llm.check_for_errors

#   # transform that raw data into something that will be manipulated throughout the app

#   nixResponse.to_json
# end

Kemal.run
