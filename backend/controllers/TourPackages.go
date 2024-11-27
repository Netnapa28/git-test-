package controllers

import (
	"net/http"
	"fmt"
	"toursystem/config"
	"toursystem/entity"
	"github.com/gin-gonic/gin"
	
)
type TourPackageResponse struct {
    ID          uint                   `json:"id"`
    PackageCode string                 `json:"package_code"`
    TourName    string                 `json:"tour_name"`
    Duration    string                 `json:"duration"`
    Province    *ProvinceResponse      `json:"province,omitempty"`
    TourPrices  []TourPriceResponse    `json:"tour_prices,omitempty"`
    TourImages  []TourImageResponse    `json:"tour_images,omitempty"`
    Activities  []ActivityResponse     `json:"activities,omitempty"`
}

type ProvinceResponse struct {
    ID           uint   `json:"id"`
    ProvinceName string `json:"province_name"`
}

type TourPriceResponse struct {
    ID    uint    `json:"id"`
    Price float64 `json:"price"`
}

type TourImageResponse struct {
    ID       uint   `json:"id"`
    FilePath string `json:"file_path"`
}

type ActivityResponse struct {
    ID           uint   `json:"id"`
    ActivityName string `json:"activity_name"`
    Description  string `json:"description"`
    DateTime     string `json:"date_time"`
}

// GET /tour-packages
/*func ListTourPackages(c *gin.Context) {
	var tourPackages []entity.TourPackages

	db := config.DB()

	if err := db.Preload("Province").Preload("TourPrices").Preload("TourImages").Preload("TourDescriptions").Find(&tourPackages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, &tourPackages)
}*/


func GetAllTourPackages(c *gin.Context) {
    var tourPackages []entity.TourPackages
    var responses []TourPackageResponse

    db := config.DB()

    // ดึงข้อมูล tour package ทั้งหมด
    err := db.Table("tour_packages").
        Select("id, package_code, tour_name, duration, province_id").
        Find(&tourPackages).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch tour packages"})
        return
    }

    // แปลงข้อมูล tour package แต่ละรายการ
    for _, tourPackage := range tourPackages {
        response := TourPackageResponse{
            ID:          tourPackage.ID,
            PackageCode: tourPackage.PackageCode,
            TourName:    tourPackage.TourName,
            Duration:    tourPackage.Duration,
        }

        // ดึงข้อมูล province
        var province ProvinceResponse
        db.Table("provinces").
            Select("id, province_name").
            Where("id = ?", tourPackage.ProvinceID).
            First(&province)
        response.Province = &province

        // ดึงข้อมูล tour prices
        db.Table("tour_prices").
            Select("id, price").
            Where("tour_package_id = ?", tourPackage.ID).
            Scan(&response.TourPrices)

        // ดึงข้อมูล tour images
        db.Table("tour_images").
            Select("id, file_path").
            Where("tour_package_id = ?", tourPackage.ID).
            Scan(&response.TourImages)

        // ดึงข้อมูล activities
        db.Table("activities").
            Select("id, activity_name, description, date_time").
            Where("tour_package_id = ?", tourPackage.ID).
            Scan(&response.Activities)

        // เพิ่ม response ลงใน responses
        responses = append(responses, response)
    }

    // ส่งข้อมูลทั้งหมดกลับ
    c.JSON(http.StatusOK, responses)
    fmt.Println("OKkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")
}






func ListTourPackages(c *gin.Context) {
	var tourPackages []entity.TourPackages

	db := config.DB()

	if err := db.Preload("Province").
		Preload("Activities").
		Preload("TourPrices").
		Preload("TourImages").
		Preload("TourDescriptions").
		Preload("TourSchedules"). // ตรวจสอบชื่อ relation ให้ตรงกับ model
		Preload("TourSchedules.TourScheduleStatus"). // ตรวจสอบชื่อ relation ในตารางย่อย
		Preload("Accommodations.Hotel").
		Find(&tourPackages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &tourPackages)
}



// GET /tour-package/:id
/*func GetTourPackageByID(c *gin.Context) {
	var tourpackage entity.TourPackages
    id := c.Param("id")

    db := config.DB()

    if err := db.Preload("Province").Preload("TourPrices.PersonType").Preload("TourImages").Preload("TourDescriptions").Preload("Activities.Location").Preload("TourSchedules.TourScheduleStatus").First(&tourpackage, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    c.JSON(http.StatusOK, tourpackage)
}*/

/*func GetTourPackageByID(c *gin.Context) {
	var tourpackage entity.TourPackages
    id := c.Param("id")

    db := config.DB()

    if err := db.Preload("Province").Preload("TourPrices.PersonType").Preload("TourImages").Preload("TourDescriptions").Preload("Activities.Location").Preload("TourSchedules.TourScheduleStatus").Preload("Accommodations.Hotel").First(&tourpackage, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    c.JSON(http.StatusOK, tourpackage)
}*/

