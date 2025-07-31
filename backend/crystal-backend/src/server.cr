require "kemal"

Kemal.config.port = 8080

before_all do |env|
  env.response.headers["Access-Control-Allow-Origin"] = "*"
  env.response.headers["Access-Control-Allow-Methods"] = "GET, POST, PUT, DELETE, OPTIONS"
  env.response.headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
  env.response.content_type = "application/json"
end

get "/" do
  data = {
    message: "Hello, JSON!",
    success: true,
    items:   [1, 2, 3],
  }

  data.to_json
end

Kemal.run
