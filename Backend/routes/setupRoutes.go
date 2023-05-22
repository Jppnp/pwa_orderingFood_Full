package routes

import (
	"github.com/gin-gonic/gin"

	"pwaV3/controller"
)

func SetupRoutes(router *gin.Engine) {
	merchantController := &controller.MerchantController{}
	restaurantController := &controller.RestaurantController{}
	locationController := &controller.LocationController{}
	menuController := &controller.MenuController{}

	merchantRoutes := router.Group("/merchants")
	{
		merchantRoutes.GET("", merchantController.GetMerchants)
		merchantRoutes.GET(":id", merchantController.GetMerchantByID)
		merchantRoutes.GET("email/:email", merchantController.GetMerChantByEmail)
		merchantRoutes.POST("", merchantController.CreateMerchant)
		merchantRoutes.PATCH(":id", merchantController.UpdateMerchant)
		merchantRoutes.DELETE(":id", merchantController.DeleteMerchant)
	}

	// adminRoutes := router.Group("/admins")
	// {
	// 	adminRoutes.POST("", adminController.AddAdmin)
	// 	adminRoutes.PATCH(":rid/:mid", adminController.UpdateAdmin)
	// 	adminRoutes.DELETE(":rid/:mid", adminController.DeleteAdmin)
	// }

	restaurantRoutes := router.Group("/restaurants")
	{
		restaurantRoutes.GET("", restaurantController.GetRestaurants)
		restaurantRoutes.GET(":id", restaurantController.GetRestaurantByID)
		restaurantRoutes.POST("", restaurantController.CreateRestaurant)
		restaurantRoutes.PATCH(":id", restaurantController.UpdateRestaurant)
		restaurantRoutes.DELETE(":id", restaurantController.DeleteRestaurant)
	}
	resLocationRoutes := router.Group("/restaurant/locations")
	{
		resLocationRoutes.POST("", locationController.CreateRestaurantLocation)
		resLocationRoutes.GET(":restaurantID", locationController.GetRestaurantLocationsByRestaurantID)
		resLocationRoutes.GET("", locationController.GetRestaurantLocations)
		resLocationRoutes.PATCH(":id", locationController.UpdateRestaurantLocation)
		resLocationRoutes.DELETE(":id", locationController.DeleteRestaurantLocation)
		resLocationRoutes.GET("location/:id", locationController.GetLocationByLocationID)
	}

	menuRoutes := router.Group("/menus")
	{
		menuRoutes.POST(":rid", menuController.CreateMenu)
		menuRoutes.GET("", menuController.GetAllMenus)
		menuRoutes.GET("location/:rid", menuController.GetMenusByLocation)
		menuRoutes.GET(":id", menuController.GetMenu)
		menuRoutes.PATCH("update/:rid/:menuID", menuController.UpdateMenu)
		menuRoutes.DELETE("delete/:rid/:menuID", menuController.DeleteMenu)
	}

}
