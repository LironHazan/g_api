// Code generated by ent, DO NOT EDIT.

package weather

const (
	// Label holds the string label denoting the weather type in the database.
	Label = "weather"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldTimeEpoch holds the string denoting the time_epoch field in the database.
	FieldTimeEpoch = "time_epoch"
	// FieldTemp holds the string denoting the temp field in the database.
	FieldTemp = "temp"
	// FieldFeelsLike holds the string denoting the feels_like field in the database.
	FieldFeelsLike = "feels_like"
	// EdgeForecast holds the string denoting the forecast edge name in mutations.
	EdgeForecast = "forecast"
	// Table holds the table name of the weather in the database.
	Table = "weathers"
	// ForecastTable is the table that holds the forecast relation/edge.
	ForecastTable = "weathers"
	// ForecastInverseTable is the table name for the Forecast entity.
	// It exists in this package in order to avoid circular dependency with the "forecast" package.
	ForecastInverseTable = "forecasts"
	// ForecastColumn is the table column denoting the forecast relation/edge.
	ForecastColumn = "forecast_weather"
)

// Columns holds all SQL columns for weather fields.
var Columns = []string{
	FieldID,
	FieldIcon,
	FieldDate,
	FieldTime,
	FieldTimeEpoch,
	FieldTemp,
	FieldFeelsLike,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "weathers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"forecast_weather",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
