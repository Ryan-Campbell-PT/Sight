# this files job will be to handle any logic releated to analysing the users input string
# things like getting recipes, splitting and cleaning the string into an object array
# getting any errors from the string
module LLM
  def self.create_query_bit_array(userInputQuery : String) : Array(QueryBit)
    # regex to capture various ways of writing food lists
    # ([\d/.]+)? -> optional quantity (like 1, 1.5, 1/2)
    # ([a-zA-Z]+)? -> optional unit (like cup, oz, serving)
    # (.+)$ -> the rest (your recipe/food name)
    pattern = Regex.new(
      "^\\s*([\\d/.]+)?\\s*([a-zA-Z]+)?\\s*(?:of\\s+)?(.+)$"
    )
  end

  # QueryBit's are pieces of the user input, parsed into properties for use throughout the app
  # ex: "1 banana, 2 cookies, 5oz salmon" are 3 QueryBits
  class QueryBit
    property full_string : String
    property food_name : String
    property serving : Int
    property measurement : String

    def initialize(@full_string = "", @food_name = "", @serving = 0, @measurement = "")
    end
  end
end
