package slice

// Int64SliceToInterfaceSlice .
func Int64SliceToInterfaceSlice(slice []int64) []interface{} {
	iterfaceSlice := make([]interface{}, len(slice))
	for i := range slice {
		iterfaceSlice[i] = slice[i]
	}
	return iterfaceSlice
}

// StringSliceToInterfaceSlice .
func StringSliceToInterfaceSlice(slice []string) []interface{} {
	iterfaceSlice := make([]interface{}, len(slice))
	for i := range slice {
		iterfaceSlice[i] = slice[i]
	}
	return iterfaceSlice
}
