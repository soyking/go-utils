package stk

import (
	"testing"
)

func TestStructToKey(t *testing.T) {
	type Request struct {
		BeginDate   string
		EndDate     string
		AccountList []string
		Numbers     []int
	}
	request := Request{}
	request.BeginDate = "2016-03-01"
	request.EndDate = "2016-03-07"
	request.AccountList = []string{"eb12d89d-fecf-4bba-9396-94b831ce3ee3"}
	request.Numbers = []int{3, 4, 5}

	t.Log(StructToKey(&request))
}
