package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"time"

	"github.com/xuri/excelize/v2"
)

type Calculator struct {
	FullPayment        string
	AmountPaid         float64
	TenancyLength      int
	NewTenancyLength   int
	FirstInstalment    float64
	AmountInInstalment float64
	AmountInFull       float64
	RemainingAmount    float64
	RecognitionAmount  float64
	NameOfApplicant    string
	DesktopDir         string
}

func Calculate() *Calculator {
	result := &Calculator{
		DesktopDir: getOS(),
	}
	return result
}

func getOS() string {
	system := runtime.GOOS
	var desktopDir string

	switch system {
	case "windows":
		user := os.Getenv("USERPROFILE")
		desktopDir := user + "\\Desktop\\"
		return desktopDir
	case "darwin":
		user := os.Getenv("HOME")
		desktopDir := user + "/Desktop/"
		return desktopDir
	default:
		fmt.Printf("%s.\n", system)
	}
	return desktopDir
}

//Creating the function to round the float number to a specific decimal place.
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

//Creating the function that will calculate the amount for the the Instalment Invoices.
func createInstalments(amountPaid float64, tenancyLenght int, nameOfApplicant string) {
	firstInstalment := amountPaid / float64(tenancyLenght)
	amountInInstalments := amountPaid - firstInstalment

	//Converting the values to two decimal places.
	firstInstalment = toFixed(firstInstalment, 2)
	amountInInstalments = toFixed(amountInInstalments, 2)
	fmt.Printf("\nThe amount of the first invoice is: £ %.2f.\nThe amount of the remaining %d invoices is: £ %.2f.\n", firstInstalment, (tenancyLenght - 1), amountInInstalments)
	createExcel(nameOfApplicant, amountPaid, tenancyLenght, firstInstalment, amountInInstalments, 0.00, true)
}

//Creating the function that will calculate the amount for the the Bulk Invoice.
func createFullPayment(amountPaid float64, tenancyLenght int, nameOfApplicant string) {
	var remainingAmount float64
	var recognitionAmount float64
	var amountInFull float64

	if tenancyLenght < 12 {
		newtenancyLenght := 12
		amountInFull = amountPaid * float64(newtenancyLenght)
		recognitionAmount := amountInFull / float64(tenancyLenght)
		remainingAmount = amountInFull - recognitionAmount
	} else {
		amountInFull := amountPaid * float64(tenancyLenght)
		recognitionAmount := amountInFull / float64(tenancyLenght)
		remainingAmount = amountInFull - recognitionAmount
	}
	//Converting the values to two decimal places.
	recognitionAmount = toFixed(recognitionAmount, 2)
	remainingAmount = toFixed(remainingAmount, 2)
	amountInFull = toFixed(amountInFull, 2)
	fmt.Printf("\nThe full amount is: £ %.2f.\nThe remaining amount is: £ %.2f.\n", amountInFull, remainingAmount)
	fmt.Printf("The recognition amount is: £ %.2f.\n", recognitionAmount)
	createExcel(nameOfApplicant, amountPaid, tenancyLenght, amountInFull, remainingAmount, recognitionAmount, false)
}

//Creating a function that will enter all the details from the program and add them to an Excel for historic data.
func createExcel(name string, amountPaid float64, tenancyLenght int, amountInFull float64, remainingAmount float64, recognitionAmount float64, isFull bool) {

	excelFile := excelize.NewFile()
	var p Calculator
	cTime := time.Now()
	OUTPUT_DIRECTORY := p.DesktopDir + "Invoice Calculations/" + cTime.Format("06-Jan-02") + "/"
	excelName := OUTPUT_DIRECTORY + name + ".xlsx"

	if isFull {
		excelFile.SetCellValue("Sheet1", "C5", "Name of Applicant")
		excelFile.SetCellValue("Sheet1", "D5", name)
		excelFile.SetCellValue("Sheet1", "C6", "Fees amount paid")
		excelFile.SetCellValue("Sheet1", "D6", amountPaid)
		excelFile.SetCellValue("Sheet1", "C7", "Tenancy duration in months")
		excelFile.SetCellValue("Sheet1", "D7", tenancyLenght)
		excelFile.SetCellValue("Sheet1", "C8", "Fees amount paid")
		excelFile.SetCellValue("Sheet1", "D8", amountInFull)
		excelFile.SetCellValue("Sheet1", "C9", "Remaining amount is")
		excelFile.SetCellValue("Sheet1", "D9", remainingAmount)
		excelFile.SetCellValue("Sheet1", "C10", "Recognition Amount")
		excelFile.SetCellValue("Sheet1", "D10", recognitionAmount)
	} else {
		excelFile.SetCellValue("Sheet1", "C5", "Name of Applicant")
		excelFile.SetCellValue("Sheet1", "D5", name)
		excelFile.SetCellValue("Sheet1", "C6", "Fees amount paid")
		excelFile.SetCellValue("Sheet1", "D6", amountPaid)
		excelFile.SetCellValue("Sheet1", "C7", "Tenancy duration in months")
		excelFile.SetCellValue("Sheet1", "D7", tenancyLenght)
		excelFile.SetCellValue("Sheet1", "C8", "Full Amount")
		excelFile.SetCellValue("Sheet1", "D8", amountInFull)
		excelFile.SetCellValue("Sheet1", "C9", "Remaining Amount")
		excelFile.SetCellValue("Sheet1", "D9", remainingAmount)
		excelFile.SetCellValue("Sheet1", "C10", "Recognition Amount")
		excelFile.SetCellValue("Sheet1", "D10", recognitionAmount)
	}
	if err := excelFile.SaveAs(excelName); err != nil {
		log.Fatal(err)
	}
}
