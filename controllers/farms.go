package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/nigi4/fish-farm/models"
	"github.com/nigi4/fish-farm/stats"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// FarmsController operations for Farms
type FarmsController struct {
	beego.Controller
}

// URLMapping ...
func (c *FarmsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Farms
// @Param	body		body 	models.Farms	true		"body for Farms content"
// @Success 201 {int} models.Farms
// @Failure 403 body is empty
// @router / [post]
func (c *FarmsController) Post() {
	var v models.Farms
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddFarms(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
			logs.Error(err)
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
		logs.Error(err)
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Farms by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Farms
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FarmsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetFarmsById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Farms
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Farms
// @Failure 403
// @router / [get]
func (c *FarmsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = -1
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllFarms(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Farms
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Farms	true		"body for Farms content"
// @Success 200 {object} models.Farms
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FarmsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Farms{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateFarmsById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
		logs.Error(err)
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Farms
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FarmsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteFarms(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetStats ...
// @Title GetStats
// @Description get statistics for each metric on farm
// @Param	id		path 	string	true		"Id of farm you want to analyse"
// @Success 200 {object} models.Farms
// @Failure 404 id is empty
// @router /:id/stats [get]
func (c *FarmsController) GetStats() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = -1
	var offset int64

	idStr := c.Ctx.Input.Param(":id")
	query["FarmId"] = idStr

	sortby = append(sortby, "Date")
	order = append(order, "asc")

	l, err := models.GetAllConditions(query, fields, sortby, order, offset, limit)
	if err != nil || len(l) == 0 {
		c.Abort("404")
		c.TplName = "views/index.tpl"
	} else {
		conditions, err := RequestStats(l)

		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = conditions
		}
		c.ServeJSON()
	}
}

func RequestStats(l []interface{}) (*stats.ConditionsStatsResponse, error) {
	client, _ := stats.InitGrpcConnection()

	var conditions []*stats.Condition
	for i := range l {
		modelCond := l[i].(models.Conditions)
		statsCond := stats.Condition{
			MetricId:      int32(modelCond.MetricId.Id),
			Value:         float32(modelCond.Value),
			UnixTimestamp: modelCond.Date.Unix(),
		}
		conditions = append(conditions, &statsCond)
	}
	out, err := client.MyStats(conditions)

	return out, err
}
