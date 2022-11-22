package productcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vanjul123/go-restapi-mux/helper"
	"github.com/vanjul123/go-restapi-mux/models"
	"gorm.io/gorm"
)

var ResponesJson = helper.ResponesJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	err := models.DB.Find(&products).Error
	if err != nil {
		ResponseError(w,http.StatusInternalServerError,err.Error())
		return
	}

	// response, _ := json.Marshal(products)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(response)

	ResponesJson(w,http.StatusOK,products)

}


func Show(w http.ResponseWriter, r *http.Request) {
	// menggunakan mux untuk response
	vars := mux.Vars(r)
	// ambil id dari vars dan convert ke int
	id , err := strconv.ParseInt(vars["id"],10,64)
	if err != nil {
		ResponseError(w,http.StatusBadRequest,err.Error())
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			ResponseError(w,http.StatusNotFound,"Produk tidak ditemukan")
			return
		default:
			ResponseError(w,http.StatusInternalServerError,err.Error())
			return
		}
	}

	ResponesJson(w,http.StatusOK,product)
}


func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	// decode json
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil{
		ResponseError(w,http.StatusInternalServerError,err.Error())
		return
	}

	defer r.Body.Close()

	// create db
	if err := models.DB.Create(&product).Error; err != nil{
		ResponseError(w,http.StatusInternalServerError,err.Error())
		return
	}

	ResponesJson(w,http.StatusCreated,product)

}


func Update(w http.ResponseWriter, r *http.Request) {
	// menggunakan mux untuk response
	vars := mux.Vars(r)
	// ambil id dari vars dan convert ke int
	id , err := strconv.ParseInt(vars["id"],10,64)
	if err != nil {
		ResponseError(w,http.StatusBadRequest,err.Error())
		return
	}

	var product models.Product

	// decode json
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil{
		ResponseError(w,http.StatusInternalServerError,err.Error())
		return
	}

	defer r.Body.Close()

	// insert data
	if models.DB.Where("id = ?", id).Updates(&product).RowsAffected == 0{
		ResponseError(w,http.StatusBadRequest, "Produk tidak ditemukan")
		return
	}

	product.Id = id

	ResponesJson(w,http.StatusOK,product)
}


func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id":""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil{
		ResponseError(w,http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var product models.Product
	if models.DB.Delete(&product,input["id"]).RowsAffected == 0{
		ResponseError(w,http.StatusBadRequest, "Tidak dapat menghapus produk")
		return
	}

	response := map[string]string{"message":"Produk berhasil dihapus"}

	ResponesJson(w,http.StatusOK,response)
}