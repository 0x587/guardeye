package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("client_id", uuid.New()).Unique(),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("alias").Default(""),
		field.String("os"),
		field.String("os_version"),
		field.Strings("ips"),
		field.Strings("macs"),
		field.String("hostname"),
		field.String("cpu"),
		field.String("memory"),
		field.String("disk"),
		field.String("uptime"),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return nil
}
