struct Recipe
  JSON::Serializable
  DB::Serializable
  property id : Int32
  property recipe_name : String
  property food_string : String
  property serving_size : Int32
  property active : Bool
  property nutrition_id : Int32

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
  def self.from_rs(rs)
    new(
      rs.read(Int32),
      rs.read(String),
      rs.read(String),
      rs.read(Int32),
      rs.read(Bool),
      rs.read(Int32)
    )
  end
end
