package objutils

import (
	"fmt"
	"reflect"
)

func CopyProperties(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("both src and dst must be pointers")
	}

	srcVal = srcVal.Elem()
	dstVal = dstVal.Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		sFieldName := srcVal.Type().Field(i).Name
		sFieldValue := srcVal.Field(i)

		dField := dstVal.FieldByName(sFieldName)
		if !dField.IsValid() {
			continue
		}
		if dField.Type().String() != srcVal.FieldByName(sFieldName).Type().String() {
			continue
		}

		if dField.CanSet() {
			dField.Set(sFieldValue)
		}
	}

	return nil
}
