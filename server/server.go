package server

import (
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"demo/handler"
	"demo/model"
	"demo/repository"
	"demo/service"

)

func StartServer() {
	dsn := "Database Connection String"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(&model.Admin{})

	repo := repository.NewAdminRepository(db)
	svc := service.NewAdminService(repo)
	h := handler.NewAdminHandler(svc)

	http.HandleFunc("/admin", h.CreateAdmin)
	http.HandleFunc("/admin/get", h.GetAdmin)

	http.ListenAndServe(":8080", nil)
}
