package service

import (
	"github.com/aguncn/nezha/models"
	"github.com/aguncn/nezha/page"
)

//IArticleService Article接口定义
type IArticleService interface {
	//GetArticle 根据id获取Article
	GetArticle(id uint) *models.Article
	//GetTables 分页返回文章
	GetTables(pageNum, pagesize uint) *[]page.Article
	//AddArticle 新增Article
	AddArticle(article *models.Article) bool
	//GetArticles 获取文章信息
	GetArticles(PageNum, PageSize uint, total *uint64, where interface{}) *[]models.Article
	// //ExistArticleByID 根据ID判断Article是否存在
	// ExistArticleByID(id int) bool
	// //EditArticle 编辑Article
	// EditArticle(article models.Article) bool
	// //DeleteArticle 删除Article
	// DeleteArticle(id int) bool
}
