package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "unknown"
	}
	const layout = "2006/01"
	return ts.Format(layout)
}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int
	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}

// customize the json encoder by modifying the interface method
func (ts timestamp) MarshalJSON() (data []byte, _ error) {
	//  ts -> interger -> ts.Unix() -> interger
	//  data <- interger -> strconv.AppendInt(data, integer, 10)

	return strconv.AppendInt(data, ts.Unix(), 10), nil
}

// customize the json Unmarshaler by modifying the interface method
// type Unmarshaler interface {
// 	UnmarshalJSON([]byte) error
// }

func (ts *timestamp) UnmarshalJSON(data []byte) error {
	*ts = toTimestamp(string(data))
	return nil
}
