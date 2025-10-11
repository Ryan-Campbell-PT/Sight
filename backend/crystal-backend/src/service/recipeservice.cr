require "../models/databasemodels"
require "../database"
require "../models/foodquerymodels"

module RecipeService
  extend self

  def all : Array(Recipe)
    Database.db.query_all("SELECT * FROM recipe", as: Recipe)
  end

  def all_active(active : Bool) : Array(Recipe)
    Database.db.query_all("SELECT * FROM recipe WHERE active = $1", active, as: Recipe)
  end

  def get(id : Int32) : Recipe?
    Database.db.query_one?("SELECT * FROM recipe WHERE id = $1", id, as: Recipe)
  end

  def get_many(ids : Array(Int32)) : Array(Recipe)
    ret = Array(Recipe).new
    Database.db.query("SELECT * FROM recipe WHERE id IN ($1)", ids.join(",")) do |rs|
      # TODO check back on this in compile time, i kept getting the error
      # expected argument #1 to 'Array(Recipe)#<<' to be Recipe, not Array(Recipe)
      # when trying to << individual recipes, but it seems doing Recipe.from_rs(rs) is making an array from the rs, not individual objects
      # which wasnt totally my intention, but if that is indeed what its doing then it saves a few lines of code
      ret.concat(Recipe.from_rs(rs))
    end
    ret
  end

  def create(r : Recipe) : Int32?
    Database.db.query_one("
    CREATE recipe (recipe_name, food_string, serving_size, active, nutrition_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id",
      r.recipe_name,
      r.food_string,
      r.serving_size,
      true,
      r.nutrition_id,
      as: Int32)
  end

  def delete(id : Int32) : Bool
    Database.db.exec("
    UPDATE recipe
    SET active = false
    WHERE id = $1",
      id)
    true
  end

  def update(r : Recipe) : Bool
    current_data = RecipeService.get(r.id)

    # if (r.food_string != current_data.food_string)
    # TODO updating the nutrition_id is a bit more complicated,
    # you need to take the food_string, run it against the nix api again
    # and update that in the nutrition_info table

    # if this block doesnt run, then that isnt necessary, and you can just update whatever values are changed
    # end

    Database.db.exec("
    UPDATE recipe
    SET recipe_name = $1,
        serving_size = $2,
        active = $3,
    WHERE id = $4",
      r.recipe_name, r.serving_size,
      r.active, r.id
    )

    true
  end
end
