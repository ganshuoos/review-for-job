47. 全排列 II

源地址：[47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

问题描述：

>给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
>
> 
>
>示例 1：
>
>输入：nums = [1,1,2]
>输出：
>[[1,1,2],
> [1,2,1],
> [2,1,1]]
>示例 2：
>
>输入：nums = [1,2,3]
>输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
>
>
>提示：
>
>1 <= nums.length <= 8
>-10 <= nums[i] <= 10

``` go
import "fmt"
func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    res :=  make([][]int, 0)
    path := make([]int, 0)
    st := make([]bool, len(nums))

    dfs(nums, 0, &st, &path, &res)
    return res
}

func dfs(nums []int, index int, st *[]bool, path *[]int, res *[][]int) {
    if index == len(nums) {
        //fmt.Println("YES!")
        tmp := make([]int, len(*path))
        copy(tmp, *path)
        *res = append(*res, tmp)
    }

    for i := 0; i < len(nums); i++ {
        if (*st)[i] == true || (i > 0 && nums[i] == nums[i-1] && (*st)[i-1] == false) {continue}

        (*st)[i] = true
        *path = append(*path, nums[i])
        dfs(nums, index+1, st, path, res)
        (*st)[i] = false
        *path = (*path)[:len(*path)-1]
    }
    return
}
```



