require "./llm"
require "./foods"
require "./models/foodquerymodels"
require "./crystal-backend"
require "./nutritionix"
require "./util"
require "./database"
require "./models/databasemodels"
require "./service/recipeservice"

response = PostRecipeResponse.new
response.recipe_id = -1

recipeId = 6
foodQuery = "1 banana, 1 apple"
ignoreRecipe = false

llm = LLM.new(foodQuery)
# if a recipe is included in the query, confirm it by the user first (could match a recipe they didnt mean/know)
if (llm.get_only_recipe_items.size > 0 && !ignoreRecipe)
  llm.get_only_recipe_items.each do |qb|
    response.errors << AnalysisErrorObject.new(qb.full_string)
  end
end

nixResponse = Foods.natural_language_query(llm.original_query_string)
errorList = llm.check_for_errors(nixResponse)
# dont go forward with creating/updating a recipe if there are issues with the string typed in
if (errorList.size > 0)
  response.errors.concat(errorList)
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
puts totalNutrition.to_json
