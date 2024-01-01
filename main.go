package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")

        c.Header("Access-Control-Allow-Methods", "POST, PUT, HEAD, PATCH, OPTIONS, GET, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
    router := gin.Default()
    router.Use(CORSMiddleware())

    router.GET("/bills", getBills)
	  router.GET("/bill/:id", getBillByID)
	  router.POST("/bill", postBill)
    router.PUT("/bill", updateBill)
    router.DELETE("/bill/:id", deleteBillByID)
    router.Run("localhost:8080")
	
}


func postBill(c *gin.Context) {
    var newBill bill

    // Call BindJSON to bind the received JSON to
    // newBill.
    if err := c.BindJSON(&newBill); err != nil {
        return
    }

    // Add the new bill to the slice.
    bills = append(bills, newBill)
    c.IndentedJSON(http.StatusCreated, newBill)
}

// bill represents data about a record bill.
type bill struct {
  ID string
  Name string 
  Amount float64
  Category string 
  Importance string 
  Frequency string
  DueDate string 
}


//dummy bill data
var bills = []bill{
 {
	ID: "4e8fd4e8-563e-4206-849b-88ae64efb25d",
    Name: "Mortgage",
    Amount: 3245,
    Category: "living expense",
    Frequency: "monthly",
  },
  {
	ID: "4e8fd4e8-563e-4206-849b-88ae66efb25d",
    Name: "Goddard Tuition",
    Amount: 2100,
    Category: "daycare",
    Frequency: "monthly",
  },
  {
	ID: "4d8fd4e8-563e-4206-849b-88ae64efb25d",
    Name: "Groceries",
    Amount: 200,
    Category: "living expense",
    Frequency: "weekly",
  },
  {
	ID: "4e8fd4e8-583e-4206-849b-88ae64efb25d",
    Name: "Tahoe",
    Amount: 700,
    Category: "living expense",
    Frequency: "monthly",
  },
  {
	ID: "4e8fd4e8-563e-4206-849b-88as64efb25d",
    Name: "TECO",
    Amount: 300,
    Category: "Utilities",
    Frequency: "monthly",
  },
  {
	ID: "4e8fa4e8-563e-4206-849b-88ae64efb25d",
    Name: "Water",
    Amount: 80,
    Category: "Utilities",
    Frequency: "monthly",
  },
  {
	ID: "4e8fd4e8-563e-4201-849b-88ae64efb25d",
    Name: "Frontier",
    Amount: 70,
    Category: "Utilities",
    Frequency: "monthly",
  },
  {
	ID: "4e8fd4e8-563e-4206-849b-98ae64efb25d",
    Name: "Car Insurance",
    Amount: 540,
    Category: "Insurance",
    Frequency: "monthly",
  },
  {
	ID: "4e8fd4e8-563e-4906-849b-88ae64efb25d",
    Name: "Life Insurance",
    Amount: 90,
    Category: "insurance",
    Frequency: "monthly",
  },
}

// getAlbums responds with the list of all albums as JSON.
func getBills(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, bills)
}

// getBillByID locates the bill whose ID value matches the id
// parameter sent by the client, then returns that bill as a response.
func getBillByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range bills {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
}

// getBillByID locates the bill whose ID value matches the id and delete item from list
// parameter sent by the client, then deletes the bill and returns a response
func deleteBillByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range bills {
        if a.ID == id {
            bills = append(bills[:i], bills[i+1:]...)
            c.IndentedJSON(http.StatusOK,  gin.H{"message": "bill successfully deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
}

// getBillByID locates the bill whose ID value matches the id and delete item from list
// parameter sent by the client, then deletes the bill and returns a response
func updateBill(c *gin.Context) {
    
  var updatedBill bill

    if err := c.BindJSON(&updatedBill); err != nil {
      fmt.Println("error getting json")
      return
    }
    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for i, a := range bills {
        if a.ID == updatedBill.ID {
            bills[i].Amount = updatedBill.Amount
            bills[i].Name = updatedBill.Name
            c.IndentedJSON(http.StatusOK,  gin.H{"message": "bill updated successfully."})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bill not found"})
}