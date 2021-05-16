package schema

import "encoding/json"

type Origin struct {
	Key         string `json:"key"`
	Title       string `json:"title,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Description string `json:"description,omitempty"`
}

var (
	OriginPostgres   = Origin{Key: "postgres", Title: "PostgreSQL", Icon: "database", Description: "PostgreSQL database schema"}
	OriginGraphQL    = Origin{Key: "graphql", Title: "GraphQL", Icon: "star", Description: "GraphQL schema and queries"}
	OriginProtobuf   = Origin{Key: "protobuf", Title: "Protobuf", Icon: "star", Description: "File describing proto3 definitions"}
	OriginJSONSchema = Origin{Key: "jsonschema", Title: "JSON Schema", Icon: "star", Description: "JSON Schema definition files"}
	OriginMock       = Origin{Key: "mock", Title: "Mock", Icon: "star", Description: "Simple type that returns mock data"}
	OriginUnknown    = Origin{Key: "unknown", Title: "Unknown", Icon: "star", Description: "Not quite sure what this is"}
)

var AllOrigins = []Origin{OriginPostgres, OriginGraphQL, OriginProtobuf, OriginJSONSchema, OriginMock}

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
