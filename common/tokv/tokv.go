package tokv

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v3"
)

type KV map[string]string

func JsonToKv(v string) (res KV) {
	res = make(KV)
	var obj interface{}
	err := json.Unmarshal([]byte(v), &obj)
	if err != nil {
		return
	}
	return objToKv(obj)
}

func YamlToKv(v string) (res KV) {
	res = make(KV)
	var obj interface{}
	err := yaml.Unmarshal([]byte(v), &obj)
	if err != nil {
		return
	}
	return objToKv(obj)
}

func objToKv(obj interface{}) KV {
	kv := make(KV)
	extractPaths(obj, "", &kv)
	return kv
}

func extractPaths(value interface{}, prefix string, kv *KV) {
	// 获取值的类型
	val := reflect.ValueOf(value)

	// 如果值是一个map，则递归遍历
	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			// 获取键对应的值
			fieldVal := val.MapIndex(key)
			// 构建路径
			newPrefix := fmt.Sprintf("%s.%v", prefix, key)
			// 递归处理
			extractPaths(fieldVal.Interface(), newPrefix, kv)
		}
	case reflect.Struct:
		// 如果是结构体类型，获取字段信息并递归处理
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			// 获取字段值
			fieldVal := val.Field(i)
			// 跳过未导出字段
			if field.PkgPath != "" {
				continue
			}
			// 递归处理
			newPrefix := fmt.Sprintf("%s.%s", prefix, field.Name)
			extractPaths(fieldVal.Interface(), newPrefix, kv)
		}
	case reflect.Slice:
		// 如果是切片，遍历每个元素
		for i := 0; i < val.Len(); i++ {
			// 获取切片中的每个元素的值
			elementVal := val.Index(i)
			// 构建新的路径
			newPrefix := fmt.Sprintf("%s.%d", prefix, i)
			// 递归处理
			extractPaths(elementVal.Interface(), newPrefix, kv)
		}
	default:
		// 如果值是一个基本类型并且不是空值，则添加路径
		if val.IsValid() && !isEmpty(val) {
			(*kv)[prefix[1:]] = fmt.Sprintf("%v", val)
		}
	}
}

// 判断值是否为空
func isEmpty(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Map, reflect.Array, reflect.Slice:
		return val.Len() == 0
	case reflect.Ptr:
		return val.IsNil()
	default:
		return false
	}
}
