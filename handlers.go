package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Surdy-A/SMarket/models"
	"github.com/Surdy-A/SMarket/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	vars := mux.Vars(r)
	product_category_id := vars["category_id"]

	err := p.Categories.GetCategory(a.DB, product_category_id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "product category doesn't exist")
		return
	}

	var v models.Vendor
	vendor_id := vars["vendor_id"]

	err = v.GetVendor(a.DB, vendor_id)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "venor doesn't exist")
		return
	}

	if err := p.CreateProduct(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, p)
}

func (a *App) GetProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	var p models.Product
	products, err := p.GetProducts(a.DB, start, count)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, products)
}

func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	p := models.Product{ID: id}
	if err := p.GetProduct(a.DB, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, p)
}

func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	product_category_id := vars["id"]

	var p models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id
	product_category_id = vars["category_id"]

	err := p.Categories.GetCategory(a.DB, product_category_id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "product category doesn't exist")
		return
	}

	var product models.Product
	if err := product.GetProduct(a.DB, id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.UpdateProduct(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, p)
}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	p := models.Product{ID: id}
	if err := p.GetProduct(a.DB, id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.DeleteProduct(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// Blog Handlers
func (a *App) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var b models.Article
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := b.CreateArticle(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, b)
}

func (a *App) GetArticles(w http.ResponseWriter, r *http.Request) {
	var ar models.Article
	articles, err := ar.GetArticles(a.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, articles)
}

func (a *App) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Article ID")
		return
	}

	p := models.Article{ID: id}
	if err := p.GetArticle(a.DB, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Article not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, p)
}

func (a *App) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Article ID")
		return
	}

	var ar models.Article
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ar); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	ar.ID = id

	var article models.Article
	if err := article.GetArticle(a.DB, id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ar.UpdateArticle(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, ar)
}

func (a *App) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Article ID")
		return
	}

	ar := models.Article{ID: id}
	if err := ar.GetArticle(a.DB, id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ar.DeleteArticle(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// Vendor Handlers
func (a *App) CreateVendor(w http.ResponseWriter, r *http.Request) {
	var v models.Vendor
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := v.CreateVendor(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, v)
}

func (a *App) GetVendors(w http.ResponseWriter, r *http.Request) {
	var v models.Vendor
	vendors, err := v.GetVendors(a.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, vendors)
}

func (a *App) GetVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	v := models.Vendor{ID: id}
	if err := v.GetVendor(a.DB, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Vendor not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, v)
}

func (a *App) UpdateVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var v models.Vendor
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	v.ID = id

	var vendor models.Vendor
	if err := vendor.GetVendor(a.DB, id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := v.UpdateVendor(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, v)
}

func (a *App) DeleteVendor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	v := models.Vendor{ID: id}
	if err := v.GetVendor(a.DB, id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := v.DeleteVendor(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// Category Handlers
// Vendor Handlers
func (a *App) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.CreateCategory(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, c)
}

func (a *App) GetCategories(w http.ResponseWriter, r *http.Request) {
	var c models.Category
	vendors, err := c.GetCategories(a.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, vendors)
}

func (a *App) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	c := models.Category{ID: id}
	if err := c.GetCategory(a.DB, id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product Catgeory not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, c)
}

func (a *App) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var c models.Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	c.ID = id

	var category models.Category
	if err := category.GetCategory(a.DB, id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.UpdateCategory(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, c)
}

func (a *App) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	c := models.Category{ID: id}
	if err := c.GetCategory(a.DB, id); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.DeleteCategory(a.DB); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
