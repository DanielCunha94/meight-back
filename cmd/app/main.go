package main

import (
	"context"
	"errors"
	"github.com/DanielCunha94/Meight-backend/internal/app/api"
	"github.com/DanielCunha94/Meight-backend/internal/app/repository"
	"github.com/DanielCunha94/Meight-backend/internal/app/service"
	appErrors "github.com/DanielCunha94/Meight-backend/pkg/errors"
	"github.com/DanielCunha94/Meight-backend/pkg/routes"
	"github.com/DanielCunha94/Meight-backend/pkg/sse"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server *http.Server
}

func main() {
	// TODO split in small functions
	app := NewApp()
	db, err := NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	repo, err := repository.NewSQLRepository(db)
	if err != nil {
		log.Fatalln(err)
	}

	routesFinderService := routes.NewRouteMock()
	sseServer := sse.NewServer()

	s := service.NewService(repo, routesFinderService, sseServer)
	appAPI := api.NewAPI(s, sseServer)

	r := gin.Default()
	r.Use(appErrors.ErrorHandlingMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))
	appAPI.InitRouter(r.Group("/"))
	app.server.Handler = r

	go func() {
		if err := app.server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Fatal("Server successfully shutdown...")
			}
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}

func NewApp() *App {
	server := &http.Server{
		Addr: ":8080",
	}
	return &App{server: server}
}

func NewDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
