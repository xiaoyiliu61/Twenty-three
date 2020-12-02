package main

import "fmt"

/*func main(){
	arr:=[]int{1,6,9,4}
	fanzhuan(arr)
}
func fanzhuan(arr[] int)  {
	for i := 0; i < len(arr)/2;i++  {
		temp:=arr[i]
		arr[i]=arr[len(arr)-1-i]
		arr[len(arr)-1-i]=temp
	}
	fmt.Println(arr)

}*/
func main() {
  arr1:=[5]int{1,4,6,9,3}
  rs1:=reversr(arr1)
  fmt.Println(rs1)
}
func reversr(arr [5]int) [5]int {
	for index := 0; index < len(arr)/2; index++ {
		//a=arr[index]
		//arr[index]=arr[len(arr)-1-index]
		//arr[len(arr)-1-index]=a
		arr[index],arr[len(arr)-1-index]=arr[len(arr)-1-index],arr[index]
	}
	return arr
}
