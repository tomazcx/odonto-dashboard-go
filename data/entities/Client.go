package entities

type Client struct {
	ID                string
	Name              string
	Age               int
	Telephone         string
	City              string
	Address           string
	District          string
	Budget            string
	BudgetDescription string
	Anamnese          string
	Appointments      []Appointment
}
