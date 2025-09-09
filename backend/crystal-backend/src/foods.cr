require "./models/nutritionixmodels"

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
    response = NutritionixNaturalLangaugeResponse.from_json(r.body)
    return response
  end

  # This class is designed to handle parsing of information related to the food list passed in by the user
  class LLM
    # Analyse the user food query, returning any recipes or errors found in the string
    # then returning the modified string with just standard foods to pass into nutritionix
    def self.analyseUserFoodQuery(userFoodQuery : String, foods : CustomFoodItem) : Array(AnalysisErrorObject)
      errorList = [] of AnalysisErrorObject
      userQuerySplitArray = userFoodQuery.split(",").map(&.strip.downcase)

      # If there are more items in the split array, that means some foods were not analysed and likely errored
      return errorList if userQuerySplitArray.size <= foods.size

      foodsIndex = 0
      userQuerySplitArray.each do |query|
        # Edgecase: if all knownfoods have been matched, then everything else is an error
        if foodsIndex >= foods.len
          errorList << AnalysisErrorObject.new(query)
          next
        end
        # if string typed by user was handled by the api
        expectedFood = foods[foodsIndex].FoodName
        if (query.includes?(expectedFood))
          foodsIndex += 1
        else
          errorList << AnalysisErrorObject.new(query)
        end
      end

      return errorList
    end
  end
end
