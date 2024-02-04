package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/analopesdev/goapi/internal/database"
	"github.com/analopesdev/goapi/internal/service"
	"github.com/analopesdev/goapi/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)
 
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)
	
	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	WebCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	WebProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()

  c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category/{id}", WebCategoryHandler.GetCategory)
	c.Get("/category", WebCategoryHandler.GetCategories)
	c.Post("/category", WebCategoryHandler.CreateCategory)

	c.Get("/product/{id}", WebProductHandler.GetProduct)
	c.Get("/product", WebProductHandler.GetProducts)
	c.Get("/product/category/{categoryId}", WebProductHandler.GetProductByCategoryId)
	c.Post("/product", WebProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
} 