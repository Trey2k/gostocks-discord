package mysql

import (
	"database/sql"
	"encoding/json"

	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	_ "github.com/go-sql-driver/mysql"
)

var connectionString string
var timeFormat string = "2006-01-02 15:04:05"

//Init init MySQL Connectiong
func Init() {
	config := &utils.Config.MySQL
	connectionString = config.Username + ":" + config.Password + "@tcp(" + config.IP + ":" + config.Port + ")/" + config.Database
}

//StoreOrder g
func StoreOrder(order utils.OrderStruct, optionData td.ExpDateOption) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	queryStr := "INSERT INTO `Orders`(`buy`, `risky`, `ticker`, `expDate`, `strikePrice`, `contractType`, `price`, `currentPrice`, `stopLoss`, `sender`, `messageID`, `description`, `active`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	senderBytes, err := json.Marshal(order.Sender)
	if err != nil {
		panic(err)
	}

	insert, err := db.Query(queryStr, order.Buy, order.Risky, order.Ticker, order.ExpDate.Format(timeFormat), order.StrikPrice, order.ContractType, order.Price, optionData.Ask, order.StopLoss, senderBytes, order.MessageID, optionData.Description, true)
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}
