package main

import (
	"net/http"
	"toursystem/config"
	controllers "toursystem/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT = "8000"

func main() {
	config.ConnectionDB()
	config.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// r.POST("/signin-customer", controller.SignInForCustomer)
	// r.POST("/signin-owner", controller.SignInForOwner)

	r.Static("/images", "./images")

	router := r.Group("/")

	{
		// ตั้งค่า CORS
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:5173"}, // พอร์ตของ Vite
			AllowMethods: []string{"POST", "GET", "OPTIONS", "PATCH"},
			AllowHeaders: []string{"Content-Type", "Authorization"},
		}))

		router.POST("/generate-qr", controllers.GenerateQR)
		
		// BookingDetails
		router.GET("/booking-details", controllers.ListBookingDetails)
		router.POST("/booking-detail", controllers.CreateBookingDetail)

		// Bookings
		router.GET("/bookings", controllers.ListBookings)
		router.GET("/booking/:id", controllers.GetBookingByID)
		router.GET("/bookings/:customerId", controllers.GetBookingByCustomerID)
		router.POST("/booking", controllers.CreateBooking)
		router.PATCH("/booking/:id", controllers.UpdateBooking)

		// Customers
		router.GET("/customer/:id", controllers.GetCustomerByID)

		// PersonTypes
		router.GET("/person-types", controllers.ListPersonTypes)

		// Promotions
		router.GET("/promotion/:code", controllers.GetPromotionByCode)

		// Provinces
		router.GET("/provinces", controllers.ListProvinces)

		// RoomTypes
		router.GET("/room-types", controllers.ListRoomTypes)

		// TourImages
		router.GET("/tour-image/:tourpackageId", controllers.GetTourImageByTourPackageID)

		// TourPackages
		router.GET("/tour-packages", controllers.ListTourPackages)

		router.GET("/tour-package/:id", controllers.GetTourPackageByID)/**/
		router.GET("/all-tour-packages", controllers.GetAllTourPackages)  /**/

		// TourSchedule
		router.GET("/tour-schedule/:id", controllers.GetTourScheduleByID)
		router.PATCH("/tour-schedule/:id", controllers.UpdateTourSchedule)

	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	r.Run("localhost:" + PORT)

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}