/*func GetTourPackageByID(c *gin.Context) {
    var tourPackage entity.TourPackages
    id := c.Param("id")
    fmt.Println("Requested ID:", id) // ตรวจสอบค่า id

    db := config.DB()
    if db == nil {
        fmt.Println("Failed to connect to the database")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
        return
    }

    // เลือกคอลัมน์ที่ต้องการจาก TourPackages
    selectFields := "id, package_code, tour_name, duration, province_id"

    err := db.Select(selectFields).
        Preload("Province", func(db *gorm.DB) *gorm.DB {
            return db.Select("id, province_name") // เลือกคอลัมน์จาก Province
        }).
        Preload("TourPrices", func(db *gorm.DB) *gorm.DB {
            return db.Select("id, price")
        }).
        Preload("TourImages", func(db *gorm.DB) *gorm.DB {
            return db.Select("id, file_path")
        }).
        Preload("TourDescriptions", func(db *gorm.DB) *gorm.DB {
            return db.Select("id, intro, package_detail, trip_highlight, places_highlight")
        }).
        Preload("Activities", func(db *gorm.DB) *gorm.DB {
            return db.Select("id, activity_name, description, date_time")
        }).
        Preload("TourSchedules", func(db *gorm.DB) *gorm.DB {
            return db.Select("id, start_date, end_date, available_slots, tour_schedule_status_id").
                Preload("TourScheduleStatus", func(db *gorm.DB) *gorm.DB {
                    return db.Select("id, status_name") // เลือกคอลัมน์จาก TourScheduleStatus
                })
        }).
        Preload("Accommodations", func(db *gorm.DB) *gorm.DB {
            return db.Select("id, hotel_id").
                Preload("Hotel", func(db *gorm.DB) *gorm.DB {
                    return db.Select("id, hotel_name") // เลือกคอลัมน์จาก Hotel
                })
        }).
        First(&tourPackage, "id = ?", id).Debug().Error // เพิ่ม Debug()

    if err != nil {
        fmt.Println("Error:", err) // แสดงข้อผิดพลาด
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    c.JSON(http.StatusOK, tourPackage)
}*/


/*func GetTourPackageByID(c *gin.Context) {
    var tourPackage entity.TourPackages
    id := c.Param("id")
    fmt.Println("Requested ID:", id) // ตรวจสอบค่า id

    db := config.DB()
    if db == nil {
        fmt.Println("Failed to connect to the database")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
        return
    }

    // ใช้ Joins แทน Preload
    err := db.Table("tour_packages").
        Select("tour_packages.id, tour_packages.package_code, tour_packages.tour_name, tour_packages.duration, tour_packages.province_id, " +
            "provinces.id AS province_id, provinces.province_name, " +
            "tour_prices.id AS tour_prices_id, tour_prices.price, " +
            "tour_images.id AS tour_images_id, tour_images.file_path, " +
            "tour_descriptions.id AS tour_descriptions_id, tour_descriptions.intro, tour_descriptions.package_detail, tour_descriptions.trip_highlight, tour_descriptions.places_highlight, " +
            "activities.id AS activities_id, activities.activity_name, activities.description, activities.date_time, " +
            "tour_schedules.id AS tour_schedules_id, tour_schedules.start_date, tour_schedules.end_date, tour_schedules.available_slots, tour_schedules.tour_schedule_status_id, " +
            "tour_schedule_statuses.id AS tour_schedule_status_id, tour_schedule_statuses.status_name, " +
            "accommodations.id AS accommodations_id, accommodations.hotel_id, " +
            "hotels.id AS hotel_id, hotels.hotel_name").
        Joins("JOIN provinces ON provinces.id = tour_packages.province_id").
        Joins("JOIN tour_prices ON tour_prices.tour_package_id = tour_packages.id").
        Joins("JOIN tour_images ON tour_images.tour_package_id = tour_packages.id").
        Joins("JOIN tour_descriptions ON tour_descriptions.tour_package_id = tour_packages.id").
        Joins("JOIN activities ON activities.tour_package_id = tour_packages.id").
        Joins("JOIN tour_schedules ON tour_schedules.tour_package_id = tour_packages.id").
        Joins("JOIN tour_schedule_statuses ON tour_schedule_statuses.id = tour_schedules.tour_schedule_status_id").
        Joins("JOIN accommodations ON accommodations.tour_package_id = tour_packages.id").
        Joins("JOIN hotels ON hotels.id = accommodations.hotel_id").
        Where("tour_packages.id = ?", id).
        First(&tourPackage).Debug().Error

    if err != nil {
        fmt.Println("Error:", err) // แสดงข้อผิดพลาด
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    c.JSON(http.StatusOK, tourPackage)
}*/

