package database

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/danilodcn/go-intensivo/internal/order/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	fmt.Println("Inicio do SetupSuite")
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.NoError(err)
	suite.Db = db
	fmt.Println("Fim do SetupSuite")
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	fmt.Println("Inicio do TearDownTest")
	suite.Db.Close()
	fmt.Println("Fim do TearDownTest")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAOrder_WhenSave_ThenShoulSaveOrder() {
	order, err := entity.NewOrder(234, 34.081)

	suite.NoError(err)
	repo := NewOrderRepository(suite.Db)

	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("SELECT * from orders WHERE id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)
	suite.NoError(err)

	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
