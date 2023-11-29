package repositories

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/tomazcx/odonto-dashboard-go/data/dto"
	"github.com/tomazcx/odonto-dashboard-go/data/entities"
	"github.com/tomazcx/odonto-dashboard-go/infra/db"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository() ClientRepository {
	dbInstance := db.GetConn()
	return ClientRepository{db: dbInstance}
}

func (c ClientRepository) GetCount() (int, error) {
	stmt, err := c.db.Prepare("SELECT COUNT(*) FROM clients")

	if err != nil {
		log.Printf("Error preparing the clients count statement: %v", err)
		return 0, err
	}
	defer stmt.Close()

	var clientsCount int
	err = stmt.QueryRow().Scan(&clientsCount)

	if err != nil {
		log.Printf("Error fetching the clients count statement: %v", err)
		return 0, err
	}

	return clientsCount, nil
}

func (c ClientRepository) GetMany(skip int, orderByField string, ascendingSort bool, name string) ([]entities.Client, int, error) {
	queryStr := "SELECT COUNT(*) OVER () AS totalEntities, c.id, c.name, c.age, c.phoneNumber, a.city FROM clients c INNER JOIN addresses a ON a.clientId = c.id WHERE c.name LIKE ?"

	if len(orderByField) > 0 {
		queryStr += " ORDER BY c." + orderByField

		if !ascendingSort {
			queryStr += " DESC"
		}

	}

	queryStr += " LIMIT 15 OFFSET ?"

	stmt, err := c.db.Prepare(queryStr)

	if err != nil {
		log.Printf("Error preparing the get many clients statement: %v", err)
		return []entities.Client{}, 0, err
	}
	defer stmt.Close()

	var clients []entities.Client
	var totalEntities int

	nameFilterStr := "%" + name + "%"
	rows, err := stmt.Query(nameFilterStr, skip)

	if err != nil {
		log.Printf("Error fetching the clients: %v", err)
		return []entities.Client{}, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var client entities.Client
		if err = rows.Scan(&totalEntities, &client.ID, &client.Name, &client.Age, &client.Telephone, &client.City); err != nil {
			log.Printf("Error scanning the client: %v", err)
			return []entities.Client{}, 0, err
		}
		clients = append(clients, client)
	}

	return clients, totalEntities, nil
}

func (c ClientRepository) Create(dto dto.CreateClientDTO) error {
	stmt, err := c.db.Prepare("INSERT INTO clients (id, name, age, phoneNumber, budget, budgetDescription, anamnese) VALUES (?, ?, ?, ?, ?, ?, ?)")

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

	stmtAddrs, err := c.db.Prepare("INSERT INTO addresses (id, city, district, streetAndNumber, clientId) VALUES (?, ?, ?, ?, ?)")

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
