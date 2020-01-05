package commentService

import (
	mydb "blog/db/mysql"
	model "blog/models/comment"
	"blog/util"
	"log"

	"github.com/gohouse/gorose"
)

var comment model.Comment

// var commentList []model.Comment
func CreateComment(commentUSerId string, byCommentUserId string, commentContent string, commentArticleId string) (model.Comment, error) {
	sqlConn := mydb.Conn()
	commentId := util.GetRowId()
	_, err := sqlConn.Execute("insert into comments(id, articleId, parentId, creatorId, content) values(?,?,?,?,?)", commentId, commentArticleId, byCommentUserId, commentUSerId, commentContent)
	if err != nil {
		log.Println("Join in a new comment --> %s", err.Error())
		return comment, err
	}
	_, err = sqlConn.Table(&comment).Where("id", commentId).First()
	if err != nil {
		return comment, err
	}
	return comment, nil
}
func GetCommentListByArticleId(articleId string) ([]gorose.Data, error) {
	sqlConn := mydb.Conn()
	// commentList = commentList[:0]
	// err := sqlConn.Table(&commentList).Where("articleId", articleId).Select()
	data, err := sqlConn.Query("SELECT comments.*,users.username,users.id AS userId FROM "+
		"comments Left Join users On users.id=comments.creatorId WHERE comments.articleId=?", articleId)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}
