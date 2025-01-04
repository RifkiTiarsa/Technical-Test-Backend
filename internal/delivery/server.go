package delivery

import (
	"fmt"
	"technical-test/internal/config"
	"technical-test/internal/delivery/handler"
	"technical-test/internal/middleware"
	"technical-test/internal/repository"
	"technical-test/internal/shared/service"
	"technical-test/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	authUC     usecase.AuthUsecase
	productUC  usecase.ProductUsecase
	cartUC     usecase.CartUsecase
	jwtService service.JwtService
	engine     *gin.Engine
	host       string
}

func (s Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	handler.NewAuthHandler(s.authUC, rg).Route()
	handler.NewProductHandler(s.productUC, rg).Route()
	handler.NewCartHandler(s.cartUC, rg, authMiddleware).Route()
}

func (s Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	jwtService := service.NewJwtService(cfg.TokenConfig)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartRepo := repository.NewCartRepository(db)
	userUc := usecase.NewUserUsecase(userRepo)
	authUc := usecase.NewAuthUsecase(userUc, jwtService)
	productUc := usecase.NewProductUsecase(productRepo)
	cartUc := usecase.NewCartUsecase(cartRepo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		authUC:     authUc,
		productUC:  productUc,
		cartUC:     cartUc,
		jwtService: jwtService,
		engine:     engine,
		host:       host,
	}
}
