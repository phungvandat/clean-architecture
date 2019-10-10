package engine

import (
	"go.mongodb.org/mongo-driver/bson"
)

func TranslateQueryToMongoFilter(query *Query) interface{} {
	m := bson.M{}
	for _, filter := range query.Filters {
		switch filter.Condition {
		case Equal:
			m[filter.Property] = filter.Value
		case LessThan:
			m[filter.Property] = bson.M{"$lt": filter.Value}
		case LessThanOrEqual:
			m[filter.Property] = bson.M{"$lte": filter.Value}
		case GreaterThan:
			m[filter.Property] = bson.M{"$gt": filter.Value}
		case GreaterThanOrEqual:
			m[filter.Property] = bson.M{"$gte": filter.Value}
		case And:
			deepQuerys := (filter.Value).([]*Query)
			andQuery := []interface{}{}
			andValue, ok := m["$and"]
			if ok {
				andQuery = (andValue).([]interface{})
			}

			for _, deepQuery := range deepQuerys {
				deepM := TranslateQueryToMongoFilter(deepQuery)
				andQuery = append(andQuery, deepM)
			}
			m["$and"] = andQuery
		case Or:
			deepQuerys := (filter.Value).([]*Query)
			orQuery := []interface{}{}
			orValue, ok := m["$or"]
			if ok {
				orQuery = (orValue).([]interface{})
			}

			for _, deepQuery := range deepQuerys {
				deepM := TranslateQueryToMongoFilter(deepQuery)
				orQuery = append(orQuery, deepM)
			}
			m["$or"] = orQuery
		}
	}
	return m
}
