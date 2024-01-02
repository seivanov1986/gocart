package helpers

func StringPointer(value string) *string {
	return &value
}

func IntegerPointer(value int64) *int64 {
	return &value
}

func BoolPointer(value bool) *bool {
	return &value
}