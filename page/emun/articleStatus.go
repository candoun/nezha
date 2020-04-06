package emun

var ArticleStatus = map[uint]string{
	0: "error",     //出现未定义状态
	1: "published", //已发布
	2: "draft",     //编辑中
	3: "deleted",   //已删除
}

func GetArticleStatus(code uint) string {
	msg, ok := ArticleStatus[code]
	if ok {
		return msg
	}

	return ArticleStatus[0]
}
