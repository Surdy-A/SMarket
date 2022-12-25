package main

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/product", a.CreateProduct).Methods("POST")
	a.Router.HandleFunc("/products", a.GetProducts).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.GetProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.UpdateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.DeleteProduct).Methods("DELETE")

	//Blog Routes
	a.Router.HandleFunc("/article", a.CreateArticle).Methods("POST")
	a.Router.HandleFunc("/articles", a.GetArticles).Methods("GET")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.GetArticle).Methods("GET")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.UpdateArticle).Methods("PUT")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.DeleteArticle).Methods("DELETE")

	//Vendor Routes
	a.Router.HandleFunc("/vendor", a.CreateVendor).Methods("POST")
	a.Router.HandleFunc("/vendors", a.GetVendors).Methods("GET")
	a.Router.HandleFunc("/vendor/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", a.GetVendor).Methods("GET")
	a.Router.HandleFunc("/vendor/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", a.UpdateVendor).Methods("PUT")
	a.Router.HandleFunc("/vendor/{id:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}}", a.DeleteVendor).Methods("DELETE")

	//Category Routes
	a.Router.HandleFunc("/category", a.CreateCategory).Methods("POST")
	a.Router.HandleFunc("/categories", a.GetCategories).Methods("GET")
	a.Router.HandleFunc("/category/{id:[0-9]+}", a.GetCategory).Methods("GET")
	a.Router.HandleFunc("/category/{id:[0-9]+}", a.UpdateCategory).Methods("PUT")
	a.Router.HandleFunc("/category/{id:[0-9]+}", a.DeleteCategory).Methods("DELETE")
}
