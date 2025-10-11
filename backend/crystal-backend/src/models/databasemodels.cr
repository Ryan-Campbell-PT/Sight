require "db"
require "json"

struct Recipe
  include JSON::Serializable
  include DB::Serializable

  property id : Int32
  property recipe_name : String
  property food_string : String
  property serving_size : Int32
  property active : Bool
  property nutrition_id : Int32

  # adding color feels like a good idea to let you categorize recipes in a way
  # property color : String

  # this will eventually have to be added, not that many people will use this but to decern recipes from user
  # property user_id : Int32

  def initialize(
    @id : Int32 = 0,
    @recipe_name : String = "",
    @food_string : String = "",
    @serving_size : Int32 = 0,
    @active : Bool = false,
    @nutrition_id : Int32 = 0,
  )
  end

  # result set, aka db row
  # def self.from_rs(rs)
  #   new(
  #     id: rs.read("id", Int32),
  #     recipe_name: rs.read("recipe_name", String),
  #     food_string: rs.read("food_string", String),
  #     serving_size: rs.read("serving_size", Int32),
  #     active: rs.read("active", Bool),
  #     nutrition_id: rs.read("nutrition_id", Int32)
  #   )
  # end
end

# how this class will be used is still in the air,
# as it may just be fine getting back the FoodInfo and using that

struct NutritionInfo
  include JSON::Serializable
  include DB::Serializable

  property id : Int32
  property calories : Int32
  property total_fat : Int32
  property saturated_fat : Int32
  property poly_fat : Int32
  property mono_fat : Int32
  property cholesterol : Int32
  property sodium : Int32
  property carbs : Int32
  property fiber : Int32
  property sugar : Int32
  property protein : Int32
  property potassium : Int32
  property phosphorus : Int32

  def initialize(
    @id : Int32 = 0,
    @calories : Int32 = 0,
    @total_fat : Int32 = 0,
    @saturated_fat : Int32 = 0,
    @poly_fat : Int32 = 0,
    @mono_fat : Int32 = 0,
    @cholesterol : Int32 = 0,
    @sodium : Int32 = 0,
    @carbs : Int32 = 0,
    @fiber : Int32 = 0,
    @sugar : Int32 = 0,
    @protein : Int32 = 0,
    @potassium : Int32 = 0,
    @phosphorus : Int32 = 0,
  )
  end
end
