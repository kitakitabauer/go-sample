package cast

import (
	"fmt"
	"reflect"
)

func main() {
	interfaceMap := map[string]interface{}{
		"a": 3600 * 24 * 30,
	}
	d := interfaceMap["a"]
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))
	dd := d.(int64)
	fmt.Println(reflect.TypeOf(dd))
}
