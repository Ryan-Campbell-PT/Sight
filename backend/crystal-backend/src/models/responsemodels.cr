# Response object containing all information relevant to a userFoodQuery request
class UserFoodQueryResponse < DefaultErrorResponse
  include JSON::Serializable

  property list_of_foods : ListOfFoods
  property total_nutrition_data : FoodItem
  property error_list : Array(AnalysisErrorObject)

  def initialize(nix : NutritionixNaturalLangaugeResponse)
    @list_of_foods = ListOfFoods.new(nix)
    @total_nutrition_data = FoodItem.new
    @error_list = Array(AnalysisErrorObject).new
  end
end

# Response object containing information relevant to the creation and updating of a recipe
class SaveRecipeResponse < DefaultErrorResponse
  include JSON::Serializable

  property recipe_id : Int32 # Recipe_id for the created or updated recipe
  property error_list : Array(AnalysisErrorObject)

  def initialize(@recipe_id : Int32 = -1, @error_list : Array(AnalysisErrorObject) = Array(AnalysisErrorObject).new)
  end
end

class GetActiveRecipesResponse < DefaultErrorResponse
  include JSON::Serializable

  property recipe_list : Array(Recipe)

  def initialize(@recipe_list : Array(Recipe) = Array(Recipe).new)
  end
end
