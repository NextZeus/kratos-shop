package pagination

func GetPagOffset(pageNum, pageSize int64) int64 {
	return (pageNum - 1) * pageSize
}
