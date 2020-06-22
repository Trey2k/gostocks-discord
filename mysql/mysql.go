package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	_ "github.com/go-sql-driver/mysql"
)

var connectionString string
var dateFormat string = "2006-01-02"
var dateTimeFormat string = "2006-01-02 15:04:05"

//StoredOrder stuff
type StoredOrder struct {
	ID            int
	Symbol        string
	PurchasePrice float64
	Contracts     int
	Status        string
	CreatedDate   time.Time
	UpdatedDate   time.Time
	Order         utils.OrderStruct
}

//Init init MySQL Connectiong
func Init() {
	config := &utils.Config.MySQL
	connectionString = config.Username + ":" + config.Password + "@tcp(" + config.IP + ":" + config.Port + ")/" + config.Database
}

//NewOrder stuff
func NewOrder(order utils.OrderStruct, optionData td.ExpDateOption, contracts int64) error {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	queryStr := "INSERT INTO `Orders`(`risky`, `ticker`, `symbol, `expDate`, `strikePrice`, `contractType`, `reportedPrice`, `purchasePrice`, `contracts`, `stopLoss`, `sender`, `messageID`, `message`, `status`, `orderID`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?.?)"

	senderBytes, err := json.Marshal(order.Sender)
	if err != nil {
		return err
	}

	insert, err := db.Query(queryStr, order.Risky, order.Ticker, optionData.Symbol, order.ExpDate.Format(dateFormat), order.StrikPrice, order.ContractType, order.Price, optionData.Last, contracts, order.StopLoss, senderBytes, order.MessageID, order.Message, "FILLED", 0) //TODO: default status should be pending. This is here for offline test runs
	if err != nil {
		return err
	}
	defer insert.Close()

	return nil
}

//FailedOrder stuff
func FailedOrder(order utils.OrderStruct, failCode int, failMessage string) error {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	queryStr := "INSERT INTO `FailedOrders`(`messageID`, `message`, `failCode`, `failMessage`) VALUES (?,?,?,?)"

	insert, err := db.Query(queryStr, order.MessageID, order.Message, failCode, failMessage)
	if err != nil {
		return err
	}
	defer insert.Close()

	return nil
}

//AlreadyOwn stuff
func AlreadyOwn(symbol string) (bool, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return true, err
	}
	defer db.Close()

	queryStr := "SELECT * FROM `Orders` WHERE `symbol`='" + symbol + "' and `status`<>'sold'"

	results, err := db.Query(queryStr)
	if err != nil {
		return true, err
	}
	defer results.Close()

	return results.Next(), nil
}

//RetriveActiveOrder stuff
func RetriveActiveOrder(symbol string) (StoredOrder, error) {
	var response StoredOrder
	var sender []byte
	db, err := sql.Open("mysql", connectionString+"?parseTime=true")
	if err != nil {
		return response, err
	}
	defer db.Close()

	queryStr := "SELECT * FROM `Orders` WHERE `ticker`='" + symbol + "' and `status`<>'sold'"

	results, err := db.Query(queryStr)
	if err != nil {
		return response, err
	}
	defer results.Close()

	for results.Next() {
		err = results.Scan(&response.ID, &response.Order.Risky, &response.Order.Ticker, &response.Symbol, &response.Order.ExpDate, &response.Order.StrikPrice, &response.Order.ContractType,
			&response.Order.Price, &response.PurchasePrice, &response.Contracts, &response.Order.StopLoss, &sender, &response.Order.MessageID,
			&response.Order.Message, &response.Status, &response.CreatedDate, &response.UpdatedDate)
		if err != nil {
			return response, err
		}
	}

	err = json.Unmarshal(sender, &response.Order.Sender)
	if err != nil {
		return response, err
	}

	return response, nil
}

//SellContract stuff
func SellContract(order utils.OrderStruct) error {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	queryStr := "UPDATE `Orders` SET `status`='sold',`updatedDate`='" + time.Now().Format(dateTimeFormat) + "' WHERE `ticker`='" + order.Ticker + "' AND `expDate`='" +
		order.ExpDate.Format(dateFormat) + "' AND `strikePrice`='" + fmt.Sprint(order.StrikPrice) + "' and `contractType`='" + order.ContractType + "'"

	insert, err := db.Query(queryStr)
	if err != nil {
		return err
	}
	defer insert.Close()

	return nil
}
