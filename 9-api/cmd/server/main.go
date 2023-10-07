package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valoirmoura/goexpert/curso-go-api/config"
	_ "github.com/valoirmoura/goexpert/curso-go-api/docs"
	"github.com/valoirmoura/goexpert/curso-go-api/internal/entity"
	"github.com/valoirmoura/goexpert/curso-go-api/internal/infra/database"
	"github.com/valoirmoura/goexpert/curso-go-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title           VMCOD3R API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API VMCod3r GO
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  VMCod3r
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, cfg.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer) //Não deixa a aplicação morrer, mesmo se um handler tiver um panic
	r.Use(middleware.Logger)
	//r.Use(LogRequest) //Apenas para exemplicar os Middlewares
	r.Route("/products", func(r chi.Router) {
		//Fazendo Autenticação nas Rotas....
		r.Use(jwtauth.Verifier(cfg.TokenAuthKey))
		r.Use(jwtauth.Authenticator)
		r.Use(middleware.WithValue("jwt", cfg.TokenAuthKey))

		//Rotas
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_token", userHandler.GetJWT)

	log.Println("Listening Server Port: " + cfg.DBPort)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}

// Exemplificando um Middleware
//func LogRequest(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("Request: %s %s", r.Method, r.URL.Path)
//		next.ServeHTTP(w, r)
//	})
//}

//swag init -g cmd/server/main.go (COMANDO PARA INICIAR A DOCUMENTAÇÃO)
