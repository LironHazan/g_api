// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"g_api/libs/weather-lib/ent/forecast"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Forecast is the model entity for the Forecast schema.
type Forecast struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Localtime holds the value of the "localtime" field.
	Localtime time.Time `json:"localtime,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Temp holds the value of the "temp" field.
	Temp float64 `json:"temp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ForecastQuery when eager-loading is set.
	Edges ForecastEdges `json:"edges"`
}

// ForecastEdges holds the relations/edges for other nodes in the graph.
type ForecastEdges struct {
	// Weather holds the value of the weather edge.
	Weather []*Weather `json:"weather,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WeatherOrErr returns the Weather value or an error if the edge
// was not loaded in eager-loading.
func (e ForecastEdges) WeatherOrErr() ([]*Weather, error) {
	if e.loadedTypes[0] {
		return e.Weather, nil
	}
	return nil, &NotLoadedError{edge: "weather"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Forecast) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case forecast.FieldTemp:
			values[i] = new(sql.NullFloat64)
		case forecast.FieldID:
			values[i] = new(sql.NullInt64)
		case forecast.FieldCountry, forecast.FieldRegion, forecast.FieldIcon:
			values[i] = new(sql.NullString)
		case forecast.FieldDate, forecast.FieldLocaltime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Forecast", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Forecast fields.
func (f *Forecast) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case forecast.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case forecast.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				f.Country = value.String
			}
		case forecast.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				f.Region = value.String
			}
		case forecast.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				f.Date = value.Time
			}
		case forecast.FieldLocaltime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field localtime", values[i])
			} else if value.Valid {
				f.Localtime = value.Time
			}
		case forecast.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				f.Icon = value.String
			}
		case forecast.FieldTemp:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field temp", values[i])
			} else if value.Valid {
				f.Temp = value.Float64
			}
		}
	}
	return nil
}

// QueryWeather queries the "weather" edge of the Forecast entity.
func (f *Forecast) QueryWeather() *WeatherQuery {
	return NewForecastClient(f.config).QueryWeather(f)
}

// Update returns a builder for updating this Forecast.
// Note that you need to call Forecast.Unwrap() before calling this method if this Forecast
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Forecast) Update() *ForecastUpdateOne {
	return NewForecastClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Forecast entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Forecast) Unwrap() *Forecast {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Forecast is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Forecast) String() string {
	var builder strings.Builder
	builder.WriteString("Forecast(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("country=")
	builder.WriteString(f.Country)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(f.Region)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(f.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("localtime=")
	builder.WriteString(f.Localtime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("icon=")
	builder.WriteString(f.Icon)
	builder.WriteString(", ")
	builder.WriteString("temp=")
	builder.WriteString(fmt.Sprintf("%v", f.Temp))
	builder.WriteByte(')')
	return builder.String()
}

// Forecasts is a parsable slice of Forecast.
type Forecasts []*Forecast
