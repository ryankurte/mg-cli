package main

import mailgun "github.com/mailgun/mailgun-go"

type List struct {
	Address     string
	Name        string
	Description string
}

func NewListArr() []List {
	return make([]List, 0)
}

func ListFromMailgun(m mailgun.List) List {
	return List{Address: m.Address, Name: m.Name, Description: m.Description}
}

func ListToMailgun(l List) mailgun.List {
	return mailgun.List{Address: l.Address, Name: l.Name, Description: l.Description}
}

func ListArrFromMailgun(m []mailgun.List) []List {
	l := make([]List, len(m))
	for i, v := range m {
		l[i] = List{Address: v.Address, Name: v.Name, Description: v.Description}
	}
	return l
}

func ListArrToMailgun(l []List) []mailgun.List {
	m := make([]mailgun.List, len(l))
	for i, v := range l {
		m[i] = mailgun.List{Address: v.Address, Name: v.Name, Description: v.Description}
	}
	return m
}

func Difference(l1, l2 []List) (ok, add, rem, update []List) {
	for _, v1 := range l1 {
		found := false
		for _, v2 := range l2 {
			if v1 == v2 {
				ok = append(ok, v1)
				found = true
			} else if v1.Address == v2.Address {
				update = append(update, v1)
				found = true
			}
		}
		if !found {
			add = append(add, v1)
		}
	}

	for _, v2 := range l2 {
		found := false
		for _, v1 := range l1 {
			if v1 == v2 || v1.Address == v2.Address {
				found = true
			}
		}
		if !found {
			rem = append(rem, v2)
		}
	}

	return ok, add, rem, update
}
