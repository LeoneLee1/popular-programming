package main

import "fmt"

func main() {
   array1 := [3]string{"a", "c", "d"}
   array2 := [3]string{"a", "d", "c"}

   for i := 0; i < len(array1); i++ {
      if i == 1 || i == 2 {
         if array1[i] != array2[i] {
            fmt.Printf("Unsur index %d berbeda: %s vs %s\n", i, array1[i], array2[i])
         }
      }
   }
}
