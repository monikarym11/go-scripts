import (
        "fmt"
        )
        
func palindromeRearranging(inputString string) bool {
    count := make(map[rune]int)
    for _, ch := range inputString {
        count[ch]++
    }
    np := 0
    for _, v := range count {
        if v%2 != 0 {
            np++
        }
    }
    if np > 1 {
        return false        
    } else {
        return true
    }

}
func checkPalindrome(inputString string) bool {
  return inputString==reverse(inputString) 
}
func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}
func main() {
  inputString1 := "aabaa"
  inputString2 := "abac"
  inputString3 := "abbcabb"
  inputString4 := "abca"

  fmt.Println(checkPalindrome(inputString1), checkPalindrome(inputString2))
  fmt.Println(palindromeRearranging(inputString3), palindromeRearranging(inputString4))
  
}
