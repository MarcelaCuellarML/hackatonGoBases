package main

import (
	"fmt"
	"strconv"

	"github.com/MarcelaCuellarML/hackatonGoBases/internal/file"
	"github.com/MarcelaCuellarML/hackatonGoBases/internal/service"
)

func main() {
	b := &service.Bookings{}

	file := &file.File{Path: "tickets.csv"}
	b.Tickets, _ = file.Read()

	t := service.Ticket{
		Names:       "Maria,",
		Email:       "Maria@gmail.com,",
		Destination: "Cartagena,",
		Date:        "02 10 2022,",
		Price:       "100",
	}
	newId := len(b.Tickets) + 1
	s := strconv.Itoa(newId)
	t.Id = "\n" + s + ","
	fmt.Println("Id generdo: ", t.Id)

	//b.Create(t)
	//b.ShowBookins()
	//file.Write(b.Tickets)
	//ConsultaID := b.ReadById("7,")
	//if ConsultaID != nil {}

	//b.Delete("4,")
	file.Write(b.Tickets)
	b.ShowBookins()
	b.Tickets, _ = b.Update("4,", t)
	b.ShowBookins()

}
