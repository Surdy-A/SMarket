package main

func (a *App) InitializeRoutes() {
	router := a.Router.PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("/product/{category_id}/{vendor_id}", a.CreateProduct).Methods("POST")
	router.HandleFunc("/products", a.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", a.GetProduct).Methods("GET")
	router.HandleFunc("/product/{id}/{category_id}", a.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", a.DeleteProduct).Methods("DELETE")

	//Blog Routes
	router.HandleFunc("/article", a.CreateArticle).Methods("POST")
	router.HandleFunc("/articles", a.GetArticles).Methods("GET")
	router.HandleFunc("/article/{id:[0-9]+}", a.GetArticle).Methods("GET")
	router.HandleFunc("/article/{id:[0-9]+}", a.UpdateArticle).Methods("PUT")
	router.HandleFunc("/article/{id:[0-9]+}", a.DeleteArticle).Methods("DELETE")

	//Vendor Routes
	router.HandleFunc("/vendor", a.CreateVendor).Methods("POST")
	router.HandleFunc("/vendors", a.GetVendors).Methods("GET")
	router.HandleFunc("/vendor/{id}", a.GetVendor).Methods("GET")
	router.HandleFunc("/vendor/{id}", a.UpdateVendor).Methods("PUT")
	router.HandleFunc("/vendor/{id}", a.DeleteVendor).Methods("DELETE")

	//Product Category Routes
	router.HandleFunc("/category", a.CreateCategory).Methods("POST")
	router.HandleFunc("/categories", a.GetCategories).Methods("GET")
	router.HandleFunc("/category/{id}", a.GetCategory).Methods("GET")
	router.HandleFunc("/category/{id}", a.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", a.DeleteCategory).Methods("DELETE")

	//Article Category Routes
	router.HandleFunc("article/category", a.CreateArticleCategory).Methods("POST")
	router.HandleFunc("article/categories", a.GetArticleCategories).Methods("GET")
	router.HandleFunc("article/category/{id}", a.GetArticleCategory).Methods("GET")
	router.HandleFunc("article/category/{id}", a.UpdateArticleCategory).Methods("PUT")
	router.HandleFunc("article/category/{id}", a.DeleteArticleCategory).Methods("DELETE")
}
