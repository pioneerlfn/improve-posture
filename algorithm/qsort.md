
## Go语言实现

> 实现参考`严蔚敏`版`数据结构与算法分析`

```Go

func quickSort(a []int, low, high int) {
	if low >= high {
		return
	}
	index := partition(a, low, high)
	quickSort(a, 0, index-1)
	quickSort(a, index+1, high)
}

func partition(a []int, low, high int) int {
	pivot := a[low] // 1
	for low < high {
		// 注意，如果选了low为pivot index, 则先从high开始
		for low < high && a[high] >= pivot { // 注意：相等的情况不停
			high--
		}

		// 填坑，low现在是坑: 来自1和3
		a[low] = a[high] // 2
		// 现在high是坑了

		for low < high && a[low] <= pivot {
			low++
		}

		// 开始填坑，坑来自2
		a[high] = a[low] // 3
		// 现在low又是坑了
	}
	a[low] = pivot // 最后注意填坑
	return low
}

```

## 参考
- 王道考研2019年数据结构考研复习指导
- [深入解析快速排序(Quick Sort)](http://www.yebangyu.org/blog/2016/03/09/quicksort/)
