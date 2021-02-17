import "fmt"

func twoSum(nums []int, target int) []int {
   var a [] int
    for i:=0; i<len(nums); i++ {
        for j:=i; j<len(nums); j++ {
            if target == nums[i] + nums[j]{
                a = append(a,i)
                a = append(a,j)
            }
        }
    }
    printSlice(a)
    return a
}
func printSlice(s []int) {
	fmt.Printf("%v\n", s)
}
