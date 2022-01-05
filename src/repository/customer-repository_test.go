package repository

import (
	"github.com/go-playground/assert/v2"
	"github.com/nuno/nunes-jumia/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os/exec"
	"testing"
)

func TestCustomerRepository(t *testing.T) {
	scenarios := []testScenario{
		*MakeScenarioReturnsTwoCustomers(),
		*MakeScenarioReturnsOneCustomer(),
		*MakeScenarioReturnsNoneCustomer(),
	}

	dbName := "sample_test.db"
	defer removeTestDB(dbName)

	for _, scenario := range scenarios {
		t.Run(scenario.TestName, func(t *testing.T) {
			session := beforeEach(dbName)
			repository := NewCustomerRepository(session)

			if len(scenario.Rows) > 0 {
				for _, row := range scenario.Rows {
					session.Create(&row)
				}
			}

			result, err := repository.GetCustomers()

			assert.Equal(t, len(result), scenario.ExpectLengthResult)
			assert.Equal(t, err, scenario.ExpectError)
		})
	}
}

func beforeEach(dbName string) *gorm.DB {
	removeTestDB(dbName)

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("error to connect testDB. Err: %v", err)
	}

	_ = db.AutoMigrate(&model.Customer{})

	return db
}

func removeTestDB(dbName string) {
	cmd := exec.Command("rm", "-f", dbName)
	if err := cmd.Run(); err != nil {
		log.Fatalf("error to remove previous DB. Err: %v", err)
	}
}
