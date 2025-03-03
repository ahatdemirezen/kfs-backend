package utils

import "reflect"

// GetFileKey, bir nesnenin FileKey özelliğini alır.
func GetFileKey(record interface{}) (string, bool) {
	val := reflect.ValueOf(record)

	// Eğer record bir pointer ise, pointer'ın gösterdiği değeri al
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// FileKey alanını kontrol et
	field := val.FieldByName("FileKey")
	if field.IsValid() && field.Kind() == reflect.String {
		return field.String(), true
	}

	return "", false
}
