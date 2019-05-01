package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
)

// Invoices contains a slice of Invoices
type Invoices struct {
	InvoiceList []Invoice
}

// Invoice contains the top level structure containing all information
type Invoice struct {
	Info  Bin     `json:"bininfo,omitempty"`
	List  []Items `json:"items,omitempty"`
	Total Totals  `json:"totals,omitempty"`
}

// Bin contains a row in BinInfo
type Bin struct {
	BinID            int    `json:"bin_id,omitempty"`
	JoinName         string `json:"join_name,omitempty"`
	ClubID           string `json:"club_id,omitempty"`
	LiquidatorName   string `json:"liquidator_name,omitempty"`
	LiquidatorNumber int    `json:"liquidator_number,omitempty"`
}

// Items contains the actual item in the list
type Items struct {
	BinItemID                      int    `json:"bin_item_id,omitempty"`
	ItemDescription                string `json:"item_description,omitempty"`
	ItemCategory                   string `json:"item_category,omitempty"`
	ItemNumber                     int    `json:"item_number,omitempty"`
	ItemQuantity                   int    `json:"item_quantity,omitempty"`
	ItemRetailAmount               string `json:"item_retail_amount,omitempty"`
	LiquidatorCategoryRecoveryRate string `json:"liquidator_category_recovery_rate,omitempty"`
	ItemRetailRate                 string `json:"item_retail_rate,omitempty"`
	TotalItemRate                  string `json:"total_item_rate,omitempty"`
}

// Totals contains the sum total of all items in the list
type Totals struct {
	TotalQuantity int    `json:"total_quantity,omitempty"`
	TotalRetail   string `json:"total_retail,omitempty"`
}

func main() {

	file, err := ioutil.ReadFile("singleInvoice.json")

	if err != nil {
		panic(err)
	}

	// fmt.Println(file)

	var invoice Invoice

	json.Unmarshal(file, &invoice)

	// fmt.Println(invoice)

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, "Invoice")

	pdf.Cell(10, 20, string(invoice.Info.BinID))
	pdf.Cell(20, 30, invoice.Info.ClubID)
	pdf.Cell(30, 40, invoice.Info.JoinName)
	pdf.Cell(40, 50, invoice.Info.LiquidatorName)
	pdf.Cell(50, 60, string(invoice.Info.LiquidatorNumber))

	err = pdf.OutputFileAndClose("Hello.pdf")
	if err != nil {
		panic(err)
	}
}
