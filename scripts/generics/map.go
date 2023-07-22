package main

import (
	"fmt"
)

func GetValue[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultVal
}

func GetOrSetValue[K comparable, V any](m *map[K]V, key K, defaultVal V) V {
	ref := *m
	if v, ok := ref[key]; ok {
		return v
	} else {
		ref[key] = defaultVal
		return defaultVal
	}
}

func main() {

	serverStats := map[string]int{
		"port":      8000,
		"pings":     47,
		"status":    1,
		"endpoints": 13,
	}
	v := GetValue(serverStats, "status", -1)
	fmt.Println(v)
	v = GetValue(serverStats, "cpu", 4)
	fmt.Println(v)
	fmt.Println(serverStats)

	v = GetOrSetValue(&serverStats, "cpu", 4)
	fmt.Println(v)
	fmt.Println(serverStats)


}
