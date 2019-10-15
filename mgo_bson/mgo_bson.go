package mgo_bson

import (
	. "fmt"
	"gopkg.in/mgo.v2/bson"
)

type MyType struct {
	ID string
	plan int
	platform int
}

func main() {
	selector := bson.M{
		"userId": bson.M{
			"$in": []string{
				"hoge",
				"fuga",
			},
		},
	}

	selector["platform"] = 1

	checkStruct()
}

func checkStruct() {
	struc1 := MyType{
		ID: "ID1",
	}

	Println("struc1.ID: ", struc1.ID)
	Println("struc1.platform: ", struc1.platform)

	struc2 := MyType{
		ID: "ID2",
		plan: 0,
		platform: 1,
	}

	Println("struc2.ID: ", struc2.ID)
	Println("struc2.plan: ", struc2.plan)
	Println("struc2.platform: ", struc2.platform)
}


