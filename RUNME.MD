# CALCULATE ACCUMULATED CLAIM

## Getting Started

### Prerequistes of Program

To start using this application these prerequistes must be followed sequentially
- Install **[Go](https://go.dev/)**: (https://go.dev/doc/devel/release) (now version **1.16+** is required).

- Clone repository on the local system using any of these address 

```sh
$ git clone https://github.com/wtwictrecruitment/intern-developer-technicaltest-Intern004.git
```

```sh 
$ git clone git@github.com:wtwictrecruitment/intern-developer-technicaltest-Intern004.git
```
### Usage of Program
- Update and install all package dependencies by running this command

```sh
go mod tidy
```
- Compile and run the main application by running this command
```sh
go run main.go
```
- Check test cases : cd into "testFiles" folder then invoke test cases by running this command
```sh 
go test
```
## To manipulate or change input data and view output data
- The input data is stored in input_data.txt
- The output result is stored in out_data.txt
```sh 
input_data.txt
output_data.txt
```

### Main Package , Main Entry and Exist of Program
```sh
func main() {

	app.ClaimApp("input_data.txt")

}

```
### Helper Funtions
```sh

- func GetMinYearAndDevelopementYears(inputData [][]string) (int, int)

- func Convert2dDataToMap(dataWithoutHeader [][]string) map[string]models.Triangle {}

- func ConvertDataToMatix(results map[string][]models.FinalValue) [][]string {}

- func FinalOutPut(results map[string][]models.FinalValue, minOriginYear int, totalDevYear int) [][]string{}

```

### Models And Datastructures
```sh

type ProductTriangle struct {
	Product         string
	Origin          int
	DevelopmentYear int
	Value           float64
}

type Triangle map[string]ProductTriangle


type FinalValue struct {
	Origin           int
	AccumulatedClaim float64
}

```

### Main Program Function
```sh
- func ClaimApp(filename string) [][]string {}

```

### Time Constraint
```sh
- If i had more time, i would have written more unit test cases and converted the program to a web application using golang default net/http package and design the user interface with the help of tools like html, css

```
