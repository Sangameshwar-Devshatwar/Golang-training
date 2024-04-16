package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter string1 for anagram comparison")
	fmt.Scan(&str)
	var str1 string
	fmt.Println("Enter string2 for anagram comparison")
	fmt.Scan(&str1)

	result := isAnagram(str, str1)
	fmt.Println(result)
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) && len(t) != len(s) {
		return false
	}
	var cnt int
	var cnts int
	for i := 0; i < len(s); i++ {

		for j := 0; j < len(t); j++ {

			if s[i] == t[j] && t[j] == s[i] {
				cnt++
				break

			}
		}
	}
	for k := 0; k < len(s); k++ {
		for l := k + 1; l < len(t); l++ {
			if s[k] == s[l] {
				cnts++
			}
			if t[k] == t[l] {
				cnts++
			}
		}
	}
	if cnt-cnts == len(s) {
		return true
	} else {
		return false
	}

}
