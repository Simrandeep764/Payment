package Controllers

import (
	"Checkout/Config"
	"Checkout/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func DoCheckout(c *gin.Context) {
	var pay Models.Checkout
	c.BindJSON(&pay)
	fmt.Println(pay)
	cid := pay.CartId
	fmt.Println(cid)
	var m sync.Mutex
	err := Models.DoCheckout(&pay)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pay)
		m.Lock()
		AddOrder(uint(cid), pay.CheckoutId)
		m.Unlock()
	}
}

func GetCheckoutDetailsById(c *gin.Context) {
	id := c.Params.ByName("pid")
	var pay Models.Checkout
	err := Models.GetCheckoutDetailsById(&pay, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pay)
	}
}

func AddOrder(cid, Checkout_id uint) {
	var cart []Models.Cart
	Config.DB.Raw("SELECT * FROM Cartdb.cart WHERE Cartdb.cart.cart_id = ?", cid).Scan(&cart)
	fmt.Println(cart)

	var bill_amount int
	var pid uint
	var purchased_qty int
	var inv Models.Inventory
	var available_qty int
	var customer_id uint
	for _, v := range cart {

		pid = uint(v.ProductId)
		customer_id = uint(v.CustomerId)
		purchased_qty = v.Qty
		Config.DB.Raw("SELECT * FROM Productdb.inventory WHERE Productdb.inventory.product_id = ?", pid).Scan(&inv)
		available_qty = inv.ProductQty

		if available_qty >= purchased_qty {
			new_qty := available_qty - purchased_qty
			Config.DB.Exec("UPDATE Productdb.inventory SET Productdb.inventory.product_qty = ? WHERE Productdb.inventory.product_id = ?", new_qty, pid)
			bill_amount += (inv.ProductPrice * purchased_qty)
		} else {
			fmt.Printf("No enough stocks for product %d\n", pid)
			if len(cart) > 1 {
				fmt.Println("Billing is continued for other purchased items")
			}
		}
	}
	postBody, _ := json.Marshal(map[string]uint{
		"pid":    Checkout_id,
		"custid": customer_id,
		"amt":    uint(bill_amount),
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, _ := http.Post("http://localhost:8004/v1/order/order", "application/json", responseBody)
	fmt.Println(resp)
}
