package json

import (
	"fmt"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(data interface{}) ([]byte, error) {
	result, err := json.Marshal(&data)
	if err != nil{
		fmt.Println(err)
	}
	return result, err
}

func Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil{
		fmt.Println(err)
	}
	return err
}