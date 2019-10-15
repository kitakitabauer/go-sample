package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// SubscriptionStatus has user's subscription status
type SubscriptionStatus struct {
	ID            string `bson:"_id"`
	Status        int    `bson:"status"`
	Type          int    `bson:"type"`
	TimeLeft      int    `bson:"timeLeft"`
	PlaybackTime  int    `bson:"playbackTime"`
	ResetAt       int64  `bson:"resetAt"`
	PlanChangedAt int64  `bson:"planChangedAt"`
	ExpiresAt     int64  `bson:"expiresAt"`
	IsOverlapped  bool   `bson:"isOverlapped,omitempty"`
	IsTrialUsed   bool   `bson:"isTrialUsed,omitempty"`
	UpdatedAt     int64  `bson:"updatedAt"`
	CreatedAt     int64  `bson:"createdAt"`
}

func main() {
	field := "isTrial"
	// check field name is exists
	names, _ := GetFieldNames("bson", SubscriptionStatus{})
	fmt.Printf("%#v", names)
	_, ok := names[field]
	if !ok {
		panic(errors.New("bad field name"))
	}
}

// GetFieldNames is extract field names in given struct
// Why the return type is map, because simplify check of contains value.
func GetFieldNames(tagName string, src interface{}) (map[string]bool, error) {
	if src == nil {
		return nil, errors.New("src blank")
	}

	srcVal := reflect.Indirect(reflect.ValueOf(src))
	if srcVal.Kind() != reflect.Struct {
		return nil, errors.New("not struct")
	}

	// can't determine number of key beforehand
	m := make(map[string]bool, 50)

	getFieldNamesInternal(&m, tagName, src, "")

	return m, nil
}

// internal function for GetFieldNames
func getFieldNamesInternal(m *map[string]bool, tagName string, src interface{}, parentFieldName string) {
	srcVal := reflect.Indirect(reflect.ValueOf(src))
	srcType := srcVal.Type()

	for i := 0; i < srcType.NumField(); i++ {
		typeField := srcType.Field(i)
		valField := srcVal.Field(i)

		var name string
		if tagName != "" {
			name = typeField.Tag.Get(tagName)
			// for bad tagname
			if name == "" {
				name = typeField.Name
			} else {
				// ex) `json:"hoge,omitempty"`
				sname := strings.Split(name, ",")
				name = sname[0]

				// ex) `json:",string"`
				if name == "" {
					name = typeField.Name
				}
			}
		} else {
			name = typeField.Name
		}

		var fieldName string
		if parentFieldName != "" {
			fieldName = parentFieldName + "." + name
		} else {
			fieldName = name
		}

		if valField.Kind() == reflect.Struct {
			// recursion
			getFieldNamesInternal(m, tagName, valField.Interface(), fieldName)
		}

		(*m)[fieldName] = true
	}

	return
}
