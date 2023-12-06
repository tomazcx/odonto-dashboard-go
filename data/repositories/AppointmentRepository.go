package repositories

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/tomazcx/odonto-dashboard-go/data/dto"
	"github.com/tomazcx/odonto-dashboard-go/data/entities"
	"github.com/tomazcx/odonto-dashboard-go/infra/db"
)

type AppointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository() AppointmentRepository {
	dbInstance := db.GetConn()
	return AppointmentRepository{
		db: dbInstance,
	}
}

func (a AppointmentRepository) Exists(id string) bool {
	stmt, err := a.db.Prepare("SELECT id FROM appointments WHERE id = ?")

	if err != nil {
		log.Printf("Error preparing the appointments exists statement: %v", err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		log.Printf("Error fetching the appointment: %v", err)
		return false
	}

	return true
}

func (a AppointmentRepository) FindByClientId(id string, skip int) (int, []entities.Appointment, error) {
	stmt, err := a.db.Prepare("SELECT COUNT(*) OVER() AS totalEntities, id, date, teeth, proccedure FROM appointments WHERE clientId = ? LIMIT 15 OFFSET ?")

	if err != nil {
		log.Printf("Error preparing the get appointments statement: %v", err)
		return 0, []entities.Appointment{}, err
	}
	defer stmt.Close()

	var appointments []entities.Appointment
	var totalEntities int

	rows, err := stmt.Query(id, skip)

	for rows.Next() {
		var appointment entities.Appointment
		if err := rows.Scan(&totalEntities, &appointment.ID, &appointment.Date, &appointment.Teeth, &appointment.Proccedure); err != nil {
			log.Printf("Error scanning the appointment: %v", err)
			return 0, []entities.Appointment{}, err
		}

		appointments = append(appointments, appointment)
	}

	return totalEntities, appointments, nil
}

func (a AppointmentRepository) Create(clientId string, dto dto.CreateAppointmentDTO) error {
	stmt, err := a.db.Prepare("INSERT INTO appointments (id, date, teeth, proccedure, clientId) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error preparing the create appointment statement: %v", err)
		return err
	}
	defer stmt.Close()

	appointmentId := uuid.New()
	_, err = stmt.Exec(appointmentId, dto.Date, dto.Teeth, dto.Proccedure, clientId)

	if err != nil {
		log.Printf("Error creating the appointment: %v", err)
		return err
	}

	return nil
}

func (a AppointmentRepository) Delete(id string) error {
	stmt, err := a.db.Prepare("DELETE FROM appointments WHERE id = ?")

	if err != nil {
		log.Printf("Error preparing the delete appointment statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		log.Printf("Error deleting the appointment: %v", err)
		return err
	}

	return nil
}
