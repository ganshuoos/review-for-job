40. 组合总和 II

源地址：[40. 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/)

问题描述：

>给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
>
>candidates 中的每个数字在每个组合中只能使用一次。
>
>说明：
>
>所有数字（包括目标数）都是正整数。
>解集不能包含重复的组合。 
>示例 1:
>
>输入: candidates = [10,1,2,7,6,1,5], target = 8,
>所求解集为:
>[
>  [1, 7],
>  [1, 2, 5],
>  [2, 6],
>  [1, 1, 6]
>]
>示例 2:
>
>输入: candidates = [2,5,2,1,2], target = 5,
>所求解集为:
>[
>  [1,2,2],
>  [5]
>]

``` go
func combinationSum2(candidates []int, target int) [][]int {
    res := make([][]int, 0)
    path := make([]int, 0)
    sort.Ints(candidates)
    dfs(candidates, 0, target, &res, &path)
    return res
}

func dfs(candidates []int, index int, target int, res *[][]int, path *[]int) {
    if target == 0 {
        tmp := make([]int, len(*path))
        copy(tmp, *path)
        *res = append(*res, tmp)
        return
    }

    if index == len(candidates) {return}

    k := index + 1
    for k < len(candidates) && candidates[k] == candidates[index] {k += 1}
    cnt := k - index

    for i := 0; i * candidates[index] <= target && i <= cnt; i++ {
        dfs(candidates, k, target - i * candidates[index], res, path)
        *path = append(*path, candidates[index])
    }

    for i := 0; i * candidates[index] <= target && i <= cnt; i++ {
        *path = (*path)[:len(*path)-1]
    }
}
```



