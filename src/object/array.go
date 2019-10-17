// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Value []Object
}

func (this Array) Type() TYPE {
	return ARRAY
}

func (this Array) StringValue() string {
	var out bytes.Buffer
	out.WriteString("[")
	result := []string{}
	for _, v := range this.Value {
		if v.Type() == STRING {
			result = append(result, "\""+v.StringValue()+"\"")
		} else {
			result = append(result, v.StringValue())
		}
	}
	out.WriteString(strings.Join(result, ", "))
	out.WriteString("]")
	return out.String()
}

func (this *Array) Push(obj Object) []Object {
	return append(this.Value, obj)
}
