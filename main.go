package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")

        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

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
    router.Run("localhost:8080")
	
}



// bill represents data about a record bill.
type bill struct {
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
    Name: "Mortgage",
    Amount: 3245,
    Category: "living expense",
    Frequency: "monthly",
  },
  {
    Name: "Goddard Tuition",
    Amount: 2100,
    Category: "daycare",
    Frequency: "monthly",
  },
  {
    Name: "Groceries",
    Amount: 200,
    Category: "living expense",
    Frequency: "weekly",
  },
  {
    Name: "Tahoe",
    Amount: 700,
    Category: "living expense",
    Frequency: "monthly",
  },
  {
    Name: "TECO",
    Amount: 300,
    Category: "Utilities",
    Frequency: "monthly",
  },
  {
    Name: "Water",
    Amount: 80,
    Category: "Utilities",
    Frequency: "monthly",
  },
  {
    Name: "Frontier",
    Amount: 70,
    Category: "Utilities",
    Frequency: "monthly",
  },
  {
    Name: "Car Insurance",
    Amount: 540,
    Category: "Insurance",
    Frequency: "monthly",
  },
  {
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