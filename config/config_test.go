package config

import (
	"os"
	"testing"
)

var test = `

#

test=123
     ddd=    1
     dddd=333:2
     test1=
     =       


`

func Test_NewConfig(t *testing.T) {
	f, err := os.Create("test.conf")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(test)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove("test.conf")

	c, err := NewConfig("test.conf")
	if err != nil {
		t.Fatal(err)
	}
	if c.String("test") != "123" {
		t.Fatal("expect 123")
	}
	if c.String("ddd") != "1" {
		t.Fatal("expect 1")
	}
	if c.String("dddd") != "333:2" {
		t.Fatal("expect 333:2")
	}
	if c.String("test1") != "" {
		t.Fatal("expect nil string")
	}
}
