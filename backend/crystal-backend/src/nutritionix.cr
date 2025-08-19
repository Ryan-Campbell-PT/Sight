require "json"
require "http/client"

module Nutritionix
  # extend self lets you create functions for a module, like making static functions
  extend self

  # Takes a users foodList query, reaches out to the nutritionix api, returns the json from the api
  def naturalLanguageQuery(userFoodQuery : String)
    cfg = Config.from_env
    body = {
      query: userFoodQuery,
    }.to_json
    url = cfg.nutritionix_domain + cfg.nutritionix_natural_language
    headers = HTTP::Headers{
      "Content-Type" => cfg.nutritionix_content_type,
      "x-app-id"     => cfg.nutritionix_appid,
      "x-app-key"    => cfg.nutritionix_appkey,
    }
    r = HTTP::Client.post(
      url,
      headers: headers,
      body: body
    )

    # turn into a custom object
    response = NaturalLangaugeResponse.from_json(r.body)
    return response
  end

  struct NaturalLangaugeResponse
    include JSON::Serializable
    property foods : Array(Food)
  end

  struct Food
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

    property full_nutrients : Array(Nutrient)
    # property nix_brand_name : String?
    # property nix_brand_id : String?
    # property nix_item_name : String?
    # property nix_item_id : String?
    # property upc : String?

    # property consumed_at : String
    # property metadata : Metadata

    # property source : Int32
    # property ndb_no : Int32?
    # property tags : Tags
    # property alt_measures : Array(AltMeasure)?
    # property lat : Float64?
    # property lng : Float64?
    # property meal_type : Int32
    # property photo : Photo
    # property sub_recipe : Bool?
    # property class_code : String?
    # property brick_code : String?
    # property tag_id : String?
  end

  struct Nutrient
    include JSON::Serializable

    property attr_id : Int32
    property value : Float64
  end

  struct Tags
    include JSON::Serializable

    property item : String
    property measure : String?
    property quantity : String?
    property food_group : Int32?
    property tag_id : Int32
  end

  struct AltMeasure
    include JSON::Serializable

    property serving_weight : Float64
    property measure : String
    property seq : Int32
    property qty : Float64
  end

  struct Photo
    include JSON::Serializable

    property thumb : String
    property highres : String?
    property is_user_uploaded : Bool?
  end
end
