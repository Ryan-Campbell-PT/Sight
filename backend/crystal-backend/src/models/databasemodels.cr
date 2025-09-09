struct Recipe
  JSON::Serializable
  DB::Serializable
  property id : Int32
  property recipe_name : String
  property food_string : String
  property serving_size : Int32
  property active : Bool
  property nutrition_id : Int32

  def initialize(
    @id : Int32 = 0,
    @recipe_name : String = "",
    @food_string : String = "",
    @serving_size : Int32 = 0,
    @active : Bool = false,
    @nutrition_id : Int32 = 0,
  )
  end
end
