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

type resultResp struct{
	Data interface{} `json:"data" form:"data" query:"data"`
	ErrMsg string `json:"err_msg" form:"err_msg" query:"err_msg"`
	Code int `json:"code" form:"code" query:"code"`
	Page int `json:"page" form:"page" query:"page"`
	PageSize int `json:"page_size" form:"page_size" query:"page_size"`
	Count int `json:"count" form:"count" query:"count"`
	LastPage int `json:"last_page" form:"last_page" query:"last_page"`
}

const (
	succeesCode = 200
	// 参数校验异常
	fali_param = 412
	// 服务器错误
	fail_service = 500
)

func succeesDataByPage(data interface{},page,pageSize,count int)resultResp  {
	var lastPage int
	lastPage = count/pageSize
	if count%pageSize>0{
		lastPage++
	}
	return resultResp{
		Data:data,
		Code:succeesCode,
		Count:count,
		LastPage:lastPage,
		Page:page,

	}
}

func succeesData(data interface{})resultResp  {
	return resultResp{
		Data:data,
		Code:succeesCode,
	}
}

func failError(err error,code int)resultResp  {
	return resultResp{
		ErrMsg:err.Error(),
		Code:code,
	}
}
