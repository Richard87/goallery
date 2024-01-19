package pointers

func Int32(i int32) *int32    { return &i }
func Int64(i int64) *int64    { return &i }
func String(s string) *string { return &s }
