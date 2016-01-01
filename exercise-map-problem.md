[https://tour.go-zh.org/moretypes/20](https://tour.go-zh.org/moretypes/20 "map-practice")
#**练习：map**
实现 `WordCount`。它应当返回一个含有 `s` 中每个 “词” 个数的 `map`。函数 `wc.Test` 针对这个函数执行一个测试用例，并输出成功还是失败。

你会发现 `strings.Fields` 很有帮助。

#**运行结果**
	PASS
    f("I am learning Go!") = 
  	map[string]int{"am":1, "learning":1, "Go!":1, "I":1}
	PASS
 	f("The quick brown fox jumped over the lazy dog.") = 
  	map[string]int{"jumped":1, "lazy":1, "The":1, "brown":1, "dog.":1, "fox":1, "over":1, "quick":1, "the":1}
	PASS
 	f("I ate a donut. Then I ate another donut.") = 
  	map[string]int{"another":1, "ate":2, "donut.":2, "I":2, "Then":1, "a":1}
	PASS
 	f("A man a plan a canal panama.") = 
  	map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}

