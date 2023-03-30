package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// daily_forecast{} --> [Weather{}, Weather{}, Weather{}]

// Forcast holds the schema definition for the Forcast entity.
type Forcast struct {
	ent.Schema
}

// Fields of the Forcast.
func (Forcast) Fields() []ent.Field {
	return []ent.Field{
		field.Text("country").
			NotEmpty(),
		field.Text("region").
			Optional(),
		field.Time("date"), //day
		field.Time("localtime").Optional(),
		field.Text("icon").
			Optional(),
		field.Float("max_temp").
			Default(0.0),
		field.Float("min_temp").
			Default(0.0),
		field.Float("avg_temp").
			Default(0.0),
	}
}

// Edges of the Forcast.
func (Forcast) Edges() []ent.Edge {
	return []ent.Edge{ //O2M
		edge.To("weather", Weather.Type),
	}
}
