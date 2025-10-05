require "./database"
require "./models/errormodels"

# this class' job is to handle any logic releated to analysing the users input string
# things like getting recipes, splitting and cleaning the string into an object array
# getting any errors from the string
class LLM
  # the original query string split into an array of custom objects
  getter user_query_bits : Array(QueryBit)
  getter original_query_string : String
  getter user_string_no_recipes : String

  # for just the constructor, break the string into query bits
  def initialize(@original_query_string : String = "")
    # regex to capture various ways of writing food lists
    # ([\d/.]+)? -> optional quantity (like 1, 1.5, 1/2)
    # ([a-zA-Z]+)? -> optional unit (like cup, oz, serving)
    # (.+)$ -> the rest (your recipe/food name)
    regex = /\A\s*(?:(\d+(?:[\/.]?\d+)?)\s*)?(?:(\w+)\s+(?:of\s+)?)?(.+?)\s*\z/x

    # count_regex = \b\d*\s?(?:[a-zA-Z]{1,5}s?)\b
    # food_regex = \b[a-zA-Z][a-zA-Z'\- ]+\b

    @user_query_bits = original_query_string.split(',').map_with_index do |str, index|
      trimmed = str.strip.downcase
      next unless match = regex.match(trimmed)

      count = match[1]?.try &.to_f # nil if not present
      unit = match[2]?             # nil if not present
      food = match[3].strip

      # recipe_id =
      # if recipe_id == nil
      #   @user_string_no_recipes += "#{trimmed},"
      # end

      if food
        QueryBit.new(trimmed, food, count, unit, Database.check_for_active_recipe(food), index)
      end
    end.compact # compact gets rid of any empty or nil array entries

    @user_string_no_recipes = get_no_recipe_items # get only the objects without recipe
      .map(&.full_string)                         # return only their property full_string
      .join(", ")                                 # join them all with a comma seperating them
  end

  # Analyse the user food query, returning any recipes or errors found in the string
  # then returning the modified string with just standard foods to pass into nutritionix
  # DEPRECIATED
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

  # TODO I was thinking this could be used to access either specific recipes passed in
  # or all of them that are in the list
  # that way you can get nutrition information for specific recipes or all of them (to add up)
  def self.get_recipe_names
  end

  # sadly no real better way to write these
  def get_only_recipe_items : Array(QueryBit)
    @user_query_bits.select { |m| m.recipe_id != nil }
  end

  def get_no_recipe_items : Array(QueryBit)
    @user_query_bits.select { |m| m.recipe_id == nil }
  end

  # error analysis happens in the LLM (i think)
  def check_for_errors(nixResponse : NutritionixNaturalLangaugeResponse) : Array(AnalysisErrorObject)
    retList = Array(AnalysisErrorObject).new
    if @user_query_bits.size == nixResponse.foods.size
      # if the number of querybits == how many foods are in the nix response,
      # then everything the user typed in is accepted by the api, so no errors
      puts "equal size"
      retList
    elsif get_no_recipe_items.size == nixResponse.foods.size
      # if the nix response contains all the foods requested by the user
      # besides the recipes, no errors
      puts "equal recipe size"
      retList
    end

    # if neither above is true, then you can assume there is an error somewhere
    nixIndex = 0
    @user_query_bits.each do |bit|
      puts "in loop: #{nixIndex}"
      if bit.recipe_id != nil
        puts "contains recipe"
        # if its a recipe, its been analyised and not an error
        nixIndex += 1
        next
      end

      # If all known foods have been matched, everything else is an error
      if nixIndex >= nixResponse.foods.size
        retList << AnalysisErrorObject.new(bit.full_string)
        next
      end

      nixItem = nixResponse.foods[nixIndex]

      if bit.full_string.includes?(nixItem.food_name)
        puts "no error with  #{bit.full_string}"
        # this index is fine, no error, continue
        nixIndex += 1
        next
      else
        puts "error with #{bit.full_string}"
        # if we are at the same index, but the response from nix is different, error
        retList << AnalysisErrorObject.new(bit.full_string)
        # nixIndex += 1
      end
    end
    retList
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
  # keeps track of the index this Bit was found in the original FoodQuery
  property index : Int32

  def initialize(@full_string = "",
                 @food_name = "",
                 @serving = 0,
                 @measurement = "",
                 @recipe_id = nil,
                 @index = -1)
  end
end

# llm = LLM.new("1 banana, 2 cookies, 3oz grapes, 1 test")
# puts llm.inspect
