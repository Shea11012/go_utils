package cmd

import (
	"fmt"
	"github.com/Shea11012/go_utils/internal/json2struct"
	"log"
	"reflect"
	"testing"
)

func TestJsonParser(t *testing.T) {
	str := `{"username":"mxy","age":12}`
	if res,err := json2struct.NewParser(str); err != nil {
		t.Fatal(err)
	} else {
		log.Println(res)
	}
}

func TestBase(t *testing.T) {
	var s interface{}
	fmt.Println(reflect.TypeOf(&s).String())
}