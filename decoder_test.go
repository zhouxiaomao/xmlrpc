package xmlrpc

import (
	"reflect"
	"testing"
	"time"
)

var unmarshalTests = []struct {
	value interface{}
	xml   string
}{
	{100, "<value><int>100</int></value>"},
	{"Once upon a time", "<value><string>Once upon a time</string></value>"},
	{"Mike & Mick <London, UK>", "<value><string>Mike &amp; Mick &lt;London, UK&gt;</string></value>"},
	{"Once upon a time", "<value>Once upon a time</value>"},
	{true, "<value><boolean>1</boolean></value>"},
	{false, "<value><boolean>0</boolean></value>"},
	{12.134, "<value><double>12.134</double></value>"},
	{-12.134, "<value><double>-12.134</double></value>"},
	{time.Unix(1386622812, 0).UTC(), "<value><dateTime.iso8601>20131209T21:00:12</dateTime.iso8601></value>"},
	{[]int{1, 5, 7}, "<value><array><data><value><int>1</int></value><value><int>5</int></value><value><int>7</int></value></data></array></value>"},
	{struct {
		Title  string
		Amount int
	}{"War and Piece", 20}, "<value><struct><member><name>Title</name><value><string>War and Piece</string></value></member><member><name>Amount</name><value><int>20</int></value></member></struct></value>"},
}

func Test_unmarshal(t *testing.T) {
	for _, tt := range unmarshalTests {
		v := reflect.New(reflect.TypeOf(tt.value))
		if err := unmarshal([]byte(tt.xml), v.Interface()); err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		v = v.Elem()

		if v.Kind() == reflect.Slice {
			vv := reflect.ValueOf(tt.value)
			if vv.Len() != v.Len() {
				t.Fatalf("unmarshal error:\nexpected: %v\n     got: %v", tt.value, v.Interface())
			}
			for i := 0; i < v.Len(); i++ {
				if v.Index(i).Interface() != vv.Index(i).Interface() {
					t.Fatalf("unmarshal error:\nexpected: %v\n     got: %v", tt.value, v.Interface())
				}
			}
		} else {
			if v.Interface() != interface{}(tt.value) {
				t.Fatalf("unmarshal error:\nexpected: %v\n     got: %v", tt.value, v.Interface())
			}
		}
	}
}
