package main

func Repeat(pattern string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += pattern
	}
	return repeated
}
