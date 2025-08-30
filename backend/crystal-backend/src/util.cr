require "dotenv"

struct EnvConfig
  include JSON::Serializable

  property user : String
  property password : String
  property nutritionix_appid : String
  property nutritionix_appkey : String
  property nutritionix_domain : String
  property nutritionix_natural_language : String
  property nutritionix_content_type : String

  property azure_user : String
  property azure_password : String
  property azure_database : String
  property azure_server : String
  property azure_port : Int64

  property postgres_user : String
  property postgres_password : String
  property postgres_database : String

  def initialize(
    @user : String,
    @password : String,
    @nutritionix_appid : String,
    @nutritionix_appkey : String,
    @nutritionix_domain : String,
    @nutritionix_natural_language : String,
    @nutritionix_content_type : String,
    @azure_user : String,
    @azure_password : String,
    @azure_database : String,
    @azure_server : String,
    @azure_port : Int64,
    @postgres_user : String,
    @postgres_password : String,
    @postgres_database : String,
  )
  end

  def self.from_env : EnvConfig
    Dotenv.load

    new(
      user: ENV["DBUSER"]? || "",
      password: ENV["DBPASS"]? || "",
      nutritionix_appid: ENV["nutrition__appid"]? || "",
      nutritionix_appkey: ENV["nutrition__appkey"]? || "",
      nutritionix_domain: ENV["nutrition__domain"]? || "",
      nutritionix_natural_language: ENV["nutrition__naturalLanguage"]? || "",
      nutritionix_content_type: ENV["nutrition__contentType"]? || "",
      azure_user: ENV["Azure_User"]? || "",
      azure_password: ENV["Azure_Password"]? || "",
      azure_database: ENV["Azure_Database"]? || "",
      azure_server: ENV["Azure_Server"]? || "",
      azure_port: (ENV["Azure_Port"]? || "0").to_i64,
      postgres_user: ENV["PostgreSQL_User"]? || "",
      postgres_password: ENV["PostgreSQL_Password"]? || "",
      postgres_database: ENV["PostgreSQL_Database"]? || "",
    )
  end
end
