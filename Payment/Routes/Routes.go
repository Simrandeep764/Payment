package Routes

import (
	"Checkout/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() {

	r3 := gin.Default()
	grp4 := r3.Group("v1/")
	{
		grp4.GET("Checkout/:pid", Controllers.GetCheckoutDetailsById)
		grp4.POST("Checkout", Controllers.DoCheckout)
	}
	r3.Run(":7003")
}
