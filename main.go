package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pruthvij008/backendgolang/controllers"
	"github.com/Pruthvij008/backendgolang/models"
	"github.com/Pruthvij008/backendgolang/templates"
	"github.com/Pruthvij008/backendgolang/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFs(templates.EmbeddedFS, "templatepage.gohtml", "home.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFs(templates.EmbeddedFS, "templatepage.gohtml", "contact.gohtml"))))
	r.Get("/faq", controllers.Faq(views.Must(views.ParseFs(templates.EmbeddedFS, "templatepage.gohtml", "faq.gohtml"))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}
	userc := controllers.Users{
		UserService: &userService,
	}
	userc.Templates.New = views.Must(views.ParseFs(templates.EmbeddedFS, "templatepage.gohtml", "signup.gohtml"))
	userc.Templates.SignIn = views.Must(views.ParseFs(templates.EmbeddedFS, "templatepage.gohtml", "signIn.gohtml")) // Use the exported field
	r.Get("/signup", userc.New)
	r.Get("/signin", userc.SignIn)
	r.Post("/signin", userc.ProcessSignIn)
	r.Post("/users", userc.Create)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// Start server
	fmt.Println("Starting the server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
