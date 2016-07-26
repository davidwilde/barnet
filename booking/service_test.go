package booking

import (
	"testing"
	"time"

	"github.com/davidwilde/barnet/client"
	"github.com/davidwilde/barnet/repository"
	"github.com/davidwilde/barnet/stylist"
)

func TestBookNewAppointment(t *testing.T) {

	var (
		clientRepository      = repository.NewInMemClient()
		stylistRepository     = repository.NewInMemStylist()
		appointmentRepository = repository.NewInMemAppointment()
	)

	var bookingService = NewService(appointmentRepository, clientRepository, stylistRepository)
	stylist := stylist.New("Alice", "I love cutting hair")
	client := client.New("bob@testing.com", "Bob McTest", "Bob", "Male", "555-1234")

	appointmentTime := time.Date(2016, time.November, 10, 11, 0, 0, 0, time.UTC)

	appointmentId, err := bookingService.BookNewAppointment(*client, *stylist, appointmentTime)

	if err != nil {
		t.Error("Could not save an appointment")
	}

	appointment, err := appointmentRepository.Find(appointmentId)

	if err != nil {
		t.Error("Saved appointment cannot be found in repository")
	}

	if appointmentTime != appointment.AppointmentTime {
		t.Errorf("BookNewAppointment appointment time expected %q, got %q", appointmentTime, appointment.AppointmentTime)
	}
}

func TestBookNewAppointmentTimeUnavailable(t *testing.T) {

	var (
		clientRepository      = repository.NewInMemClient()
		stylistRepository     = repository.NewInMemStylist()
		appointmentRepository = repository.NewInMemAppointment()
	)

	var bookingService = NewService(appointmentRepository, clientRepository, stylistRepository)
	stylist := stylist.New("Alice", "I love cutting hair")
	firstClient := client.New("bob@testing.com", "Bob McTest", "Bob", "Male", "555-1234")
	secondClient := client.New("charles@testing.com", "Charles Von Test", "Charlie", "Male", "555-4321")

	appointmentTime := time.Date(2016, time.November, 10, 11, 0, 0, 0, time.UTC)

	_, err := bookingService.BookNewAppointment(*firstClient, *stylist, appointmentTime)
	_, err = bookingService.BookNewAppointment(*secondClient, *stylist, appointmentTime)

	if err == nil {
		t.Error("Duplicate appointment should not be saved")
	}
}
