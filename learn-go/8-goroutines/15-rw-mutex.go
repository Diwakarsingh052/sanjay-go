package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Theater represents a theater with a specific number of seats
type Theater struct {
	seats   int          // Total seats available in the theater
	invoice chan string  // Channel to store the name of the person whose seat is booked
	rw      sync.RWMutex // Read-write Mutex for synchronizing shared access to Seats
}

var wg = new(sync.WaitGroup)
var wgBook = new(sync.WaitGroup)

func (t *Theater) bookSeat(name string) {
	defer wgBook.Done() // Decrement the counter when bookSeat is done executing

	// when Write lock is acquired, no other read or writes are allowed
	t.rw.Lock()
	// Releases the write lock when func completes
	defer t.rw.Unlock()

	// If there are seats available
	if t.seats > 0 {
		// Simulate a seat booking-making process
		fmt.Println("Seat is available for", name)
		time.Sleep(4 * time.Second)
		fmt.Println("Booking confirmed", name)

		t.seats--         // Decrement available seats
		t.invoice <- name // Send a person's name to the invoice channel
	} else {
		fmt.Println("No seats available for", name) // Inform that no seats are available
	}
}

// printInvoice method prints the invoice for all booked seats
func (t *Theater) printInvoice() {
	defer wg.Done() // Decrement the counter when func is done executing

	for name := range t.invoice {
		fmt.Printf("Invoice is sent to %s\n", name)
	}
}

// checkSeats method checks the available Seats in the Theater
func (t *Theater) checkSeats() {
	defer wg.Done() // Decrement the counter when checkSeats is done executing
	// Acquire a lock for reading
	t.rw.RLock()
	// Releases the read lock when func completes

	//no one can write when read lock is acquired,
	// there could be unlimited number of reads
	defer t.rw.RUnlock()
	fmt.Println("Available Seats:", t.seats)
}

func main() {

	t := Theater{
		seats: 2, // With 2 seat
		// using unbuffered chan, if using buffered chan, don't use select for recv values from the channel
		invoice: make(chan string), // Create the invoice channel //
	}

	// Start checkSeat routines
	for i := 1; i <= 6; i++ {
		wg.Add(1) // Increment wait group counter
		go t.checkSeats()
	}

	// Start bookSeat routines
	for i := 1; i <= 3; i++ {
		wgBook.Add(1) // Increment wait group counter
		go t.bookSeat("User " + strconv.Itoa(i))
	}

	// Start checkSeat routines
	for i := 1; i <= 6; i++ {
		wg.Add(1) // Increment wait group counter
		go t.checkSeats()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		wgBook.Wait()
		close(t.invoice)
	}()

	wg.Add(1)
	go t.printInvoice()

	wg.Wait()

}
