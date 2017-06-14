package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"log"
	"os"

	"crypto/md5"
	"qiniupkg.com/api.v7/kodo"
	"regexp"
)

var qiniuAccessKey string
var qiniuSerectKey string
var qiniuImageBucketname string
var qiniuImageHost string
var qiniuVideoBucketname string
var qiniuVideoHost string

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type UEController struct {
	beego.Controller
}

func init() {
	iniConf, _ := config.NewConfig("ini", "conf/app.conf")

	qiniuAccessKey = iniConf.String("qiniuaccesskey")
	qiniuSerectKey = iniConf.String("qiniusercetkey")
	qiniuImageBucketname = iniConf.String("qiniuimagebucketname")
	qiniuImageHost = iniConf.String("qiniuimagehost")
	qiniuVideoBucketname = iniConf.String("qiniuvideobucketname")
	qiniuVideoHost = iniConf.String("qiniuvideohost")

	fmt.Println("qiniuAccessKey :", qiniuAccessKey)
	fmt.Println("qiniuSerectKey :", qiniuSerectKey)
	fmt.Println("qiniuImageBucketname :", qiniuImageBucketname)
	fmt.Println("qiniuImageHost :", qiniuImageHost)
	fmt.Println("qiniuVideoBucketname :", qiniuVideoBucketname)
	fmt.Println("qiniuVideoHost :", qiniuVideoHost)

}

func (c *UEController) Get() {

	op := c.Input().Get("action")
	// key := c.Input().Get("key") //这里进行判断各个页面，如果是addtopic，如果是addcategory
	switch op {
	case "config": //这里是conf/config.json
		file, err := os.Open("conf/config.json")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		fd, err := ioutil.ReadAll(file)
		src := string(fd)
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/") //参考php的$CONFIG = json_decode(preg_replace("/\/\*[\s\S]+?\*\//", "", file_get_contents("config.json")), true);

		src = re.ReplaceAllString(src, "")
		tt := []byte(src)
		var r interface{}
		json.Unmarshal(tt, &r) //这个byte要解码
		c.Data["json"] = r
		c.ServeJSON()

	case "uploadimage", "uploadfile", "uploadvideo":

	}

}

func (controller *UEController) Post() {
	op := controller.Input().Get("action")
	switch op {
	case "uploadimage":

		upload(controller, qiniuImageBucketname, qiniuImageHost)

	case "uploadvideo":
		upload(controller, qiniuVideoBucketname, qiniuVideoHost)
	}

}

func upload(controller *UEController, bucketname string, host string) {

	//保存上传的图片
	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
	fileName := saveFile(controller)

	filePath := "static/upload/" + fileName

	uploadToQiniu(controller, filePath, fileName, bucketname, host)

}

func uploadToQiniu(controller *UEController, filePath string, filename string, bucketname string, host string) {
	kodo.SetMac(qiniuAccessKey, qiniuSerectKey)
	zone := 0                     // 您空间(Bucket)所在的区域
	client := kodo.New(zone, nil) // 用默认配置创建 Client

	bucket := client.Bucket(bucketname)
	ctx := context.Background()

	var ret PutRet

	error := bucket.PutFile(ctx, &ret, filename, filePath, nil)

	fmt.Println("uploadqiu ret", ret)
	if error != nil {
		fmt.Println("uploaderor", error)
	}

	url := host + ret.Key
	controller.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": url, "title": "1", "original": "2"}

	//删除零食文件
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("delete temp file fail", err)
	}

	controller.ServeJSON()
}

//保存在本地
func saveFile(controller *UEController) string {

	file, _, err := controller.GetFile("upfile")
	if err != nil {
		beego.Error(err)
	}

	md5h := md5.New()
	io.Copy(md5h, file)

	md5h.Sum(nil)
	filemd5 := md5h.Sum(nil)

	md5str1 := fmt.Sprintf("%x", filemd5)

	controller.SaveToFile("upfile", "static/upload/"+md5str1)

	return md5str1

}
