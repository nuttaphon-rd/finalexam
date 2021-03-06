package customer

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/nuttaphon-rd/finalexam/errors"
)

type CustomerServicer interface {
	createTableCustomerDB() error
	CreateCustomerDB(*Customer) *errors.Error
	GetCustomerByIdDB(string) (*Customer, *errors.Error)
	GetAllCustomerDB() ([]*Customer, *errors.Error)
	UpdateCustomerDB(Customer, string) (*Customer, *errors.Error)
	DeleteCustomerDB(string) *errors.Error
}

type CustomerService struct {
	db *sql.DB
}

func NewService(db *sql.DB) CustomerServicer {
	customerService := &CustomerService{db}

	if err := customerService.createTableCustomerDB(); err != nil {
		log.Fatalf("can't create table customers: %s", err)
	}

	return customerService
}

func (cs *CustomerService) createTableCustomerDB() error {
	createTb := `CREATE TABLE IF NOT EXISTS customers(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
		);`
	_, err := cs.db.Exec(createTb)
	if err != nil {
		return err
	}
	fmt.Println("create table success")
	return nil
}

func (cs *CustomerService) CreateCustomerDB(c *Customer) *errors.Error {
	row := cs.db.QueryRow("INSERT INTO customers (name, email,status) values ($1, $2, $3)  RETURNING id", c.Name, c.Email, c.Status)

	err := row.Scan(&c.ID)
	if err != nil {
		return &errors.Error{
			http.StatusInternalServerError,
			"Error can't create customer " + err.Error(),
		}
	}
	return nil
}

func (cs *CustomerService) GetCustomerByIdDB(id string) (*Customer, *errors.Error) {
	stmt, err := cs.db.Prepare(`SELECT id, name, email, status FROM customers WHERE id= $1`)
	if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error can't prepare in get customer by id " + err.Error(),
		}
	}

	row := stmt.QueryRow(id)
	cus := &Customer{}

	if err := row.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status); err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error can't scan in get customer by id " + err.Error(),
		}
	}

	return cus, nil
}

func (cs *CustomerService) GetAllCustomerDB() ([]*Customer, *errors.Error) {
	stmt, err := cs.db.Prepare(`SELECT id, name, email, status FROM customers`)
	if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error can't prepare in get all customer " + err.Error(),
		}
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error can't get query in all customer " + err.Error(),
		}
	}

	customers := []*Customer{}
	for rows.Next() {
		c := &Customer{}
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Status)
		if err != nil {
			return nil, &errors.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error can't create struct in all customer " + err.Error(),
			}
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (cs *CustomerService) UpdateCustomerDB(c Customer, id string) (*Customer, *errors.Error) {
	if _, err := cs.GetCustomerByIdDB(id); err != nil {
		return nil, err
	}

	stmt, err := cs.db.Prepare(`UPDATE customers SET name=$2,email=$3,status=$4 WHERE id=$1`)
	if err != nil {
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error can't prepare statement in update customer " + err.Error(),
		}
	}

	if res, err := stmt.Exec(id, c.Name, c.Email, c.Status); err != nil {
		fmt.Printf("result update : %s", res)
		return nil, &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error can't exec statement in update customer " + err.Error(),
		}
	}

	customer := &Customer{}
	var errE *errors.Error
	if customer, errE = cs.GetCustomerByIdDB(id); errE != nil {
		return nil, errE
	}
	return customer, nil
}

func (cs *CustomerService) DeleteCustomerDB(id string) *errors.Error {
	stmt, err := cs.db.Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error can't prepare statement in delete customer " + err.Error(),
		}
	}

	if _, err := stmt.Exec(id); err != nil {
		return &errors.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error can't exec statement in delete customer " + err.Error(),
		}
	}
	return nil
}
