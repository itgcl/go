package db

import (
	"errors"
)

const (
	Gt = ">"      // 大于
	Ge = ">"      // 大于等于
	Lt   = "<"      // 小于
	Le   = "<"      // 小于等于
	Eq  = "="      // 等于
	Ne =  "!="     //不等于
	Is      = "is"     // is
	IsNot   = "is not" // is not
)
type Search struct {
	DB *DataBase
	orWheres []*OrWhereCondition
	wheres []*WhereCondition
}

type WhereCondition struct {
	ColumnName string
	Operator string
	Value interface{}
}

type OrWhereCondition struct {
	ColumnName string
	Operator string
	Value interface{}
}

func NewOrWhereCondition() *OrWhereCondition {
	orWhereCondition := new(OrWhereCondition)
	orWhereCondition.Operator = Eq
	return orWhereCondition
}


func NewWhereCondition() *WhereCondition {
	whereCondition := new(WhereCondition)
	whereCondition.Operator = Eq
	return whereCondition
}

func (search *Search) Where(condition ...interface{}) *Search {
	conditionLength := len(condition)
	whereCondition := NewWhereCondition()
	switch conditionLength {
	case 2:
		columnName,ok := condition[0].(string)
		if !ok {
			panic(errors.New("column name not string"))
		}
		whereCondition.ColumnName = columnName
		whereCondition.Value = condition[1]
	case 3:
		columnName,ok := condition[0].(string)
		if !ok {
			panic(errors.New("column name not string"))
		}
		operator,ok := condition[1].(string)
		if !ok {
			panic(errors.New("operator name not string"))
		}
		whereCondition.ColumnName = columnName
		whereCondition.Operator = operator
		whereCondition.Value = condition[2]
	default:
		panic(errors.New("condition number errors"))
	}
	search.wheres = append(search.wheres, whereCondition)
	return search
}


func (search *Search) OrWhere(condition ...interface{}) *Search {
	conditionLength := len(condition)
	orWhereCondition := NewOrWhereCondition()
	switch conditionLength {
	case 2:
		columnName,ok := condition[0].(string)
		if !ok {
			panic(errors.New("column name not string"))
		}
		orWhereCondition.ColumnName = columnName
		orWhereCondition.Value = condition[1]
	case 3:
		columnName,ok := condition[0].(string)
		if !ok {
			panic(errors.New("column name not string"))
		}
		operator,ok := condition[1].(string)
		if !ok {
			panic(errors.New("operator name not string"))
		}
		orWhereCondition.ColumnName = columnName
		orWhereCondition.Operator = operator
		orWhereCondition.Value = condition[2]
	default:
		panic(errors.New("condition number errors"))
	}
	search.orWheres = append(search.orWheres, orWhereCondition)
	return search
}


