package main

import(
    "fmt"
    "bytes"
    "context"
	"testing"
    "net/http"
    "encoding/json"
    "net/http/httptest"
    "go-service/model"
    "go-service/lib"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/jackc/pgx/v4"

    "github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {

    db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
    if err != nil {
        t.Fatalf("err not expected while open a mock db, %v", err)
    }
    defer db.Close()

    order := model.Order{
        CustomerID : 2,
        Item: "goods",
        Amount: 544.00,
    }

    // Encode the order to JSON
    orderJSON, err := json.Marshal(order)
    if err != nil {
        t.Fatalf("Failed to marshal order: %v\n", err)
    }

    t.Run("NewOrder", func(t *testing.T) {
        mock.ExpectQuery("SELECT name FROM users WHERE name = $1").WithArgs("John Doe").WillReturnError(pgx.ErrNoRows)

        w := httptest.NewRecorder()
        r := httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBuffer(orderJSON))

        ctx := context.WithValue(r.Context(), "DB", db)
        r = r.WithContext(ctx)

        // Perform the request
        lib.CreateOrder(w, r)
        fmt.Println(w.Body)

        assert.Equal(t, http.StatusOK, w.Code)
        if err := mock.ExpectationsWereMet(); err != nil {
            t.Errorf("not all expectations were met: %v", err)
        }
    })
    /*
    order := model.Order{
        CustomerID : 2,
        Item: "goods",
        Amount: 544.00,
    }

    // Encode the order to JSON
    orderJSON, err := json.Marshal(order)
    if err != nil {
        t.Fatalf("Failed to marshal order: %v\n", err)
    }

    req, err := http.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBuffer(orderJSON))
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    recorder := httptest.NewRecorder()

    // Perform the request
    lib.CreateOrder(recorder, req)
    fmt.Println(recorder.Body)

    // Check to see if the response was what you expected
    if recorder.Code == http.StatusOK {
        t.Logf("Expected to get status %d, but instead got %d\n", http.StatusOK, recorder.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, recorder.Code)
    }
        */
}