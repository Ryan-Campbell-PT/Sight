require "kemal"
require "./crystal-backend"
require "./nutritionix"
require "./util"

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
end

post "/userFoodQuery" do |env|
  foodQuery = env.params.json["userFoodQuery"].as(String)
  responseJson = Nutritionix.naturalLanguageQuery(foodQuery)
  responseJson.to_json
end

Kemal.run
