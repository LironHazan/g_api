package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// daily_forecast{} --> [Weather{}, Weather{}, Weather{}]

// Forecast holds the schema definition for the Forecast entity.
type Forecast struct {
	ent.Schema
}

// Fields of the Forecast.
func (Forecast) Fields() []ent.Field {
	return []ent.Field{
		field.Text("country").
			NotEmpty(),
		field.Text("region").
			Optional(),
		field.Time("date"), //day
		field.Time("localtime").Optional(),
		field.Text("icon").
			Optional(),
		field.Float("temp").
			Default(0.0),
	}
}

// Edges of the Forecast.
func (Forecast) Edges() []ent.Edge {
	return []ent.Edge{ //O2M
		edge.To("weather", Weather.Type),
	}
}
