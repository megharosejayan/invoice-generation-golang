package main

import (
	"fmt"

	internal "github.com/megharosejayan/invoice-generation-golang/internal"
)

func main() {
	// Generate sample invoice data
	ecommerceInvoiceData, err := internal.NewInvoiceData("カレーパン", 1, 3000.50)
	if err != nil {
		panic(err)
	}
	laptopInvoiceData, err := internal.NewInvoiceData("ドーナツ", 2, 2000.70)
	if err != nil {
		panic(err)
	}
	// Invoice Items collection
	invoiceItems := []*internal.InvoiceData{ecommerceInvoiceData, laptopInvoiceData}

	// Create single invoice
	invoice := internal.CreateInvoice("森のくまさん", "米子市", invoiceItems)
	err = internal.GenerateInvoicePdf(*invoice)
	fmt.Printf("The Total Amount is: %f", invoice.CalculateInvoiceTotalAmount())
	fmt.Printf("世界")
}
