package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	// required by GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kentpon/LetsGO/cfg"
)

var DB *gorm.DB

func TestRequest(r http.Handler, path, method, token string, body interface{}) *httptest.ResponseRecorder {
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(b))
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func init() {
	cfg.Load()
	fmt.Println("Loading configs")
	args := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s ",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Name, cfg.DB.Password, cfg.DB.SSLMode)
	if cfg.DB.SSLMode != "disable" {
		args += fmt.Sprintf("sslrootcert=%s sslcert=%s sslkey=%s",
			cfg.DB.SSLRootCert, cfg.DB.SSLClientCert, cfg.DB.SSLClientKey)
	}
	db, err := gorm.Open("postgres", args)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database: %s", err))
	}
	DB = db
	fmt.Println("DB Connection Established")
}
