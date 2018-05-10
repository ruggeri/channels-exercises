package main

func reverse(s string) string {
	result := ""
	for _, char := range s {
		// O(n**2) time but whatever.
		result = string(char) + result
	}

	return result
}
