package models

import (
	"fmt"
)




type Filter struct{
	Limit string `query:"limit"`
	Offset string `query:"offset"`
	Marks []string `query:"mark"`
	Models []string `query:"model"`
	Year []string `query:"year"`
	OwnerName []string `query:"owner_name"`
	OwnerSurname []string `query:"owner_surname"`
	OwnerPatronymic []string `query:"owner_patronymic"`
}


func (f *Filter) ParseToSql(query string) (string, []interface{}){
	counter := 1
	args := []interface{}{}
	// build sql query
	f.parseParams(&query, &args, "mark", &counter, f.Marks)
	f.parseParams(&query, &args, "model", &counter, f.Models,)
	f.parseParams(&query, &args, "year", &counter, f.Year)
	f.parseParams(&query, &args, "owner_name", &counter, f.OwnerName)
	f.parseParams(&query, &args, "owner_surname", &counter, f.OwnerSurname)
	f.parseParams(&query, &args, "owner_patronymic", &counter, f.OwnerPatronymic)
	if f.Limit == ""{
		f.Limit = "10"
	}
	if f.Offset == ""{
		f.Offset = "0"
	}
	query += fmt.Sprintf(" LIMIT %v OFFSET %v", f.Limit, f.Offset)
	return query, args
}



func (f *Filter) parseParams(query *string, args *[]interface{}, coloumn string, counter *int, params []string) {
	if len(params) != 0{
		*query += " AND ("
		for index, value := range params{
			*query += fmt.Sprintf(" %v=$%v", coloumn, *counter)
			*args = append(*args, value)
			if index + 1 != len(params){
				*query += " OR"
			}else{
				*query += ")"
			}
			*counter += 1
		}
	}
}
