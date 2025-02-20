package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Subscribe holds the schema definition for the Subscribe entity.
type Subscribe struct {
	ent.Schema
}

// Fields of the Subscribe.
func (Subscribe) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("enable").Default(true),
		field.String("query").NotEmpty(),
		field.Enum("type").Values(
			"WebHook",
			"MqttPush",
			"WebSocket",
		),
		field.String("web_hook_url").Optional(),
		field.String("mqtt_push_url").Optional(),
		field.String("mqtt_push_topic").Optional(),
		field.String("web_socket_key").Optional(),
	}
}

// Edges of the Subscribe.
func (Subscribe) Edges() []ent.Edge {
	return nil
}
