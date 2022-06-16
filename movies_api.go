package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var records = readCsvFile("./movies.csv")

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	fmt.Println(records)

}

func setupRoutes(r *gin.Engine) {
	r.GET("/movies/year/:year", year)
	r.GET("/movies/rating/:rating", rating)
	r.GET("/movies/genre/:genre", genre)
}

//Dummy function
func year(c *gin.Context) {
	year, ok := c.Params.Get("year")
	movieName := getMovieByYear(records, year)
	if ok == false {
		res := gin.H{
			"error": "year is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"year":  year,
		"movie": movieName,
	}
	c.JSON(http.StatusOK, res)
}

//Dummy function
func rating(c *gin.Context) {
	rating, ok := c.Params.Get("rating")

	movieName := getMovieByRating(records, rating)
	if ok == false {
		res := gin.H{
			"error": "rating is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"rating": rating,
		"movie":  movieName,
	}
	c.JSON(http.StatusOK, res)
}

func genre(c *gin.Context) {
	genre, ok := c.Params.Get("genre")
	movieName := getMovieByGenure(records, genre)
	if ok == false {
		res := gin.H{
			"error": "genure is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	/*
		city := ""
	*/
	res := gin.H{
		"genre": genre,
		"movie": movieName,
	}
	c.JSON(http.StatusOK, res)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func getMovieByYear(records [][]string, year string) []string {
	var movieName string
	var movieArr = []string{}
	for i := 1; i < len(records); i++ {

		//fmt.Println(records[0][0], i)
		if records[i][7] == year {
			movieName = records[i][0]
			movieArr = append(movieArr, movieName)

		}

	}
	return movieArr
}

func getMovieByRating(records [][]string, rating string) []string {
	var movieName string
	var movieArr = []string{}
	for i := 1; i < len(records); i++ {

		//fmt.Println(records[0][0], i)
		if records[i][5] >= rating {
			movieName = records[i][0]
			movieArr = append(movieArr, movieName)

		}

	}
	return movieArr
}

func getMovieByGenure(records [][]string, genure string) []string {
	var movieName string
	var movieArr = []string{}
	for i := 1; i < len(records); i++ {

		//fmt.Println(records[0][0], i)
		if records[i][1] == genure {
			movieName = records[i][0]
			movieArr = append(movieArr, movieName)

		}

	}
	return movieArr
}
