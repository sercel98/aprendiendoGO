package main

import "fmt"

func main() {

	string1 := "hi}"
	string2 := "world"

	result := twoStrings(string1, string2)
	fmt.Println(result)
}
func twoStrings(s1 string, s2 string) string {

	hashString1 := makeHash(s1)
	hashString2 := makeHash(s2)

	fmt.Println(hashString1)
	fmt.Println(hashString2)

	if s1 == s2 {
		return "YES"
	}

	for _, sub := range hashString1 {

		if hashString2[sub] != "" {
			return "YES"
		}

	}

	return "NO"

}

func makeHash(s string) map[string]string {

	resultMap := make(map[string]string)
	//iterates over the string
	for j := 0; j < len(s); j++ {

		if j+1 > len(s) {
			break
		}
		substring := s[j : j+1]
		resultMap[substring] = substring
	}

	return resultMap
}
