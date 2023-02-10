package gkit

import (
	"time"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func StringToStringValue(value *string) *wrapperspb.StringValue {
	if value != nil {
		return wrapperspb.String(*value)
	}

	return nil
}

func StringValueToString(value *wrapperspb.StringValue) *string {
	v := value.GetValue()
	if len(v) > 0 {
		return &v
	}
	return nil
}

func ParseTimeYMD(value string) *time.Time {
	if len(value) == 0 {
		return nil
	}

	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil
	}

	return &date
}

func TimeToYMDStringValue(t *time.Time) *wrapperspb.StringValue {
	if t != nil {
		value := t.Format("2006-01-02")
		return wrapperspb.String(value)
	}
	return nil
}
