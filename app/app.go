package app

import (
	"fmt"

	"github.com/Durotimicodes/wtw-intern-developer/helperfunctions"
	"github.com/Durotimicodes/wtw-intern-developer/models"
	"github.com/Durotimicodes/wtw-intern-developer/persistingdata"
)

// function ClaimApp receives the txt file and calculates the accumulated claim of each product and corresponding minimum origin and development year
func ClaimApp(filename string) [][]string {

	//read txt file
	data, _ := persistingdata.ReadTxTFile(filename)

	//eliminate the headers/titles
	dataWithoutHeader := data[1:]

	//convert2dDataToMap converts the 2d matrix to a map of key and value
	Convert2dDataToMap := helperfunctions.Convert2dDataToMap(dataWithoutHeader)

	//getsMinYearAndDevelopmentYears gets the minimum origin year and development year
	minOriginYear, totalDevYear := helperfunctions.GetMinYearAndDevelopementYears(data)

	//data-structure to store final values
	results := make(map[string][]models.FinalValue)

	//iterate over the map of unique product and its corresponding data
	for k, v := range Convert2dDataToMap {

		for originYear := minOriginYear; originYear < minOriginYear+totalDevYear; originYear++ {

			for developmentYear := originYear; developmentYear < minOriginYear+totalDevYear; developmentYear++ {

				//create the range of origin years and development years
				rangeOfYearsAsKey := fmt.Sprintf("%v-%v", originYear, developmentYear)

				//check map for range of years and if found, set its value of map to current origin year
				productTriangle := v[rangeOfYearsAsKey]

				//set a final value for every unique product
				finalValue := results[k]

				//check if value exist //TO SOLVE THE FIRST PART
				if _, ok := v[rangeOfYearsAsKey]; !ok { //contingency check

					/*
						if the rangeOfYearsAsKey does not exist, however if the final value object has a value and its last origin year is equal
						to the current origin year append the last object into the slice of product
					*/
					if len(finalValue) != 0 && finalValue[len(finalValue)-1].Origin == originYear {
						results[k] = append(results[k], finalValue[len(finalValue)-1])
					} else {
						/*
							incase if the final value exist but the last orgin year is different from the current origin year then
							the accumlated claim is zero
						*/
						results[k] = append(results[k], models.FinalValue{
							Origin:           originYear,
							AccumulatedClaim: 0,
						})

					}
					//TO SOLVE SECOND PART
				} else if originYear == productTriangle.Origin && len(finalValue) != 0 && originYear != productTriangle.DevelopmentYear {
					/*
						else if the origin years match and the final value is not empty and the origin year does not match the dev year
						get the total claim payment
					*/
					results[k] = append(results[k], models.FinalValue{
						Origin:           originYear,
						AccumulatedClaim: finalValue[len(finalValue)-1].AccumulatedClaim + productTriangle.Value,
					})
				} else {
					/*
						if the origin year and dev year match just retain the claim
					*/
					results[k] = append(results[k], models.FinalValue{
						Origin:           originYear,
						AccumulatedClaim: productTriangle.Value,
					})
				}

			}
		}
	}

	//FinalOutput converts the results to a 2d matrix and save the returned value into finalResult
	finalResult := helperfunctions.FinalOutPut(results, minOriginYear, totalDevYear)

	//write data to disc
	// persistingdata.WriteTxtFile(finalResult)

	//if successful return the final result
	return finalResult

}
