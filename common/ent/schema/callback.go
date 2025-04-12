package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Callback holds the schema definition for the Callback entity.
type Callback struct {
	ent.Schema
}

// Fields of the Callback.
func (Callback) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("client_id", uuid.New()).Unique(),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("cfg").Default("{}"),
	}
}

// Edges of the Callback.
func (Callback) Edges() []ent.Edge {
	return nil
}
