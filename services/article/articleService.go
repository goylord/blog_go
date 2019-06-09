package articleService

import(
	// "blog/models/article"
	"blog/util"
	"strconv"
	mydb"blog/db/mysql"
	"fmt"
)

func CreateNewArticle( title string, content string, previewContent string, creatorId string, classId string) (string, error) {
	var sqlConn = mydb.Conn()
	var articleId = util.GetRowId()
	_,err := sqlConn.Execute("insert into articles(title, content, previewContent, creatorId, classId, id) values(?,?,?,?,?,?)", title, content, previewContent, creatorId, classId, articleId)
	if err != nil {
		fmt.Println("Create article failed, error: %s\n", err.Error())
		return "", err
	}
	return articleId, nil
}
func GetList(pageNo string, pageSize string) ([]map[string]interface{}, error){
	// var articleList []model.Articles
	var sqlConn = mydb.Conn()
	pageNoInt, err := strconv.Atoi(pageNo)
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		fmt.Println("Get article list failed,error:", err.Error())
		pageNoInt = 1
		pageSizeInt = 10
	}
	// err = sqlConn.Table(&articleList).Join("users","articles.creatorId","=","users.id").Limit(pageSizeInt).Offset((pageNoInt - 1) * pageSizeInt).Select()
	data,err := sqlConn.Query("SELECT articles.*,users.username,files.filename FROM " +
	"articles Left Join users On users.id=articles.creatorId Left Join files On files.assoId=articles.id LIMIT ?" + 
	" OFFSET ?", pageSizeInt, (pageNoInt - 1) * pageSizeInt)
	if err != nil {
		fmt.Println("Get article list failed,error:", err.Error(), pageNoInt, pageSizeInt)
		return nil, err
	}
	return data, nil
}
func DeleteArticle(id string) (bool, error) {
	var sqlConn = mydb.Conn()
	_, err := sqlConn.Table("articles").Where("id", id).Delete()
	if err != nil {
		fmt.Println("Delete article failed!error:%s", err.Error())
		return false, err
	}
	return true, nil
}