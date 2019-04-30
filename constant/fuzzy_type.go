package constant

const(
	// 全文匹配
	FUZZY_FALSE = 0
	// 模糊
	FUZZY_TRUE = 1
)

var FuzzyNameMap = map[int]string{
	FUZZY_FALSE:"全文搜索",
	FUZZY_TRUE:"模糊匹配",
}
