package service

import (
	"errors"
	"fmt"
)

type BookingsInterface interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	ReadById(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	// Show list of bookings
	ShowBookins() error
}

type Bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id, Names, Email, Destination, Date, Price string
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return Bookings{Tickets: Tickets}
}

func (b *Bookings) Create(t Ticket) ([]Ticket, error) {
	//ticketCreado := fmt.Sprintf("\n%v, %v, %v, %v, %v, %v", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
	b.Tickets = append(b.Tickets, t)
	return b.Tickets, nil
}

func (b *Bookings) ReadById(id string) error {

	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			fmt.Printf("ID %s \n Names: %s\nEmail: %s\nDestination: %s\nDate: %s\nPrice: %s\n", b.Tickets[i].Id, b.Tickets[i].Names, b.Tickets[i].Email, b.Tickets[i].Destination, b.Tickets[i].Date, b.Tickets[i].Price)
			return nil
		}
	}
	fmt.Println("No se encontro el id: ", id)
	err := errors.New("no se ha encontrado la reserva con el id solictado")
	return err
}

func (b *Bookings) Update(id string, t Ticket) ([]Ticket, error) {
	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			b.Tickets[i].Names = t.Names
			b.Tickets[i].Email = t.Email
			b.Tickets[i].Destination = t.Destination
			b.Tickets[i].Date = t.Date
			b.Tickets[i].Price = t.Price
		}
	}
	//resp := fmt.Sprintf("Se ha actualizado correctamente el registro del ID %s con los datos %v", id, Ticket{})
	return b.Tickets, nil
}

func (b *Bookings) Delete(id string) (int, error) {

	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			copy(b.Tickets[i:], b.Tickets[i+1:])
			b.Tickets[len(b.Tickets)-1] = Ticket{}
			b.Tickets = b.Tickets[:len(b.Tickets)-1]
		}
	}
	return 0, nil
}

func (b *Bookings) ShowBookins() error {
	fmt.Println("Listado de reservas")
	for i := 0; i < len(b.Tickets); i++ {
		fmt.Println(i+1, " : ", b.Tickets[i].Names)
	}
	return nil
}
