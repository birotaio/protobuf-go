package structpb

import (
	"fmt"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"gopkg.in/yaml.v3"
)

// -------------------------------------------------------------- //
// THIS IS CUSTOM CODE TO HANDLE YAML MARSHALING AND UNMARSHALING //
// -------------------------------------------------------------- //

// MarshalYAML implements the yaml.Marshaler interface.
func (x *Value) MarshalYAML() (interface{}, error) {
	if x == nil {
		return nil, nil
	}
	return x.AsInterface(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (x *Value) UnmarshalYAML(value *yaml.Node) error {
	var val interface{}
	if err := value.Decode(&val); err != nil {
		return protoimpl.X.NewError("failed to decode YAML: %v", err)
	}

	// Sanitize any potential YAML-specific types
	val = SanitizeYAMLTypes(val)

	v, err := NewValue(val)
	if err != nil {
		return err
	}
	*x = *v
	return nil
}

// SanitizeYAMLTypes converts YAML-specific types to Go types that are compatible with protobuf.
func SanitizeYAMLTypes(val interface{}) interface{} {
	switch v := val.(type) {
	case map[interface{}]interface{}:
		// Convert map[interface{}]interface{} to map[string]interface{}
		m := make(map[string]interface{})
		for key, value := range v {
			// YAML allows non-string keys, but protobuf requires string keys
			k, ok := key.(string)
			if !ok {
				// Convert key to string
				k = fmt.Sprintf("%v", key)
			}
			m[k] = SanitizeYAMLTypes(value)
		}
		return m
	case []interface{}:
		// Recursively sanitize each element in the slice
		for i, item := range v {
			v[i] = SanitizeYAMLTypes(item)
		}
		return v
	case map[string]interface{}:
		// Recursively sanitize each value in the map
		for key, value := range v {
			v[key] = SanitizeYAMLTypes(value)
		}
		return v
	default:
		return val
	}
}

// NewYAMLValue creates a new Value from YAML node.
// func NewYAMLValue(node *yaml.Node) (*Value, error) {
// 	var val interface{}
// 	if err := node.Decode(&val); err != nil {
// 		return nil, protoimpl.X.NewError("failed to decode YAML: %v", err)
// 	}

// 	// Sanitize the types before conversion to protobuf values
// 	val = SanitizeYAMLTypes(val)

// 	return NewValue(val)
// }

// -------------------------------------------------------------- //
