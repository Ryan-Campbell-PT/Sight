require "json"

# move this
struct AnalysisErrorObject
  include JSON::Serializable

  property original_string : String
  property error_level : ErrorLevel

  def initialize(@original_string : String = "", @error_level : ErrorLevel = ErrorLevel::None)
  end
end

# this could also be something to include in the frontend/backend code share file
enum ErrorLevel
  None
  Minor
  Warning
  Error
end
