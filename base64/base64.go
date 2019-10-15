package base64

import (
	"os"
	"encoding/base64"
	"fmt"
	"os/user"
	"io/ioutil"
	"unsafe"
)

func encode(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(b)

	return str, nil
}

func decode(str string) error {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	fmt.Println("decode: ", *(*string)(unsafe.Pointer(&data)))
	return nil
}

func homeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func main() {
	str, err := encode(homeDir() + "/test.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("encode: ", str)

	err = decode(str)
	if err != nil {
		panic(err)
	}
}
