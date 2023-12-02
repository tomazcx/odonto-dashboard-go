package repositories

import (
	"context"
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

func (c ClientRepository) Exists(id string) (bool, error) {
	stmt, err := c.db.Prepare("SELECT id FROM clients WHERE id = ?")

	if err != nil {
		log.Printf("Error preparing the client exists statement: %v", err)
		return false, err
	}
	defer stmt.Close()

	var clientId string
	err = stmt.QueryRow(id).Scan(&clientId)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c ClientRepository) GetClient(id string) (entities.Client, error) {
	stmt, err := c.db.Prepare("SELECT c.id, c.name, c.age, c.phoneNumber, a.city, a.streetAndNumber, a.district, c.budget, c.budgetDescription, c.anamnese FROM clients c INNER JOIN addresses a ON a.clientId = c.id WHERE c.id = ?")

	if err != nil {
		log.Printf("Error preparing the client statement: %v", err)
		return entities.Client{}, err
	}
	defer stmt.Close()

	var client entities.Client
	err = stmt.QueryRow(id).Scan(&client.ID, &client.Name, &client.Age, &client.Telephone, &client.City, &client.Address, &client.District, &client.Budget, &client.BudgetDescription, &client.Anamnese)

	if err != nil {
		log.Printf("Error fetching the client data statement: %v", err)
		return entities.Client{}, err
	}

	return client, nil
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

func (c ClientRepository) Create(ctx context.Context, dto dto.CreateClientDTO) error {

	tx, err := c.db.BeginTx(ctx, nil)

	if err != nil {
		log.Printf("Error begining the create transaction: %v", err)
		return err
	}
	defer tx.Rollback()

	clientId := uuid.New()
	if _, err = tx.ExecContext(ctx, "INSERT INTO clients (id, name, age, phoneNumber, budget, budgetDescription, anamnese) VALUES (?, ?, ?, ?, ?, ?, ?)",
		clientId,
		dto.Name,
		dto.Age,
		dto.Telephone,
		dto.Budget,
		dto.BudgetDescription,
		dto.Anamnese); err != nil {
		log.Printf("Error creating the client: %v", err)
		return err
	}

	if _, err := tx.ExecContext(ctx, "INSERT INTO addresses (id, city, district, streetAndNumber, clientId) VALUES (?, ?, ?, ?, ?)",
		uuid.New(),
		dto.City,
		dto.District,
		dto.Address,
		clientId); err != nil {
		log.Printf("Error creating the client address: %v", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Printf("Error commiting the transaction: %v", err)
		return err
	}

	return nil
}

func (c ClientRepository) Update(ctx context.Context, id string, dto dto.UpdateClientDTO) error {

	tx, err := c.db.BeginTx(ctx, nil)

	if err != nil {
		log.Printf("Error begining the update transaction: %v", err)
		return err
	}
	defer tx.Rollback()

	if _, err = tx.ExecContext(ctx, "UPDATE clients SET name = ?, age = ?, phoneNumber = ?, budget = ?, budgetDescription = ?, anamnese = ? WHERE id = ?",
		dto.Name,
		dto.Age,
		dto.Telephone,
		dto.Budget,
		dto.BudgetDescription,
		dto.Anamnese,
		id); err != nil {
		log.Printf("Error updating the client: %v", err)
		return err
	}

	if _, err = tx.ExecContext(ctx, "UPDATE addresses SET city = ?, district = ?, streetAndNumber = ? WHERE clientId = ?",
		dto.City,
		dto.District,
		dto.Address,
		id); err != nil {
		log.Printf("Error updating the client address: %v", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Printf("Error commiting the transaction: %v", err)
		return err
	}

	return nil
}

func (c ClientRepository) Delete(id string) error {
	stmt, err := c.db.Prepare("DELETE FROM clients WHERE id = ?")

	if err != nil {
		log.Printf("Error preparing a create client statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		log.Printf("Error deleting client with id %v: %v", id, err)
		return err
	}

	return nil
}
