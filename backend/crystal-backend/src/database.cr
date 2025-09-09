require "pg"
require "./util"

module Database
  @@db : DB::Database? = nil

  def self.connect : DB::Database
    return @@db.not_nil! if @@db

    cfg = Util::EnvConfig.from_env
    # Connection string format:
    # postgres://user:password@host:port/database
    conn_str = "postgres://#{cfg.postgres_user}:#{URI.encode_path(cfg.postgres_password)}@localhost:5432/#{cfg.postgres_database}"
    @@db = DB.open conn_str
  end

  def self.query(sql : String, *args, &block : DB::ResultSet ->) : Nil
    # example code
    # Database.exec("SELECT * FROM recipe") do |rs|
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
end
