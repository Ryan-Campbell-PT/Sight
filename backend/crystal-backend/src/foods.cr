require "./models/nutritionixmodels"
require "./models/foodquerymodels"

module Foods
  # Takes a users foodList query, reaches out to the nutritionix api, returns the json from the api
  def self.natural_language_query(userFoodQuery : String)
    cfg = Util::EnvConfig.from_env
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
    return NutritionixNaturalLangaugeResponse.from_json(r.body)
  end
end
