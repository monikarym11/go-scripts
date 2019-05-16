import "fmt"
func almostIncreasingSequence(sequence []int) bool {
    counter := 0
    
    if sequence[1] <= sequence[0]{
        counter++
    }
    
    for i := 2; i < len(sequence); i++{
        
        if sequence[i] <= sequence[i-1]{            
            counter ++ 
            if i < len(sequence)-1{
                if sequence[i+1] <= sequence[i-1] && sequence[i] <= sequence[i-2]{
                    counter ++
                }
            }
        }
        if counter > 1{
            fmt.Println(counter)
            return false
        }
    }
    fmt.Println(counter)
    return true
}
func main() {
  sequence1 := [1, 3, 2, 1]
  sequence2 := [1, 3, 2]
  fmt.Println(almostIncreasingSequence(sequence1), almostIncreasingSequence(sequence2))
}
