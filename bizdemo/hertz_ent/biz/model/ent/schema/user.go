package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive().Unique().Comment("PK"),
		field.String("name").Default("unknown").Comment("user name"),
		field.Int("gender").Default(0).Comment("gender"),
		field.Int("age").Positive().Comment("age"),
		field.String("introduce").Optional().Comment("introduce"),
		field.Time("createdAt").StorageKey("created_at"),
		field.Time("updatedAt").StorageKey("updated_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
