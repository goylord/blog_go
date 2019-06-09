package fileService
import (
	"blog/util"
	"log"
	"os"
	"io"
	"mime/multipart"
	mydb"blog/db/mysql"	
	"path"
)
var dst = util.GetCurrentDirectory() + "/publish/"
func UploadFile(file *multipart.FileHeader, assoId string) error {
	log.Println(dst)
	src, err := file.Open()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer src.Close()
	_, err = os.Stat(dst)
	if err !=nil && os.IsNotExist(err) {
		err := os.Mkdir(dst, os.ModePerm)
		if err != nil {
			log.Println("创建文件上传目录出错：", err.Error())
		}
	} else if err != nil {
		log.Println("获取文件夹状态失败：", err.Error())
		return err
	}
	var id = util.GetRowId()
	out, err := os.Create(dst + id + path.Ext(file.Filename))
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}
	err = CreateFileInfo(file.Filename, assoId, id)
	if err != nil {
		os.Remove(dst + id + path.Ext(file.Filename))
		return err
	}
	return nil
}
func CreateFileInfo(orgName string, assoId string, id string) error {
	sqlConn := mydb.Conn()
	_, err := sqlConn.Execute("insert into files(id, filename, assoId, orgName) values(?,?,?,?)", id, id + path.Ext(orgName), assoId, orgName)
	return err
}