package gu

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"testing"
	"time"
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

func TestMidnight(t *testing.T) {
	currentTime := time.Now()
	// 凌晨
	t0 := GetMidnightBeforeDays(currentTime, 0)
	t.Logf("%s", t0.String())
	// 今晚凌晨
	t1 := GetMidnightBeforeDays(currentTime, -1)
	t.Logf("%s", t1.String())
	// 昨天凌晨
	t2 := GetMidnightBeforeDays(currentTime, 1)
	t.Logf("%s", t2.String())
}

func TestGetConsequentTime(t *testing.T) {
	dates, err := GetConsequentTime("2016-03-01", "2016-03-07")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", dates)
}

func TestToArrayInterface(t *testing.T) {
	n := []int{1, 2, 3}
	ni, err := ToArrayInterface(n)
	if err != nil {
		t.Fatal(err)
	}
	for i := range ni {
		t.Log(ni[i].(int))
	}
}

func TestIntn(t *testing.T) {
	count := map[int]int{0: 0, 1: 0}
	for i := 0; i < 1000; i++ {
		n := Intn(2)
		count[n]++
	}
	t.Log(count[0])
	t.Log(count[1])
}

func TestBTime(t *testing.T) {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		t.Error("set up your MongoDB")
	} else {
		db := "__gu_test__"
		collection := "btime"
		id := "__gu_test_btime__"

		type Test struct {
			ID   string `json:"_id" bson:"_id"`
			Time *BTime `json:"time" bson:"time"`
		}

		now := NowBTime()
		test1 := &Test{ID: id, Time: now}
		_, err := sess.DB(db).C(collection).Upsert(bson.M{"_id": id}, test1)
		if err != nil && !mgo.IsDup(err) {
			t.Error(err)
		} else {
			var test2 Test
			err = sess.DB(db).C(collection).Find(bson.M{"_id": id}).One(&test2)
			if err != nil {
				t.Error(err)
			} else {
				t.Logf("now: %s", now)
				t.Logf("db:  %s", test2.Time)
				b, _ := json.Marshal(test2)
				t.Logf("parse to json: %s", string(b))
			}
		}
	}
}

func checkEmpty(i interface{}) bool {
	return CheckEmpty(reflect.ValueOf(i))
}

func assertBool(t *testing.T, target, result bool) {
	if result != target {
		t.Errorf("should be %t but %t\n", target, result)
	}
}

func TestCheckEmptyString(t *testing.T) {
	assertBool(t, true, checkEmpty(""))
	assertBool(t, false, checkEmpty("abc"))
}

func TestCheckEmptySlice(t *testing.T) {
	assertBool(t, true, checkEmpty([]int{}))
	assertBool(t, false, checkEmpty([]int{1, 2, 3}))
}

func TestCheckEmptyPtr(t *testing.T) {
	var p *int = nil
	assertBool(t, true, checkEmpty(p))
	v := 1
	p = &v
	assertBool(t, false, checkEmpty(p))
}

func TestCheckEmptyStruct(t *testing.T) {
	type Value1 struct {
		String   string
		IntSlice []int
	}

	type Value2 struct {
		String   string
		IntSlice []int
		V1       *Value1
	}

	v1 := Value1{}
	v1.String = "abc"
	assertBool(t, true, checkEmpty(v1))
	v1.IntSlice = []int{1, 2, 3}
	assertBool(t, false, checkEmpty(v1))

	v2 := Value2{}
	v2.String = "abc"
	v2.IntSlice = []int{1, 2, 3}
	v2.V1 = nil
	assertBool(t, true, checkEmpty(v2))
	v2.V1 = &Value1{String: "abc"}
	assertBool(t, true, checkEmpty(v2))
	v2.V1.IntSlice = []int{1, 2, 3}
	assertBool(t, false, checkEmpty(v2))
}
