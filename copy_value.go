package copy_value

import "reflect"

// Copy from src to dst assuming both are pointer to something.
func CopyValue(dst, src interface{}) {
	copyRecursively(reflect.ValueOf(dst).Elem(), reflect.ValueOf(src).Elem())
}

func copyRecursively(dst, src reflect.Value) {
	if dst.Kind() != src.Kind() {
		return
	}
	switch dst.Kind() {
	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			s := src.Field(i)
			d := dst.FieldByName(src.Type().Field(i).Name)
			if d.IsValid() {
				copyRecursively(d, s)
			}
		}
	case reflect.Slice, reflect.Array:
		length := src.Len()
		if length > dst.Len() {
			length = dst.Len()
		}
		for i := 0; i < length; i++ {
			copyRecursively(dst.Index(i), src.Index(i))
		}
	case reflect.Map:
		if dst.Type().Key() == src.Type().Key() {
			for _, key := range src.MapKeys() {
				//copyRecursively(d, src.MapIndex(key))
				dst.SetMapIndex(key, src.MapIndex(key))
			}
		}
	case reflect.Ptr:
		copyRecursively(dst.Elem(), src.Elem())
	default:
		if dst.Type() == src.Type() {
			dst.Set(src)
		}
	}
}
