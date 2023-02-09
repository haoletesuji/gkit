package gkit

import "google.golang.org/protobuf/types/known/wrapperspb"

func StringToStringValue(value *string) *wrapperspb.StringValue {
	if value != nil {
		return wrapperspb.String(*value)
	}

	return nil
}
