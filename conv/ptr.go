package conv

func IntPtr(v int) *int {
	return &v
}

func Int32Ptr(v int32) *int32 {
	return &v
}

func Int64Ptr(v int64) *int64 {
	return &v
}

func StrPtr(v string) *string {
	return &v
}
