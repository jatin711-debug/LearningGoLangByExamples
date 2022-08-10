package main

import (
	"fmt"
	"time"
)

type Title string

type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMember(member *Member) {
	fmt.Println("Name->", member.name)
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[Book has not been return]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println("Member:", member.name, "Title:", title, " Audit:", audit.checkOut, "return:", returnTime)
	}
}

func printMemberAudit(lib *Library) {
	for _, member := range lib.members {
		printMember(&member)
	}
}

func printLibraryBooks(lib *Library) {
	for title, book := range lib.books {
		fmt.Println(book.total, "->", title, "->title", book)
	}
}

func main() {
	lib := Library{
		members: make(map[Name]Member),
		books:   make(map[Title]BookEntry),
	}
	lib.books["Alen"] = BookEntry{
		total:  3,
		lended: 2,
	}
	lib.books["aweb"] = BookEntry{
		total:  3,
		lended: 2,
	}
	lib.books["Wakiiu"] = BookEntry{
		total:  3,
		lended: 2,
	}
	lib.members["Alen"] = Member{"Alen Derulo", make(map[Title]LendAudit)}
	lib.members["Json"] = Member{"Jason Derulo", make(map[Title]LendAudit)}
	printMemberAudit(&lib)
	printLibraryBooks(&lib)
}
