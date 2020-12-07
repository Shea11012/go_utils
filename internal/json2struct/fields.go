package json2struct

import (
	"fmt"
	"github.com/Shea11012/go_utils/internal/word"
)

type Field struct {
	Name string
	Type string
}

type FieldValue struct {
	CamelCase bool
	Value string
}

type FieldSegment struct {
	Format string
	FieldValues []FieldValue
}

type Fields []*Field

func (f *Fields) removeDuplicate() Fields {
	m := make(map[string]bool)
	fields := Fields{}
	for _,v := range *f {
		t := v
		if _,value := m[v.Name];!value {
			m[v.Name] = true;
			fields = append(fields,t)
		}
	}

	return fields
}

func (f *Fields) appendSegment(name string, segment FieldSegment) {
	var s []interface{}
	for _,v := range segment.FieldValues {
		value := v.Value
		if v.CamelCase {
			value = word.UnderscoreToUpperCamelCase(v.Value)
		}

		s = append(s,value)
	}

	*f = append(*f,&Field{
		Name: word.UnderscoreToUpperCamelCase(name),
		Type: fmt.Sprintf(segment.Format,s...),
	})
}