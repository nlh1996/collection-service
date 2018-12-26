package user

import (
	"collect/model"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// SaveUser .
func SaveUser(c *gin.Context) {
	form := &model.Form{}
	err := c.Bind(form)
	if err != nil {
		fmt.Println(err)
	}
	// 图片解码，保存至文件服务器
	var (
		enc  = base64.StdEncoding
		path string
		dir  string
	)
	for index, img := range form.Pictures {
		if img.Src != "" {
			if img.Src[11] == 'j' {
				img.Src = img.Src[23:]
				path = fmt.Sprintf("d:/img/%s%s/%s.jpg", form.User.Name, form.User.Phone, img.Name)
				dir = fmt.Sprintf("d:/img/%s%s", form.User.Name, form.User.Phone)
			} else if img.Src[11] == 'p' {
				img.Src = img.Src[22:]
				path = fmt.Sprintf("d:/img/%s%s/%s.png", form.User.Name, form.User.Phone, img.Name)
				dir = fmt.Sprintf("d:/img/%s%s", form.User.Name, form.User.Phone)
			} else {
				fmt.Println("不支持该文件类型")
			}
			data, err := enc.DecodeString(img.Src)
			if err != nil {
				fmt.Println(err.Error())
			}
			//创建目录
			os.MkdirAll(dir, os.ModePerm)
			//图片写入文件
			f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()
			f.Write(data)
			//记录图片保存的地址
			form.Pictures[index].Src = path
		}
	}
	form.Save()
	c.String(http.StatusOK, "保存成功")
}

// Get 获取客户信息
func Get(c *gin.Context) {
	user := model.User{}
	result := user.GetAll()
	var userList []model.User
	for _,value := range *result {
		userList = append(userList,value.User)
	}
	c.JSON(200,gin.H{
		"userList": userList,
	})
}

// Search 查找客户
func Search(c *gin.Context) {
	form := &model.Form{}
	form.User.Name = c.Query("name")
	result := form.Find()
	var userList []model.User
	for _,value := range *result {
		userList = append(userList,value.User)
	}
	if result != nil {
		c.JSON(200,gin.H{
			"userList": userList,
		})
	}
}
