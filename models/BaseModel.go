package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"reflect"
	"errors"
	"strings"
	"blog/util"
)

var OperatorsMap map[string]string

func init() {
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	dbname := beego.AppConfig.String("mysqldb")
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local&parseTime=true", user, passwd, host, dbname))

	OperatorsMap = initOperatorsMap()
}

func initOperatorsMap() map[string]string {
	var operatorsMap = make(map[string]string)
	var prefix = "__"
	operatorsMap["="] = prefix + "exact"
	operatorsMap["like"] = prefix + "icontains"
	operatorsMap["in"] = prefix + "in"
	operatorsMap[">"] = prefix + "gt"
	operatorsMap[">="] = prefix + "gte"
	operatorsMap["<"] = prefix + "lt"
	operatorsMap["<="] = prefix + "lte"
	operatorsMap["rlike"] = prefix + "startswith"
	operatorsMap["llike"] = prefix + "endswith"
	operatorsMap["isnull"] = prefix + "isnull"
	operatorsMap["rllike"] = prefix + "icontains"
	operatorsMap["istartswith"] = prefix + "istartswith"
	return operatorsMap
}

type ConditionStruct struct {
	field string
	flag  string
	value interface{}
}

func NewConditionStruct(field string, flag string, value interface{}) *ConditionStruct {
	return &ConditionStruct{field, flag, value}
}

type BaseModal struct {
	TableName string
}

func (self *BaseModal) newOrm() orm.Ormer {
	return orm.NewOrm()
}

func (self *BaseModal) selectTable() orm.QuerySeter {
	return self.newOrm().QueryTable(self.TableName)
}

func (self *BaseModal) InsertDataToTable(insertStruct interface{}) (int64, error) {
	return self.newOrm().Insert(insertStruct)
}

func (self *BaseModal) UpdateDataInTableCanSetCondition(udpateParams orm.Params, ConditionStructs ...*ConditionStruct) (int64, error) {
	return self.AppendConditionFilter(ConditionStructs...).Update(udpateParams)
}

func (self *BaseModal) GetDataCanSetConditionAndLimit(pageUtil util.PageUtil, ConditionStructs ...*ConditionStruct) orm.QuerySeter {
	return self.AppendConditionFilter(ConditionStructs...).Limit(pageUtil.PageSize, pageUtil.PageNo)
}

func (self *BaseModal) GetCountTableCanSetCondition(ConditionStructs ...*ConditionStruct) (int64, error) {
	return self.AppendConditionFilter(ConditionStructs...).Count()
}

func (self *BaseModal) AppendConditionFilter(ConditionStructs ...*ConditionStruct) orm.QuerySeter {
	o := self.selectTable()
	if len(ConditionStructs) == 0 {
		return o
	} else {
		for _, condition := range ConditionStructs {
			suffix, ok := OperatorsMap[condition.flag]
			if ok == false {
				continue
			}
			field := strings.ToUpper(string(condition.field[0])) + condition.field[1:]
			o = o.Filter(field+suffix, condition.value)
		}
	}
	return o
}

func (self *BaseModal) MakeOrmParamsByStruct(entity interface{}, updateKeys ...string) (orm.Params, error) {
	var params orm.Params = make(orm.Params)
	value := reflect.ValueOf(entity)
	if value.Elem().Kind().String() != "struct" {
		return params, errors.New("本函数只处理struct")
	}
	num := value.Elem().NumField()
	if num <= 0 {
		return params, errors.New("struct中没有元素")
	}
	if len(updateKeys) <= 0 {
		return params, nil
	}
	for _, key := range updateKeys {
		key = strings.ToUpper(string(key[0])) + key[1:]
		/**
			elem-获取value这个interface包装的值
			fieldByName获取包装的值的某个key的value
		 */
		for i := 0; i < num; i++ {
			if value.Elem().FieldByName(key).IsValid() == true {
				params[key] = value.Elem().FieldByName(key).Interface()
				continue
			}
		}
	}
	return params, nil
}
