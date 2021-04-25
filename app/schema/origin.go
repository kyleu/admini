package schema

import "encoding/json"

type Origin struct {
	Key         string `json:"key"`
	Title       string `json:"title,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Description string `json:"description,omitempty"`
}

var (
	OriginGraphQL    = Origin{Key: "graphql", Title: "GraphQL", Icon: "social", Description: "GraphQL schema and queries"}
	OriginProtobuf   = Origin{Key: "protobuf", Title: "Protobuf", Icon: "move", Description: "File describing proto3 definitions"}
	OriginPostgres   = Origin{Key: "postgres", Title: "Database", Icon: "database", Description: "PostgreSQL database schema"}
	OriginJSONSchema = Origin{Key: "jsonschema", Title: "JSONSchema", Icon: "location", Description: "JSON Schema definition files"}
	OriginUnknown    = Origin{Key: "unknown", Title: "Unknown", Icon: "question", Description: "Not quite sure what this is"}
)

var AllOrigins = []Origin{OriginGraphQL, OriginProtobuf, OriginPostgres, OriginJSONSchema}

func OriginFromString(s string) Origin {
	for _, t := range AllOrigins {
		if t.Key == s {
			return t
		}
	}
	return OriginUnknown
}

func (t *Origin) String() string {
	return t.Key
}

func (t *Origin) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

func (t *Origin) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = OriginFromString(s)
	return nil
}
