import (
        "math/bits"
        "fmt"
        )
        
func mirrorBits(a int) int {
   var mirror int
   if a < 256 {
      mirror = int(bits.Reverse8(uint8(a)))      
   } else {
      mirror = int(bits.Reverse32(uint32(a)))
   }
   fmt.Println(mirror)
   for mirror&1 != 1 {
      mirror >>= 1
   }
   return mirror
}

func secondRightmostZeroBit(n int) int {
  return -(^(n|(n+1))|((n|(n+1))+1))
}

func main() {
  ###mirrorBits
  a := 97
  fmt.Println(a, mirrorBits(a))
  a := 8
  fmt.Println(a, mirrorBits(a))
  a := 123
  fmt.Println(a, mirrorBits(a))
  
  ###secondRightmostZeroBit
  a := 37
  fmt.Println(a, secondRightmostZeroBit(a))
  a := 1073741824
  fmt.Println(a, secondRightmostZeroBit(a))
  a := 83748
  fmt.Println(a, secondRightmostZeroBit(a))
}
