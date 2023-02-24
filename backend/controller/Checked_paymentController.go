package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut65/team11/entity"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////   		   controller Checked_payment    		////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// POST /Checked_payment
func CreateChecked_payment(c *gin.Context) {
	var Checked_payment entity.Checked_payment
	var Status_check entity.Status_check
	var Payment entity.Payment
	var Admin entity.Admin

	// 8: ผลลัพธ์ที่ได้จากขั้นตอนที่ * จะถูก bind เข้าตัวแปร Checked_payment
	if err := c.ShouldBindJSON(&Checked_payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Payment ด้วย id
	if tx := entity.DB().Where("id = ?", Checked_payment.Payment_ID).First(&Payment); tx.RowsAffected == 0 { //งงอยู่++++++++++++++++++++++++++++++++++
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment not find"})
		return
	}

	// 10: ค้นหา Status_check ด้วย id
	if tx := entity.DB().Where("id = ?", Checked_payment.Status_ID).First(&Status_check); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "โปรดตรวจสอบเหมือนคุณจะลืมเลือก สถานะ นะ"})
		return
	}

	// 11: ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", Checked_payment.Admin_ID).First(&Admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin not found"})
		return
	}

	// 12-15: แทรกการ validate
	if _, err := govalidator.ValidateStruct(Checked_payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 16: สร้าง Payment
	pm := entity.Checked_payment{
		Payment_ID: Checked_payment.Payment_ID, // โยงความสัมพันธ์กับ Entity Payment
		Status_ID:  Checked_payment.Status_ID,  // โยงความสัมพันธ์กับ Entity Status Check
		Date_time:  Checked_payment.Date_time,
		Other:      Checked_payment.Other,
		Message:    Checked_payment.Message,  //เพิ่มเข้ามาใหม่
		Admin_ID:   Checked_payment.Admin_ID, // โยงความสัมพันธ์กับ Entity Customer
	}

	// 17: บันทึก
	if err := entity.DB().Create(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pm})

	//สำหรับ การอัพเดต status เมื่อตรวจสอบแล้ว
	Update_payment_status(Checked_payment.Payment_ID, Checked_payment.Status_ID)

}

// ================================================== function List to frontend =================================================

func ListChecked_payment(c *gin.Context) {
	var Checked_payment []entity.Checked_payment
	if err := entity.DB().Preload("Admin").Preload("Status_check").Preload("Payment.OrderTech.ORDER").Find(&Checked_payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Checked_payment})
}

// fn สำหรับ List ข้อมูล CheckedPayment ทั้งหมด โดยไม่เอาสถานะ "รอตวจสอบ"
func List_only_checkedPayment(c *gin.Context) {
	var List_only_checkedPayment []entity.Checked_payment
	if err := entity.DB().Raw("SELECT * FROM checked_payments WHERE status_id != 3").Preload("Admin").Preload("Status_check").Preload("Payment.OrderTech.ORDER").Find(&List_only_checkedPayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": List_only_checkedPayment})

}

// ================================================================================================================================
// GET /Checked_payment/:id
// Get Checked_payment by id
func GetChecked_payment(c *gin.Context) {
	var Checked_payment entity.Checked_payment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM checked_payments WHERE id = ?", id).Preload("Admin").Preload("Status_check").Preload("Payment.OrderTech.ORDER").Preload("Payment").Find(&Checked_payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Checked_payment})
}

// Get Checked_payment by Payment_id
func GetCheckedpayment_by_PaymentID(c *gin.Context) {
	var Checked_payment entity.Checked_payment
	PaymentID := c.Param("id")
	if err := entity.DB().Raw("SELECT message  FROM checked_payments WHERE payment_id = ?", PaymentID).Scan(&Checked_payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Checked_payment})

}

// PATCH /UpdatePayment
func UpdateCheckedPayment(c *gin.Context) {
	var UpdateCheckedPayment entity.Checked_payment
	var Status_check entity.Status_check
	var Admin entity.Admin

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ * จะถูก bind เข้าตัวแปร Checked_payment
	if err := c.ShouldBindJSON(&UpdateCheckedPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// *: ค้นหา Status_check ด้วย id
	if tx := entity.DB().Where("id = ?", UpdateCheckedPayment.Status_ID).First(&Status_check); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "โปรดตรวจสอบเหมือนคุณจะลืมเลือก สถานะ นะ"})
		return
	}

	//*: ค้นหา admin ด้วย id
	if tx := entity.DB().Where("id = ?", UpdateCheckedPayment.Admin_ID).First(&Admin); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Admin not found"})
		return
	}

	// : แทรกการ validate
	if _, err := govalidator.ValidateStruct(UpdateCheckedPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pm := entity.Checked_payment{

		Status_ID: UpdateCheckedPayment.Status_ID, // โยงความสัมพันธ์กับ Entity Status Check
		Date_time: UpdateCheckedPayment.Date_time,
		Other:     UpdateCheckedPayment.Other,
		Message:   UpdateCheckedPayment.Message, //เพิ่มเข้ามาใหม่
	}

	if err := entity.DB().Model(pm).Where("id = ?", UpdateCheckedPayment.ID).Updates(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pm})

	//สำหรับ การอัพเดต status เมื่อตรวจสอบแล้ว
	payment_id_update := get_id_payment_for_status(UpdateCheckedPayment.ID)
	Update_payment_status(payment_id_update, UpdateCheckedPayment.Status_ID)
	//Update_payment_status(2, 2)
}

// DELETE /Checked_payment/:id
func DeleteChecked_payment(c *gin.Context) {
	id := c.Param("id")
	payment_id_for_del := get_id_payment_for_status(id)
	//if st ==0
	if tx := entity.DB().Exec("DELETE FROM Checked_payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Checked_payment not found"})
		// st =1
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

	//if st =1
	//สำหรับ การอัพเดต status เมื่อตรวจสอบแล้ว

	Update_payment_status(payment_id_for_del, 3)
	//Update_payment_status(2, 3)
}

//======================================================================================================================

// fn เพื่อ ดึงค่า payment id จาก ceckedPayment id
func get_id_payment_for_status(checkedPayment_id any) int {
	var ID_payment int
	entity.DB().Table("checked_payments").Select("payment_id").Where("id = ?", checkedPayment_id).Row().Scan(&ID_payment)
	return ID_payment
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////   		   controller Status_check    		  //////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// POST/Bank  สำหรับ สร้างข้อมูล
func CreateStatus_check(c *gin.Context) {
	var Status_check entity.Checked_payment
	if err := c.ShouldBindJSON(&Status_check); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&Status_check).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Status_check})
}

// GET /Status_check
// List status != 1
func ListStatus_check(c *gin.Context) {
	var Status_check []entity.Status_check
	if err := entity.DB().Raw("SELECT * FROM Status_checks WHERE id != 1 ").Scan(&Status_check).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Status_check})
}

// GET /Status_check/:id
// Get Status_check by id
func GetStatus_check(c *gin.Context) {
	var Status_check entity.Status_check
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&Status_check); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status_check not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Status_check})
}

// PATCH /Bank
func UpdateStatus_check(c *gin.Context) {
	var Status_check entity.Status_check
	if err := c.ShouldBindJSON(&Status_check); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Status_check.ID).First(&Status_check); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status_check not found"})
		return
	}

	if err := entity.DB().Save(&Status_check).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Status_check})
}

// DELETE /Status_check/:id
func DeleteStatus_check(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Status_checks WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status_check not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// ============================================= update status Payment =================================================================
func Update_payment_status(id, status any) {
	entity.DB().Table("payments").Where("id = ?", id).Updates(map[string]interface{}{"Status_ID": status})
}
