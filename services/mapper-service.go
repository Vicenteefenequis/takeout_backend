package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strings"
)

func mapper(data interface{}, dataUpdate *bson.D) {

	dataValue := reflect.ValueOf(data)

	for i := 0; i < dataValue.NumField(); i++ {
		field := strings.Split(dataValue.Type().Field(i).Tag.Get("bson"), ",")[0]
		value := dataValue.Field(i).Interface()

		if !reflect.DeepEqual(value, reflect.Zero(dataValue.Field(i).Type()).Interface()) {
			*dataUpdate = append(*dataUpdate, bson.E{Key: field, Value: value})
		}
	}
}
