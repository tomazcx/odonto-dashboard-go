package repositories

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/tomazcx/odonto-dashboard-go/data/dto"
)

type ClientRepository struct {
	Db *sql.DB
}

func (c ClientRepository) Create(dto dto.CreateClientDTO) error {
	stmt, err := c.Db.Prepare("INSERT INTO clients (id, name, age, phoneNumber, budget, budgetDescription, anamnese) VALUES (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error preparing a create client statement: %v", err)
		return err
	}
	defer stmt.Close()

	clientId := uuid.New()
	_, err = stmt.Exec(clientId, dto.Name, dto.Age, dto.Telephone, dto.Budget, dto.BudgetDescription, dto.Anamnese)

	if err != nil {
		log.Printf("Error creating the client: %v", err)
		return err
	}

	stmtAddrs, err := c.Db.Prepare("INSERT INTO addresses (id, city, district, streetAndNumber, clientId) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		log.Printf("Error preparing the create address statement: %v", err)
		return err
	}

	defer stmtAddrs.Close()

	_, err = stmtAddrs.Exec(uuid.New(), dto.City, dto.District, dto.Address, clientId)

	if err != nil {
		log.Printf("Error creating the client address: %v", err)
		return err
	}

	return nil
}
