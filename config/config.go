package config

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const tag = "env"

func Env(value interface{}) {
	t := reflect.TypeOf(value).Elem()
	s := reflect.ValueOf(value).Elem()
	for i := 0; i < s.NumField(); i++ {
		setEnv(t.Field(i).Name, s.Field(i), t.Field(i).Tag.Get(tag))
	}
}

func setEnv(name string, value reflect.Value, tag string) {
	if !value.CanSet() {
		log.Printf("field: %s, kind: %s not supported\n", name, value.Kind())
		return
	}
	switch value.Type().Kind() {
	case reflect.Bool:
		if !strings.EqualFold(os.Getenv(tag), "") {
			if b, err := strconv.ParseBool(os.Getenv(tag)); err == nil {
				value.SetBool(b)
			}
		}
	case reflect.Int:
		if !strings.EqualFold(os.Getenv(tag), "") {
			if n, err := strconv.Atoi(os.Getenv(tag)); err == nil {
				value.SetInt(int64(n))
			}
		}
	case reflect.String:
		if !strings.EqualFold(os.Getenv(tag), "") {
			value.SetString(os.Getenv(tag))
		}
	default:
		log.Printf("field: %s, kind: %s not supported\n", name, value.Kind())
	}
}
