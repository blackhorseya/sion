package stringx

import (
	"reflect"
)

// DifferenceSet serve caller to get difference set of two sting slice
func DifferenceSet(a []string, b []string) []string {
	var ret []string
	temp := map[string]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			ret = append(ret, val)
		}
	}

	return ret
}

// RemoveDuplicateStr serve caller to given string slice to remove duplicate string then return a new string slice
func RemoveDuplicateStr(slice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range slice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}

// Simple has complexity: O(n^2)
func Simple(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	av := reflect.ValueOf(a)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		if contains(b, el) {
			set = append(set, el)
		}
	}

	return set
}

func contains(a interface{}, e interface{}) bool {
	v := reflect.ValueOf(a)

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}

	return false
}
