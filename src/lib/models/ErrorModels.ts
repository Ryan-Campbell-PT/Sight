enum ErrorLevel {
    None,
    Minor,
    Warning,
    Error,
}

interface DefaultErrorResponse {
    success: boolean;
}

interface RecipeErrorResponse extends DefaultErrorResponse {

}

interface UserQueryErrorResponse extends DefaultErrorResponse {

}

interface AnalysisErrorObject {
    original_string: string;
    error_level: ErrorLevel
}