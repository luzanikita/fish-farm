package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"crypto/md5"
	"encoding/hex"
	
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id         int    `orm:"column(id);auto"`
	Email      string `orm:"column(email)"`
	Password   string `orm:"column(password)"`
	FirstName  string `orm:"column(first_name)"`
	SecondName string `orm:"column(second_name);null"`
	Phone      string `orm:"column(phone);null"`
}

type Credits struct {
	Email      string
	Password   string
}

func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

func AddUsers(m *Users) (id int64, err error) {
	data := []byte(m.Password)
	hash := md5.Sum(data)
	password_hash := hex.EncodeToString(hash[:])
	m.Password = password_hash

	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))

	for k, v := range query {

		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}

	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {

			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {

			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Users
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {

			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

func UpdateUsersById(m *Users) (err error) {
	data := []byte(m.Password)
	hash := md5.Sum(data)
	password_hash := hex.EncodeToString(hash[:])
	m.Password = password_hash

	o := orm.NewOrm()
	v := Users{Id: m.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteUsers(id int) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
