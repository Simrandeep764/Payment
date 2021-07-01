package Models

import (
	"Checkout/Config"
	_ "fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DoCheckout(pay *Checkout) (err error) {
	if err = Config.DB.Create(pay).Error; err != nil {
		return err
	}
	return nil
}

func GetCheckoutDetailsById(pay *Checkout, id string) (err error) {
	if err = Config.DB.Where("Checkout_id = ?", id).First(pay).Error; err != nil {
		return err
	}
	return nil
}
