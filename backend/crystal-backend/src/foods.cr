module Foods
  # Counterpart to the NutritionixNaturalLangaugeResponse, a container for FoodItems
  class ListOfFoods
    include JSON::Serializable

    # tells json the key it should expect when parsing
    @[JSON::Field(key: "foods")]
    food_list : Array(FoodItem)

    def initialize(jsonResponse : NutritionixNaturalLangaugeResponse)
      ding = ListOfFoods.from_json(jsonResponse)
    end
  end

  # Counterpart to NutritionixFood, this holds all the necessary data for the app,
  # condensed down and given more functionality
  struct FoodItem
    food_name : String
    brand_name : String?
    serving_qty : Int32
    serving_unit : String
  end
end
