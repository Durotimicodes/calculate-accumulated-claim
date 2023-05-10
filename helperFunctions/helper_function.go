package helperfunctions

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/Durotimicodes/wtw-intern-developer/models"
)

// function GetMinYearAndDevelepmentYear gets the minimum Origin year and total number of development years
func GetMinYearAndDevelopementYears(inputData [][]string) (int, int) {

	//store values without the headers into a 2d matrix
	store := make([][]string, 0)

	//delete the headers
	listWithOutTheDescription := inputData[1:]

	//iterate over the 2d matrix
	for _, eachRecord := range listWithOutTheDescription {

		//eliminate the the product name
		listWithOutTheProduct := eachRecord[1:]

		//eliminate the incremental value
		listWithOutTheIncrementalValue := listWithOutTheProduct[:len(listWithOutTheProduct)-1]

		//store values in 2d matrix
		store = append(store, listWithOutTheIncrementalValue)
	}

	//
	years := make([]int, 0)
	var maxYear int

	//iterate over the results from 2d matrix
	for _, val := range store {
		//iterate over slice
		for _, elem := range val {
			//trim extra space
			s := strings.TrimSpace(elem)

			//convert years from string to integers
			convToInt, err := strconv.Atoi(s)

			//handle and check for error
			if err != nil {
				log.Fatal("Failed converting years to integers:", err)
			}

			//store years into a slice
			years = append(years, convToInt)

			//get the maximum year
			if maxYear < convToInt {
				maxYear = convToInt
			}
		}
	}

	//copy years into an empty slice
	copy := append([]int(nil), years...)

	//sort the years in increasing order
	sort.Ints(copy)

	//pick the first element, which is the minimum year
	minYear := copy[0]

	//subtract the max year from min  to get the development year
	DevelopmentYears := (maxYear - minYear) + 1

	return minYear, DevelopmentYears

}

// function  Convert2dDataToMa takes in a 2d matrix and convertdata to a type map
func Convert2dDataToMap(dataWithoutHeader [][]string) map[string]models.Triangle {

	//store uniqueProducts into map
	uniqueProducts := map[string]models.Triangle{}
	//iterate over the data without headers
	for _, eachRecord := range dataWithoutHeader {

		//convert origin year from string to integer and trim whitespaces
		originYear, err := strconv.Atoi(strings.TrimSpace(eachRecord[1]))
		if err != nil {
			log.Fatal("Error:", err)
		}

		//convert development year from string to integer and trim whitespaces
		developmentYear, err := strconv.Atoi(strings.TrimSpace(eachRecord[2]))
		if err != nil {
			log.Fatal("Error:", err)
		}

		//convert incremental value to float and trim whitespace
		incValue, err := strconv.ParseFloat(strings.TrimSpace(eachRecord[3]), 64)
		if err != nil {
			log.Fatal("Error:", err)
		}

		//check if unique product exit if it does return the first unique product and its corresponding triangle else return the triangle
		if _, exist := uniqueProducts[eachRecord[0]]; !exist {

			//intializing/instantiating the map
			uniqueProducts[eachRecord[0]] = models.Triangle{
				fmt.Sprintf("%v-%v", strings.TrimSpace(eachRecord[1]), strings.TrimSpace(eachRecord[2])): models.ProductTriangle{
					Product:         strings.TrimSpace(eachRecord[0]),
					Origin:          originYear,
					DevelopmentYear: developmentYear,
					Value:           incValue,
				},
			}

		} else {
			uniqueProducts[eachRecord[0]][fmt.Sprintf("%v-%v", strings.TrimSpace(eachRecord[1]), strings.TrimSpace(eachRecord[2]))] = models.ProductTriangle{
				Product:         strings.TrimSpace(eachRecord[0]),
				Origin:          originYear,
				DevelopmentYear: developmentYear,
				Value:           incValue,
			}

		}

	}

	return uniqueProducts
}

// func ConvertDataTo Matrix convert a finalValue type to a 2d matrix
func ConvertDataToMatix(results map[string][]models.FinalValue) [][]string {
	finalResult := make([][]string, 0)

	for k, v := range results {
		if len(k) != 0 {
			finalResult = append(finalResult, []string{k})
			for _, f := range v {
				finalResult = append(finalResult, []string{strconv.FormatFloat(f.AccumulatedClaim, 'f', 2, 64)})
			}
		}

	}

	return finalResult
}

// FinalOutPut converts the result, minimum Origin year and development year to a 2d matrix
func FinalOutPut(results map[string][]models.FinalValue, minOriginYear int, totalDevYear int) [][]string {

	convtOrgYearToStr := strconv.Itoa(minOriginYear)
	convtTotalDevYearToStr := strconv.Itoa(totalDevYear)

	val := []string{convtOrgYearToStr, convtTotalDevYearToStr}

	finalResult := [][]string{val}

	for k, v := range results {

		finalResult = append(finalResult, []string{k})

		for _, f := range v {
			finalResult = append(finalResult, []string{strconv.FormatFloat(f.AccumulatedClaim, 'f', 2, 64)})
		}
	}

	log.Println(finalResult)
	return finalResult
}

// used for unit testing mapped valued
func CheckDataAsStrings(got, expected [][]string) bool {
	return fmt.Sprintf("%v", got) != fmt.Sprintf("%v", expected)
}
