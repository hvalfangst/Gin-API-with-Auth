package main

import (
	"encoding/json"
	"github.com/go-pg/pg/v10"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10/orm"
)

// Customer represents a customer entity.
type Customer struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Configuration struct {
	Db struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Address  string `json:"address"`
		Database string `json:"database"`
	} `json:"db"`
}

func readConfig(filePath string) (Configuration, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Configuration{}, err
	}
	defer file.Close()

	var config Configuration
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Configuration{}, err
	}

	return config, nil
}

func connectDB(config Configuration) *pg.DB {
	opts := &pg.Options{
		User:     config.Db.User,
		Password: config.Db.Password,
		Addr:     config.Db.Address,
		Database: config.Db.Database,
	}
	return pg.Connect(opts)
}

func createTables(db *pg.DB) error {
	models := []interface{}{(*Customer)(nil)}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateCustomer creates a new customer in the database.
func createCustomer(db *pg.DB, name, email, password string) error {
	customer := &Customer{
		Name:     name,
		Email:    email,
		Password: password,
	}
	_, err := db.Model(customer).Insert()
	return err
}

// GetCustomers retrieves all customers from the database.
func getCustomers(db *pg.DB) ([]Customer, error) {
	var customers []Customer
	err := db.Model(&customers).Select()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func main() {
	r := gin.Default()

	config, err := readConfig("src/config/config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Connect to the database
	db := connectDB(config)
	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error on connecting to DB: %v", err)
		}
	}(db)

	// Create the necessary tables
	err = createTables(db)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	configurateGetCustomersRoute(r, db)
	configurateCreateCustomerRoute(r, db)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func configurateCreateCustomerRoute(r *gin.Engine, db *pg.DB) {
	// Route to create a customer
	r.POST("/customers", func(c *gin.Context) {
		var input struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		err := createCustomer(db, input.Name, input.Email, input.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create customer"})
			return
		}
		c.JSON(201, gin.H{"message": "Customer created successfully"})
	})
}

func configurateGetCustomersRoute(r *gin.Engine, db *pg.DB) {
	// Route to fetch customers
	r.GET("/customers", func(c *gin.Context) {
		customers, err := getCustomers(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch customers"})
			return
		}
		c.JSON(200, customers)
	})
}
