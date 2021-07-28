package gobug

import (
	"reflect"
	"sync"
)

type typeInfo struct {
	typ    reflect.Type
	fields map[string]*fieldInfo
}

type fieldInfo struct {
	index []int
}

var types sync.Map // reflect.Type -> *typeInfo

func typeInfoFor(t reflect.Type) *typeInfo {
	if info, ok := types.Load(t); ok {
		return info.(*typeInfo)
	}

	info := &typeInfo{typ: t}
	info.fields = make(map[string]*fieldInfo)
	info.addFields(t, "", nil)
	types.Store(t, info)
	return info
}

func (ti *typeInfo) addFields(t reflect.Type, prefix string, prefixIdx []int) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.PkgPath != "" {
			continue
		}

		pre := prefix
		if !f.Anonymous {
			if len(pre) > 0 {
				pre = pre + f.Name + "."
			} else {
				pre = f.Name + "."
			}
		}
		switch f.Type.Kind() {
		case reflect.Struct:
			ti.addFields(f.Type, pre, append(prefixIdx, f.Index...))
			continue
		}

		name := f.Name
		finfo := &fieldInfo{
			index: append(prefixIdx, f.Index...),
		}

		switch f.Type.Kind() {
		case reflect.String:
			ti.fields[prefix+name] = finfo
		}
	}
}
