package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "test"
	databasePass := "password"
	databaseName := "crm"

	database, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}

	_, err = database.Exec("SELECT 1 + 1")
	if err != nil {
		panic(err.Error())
	}

	return database
}

func GetCustomers() []Customer {
	database := GetConnection()

	var error error
	var rows *sql.Rows

	rows, error = database.Query("SELECT * FROM customers ORDER BY customerId DESC")
	if error != nil {
		panic(error.Error())
	}

	customer := Customer{}
	customers := []Customer{}

	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}

		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn

		customers = append(customers, customer)
	}

	defer database.Close()
	return customers
}

func InsertCustomer(customer Customer) {
	database := GetConnection()

	insert, err := database.Prepare("INSERT INTO customers(customerName, ssn) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = insert.Exec(customer.CustomerName, customer.SSN)
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()
}

func UpdateCustomer(customer Customer) {
	database := GetConnection()

	update, err := database.Prepare("UPDATE customers SET customerName=?, SSN=? WHERE customerId=?")
	if err != nil {
		panic(err.Error())
	}

	_, err = update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	if err != nil {
		panic(err.Error())
	}

	defer database.Close()
}

func deleteCustomer(customer Customer) {
	database := GetConnection()
	delete, err := database.Prepare("DELETE FROM customers WHERE customerId=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = delete.Exec(customer.CustomerId)
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()
}

func Run() {
	GetConnection()
	customers := GetCustomers()
	// fmt.Println("Before insert: ", customers)
	// customer := Customer{
	// 	CustomerName: "Arnie Smith",
	// 	SSN:          "2386343",
	// }
	// InsertCustomer(customer)
	// customers = GetCustomers()
	// fmt.Println("After Insert: ", customers)

	// fmt.Println("Before Update: ", customers)
	// customer := Customer{
	// 	CustomerName: "George Thompson",
	// 	SSN:          "23233432",
	// 	CustomerId:   2,
	// }
	// UpdateCustomer(customer)
	// customers = GetCustomers()
	// fmt.Println("After Update: ", customers)

	fmt.Println("Before Delete: ", customers)
	deleteCustomer(Customer{CustomerId: 1})
	customers = GetCustomers()
	fmt.Println("After Delete: ", customers)
}
