require "pg"
require "./util"
require "./models/databasemodels"

# contains information regarding the PG library
# https://deepwiki.com/crystal-lang/crystal-db/5.1-result-sets#column-information
module Database
  @@db : DB::Database? = nil

  def self.db : DB::Database
    # ||= is shorthand for
    # if @@db.nil? -> do the code, always return db
    @@db ||= begin
      cfg = Util::EnvConfig.from_env
      conn_str = "postgres://#{cfg.postgres_user}:#{URI.encode_path(cfg.postgres_password)}@localhost:5432/#{cfg.postgres_database}"
      DB.open conn_str
    end
  end

  def self.query(sql : String, *args, &block : DB::ResultSet ->) : Nil
    # example code
    # Database.query("SELECT * FROM recipe") do |rs|
    #   rs.each do
    #     r = Recipe.from_rs(rs)
    #     newStr = newStr + "#{r.id} #{r.recipe_name} #{r.food_string} #{r.serving_size} #{r.active} #{r.nutrition_id} \n"
    #   end
    # end
    connect.query sql, *args, &block
  end

  def self.exec(sql : String, *args) : Nil
    connect.exec sql, *args
  end

  # checks if the string provided exists in the database
  def self.check_for_active_recipe(str : String) : Int32 | Nil
    Database.connect.query_one?("
      SELECT id
      FROM recipe
      WHERE recipe_name ilike $1
      AND active = true",
      # "%#{str}%", as: Int32) || nil
      "#{str}", as: Int32) || nil
  end
end

# Database.query("SELECT * FROM recipe") do |rs|
#   rs.each do
#     r = rs as: Recipe
#     puts "#{r.id} #{r.recipe_name} #{r.food_string} #{r.serving_size} #{r.active} #{r.nutrition_id}"
#   end
# end

# puts Database.check_for_active_recipe("Grilled Chicken Salad")