func GetTourPackageByID(c *gin.Context) {
    var tourPackage entity.TourPackages
    var response TourPackageResponse

    id := c.Param("id")
    db := config.DB()

    // ดึงข้อมูล tour package หลัก
    err := db.Table("tour_packages").
        Select("id, package_code, tour_name, duration, province_id").
        Where("id = ?", id).
        First(&tourPackage).Error
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "tour package not found"})
        return
    }

    // แปลงข้อมูล tour package
    response = TourPackageResponse{
        ID:          tourPackage.ID,
        PackageCode: tourPackage.PackageCode,
        TourName:    tourPackage.TourName,
        Duration:    tourPackage.Duration,
    }

    // ดึงข้อมูล province
    var province ProvinceResponse
    db.Table("provinces").
        Select("id, province_name").
        Where("id = ?", tourPackage.ProvinceID).
        First(&province)
    response.Province = &province

    // ดึงข้อมูล tour prices
    db.Table("tour_prices").
        Select("id, price").
        Where("tour_package_id = ?", tourPackage.ID).
        Scan(&response.TourPrices)

    // ดึงข้อมูล tour images
    db.Table("tour_images").
        Select("id, file_path").
        Where("tour_package_id = ?", tourPackage.ID).
        Scan(&response.TourImages)

    // ดึงข้อมูล activities
    db.Table("activities").
        Select("id, activity_name, description, date_time").
        Where("tour_package_id = ?", tourPackage.ID).
        Scan(&response.Activities)

    c.JSON(http.StatusOK, response)
}





/*func getTourPackages(c *gin.Context) {
	var tourPackages []entity.TourPackages
	var response []map[string]interface{}

	db := config.DB()

	// ดึงข้อมูลจากฐานข้อมูล พร้อม preload relations
	if err := db.Preload("Province").
		Preload("Activities").
		Preload("TourPrices").
		Preload("TourImages").
		Preload("TourDescriptions").
		Preload("TourTourSchedules").
		Preload("TourTourSchedules.TourScheduleStatuses"). // preload status ใน schedules
		Find(&tourPackages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// สร้าง response ในรูปแบบที่ต้องการ
	for _, tour := range tourPackages {
		// ตรวจสอบข้อมูลใน relations เพื่อหลีกเลี่ยง nil
		provinceName := ""
		if tour.Province != nil {
			provinceName = tour.Province.ProvinceName
		}

		price := 0.0
		if len(tour.TourPrices) > 0 {
			price = tourPackages.TourPrices[0].Price // สมมติว่าใช้ราคาตัวแรก
		}

		availableSlots := 0
		statusName := ""
		if len(tour.TourSchedules) > 0 {
			availableSlots = tour.TourSchedules[0].AvailableSlots
			if tour.TourSchedules[0].TourScheduleStatuses != nil {
				statusName = tour.TourSchedules[0].TourScheduleStatuses.StatusName
			}
		}

		// สร้าง item
		item := map[string]interface{}{
			"PackageCode":    tour.PackageCode,
			"TourName":       tour.TourName,
			"ProvinceName":   provinceName,
			"Duration":       tour.Duration,
			"AvailableSlots": availableSlots,
			"Price":          price,
			"StatusName":     statusName,
		}

		// เพิ่ม item ใน response
		response = append(response, item)
	}

	// ส่ง response กลับในรูป JSON
	c.JSON(http.StatusOK, response)
}*/


func GetTourPackages1(c *gin.Context) {
	var tourPackages1 []entity.TourPackages
	//var response []map[string]interface{}

	db := config.DB()

	// ดึงข้อมูลจากฐานข้อมูล พร้อม preload relations
	if err := db.Preload("Province").
		Preload("Activities").
		Preload("TourPrices").
		Preload("TourImages").
		Preload("TourDescriptions").
		Preload("TourSchedules"). // ตรวจสอบชื่อ relation ให้ตรงกับ model
		Preload("TourSchedules.TourScheduleStatus"). // ตรวจสอบชื่อ relation ในตารางย่อย
		Find(&tourPackages1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// สร้าง response ในรูปแบบที่ต้องการ
	/*for _, tour := range tourPackages {
		// ตรวจสอบข้อมูลใน relations เพื่อหลีกเลี่ยง nil
		provinceName := ""
		if tour.Province != nil {
			provinceName = tour.Province.ProvinceName
		}

		price := 0.00
		if len(tour.TourPrices) > 0 {
			price = tour.TourPrices[0].Price // ใช้ราคาตัวแรก
		}

		availableSlots := 0
		statusName := ""
		if len(tour.TourSchedules) > 0 { // ตรวจสอบ relation `TourSchedules` ว่ามีข้อมูลหรือไม่
			availableSlots = tour.TourSchedules[0].AvailableSlots
			if tour.TourSchedules[0].TourScheduleStatuses != nil { // ตรวจสอบ `TourScheduleStatuses` ว่าไม่ใช่ nil
				statusName = tour.TourSchedules[0].TourScheduleStatuses.StatusName
			}
		}

		// สร้าง item
		item := map[string]interface{}{
			"PackageCode":    tour.PackageCode,
			"TourName":       tour.TourName,
			"ProvinceName":   provinceName,
			"Duration":       tour.Duration,
			"AvailableSlots": availableSlots,
			"Price":          price,
			"StatusName":     statusName,
		}

		// เพิ่ม item ใน response
		response = append(response, item)
	}*/

	// ส่ง response กลับในรูป JSON
	c.JSON(http.StatusOK, &tourPackages1)
}
