package schema_test

import (
	"CrzayORM/schema"
	"CrzayORM/dialect"
	"testing"
)

type User struct {
	Name string `crazyorm:"PRIMARY KEY"`
	Age int 
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T){
	schema := Parse(&User{},TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2{
		t.Fatal(("Failed to parse User struct"))
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("Failed to parse primary Key")
	}
}