package main

import (
	"database/sql"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"log"
	"os"
	"shop/internal/models"
	"shop/internal/storage"
	"shop/internal/storage/db"
	"strings"
)

type ToOutModel struct {
	Category string
	Orders   []models.OrderProduct
}

func main() {
	args := os.Args[1:]
	dbConn, err := db.Dial()
	if err != nil {
		log.Fatal(err)
	}
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {

		}
	}(dbConn)

	s := storage.NewStorage(dbConn)
	orders, err := s.OrderProducts.GetManyOrderProductsByIds(args...)
	if err != nil {
		print(err.Error())
		return
	}
	var outOrders []*ToOutModel
first:
	for _, order := range *orders {
		for _, outOrder := range outOrders {
			if outOrder.Category == order.ProductCategory {
				outOrder.Orders = append(outOrder.Orders, order)
				continue first
			}
		}
		outOrders = append(outOrders, &ToOutModel{
			Category: order.ProductCategory,
			Orders:   []models.OrderProduct{order},
		})
	}

	for _, order := range outOrders {
		toTableOutput(order.Orders)
		fmt.Println()
	}
}

func toTableOutput(orders []models.OrderProduct) {
	rowHeader := table.Row{"Номер заказа", "ID товара", "Название товара", "Количество", "Описание товара"}
	tw := table.NewWriter()
	tw.AppendHeader(rowHeader)
	tw.SetTitle(strings.ToUpper(orders[0].ProductCategory))
	tw.Style().Options.DrawBorder = true
	tw.Style().Options.SeparateRows = true
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 3, Align: text.AlignLeft, WidthMin: 30},
		{Number: 5, Align: text.AlignLeft, WidthMax: 50, WidthMin: 50},
	})
	tw.SetStyle(table.StyleColoredBlueWhiteOnBlack)
	for _, order := range orders {
		tw.AppendRow(table.Row{order.OrderID, order.ProductID, order.ProductTitle, order.ProductQuantity, order.ProductDescription})
	}
	fmt.Println(tw.Render())
}
