package controller

const (
	// 访问结果key
	IRIS_DATA = "DataList"
	// 访问总条目数key
	IRIS_TOTAL = "Total"
	// 访问页数
	IRIS_PAGE = "page"
)

func ckeckPage(page,pageSize *int)  {
	if *page < 1{
		*page = 1
	}

	if *pageSize < 1{
		*pageSize = 2
	}
}
