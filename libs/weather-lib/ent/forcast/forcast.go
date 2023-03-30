// Code generated by ent, DO NOT EDIT.

package forcast

const (
	// Label holds the string label denoting the forcast type in the database.
	Label = "forcast"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldLocaltime holds the string denoting the localtime field in the database.
	FieldLocaltime = "localtime"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldMaxTemp holds the string denoting the max_temp field in the database.
	FieldMaxTemp = "max_temp"
	// FieldMinTemp holds the string denoting the min_temp field in the database.
	FieldMinTemp = "min_temp"
	// FieldAvgTemp holds the string denoting the avg_temp field in the database.
	FieldAvgTemp = "avg_temp"
	// EdgeWeather holds the string denoting the weather edge name in mutations.
	EdgeWeather = "weather"
	// Table holds the table name of the forcast in the database.
	Table = "forcasts"
	// WeatherTable is the table that holds the weather relation/edge.
	WeatherTable = "weathers"
	// WeatherInverseTable is the table name for the Weather entity.
	// It exists in this package in order to avoid circular dependency with the "weather" package.
	WeatherInverseTable = "weathers"
	// WeatherColumn is the table column denoting the weather relation/edge.
	WeatherColumn = "forcast_weather"
)

// Columns holds all SQL columns for forcast fields.
var Columns = []string{
	FieldID,
	FieldCountry,
	FieldRegion,
	FieldDate,
	FieldLocaltime,
	FieldIcon,
	FieldMaxTemp,
	FieldMinTemp,
	FieldAvgTemp,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
	// DefaultMaxTemp holds the default value on creation for the "max_temp" field.
	DefaultMaxTemp float64
	// DefaultMinTemp holds the default value on creation for the "min_temp" field.
	DefaultMinTemp float64
	// DefaultAvgTemp holds the default value on creation for the "avg_temp" field.
	DefaultAvgTemp float64
)
