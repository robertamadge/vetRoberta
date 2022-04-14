package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func signin(w http.ResponseWriter, r *http.Request) {
	var fileName = "signin.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error parsing file")
		return
	}
	type VetClinicName struct {
		Name string
	}
	err = t.ExecuteTemplate(w, fileName, VetClinicName{"Bobby's vet clinic"})
	if err != nil {
		fmt.Println("Error when executing template")
		return
	}
}

var userDB = map[string]string{
	"Roberta": "123456",
}

func signinSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		http.Redirect(w,r, "/booking", http.StatusFound)
		fmt.Fprint(w, "Ok")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error")
	}
	fmt.Println(username, password)
}

//estrutura de dados
type Queue struct {
	Tickets []int
}

//Add tiket at the end
func (q *Queue) Enqueue(ticket int) {
	q.Tickets = append(q.Tickets, ticket)
}

//Returns the first ticket
func (q *Queue) Dequeue() (int, error) {
	if len(q.Tickets) == 0 {
		return 0, errors.New("Empty queue")
	}
	var firstTicket int
	firstTicket, q.Tickets = q.Tickets[0], q.Tickets[1:]
	return firstTicket, nil
}

func bookingWithTicket(w http.ResponseWriter, r *http.Request) {
	queue := Queue{}
	queue.Enqueue(1)
	//ticket, _ := queue.Dequeue()
	//fmt.Println(ticket)
	//queue.Enqueue(2)
	//ticket, _ = queue.Dequeue()
	//fmt.Println(ticket)
	//queue.Enqueue(3)
	//ticket, _ = queue.Dequeue()
	//fmt.Println(ticket)

	var fileName = "booking.html"
	t, err := template.ParseFiles(fileName)
	err = t.ExecuteTemplate(w, fileName, queue)

	if err != nil {
		fmt.Println("Error when executing template booking")
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello")
	case "/signin":
		signin(w, r)
	case "/signin-submit":
		signinSubmit(w, r)
	case "/booking":
		bookingWithTicket(w, r)
	}
	fmt.Printf("Handling function %s request\n", r.Method)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}

//function with html as an example
/*func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello</h1>")
}*/
