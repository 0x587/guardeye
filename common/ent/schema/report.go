package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("client_id", uuid.New()),
		field.String("message"),
		field.String("provider_type"),
		field.Strings("provider_args"),
		field.Time("report_at").Annotations(),
	}
}

func (Report) Indexes() []ent.Index {
	return []ent.Index{
		//index.Fields("client_id").StorageKey("index_client_id" + f()),
		//index.Fields("report_at").StorageKey("index_report_at" + f()),
		//index.Fields("client_id", "report_at").StorageKey("index_client_id_report_at" + f()),
		index.Fields("client_id").StorageKey("index_client_id"),
		index.Fields("report_at").StorageKey("index_report_at"),
		index.Fields("client_id", "report_at").StorageKey("index_client_id_report_at"),
	}
}

//func (Report) Annotations() []schema.Annotation {
//	return []schema.Annotation{
//		schema.Comment(),
//	}
//}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return nil
}

//func f() string {
//	return fmt.Sprintf("report_%s", time.Now().Format("2006-01-02"))
//}
