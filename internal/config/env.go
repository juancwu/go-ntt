package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/juancwu/go-ntt/util"
)

type env struct {
	GONTT_LOG_LEVEL string
}

var (
	optional_var []string = []string{"GONTT_LOG_LEVEL"}
	required_var []string = []string{}
	Env          env      = env{}
)

func InitEnv() error {
	for _, k := range required_var {
		v, exists := os.LookupEnv(k)
		if !exists {
			return fmt.Errorf("required environment variable '%s' not set!", k)
		}

		err := setField(&Env, k, v)
		if err != nil {
			return err
		}
	}

	for _, k := range optional_var {
		v, exists := os.LookupEnv(k)
		if exists {
			err := setField(&Env, k, v)
			if err != nil {
				return err
			}
		} else {
			util.Log().Debug("optional environment variable not found.", "KEY", k)
		}
	}

	return nil
}

func setField(v interface{}, key, value string) error {
	structValue := reflect.ValueOf(v).Elem()
	fieldVal := structValue.FieldByName(key)

	if !fieldVal.IsValid() {
		return fmt.Errorf("no such field: %s in env config struct", key)
	}

	if !fieldVal.CanSet() {
		return fmt.Errorf("cannot set field %s", key)
	}

	fieldVal.SetString(value)
	util.Log().Debug("environment variable set.", "KEY", key, "VALUE", value)
	return nil
}
