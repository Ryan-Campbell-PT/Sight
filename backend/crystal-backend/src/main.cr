require "./llm"
require "./foods"
require "./models/foodquerymodels"

lmao = LLM.new("1 banana, 3 apple, pizza")
# puts lmao.inspect

query_response = Foods.natural_language_query(lmao.original_query_string)
# puts query_response.inspect
food_item = ListOfFoods.new(query_response)
puts food_item.inspect
