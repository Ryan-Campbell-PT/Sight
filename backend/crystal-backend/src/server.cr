require "pg"
require "kemal"
require "./crystal-backend"
require "./nutritionix"
require "./util"
require "./database"
require "./models/databasemodels"
require "./models/foodquerymodels"
require "./foods"
require "./llm"

Kemal.config.port = 8080

before_all do |env|
  env.response.headers["Access-Control-Allow-Origin"] = "*"
  env.response.headers["Access-Control-Allow-Methods"] = "GET, POST, PUT, DELETE, OPTIONS"
  env.response.headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
  env.response.content_type = "application/json"
end

# Handle preflight OPTIONS requests
options "/*" do |env|
  env.response.status_code = 204
  ""
end

get "/coolStuff" do
  cfg = Util::EnvConfig.from_env
  newStr = ""
  r = Recipe.new
  r.recipe_name = "Test"
  r.food_string = "1 banana"
  r.serving_size = 1
  r.active = true
  r.nutrition_id = 1
  r.id = 0
  Database.exec("INSERT INTO recipe (recipe_name, food_string, serving_size, active, nutrition_id) VALUES ($1, $2, $3, $4, $5)",
    r.recipe_name,
    r.food_string,
    r.serving_size,
    r.active,
    r.nutrition_id)
  # Connection string format:
  # postgres://user:password@host:port/database
  # DB.open "postgres://#{cfg.postgres_user}:#{URI.encode_path(cfg.postgres_password)}@localhost:5432/#{cfg.postgres_database}" do |db|
  #   newStr = ""
  #   db.query "SELECT * from recipe" do |rs|
  #     rs.each do
  #       id = rs.read(Int32)
  #       recipe_name = rs.read(String)
  #       food_string = rs.read(String)
  #       serving_size = rs.read(Int32)
  #       active = rs.read(Bool)
  #       nutrition_id = rs.read(Int32)
  #       newStr = newStr + "#{id} #{recipe_name} #{food_string} #{serving_size} #{active} #{nutrition_id} \n"
  #     end
  #   end

  #   next newStr
  # end
end

post "/userFoodQuery" do |env|
  response = UserFoodQueryResponse.new

  # grab the users food string
  foodQuery = env.params.json["userFoodQuery"].as(String)

  # analyse the string provided and get whatever information it catches
  llm = LLM.new(foodQuery)

  # the food string may contain food, recipes, or errors. handle all three
  # send the food string to the nutritionix api, get back the raw nutrition info
  nixResponse = Foods.natural_language_query(llm.original_query_string)

  # errorList = llm.check_for_errors

  # transform that raw data into something that will be manipulated throughout the app

  nixResponse.to_json
end

Kemal.run
