package model

import (
	"collect/database"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Form .
type Form struct {
	User     `json:"form1"`
	Pictures []Picture `json:"form2"`
}

// User .
type User struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	Phone        string `json:"phone"`
	Chepai       string `json:"chepai"`
	Company      string `json:"company"`
	Chexing      string `json:"chexing"`
	Color        string `json:"color"`
	GPS          string `json:"gps"`
	Part         string `json:"part"`
	Jinji        string `json:"jinji"`
	Jinjiphone   string `json:"jinjiphone"`
	Jinrong      string `json:"jinrong"`
	Koufen       int    `json:"koufen"`
	Shangpaidate string `json:"shangpaidate"`
	Baoxiandate  string `json:"baoxiandate"`
	Birth        string `json:"birth"`
	Beizhu       string `json:"beizhu"`
}

// Picture .
type Picture struct {
	Name string `json:"name"`
	Src  string `json:"src"`
}

// Save 保存用户信息。
func (p *Form) Save() {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("user")
	err := c.Insert(p)
	if err != nil {
		log.Fatal(err)
	}
}

// Find 查找用户。
func (p *Form) Find() *[]Form {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("user")
	users := &[]Form{}
	err := c.Find(bson.M{"user.name": p.User.Name}).All(users)
	if err != nil {
		return nil
	}
	return users
}

// GetAll 获取所有客户相关信息。
func (p *User) GetAll() *[]Form {
	session := database.Session.Clone()
	defer session.Close()
	c := session.DB("test").C("user")
	result := &[]Form{}
	err := c.Find(nil).Select(bson.M{
		"user.name": 1,"user.phone": 1,"user.chepai": 1,"user.chexing": 1,"user.beizhu": 1}).All(result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
