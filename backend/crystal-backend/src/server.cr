require "kemal"
require "./crystal-backend"
require "./nutritionix"
require "./util"
require "pg"

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
  cfg = EnvConfig.from_env
  # Connection string format:
  # postgres://user:password@host:port/database
  DB.open "postgres://#{cfg.postgres_user}:#{URI.encode_path(cfg.postgres_password)}@localhost:5432/#{cfg.postgres_database}" do |db|
    newStr = ""
    db.query "SELECT * from recipe" do |rs|
      rs.each do
        id = rs.read(Int32)
        recipe_name = rs.read(String)
        food_string = rs.read(String)
        serving_size = rs.read(Int32)
        active = rs.read(Bool)
        nutrition_id = rs.read(Int32)
        newStr = newStr + "#{id} #{recipe_name} #{food_string} #{serving_size} #{active} #{nutrition_id} \n"
      end
    end

    next newStr
  end
end

post "/userFoodQuery" do |env|
  foodQuery = env.params.json["userFoodQuery"].as(String)
  # recipeList, newFoodQueryString, errors =
  responseJson = Nutritionix.naturalLanguageQuery(foodQuery)
  responseJson.to_json
end

Kemal.run
