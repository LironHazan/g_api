// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"g_api/libs/weather-lib/ent/forcast"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Forcast is the model entity for the Forcast schema.
type Forcast struct {
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
	// MaxTemp holds the value of the "max_temp" field.
	MaxTemp float64 `json:"max_temp,omitempty"`
	// MinTemp holds the value of the "min_temp" field.
	MinTemp float64 `json:"min_temp,omitempty"`
	// AvgTemp holds the value of the "avg_temp" field.
	AvgTemp float64 `json:"avg_temp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ForcastQuery when eager-loading is set.
	Edges ForcastEdges `json:"edges"`
}

// ForcastEdges holds the relations/edges for other nodes in the graph.
type ForcastEdges struct {
	// Weather holds the value of the weather edge.
	Weather []*Weather `json:"weather,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WeatherOrErr returns the Weather value or an error if the edge
// was not loaded in eager-loading.
func (e ForcastEdges) WeatherOrErr() ([]*Weather, error) {
	if e.loadedTypes[0] {
		return e.Weather, nil
	}
	return nil, &NotLoadedError{edge: "weather"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Forcast) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case forcast.FieldMaxTemp, forcast.FieldMinTemp, forcast.FieldAvgTemp:
			values[i] = new(sql.NullFloat64)
		case forcast.FieldID:
			values[i] = new(sql.NullInt64)
		case forcast.FieldCountry, forcast.FieldRegion, forcast.FieldIcon:
			values[i] = new(sql.NullString)
		case forcast.FieldDate, forcast.FieldLocaltime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Forcast", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Forcast fields.
func (f *Forcast) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case forcast.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case forcast.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				f.Country = value.String
			}
		case forcast.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				f.Region = value.String
			}
		case forcast.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				f.Date = value.Time
			}
		case forcast.FieldLocaltime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field localtime", values[i])
			} else if value.Valid {
				f.Localtime = value.Time
			}
		case forcast.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				f.Icon = value.String
			}
		case forcast.FieldMaxTemp:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_temp", values[i])
			} else if value.Valid {
				f.MaxTemp = value.Float64
			}
		case forcast.FieldMinTemp:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_temp", values[i])
			} else if value.Valid {
				f.MinTemp = value.Float64
			}
		case forcast.FieldAvgTemp:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field avg_temp", values[i])
			} else if value.Valid {
				f.AvgTemp = value.Float64
			}
		}
	}
	return nil
}

// QueryWeather queries the "weather" edge of the Forcast entity.
func (f *Forcast) QueryWeather() *WeatherQuery {
	return NewForcastClient(f.config).QueryWeather(f)
}

// Update returns a builder for updating this Forcast.
// Note that you need to call Forcast.Unwrap() before calling this method if this Forcast
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Forcast) Update() *ForcastUpdateOne {
	return NewForcastClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Forcast entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Forcast) Unwrap() *Forcast {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Forcast is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Forcast) String() string {
	var builder strings.Builder
	builder.WriteString("Forcast(")
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
	builder.WriteString("max_temp=")
	builder.WriteString(fmt.Sprintf("%v", f.MaxTemp))
	builder.WriteString(", ")
	builder.WriteString("min_temp=")
	builder.WriteString(fmt.Sprintf("%v", f.MinTemp))
	builder.WriteString(", ")
	builder.WriteString("avg_temp=")
	builder.WriteString(fmt.Sprintf("%v", f.AvgTemp))
	builder.WriteByte(')')
	return builder.String()
}

// Forcasts is a parsable slice of Forcast.
type Forcasts []*Forcast
