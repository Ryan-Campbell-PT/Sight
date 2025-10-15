// Response Objects
interface DefaultErrorResponse {
    success: boolean;
}

interface RecipeErrorResponse extends DefaultErrorResponse {

}

interface UserQueryErrorResponse extends DefaultErrorResponse {

}

interface NutritionErrorObject {
    error_string: string;
}


interface AnalysisErrorObject {
    error_string: string;
}