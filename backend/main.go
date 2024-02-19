package main

import (
    "fmt"
    "sort"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin" 
    "github.com/gin-contrib/cors"
    logger "github.com/sirupsen/logrus"
)

func getIndex(c *gin.Context) {
	numberStr := c.Param("number")

	number, err := strconv.Atoi(numberStr)

	logger.WithFields(logger.Fields{
            "number": numberStr,
	}).Debug("Processing number")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
		
		logger.Warning(err)

		return
	}

	index := getNumberIndex(number, 30) // tolerance = 30%

	c.JSON(http.StatusOK, gin.H{
		"index": index,
	})	
}

func main() {
	loadConfig()
	loadNumbers()

	sort.Ints(numbers[:])	

	var logLevel = getLogLevel(config.LogLevel)
	logger.SetLevel(logLevel)

	router := gin.Default()
	// Ah, CORS!
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
    		AllowMethods:     []string{"GET"},
	}))

	router.GET("/numbers/:number", getIndex)
	router.Run(fmt.Sprintf("0.0.0.0:%d", config.Port))
}
