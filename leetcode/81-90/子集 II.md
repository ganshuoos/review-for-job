90. 子集 II

源地址：[90. 子集 II](https://leetcode-cn.com/problems/subsets-ii/)

问题描述：

>给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
>
>说明：解集不能包含重复的子集。
>
>示例:
>
>输入: [1,2,2]
>输出:
>[
>  [2],
>  [1],
>  [1,2,2],
>  [2,2],
>  [1,2],
>  []
>]

``` go
var res [][]int
var path []int
var std  []bool
func subsetsWithDup(nums []int) [][]int {
    res = make([][]int, 0)
    path = make([]int, 0)
    std = make([]bool, len(nums))

    sort.Ints(nums)
    dfs(nums, 0)
    return res
}

func dfs(nums []int, i int) {
    //fmt.Println("---------------------")
    //fmt.Printf("i: %d, path: %v, std: %v, res: %v\n", i, path, std, res)
    //fmt.Println("---------------------")
    //fmt.Println("")
    if i == len(nums) {
        //fmt.Println("here!")
        //fmt.Println("path: ", path)
        tmp := make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)
        return
    }


    if i > 0 && nums[i] == nums[i-1] && std[i-1] == false {
        dfs(nums, i+1)
        return
    }

    if std[i] == false {
        std[i] = true
        path = append(path, nums[i])
        dfs(nums, i+1)
        path =  path[:len(path)-1]
        std[i] = false
        dfs(nums, i+1)
    } 
    
}
```



