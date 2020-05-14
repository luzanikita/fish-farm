package controllers

import (
	"encoding/json"
	"crypto/md5"
	"encoding/hex"

	"github.com/nigi4/fish-farm/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @Title Post
// @Description Login
// @Param	body		body 	models.Credits	true		"body for Credits content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {
	var v models.Credits
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		user, err := CheckLogin(v.Email, v.Password)
		if err != nil {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
			logs.Error(err)
		}
		if user == nil {
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = "Wrong email or password."
		} else {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = user
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
		logs.Error(err)
	}
	c.ServeJSON()
}

// Get ...
// @Title Get
// @Description get UserId from Session
// @Success 200 {int}
// @Failure 403
// @router / [get]
func (c *LoginController) Get() {
	userId := c.GetSession("userId")
	if userId == nil {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Not authorized."
	} else {
		c.Data["json"] = userId
	}
	c.ServeJSON()
}

func CheckLogin(email, password string) (*models.Users, error){
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = -1
	var offset int64
	
	data := []byte(password)
	hash := md5.Sum(data)
	password_hash := hex.EncodeToString(hash[:])

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		return nil, err
	}
	for _, u := range l {
		user := u.(models.Users)
		if user.Email == email && user.Password == password_hash {
			return &user, nil
		}
	}
	return nil, err
}
