package main

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User), new(Profile))
}

// surruption
// blood incantation
// genocide pact
// mammoth grinder
// newest decrepit
// odious construct
// symbolic sacramento
// to violenty vomit
// severed savior bus fire