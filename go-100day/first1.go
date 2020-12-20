package main

import "fmt"

func main() {
   var arr1 []int
   fmt.Printf("%v - %T - 长度:%v\n",arr1,arr1,len(arr1))

   var arr2 =[]int{54,12,4,6}
   fmt.Printf("%v - %T - 长度:%v\n",arr2,arr2,len(arr2))

   var arr3 = []int{1: 2,2: 4, 5: 6}
   fmt.Printf("%v - %T - 长度:%v\n", arr3, arr3 ,len(arr3))
   fmt.Println("--------------------------------")
   var arr5 []int
   fmt.Println(arr5==nil)
   var str = []string{"php", "java", "node.js", "golang"}
   for i := 0; i < len(str); i++ {
      fmt.Println(str[i])
      for k, v :=range str{
         fmt.Println(k,v)
         
         
      }
      
   }
   



}