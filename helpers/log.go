package helpers

import (
	"fmt"
	"os"

	"github.com/gugadev/cryptogo/models"
	"github.com/olekukonko/tablewriter"
)

// Logger Util methods
type Logger struct{}

// Print print information about a coin
func (m *Logger) Print(coin models.Coin) {
	fmt.Printf("\n✓ Currency information:\n\n")
	data := [][]string{
		[]string{
			coin.Name,
			coin.Symbol,
			coin.Rank,
			coin.PriceUSD,
			coin.PriceEUR,
		},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(1)
	// table.SetBorder(false)
	table.SetHeader([]string{
		"Name",
		"Symbol",
		"Rank",
		"USD Price",
		"EUR Price",
	})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
	)
	table.AppendBulk(data)
	table.Render()
}
