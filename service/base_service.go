package service

// 组装分页
func assemblePage(page,pageSize *int)  {
	*page = (*page-1)* (*pageSize)
}
