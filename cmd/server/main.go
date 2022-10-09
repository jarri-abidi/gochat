package main

import (
	"fmt"
	"time"

	"github.com/jarri-abidi/gochat"
)

func main() {
	faisal, _ := gochat.NewUser("markhaur", "Faisal Nisar")
	jarri, _ := gochat.NewUser("jarri-abidi", "Jarri Abidi")
	sarim, _ := gochat.NewUser("sa41m", "sarim")

	fContact, _ := gochat.NewContact(faisal.ID(), faisal.FullName(), faisal.UserName())
	jContact, _ := gochat.NewContact(jarri.ID(), jarri.FullName(), jarri.UserName())
	sContact, _ := gochat.NewContact(sarim.ID(), sarim.FullName(), sarim.UserName())
	jarri.AddContacts(*fContact, *sContact)
	faisal.AddContacts(*jContact)
	sarim.AddContacts(*jContact)

	group, _ := gochat.NewGroup("gochat", *fContact, *jContact, *sContact)

	sm, rms, _ := gochat.NewMessage(*faisal, []gochat.Group{*group}, []gochat.User{*jarri}, []byte("message"), time.Now())

	fmt.Printf("sm: %+v\n", sm)
	fmt.Printf("rms: %+v\n", rms)
	// TODO: fetch jarri-abidi messages using Repository
}
