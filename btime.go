package gu

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

// time.Time 可以在 MongoDB 中存储成 Date 格式便于查询
// 但在转换成 json 的时候可处理性不好, 这里重新包裹一层
// 转换 json 格式是时间戳数字
type BTime struct {
	time.Time
}

// 解析成 json
func (t BTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil
}

// 解析成 bson
func (t BTime) GetBSON() (interface{}, error) {
	return t.Time, nil
}

// 从 bson 解析出来
func (t *BTime) SetBSON(raw bson.Raw) error {
	var doc time.Time
	err := raw.Unmarshal(&doc)
	*t = BTime{Time: doc}
	return err
}

func NowBTime() *BTime {
	return &BTime{Time: time.Now()}
}
