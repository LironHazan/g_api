package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Weather holds the schema definition for the Weather entity.
type Weather struct {
	ent.Schema
}

// Fields of the Weather.
func (Weather) Fields() []ent.Field {
	return []ent.Field{
		field.Text("icon").
			Optional(),
		field.Int("time"),
		field.Time("time_epoch").Optional(),
		field.Float("temp"),
		field.Float("feels_like"),
	}
}

// Edges of the Weather.
func (Weather) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("forecast", Forecast.Type).
			Ref("weather").
			Unique(),
	}
}
