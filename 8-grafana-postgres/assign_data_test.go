package grafanapostgres_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
	DB_PORT     = 5439
)

func TestInputDataWithRandomValues(t *testing.T) {
	ctx := context.Background()
	numberIteration := 10

	// Define the connection string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	// Connect to the database
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Insert data with iteration and random values
	for i := 0; i < numberIteration; i++ {
		insertQuery := `
		INSERT INTO metrics (time, value, status_2xx, status_3xx, status_4xx, status_5xx)
		VALUES ($1, $2, $3, $4, $5, $6);
		`

		// Randomized values
		timeValue := time.Now().Add(time.Duration(-i) * time.Minute) // Time decreases by i minutes
		value := rand.Float64()*100 + 1                              // Random float between 1 and 100
		status2xx := rand.Float64() * 100                            // Random float between 0 and 50
		status3xx := rand.Float64() * 100                            // Random float between 0 and 10
		status4xx := rand.Float64() * 100                            // Random float between 0 and 20
		status5xx := rand.Float64() * 100                            // Random float between 0 and 5

		_, err := pool.Exec(ctx, insertQuery, timeValue, value, status2xx, status3xx, status4xx, status5xx)
		if err != nil {
			t.Fatalf("Failed to insert data: %v\n", err)
		}

		time.Sleep(10 * time.Second) // Delay between inserts
	}

	t.Log("Random data inserted successfully")
}
