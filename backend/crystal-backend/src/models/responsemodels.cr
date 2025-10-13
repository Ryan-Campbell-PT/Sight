# Response object containing all information relevant to a userFoodQuery request
class UserFoodQueryResponse
  include JSON::Serializable

  property list_of_foods : ListOfFoods
  property errors : Array(AnalysisErrorObject)
end

# Response object containing information relevant to the creation and updating of a recipe
class PostRecipeResponse
  include JSON::Serializable

  property recipe_id : Int32 # Recipe_id for the created or updated recipe
  property errors : Array(AnalysisErrorObject)

  def initialize(@recipe_id : Int32 = -1, @errors : Array(AnalysisErrorObject) = Array(AnalysisErrorObject).new)
  end
end

class GetActiveRecipesResponse
  include JSON::Serializable

  property recipe_list : Array(Recipe)

  def initialize(@recipe_list : Array(Recipe) = Array(Recipe).new)
  end
end
