package blackbox

import (
	"fmt"

	"github.com/cucumber/godog"
)

func GetSliceOfMapFromTable(data *godog.Table) []map[string]interface{} {
	var som = make([]map[string]interface{}, 0)
	var mKeys = make(map[int]string)
	for k, v := range data.Rows[0].Cells {
		mKeys[k] = v.Value
	}
	for rk, rv := range data.Rows {
		if rk == 0 {
			continue
		}
		m := make(map[string]interface{})
		for k, v := range rv.Cells {
			m[mKeys[k]] = v.Value
		}
		som = append(som, m)
	}
	return som
}

func ToString(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func ErrGotWant(got, want interface{}) error {
	return fmt.Errorf("got: %v, want: %v", got, want)
}
