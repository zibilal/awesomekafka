package helpers

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// StructIterator is used as a helper for converting between two different struct type
// input must be a struct type
// output must be a pointer of a struct typed value
// field of the input argument that need to be inserted to output argument need to have
// both same type and same field name or tag name
func StructIterator(input interface{}, output interface{}, acceptedTags ...string) {

	if reflect.TypeOf(output).Kind() != reflect.Ptr {
		log.Fatal("only accepts addressable struct type for output")
		return
	}

	ivalue := reflect.Indirect(reflect.ValueOf(input))
	itype := ivalue.Type()

	ovalue := reflect.Indirect(reflect.ValueOf(output))

	for i := 0; i < ivalue.NumField(); i++ {
		ifvalue := ivalue.Field(i)
		iftype := itype.Field(i)

		if len(acceptedTags) > 0 {
			for t := 0; t < len(acceptedTags); t++ {
				tagName := strings.Split(iftype.Tag.Get(acceptedTags[t]), ",")
				if fOutput, err := getFieldByTag(output, acceptedTags[t], tagName[0]); err != nil {
					log.Fatal("[StructIterator]", " ", err.Error())
				} else {
					if fOutput.Type() == ifvalue.Type() {
						fOutput.Set(ifvalue)
					}
				}
			}
		} else {
			fOutput := ovalue.FieldByName(iftype.Name)
			if fOutput.IsValid() {
				if fOutput.Type() == ifvalue.Type() {
					fOutput.Set(ifvalue)
				}
			} else {
				log.Fatal("[StructIterator]unable to find filed by name", " ", iftype.Name)
			}
		}
	}
}

// IsEmpty is an helper function to decide whether a value is empty or not
// This function is mean to be used to decide whether a struct variable is empty or not
func IsEmpty(t interface{}) bool {
	return reflect.DeepEqual(t, reflect.Zero(reflect.TypeOf(t)).Interface())
}

func getFieldByTag(data interface{}, tagName, fieldTag string) (reflect.Value, error) {
	dvalue := reflect.Indirect(reflect.ValueOf(data))
	dtype := dvalue.Type()

	for i := 0; i < dvalue.NumField(); i++ {
		fvalue := dvalue.Field(i)
		ftype := dtype.Field(i)

		dtag := strings.Split(ftype.Tag.Get(tagName), ",")

		if dtag[0] == fieldTag {
			return fvalue, nil
		}
	}

	return reflect.Value{}, fmt.Errorf("unable to find field of tag %s, field tag %s", tagName, fieldTag)
}
