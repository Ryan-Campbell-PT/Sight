struct NutritionixNaturalLangaugeResponse
  include JSON::Serializable
  property foods : Array(NutritionixFood)
end

struct NutritionixFood
  include JSON::Serializable

  property food_name : String
  property brand_name : String?
  property serving_qty : Int32
  property serving_unit : String
  property serving_weight_grams : Float64
  property nf_calories : Float64
  property nf_total_fat : Float64
  property nf_saturated_fat : Float64
  property nf_cholesterol : Float64
  property nf_sodium : Float64
  property nf_total_carbohydrate : Float64
  property nf_dietary_fiber : Float64
  property nf_sugars : Float64
  property nf_protein : Float64
  property nf_potassium : Float64
  property nf_p : Float64?

  property full_nutrients : Array(NutritionixNutrient)
  property nix_brand_name : String?
  property nix_brand_id : String?
  property nix_item_name : String?
  property nix_item_id : String?
  property upc : String?

  property consumed_at : String

  property source : Int32
  property ndb_no : Int32?
  property tags : NutritionixTags
  property alt_measures : Array(NutritionixAltMeasure)?
  property lat : Float64?
  property lng : Float64?
  property meal_type : Int32
  property photo : NutritionixPhoto
  property sub_recipe : Bool?
  property class_code : String?
  property brick_code : String?
  property tag_id : String?
end

struct NutritionixNutrient
  include JSON::Serializable

  property attr_id : Int32
  property value : Float64
end

struct NutritionixTags
  include JSON::Serializable

  property item : String
  property measure : String?
  property quantity : String?
  property food_group : Int32?
  property tag_id : Int32
end

struct NutritionixAltMeasure
  include JSON::Serializable

  property serving_weight : Float64
  property measure : String
  property seq : Int32
  property qty : Float64
end

struct NutritionixPhoto
  include JSON::Serializable

  property thumb : String
  property highres : String?
  property is_user_uploaded : Bool?
end
