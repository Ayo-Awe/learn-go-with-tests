package reflection

import "reflect"

func walk(x interface{}, fn func(string)) {
	v := getValue(x)

	walkValue := func(v reflect.Value) {
		walk(v.Interface(), fn)
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := range v.Len() {
			walkValue(v.Index(i))
		}
	case reflect.Struct:
		for i := range v.NumField() {
			walkValue(v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			walkValue(v.MapIndex(key))
		}
	case reflect.Chan:
		for val, ok := v.Recv(); ok; val, ok = v.Recv() {
			walkValue(val)
		}
	case reflect.Func:
		for _, val := range v.Call(nil) {
			walkValue(val)
		}
	case reflect.String:
		fn(v.String())
	}

}

func getValue(x any) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	return v
}
