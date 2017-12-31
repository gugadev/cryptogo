package menu

import (
	"os"

	"github.com/abiosoft/ishell"
	"github.com/gugadev/cryptogo/helpers"
	"github.com/gugadev/cryptogo/request"
)

var art = `
   ______                 __        ______
  / ____/______  ______  / /_____  / ____/___
 / /   / ___/ / / / __ \/ __/ __ \/ / __/ __ \
/ /___/ /  / /_/ / /_/ / /_/ /_/ / /_/ / /_/ /
\____/_/   \__, / .___/\__/\____/\____/\____/
          /____/_/

✓ Select a coin to get information:
`

/*
ShowMain show the main menu
*/
func ShowMain() {
	var coin string
	var logger helpers.Logger
	shell := ishell.New()
	choice := shell.MultiChoice([]string{
		"Bitcoin",
		"Ethereum",
		"Monero",
		"Litecoin",
		"Ripple",
		"ZCash",
		"Dash",
		"Exit",
	}, art)

	if choice == 7 {
		os.Exit(0)
	} else {
		switch choice {
		case 0:
			coin = "bitcoin"
		case 1:
			coin = "ethereum"
		case 2:
			coin = "monero"
		case 3:
			coin = "litecoin"
		case 4:
			coin = "ripple"
		case 5:
			coin = "zcash"
		case 6:
			coin = "dash"
		}
		info := request.GetInfo(coin)[0]
		logger.Print(info)
		os.Exit(0)
	}
	shell.Run()
	shell.Close()
}
