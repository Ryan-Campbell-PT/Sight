package llm

// this file is meant to handle the parsing of user input
// turning it into custom objects representing the data it contains
// things like recipes, string lists, etc

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Ryan-Campbell-PT/Sight/backend/nutrition"
	"github.com/Ryan-Campbell-PT/Sight/backend/recipe"
	"github.com/gin-gonic/gin"
)

// take a the food list string provided by the user
// and turn it into an object representing the important info
func parseCustomRecipe(userInputArr UserInputParse) (*CustomRecipeParse, error) {
	// string: 1.5 servings of moms chocolate cake
	// Match pattern: number + "servings of" OR "serving of" + the rest
	re := regexp.MustCompile(`(?i)^\s*([\d.]+)\s+servings?\s+of\s+(.+)$`)

	matches := re.FindStringSubmatch(userInputArr.FoodString)
	if len(matches) != 3 {
		// return 0, "", fmt.Errorf("input did not match expected format")
		return nil, gin.Error{}
	}

	servingsStr := matches[1]
	foodName := strings.TrimSpace(matches[2])

	servings, err := strconv.ParseFloat(servingsStr, 64)
	if err != nil {
		// return 0, "", fmt.Errorf("invalid serving number: %v", err)
		return nil, err
	}

	return &CustomRecipeParse{RecipeName: foodName, NumServings: servings, FoodString: userInputArr.FoodString}, nil
}

func parseUserInput(foodString string) []UserInputParse {
	trimmedFoodString := strings.ToLower(strings.TrimSpace(foodString))
	splitStr := strings.Split(trimmedFoodString, ",")

	ret := make([]UserInputParse, 0, len(splitStr))
	for index, food := range splitStr {
		trimmed := strings.TrimSpace(food)
		if trimmed == "" {
			continue // skip possible empty: "1 avocado, , 1 banana"
		}

		userParse := UserInputParse{OriginalIndex: int64(index), FoodString: trimmed}
		// first run the food string against a local function, that matches it against regex
		recipeParse, err := parseCustomRecipe(userParse)
		if err == nil { // if there isnt an error, that means the string passes the local requirements for what a recipe is
			// if success, check db that it is indeed a recipe
			recipeId := recipe.IsActiveRecipeItem(recipeParse.RecipeName)
			if recipeId > 0 { // confirmed is real recipe, assign it
				userParse.RecipeId = int64(recipeId)
			}
		}

		ret = append(ret, userParse)
	}

	return ret
}

// this function name is a bit of a joke
// its intention is to handle all other parts of the foodListString
// that the api didnt handle
// that includes errors in the string: 1 apfel
// or recipes: 1 serving of moms chocolate cake
// ideally in the future, this will be done by a python script that does
// real language parsing. but for now, its simply string matching
func LLM(userInputString string) *LLMReturnResponse {
	// parse the food string into something that can be manipulated and modified

	// TODO should nlResponse also have a recipeList property?
	// TODO is this function doing too much, initially just meant to parse info
	// but is now just essentially another point of contact with the userInputString
	// BUT also it makes sense that you just want one point of contact on when getting user input
	nlResponse := nutrition.GetNaturalLanguageResponse(userInputString)
	ret := LLMReturnResponse{
		OriginalUserInput: userInputString,
		ParsedUserInput:   parseUserInput(userInputString),
		ListOfFoods:       nlResponse.Foods,
		ListOfErrors:      nlResponse.Errors,
	}

	return &ret
}
