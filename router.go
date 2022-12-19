package main

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/products", a.GetProducts).Methods("GET")
	a.Router.HandleFunc("/product", a.CreateProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.GetProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.UpdateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", a.DeleteProduct).Methods("DELETE")

	//Blog Routes
	a.Router.HandleFunc("/article", a.CreateArticle).Methods("POST")
	a.Router.HandleFunc("/articles", a.GetArticles).Methods("GET")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.GetArticle).Methods("GET")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.UpdateArticle).Methods("PUT")
	a.Router.HandleFunc("/article/{id:[0-9]+}", a.DeleteArticle).Methods("DELETE")
}
