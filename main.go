package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"strconv"


	h "microservice-shortener/api"
	mr "microservice-shortener/repository/mongo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"microservice-shortener/shortener"
)

func main()  {
	repo := chooseRepo()
	//repo:= ChooseRepo()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)
	r :=chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)


	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)

	go func() {
	fmt.Println("Listening on port :8000")
	errs <- http.ListenAndServe(httpPort(), r)
	}()



	go func() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)


	

}

func httpPort() string  {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}




func chooseRepo() shortener.RedirectRepository {
	if os.Getenv("URL_DB") == "mongo"{
		mongoURL:= os.Getenv("MONGO_URL")
		mongodb:=os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, errr := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
	

	if errr != nil {
		log.Fatal(errr)
	}
	return repo
}
return nil
}




