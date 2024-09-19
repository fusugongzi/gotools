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

func PtrInt(v *int, dv int) int {
	if v == nil {
		return dv
	}
	return *v
}

func PtrInt32(v *int32, dv int32) int32 {
	if v == nil {
		return dv
	}
	return *v
}

func PtrInt64(v *int64, dv int64) int64 {
	if v == nil {
		return dv
	}
	return *v
}

func PtrStr(v *string, dv string) string {
	if v == nil {
		return dv
	}
	return *v
}
