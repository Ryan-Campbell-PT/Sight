require "json"
require "./errormodels"

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
end

# Counterpart to NutritionixFood, this holds all the necessary data for the app,
# condensed down and given more functionality
struct FoodItem
  include JSON::Serializable

  property food_name : String
  property brand_name : String?
  property serving_qty : Int32
  property serving_unit : String
  property is_recipe : Bool
  property full_nutrient_dict : Hash(Int32, Float32) # https://crystal-lang.org/api/1.17.1/Hash.html

  def initialize(
    @food_name = "",
    @brand_name = "",
    @serving_qty = 0,
    @serving_unit = "",
    @is_recipe = false,
    full_nutrients : Array(NutritionixNutrient) = [] of Array(NutritionixNutrient),
  )
    dict = Hash(Int32, Float32).new
    full_nutrients.each do |nut|
      dict[nut.attr_id] = nut.value.to_f32 # have to cast it cause it returns as a float64
    end
    @full_nutrient_dict = dict
  end
end

# Counterpart to the NutritionixNaturalLangaugeResponse, a container for FoodItems
class ListOfFoods
  include JSON::Serializable

  # tells json the key it should expect when parsing
  @[JSON::Field(key: "foods")]
  property food_list : Array(FoodItem)

  def initialize(jsonResponse : NutritionixNaturalLangaugeResponse)
    @food_list = Array(FoodItem).new
    jsonResponse.foods.each do |food|
      foodListItem = FoodItem.new(
        food.food_name,
        food.brand_name,
        food.serving_qty,
        food.serving_unit,
        food.full_nutrients
      )
      @food_list << foodListItem
    end
  end

  def get_total_nutrition_data : FoodItem
    ret = FoodItem.new
    food_list.each do |food|
      food.full_nutrient_dict.each do |key, value|
        ret.full_nutrient_dict[key] ||= 0.0 # set the default value if the key doesnt exist
        ret.full_nutrient_dict[key] += value
      end
    end
    ret
  end
end

# foodList = Array(FoodItem).new
# foodItem = FoodItem.new
# foodItem.food_name = "ding"
# foodItem.brand_name = "dong"
# foodItem.serving_qty = 1
# foodItem.serving_unit = ""
# foodItem.full_nutrient_dict[0] = 2.5
# foodItem.full_nutrient_dict[1] = 1.5

# foodItem1 = FoodItem.new
# foodItem1.food_name = "dong"
# foodItem1.brand_name = "ding"
# foodItem1.serving_qty = 1
# foodItem1.serving_unit = ""
# foodItem1.full_nutrient_dict[0] = 2.5
# foodItem1.full_nutrient_dict[1] = 1.5

# foodList = [foodItem, foodItem1]

# lister = ListOfFoods.new
# lister.food_list = foodList

# # puts lister.inspect
# puts lister.get_total_nutrition_data.inspect
