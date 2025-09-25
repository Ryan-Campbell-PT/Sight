require "./database"

# this class' job is to handle any logic releated to analysing the users input string
# things like getting recipes, splitting and cleaning the string into an object array
# getting any errors from the string
class LLM
  # the original query string split into an array of custom objects
  getter user_query_bits : Array(QueryBit)
  getter original_query_string : String

  # for just the constructor, break the string into query bits
  def initialize(@original_query_string : String = "")
    # regex to capture various ways of writing food lists
    # ([\d/.]+)? -> optional quantity (like 1, 1.5, 1/2)
    # ([a-zA-Z]+)? -> optional unit (like cup, oz, serving)
    # (.+)$ -> the rest (your recipe/food name)
    regex = /\A\s*(?:(\d+(?:[\/.]?\d+)?)\s*)?(?:(\w+)\s+(?:of\s+)?)?(.+?)\s*\z/x

    # count_regex = \b\d*\s?(?:[a-zA-Z]{1,5}s?)\b
    # food_regex = \b[a-zA-Z][a-zA-Z'\- ]+\b

    @user_query_bits = original_query_string.split(',').map do |str|
      trimmed = str.strip.downcase
      next unless match = regex.match(trimmed)

      count = match[1]?.try &.to_f # nil if not present
      unit = match[2]?             # nil if not present
      food = match[3].strip

      if food
        QueryBit.new(trimmed, food, count, unit, Database.check_for_active_recipe(food))
      end
    end.compact # compact gets rid of any empty or nil array entries
  end

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

# QueryBit's are pieces of the user input, parsed into properties for use throughout the app
# ex: "1 banana, 2 cookies, 5oz salmon" are 3 QueryBits
class QueryBit
  property full_string : String
  property food_name : String
  property serving : Float64 | Nil
  property measurement : String | Nil
  property recipe_id : Int32 | Nil

  def initialize(@full_string = "",
                 @food_name = "",
                 @serving = 0,
                 @measurement = "",
                 @recipe_id = nil)
  end
end

# llm = LLM.new("1 banana")
# puts llm.inspect
