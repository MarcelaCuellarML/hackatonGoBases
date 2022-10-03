package file

import (
	"fmt"
	"os"
	"strings"

	"github.com/MarcelaCuellarML/hackatonGoBases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	var bookings []service.Ticket
	bookings = []service.Ticket{}
	path, err := os.ReadFile(f.Path)
	books := strings.SplitAfter(string(path), "\n")
	for i := 0; i < len(books); i++ {
		book := books[i]
		booking := strings.SplitAfter(book, ",")
		var bookRegister service.Ticket
		//fmt.Println("booking: ", booking)
		bookRegister.Id = booking[0]
		bookRegister.Names = booking[1]
		bookRegister.Email = booking[2]
		bookRegister.Destination = booking[3]
		bookRegister.Date = booking[4]
		bookRegister.Price = booking[5]
		bookings = append(bookings, bookRegister)
	}
	if err != nil {
		println(err)
	}

	return bookings, nil
}

func (f *File) Write(bookingsExists []service.Ticket) error {
	var bookingsString string
	for i := 0; i < len(bookingsExists); i++ {
		bookingsString += fmt.Sprintf("%v %v %v %v %v %v", bookingsExists[i].Id, bookingsExists[i].Names, bookingsExists[i].Email, bookingsExists[i].Destination, bookingsExists[i].Date, bookingsExists[i].Price)
		//fmt.Println("Names: ", bookingsExists[i].Names)
	}
	//ticketsConcatenados := fmt.Sprintln(bookingsString, crearTicket)
	convertirTicket := []byte(bookingsString)
	os.WriteFile(f.Path, convertirTicket, 0777)
	return nil
}
