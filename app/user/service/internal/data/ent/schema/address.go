package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Address holds the schema definition for the Card entity.
type Address struct {
	ent.Schema
}

func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("mobile"),
		field.String("address"),
		field.String("post_code"),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("server", User.Type).
			Ref("addresses").
			Unique(),
	}
}
