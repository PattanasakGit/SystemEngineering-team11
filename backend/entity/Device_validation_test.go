package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_CPU_not_blank(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "",
		Monitor:    "test",
		GPU:        "test",
		RAM:        "test",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรุณากรอกข้อมูล CPU"))
}

func Test_CPU_not_specialCharacter(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "##",
		Monitor:    "test",
		GPU:        "test",
		RAM:        "test",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรอกข้อมูล CPU ไม่ถูกต้อง"))
}

func Test_Monitor_not_blank(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "",
		GPU:        "test",
		RAM:        "test",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรุณากรอกข้อมูล Monitor"))
}

func Test_Monitor_not_specialCharacter(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "#@",
		GPU:        "test",
		RAM:        "test",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรอกข้อมูล Monitor ไม่ถูกต้อง"))
}

func Test_GPU_not_blank(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "test",
		GPU:        "",
		RAM:        "test",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรุณากรอกข้อมูล GPU"))
}

func Test_GPU_not_specialCharacter(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "test",
		GPU:        "@#@#",
		RAM:        "test",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรอกข้อมูล GPU ไม่ถูกต้อง"))
}

func Test_RAM_not_blank(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "test",
		GPU:        "test",
		RAM:        "",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรุณากรอกข้อมูล RAM"))
}

func Test_RAM_not_specialCharacter(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "test",
		GPU:        "test",
		RAM:        "@@",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรอกข้อมูล RAM ไม่ถูกต้อง"))
}

func Test_Harddisk_not_blank(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "test",
		GPU:        "test",
		RAM:        "test",
		Harddisk:   "",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรุณากรอกข้อมูล Harddisk"))
}

func Test_Harddisk_not_specialCharacter(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "test",
		GPU:        "test",
		RAM:        "test",
		Harddisk:   "****",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรอกข้อมูล Harddisk ไม่ถูกต้อง"))
}

func Test_Datetime_not_blank(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "test",
		GPU:        "test",
		RAM:        "test",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Time{},
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรุณาใส่วันที่ และ เวลา"))
}

func Test_Datetime_notPast(t *testing.T) {
	g := NewGomegaWithT(t)

	Device := Device{
		CPU:        "test",
		Monitor:    "test",
		GPU:        "test",
		RAM:        "test",
		Harddisk:   "test",
		Problem:    "test",
		Customer:   Customer{},
		CustomerID: new(uint),
		TypeID:     new(uint),
		Type:       Type{},
		WindowsID:  new(uint),
		Windows:    Windows{},
		Save_Time:  time.Now().Add(1 * time.Second),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(Device)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("วันที่ และ เวลา ไม่ถูกต้อง"))
}
