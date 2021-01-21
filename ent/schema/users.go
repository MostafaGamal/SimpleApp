package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Optional(),
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty().Sensitive(),
		field.String("role").Default("UserA"),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
