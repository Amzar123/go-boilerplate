package app

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	FiberApp        *fiber.App
	GormDb          *gorm.DB
	ViperConfig     *viper.Viper
	OtelTracer      *sdktrace.TracerProvider
	Zap             *zap.Logger
	WatermillRouter *message.Router
	RedisClient     *redis.Client
}

func NewApp(ctx context.Context) (*App, error) {
	// Initialize Viper Config
	viperConfig := viper.New()
	viperConfig.SetConfigName("config") // Use your config file name
	viperConfig.AddConfigPath(".")
	if err := viperConfig.ReadInConfig(); err != nil {
		return nil, errors.New("error reading config file: " + err.Error())
	}

	// Initialize Zap Logger
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, errors.New("failed to initialize logger: " + err.Error())
	}

	// Initialize Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     viperConfig.GetString("redis.address"), // e.g., "localhost:6379"
		Password: "",                                     // No password by default
		DB:       0,                                      // Default DB
	})

	// Check Redis Connection with Context
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		return nil, errors.New("failed to connect to Redis: " + err.Error())
	}

	// Initialize Fiber App
	fiberApp := fiber.New()

	// Return Initialized App Struct
	return &App{
		FiberApp:        fiberApp,
		GormDb:          nil, // Initialize DB connection separately
		ViperConfig:     viperConfig,
		OtelTracer:      nil, // Initialize OpenTelemetry separately
		Zap:             logger,
		WatermillRouter: nil, // Initialize Watermill Router separately
		RedisClient:     redisClient,
	}, nil
}

func (a *App) Serve(ctx context.Context) error {
	// Run Watermill router on separate goroutine (todo: check if this is necessary)
	go func() {
		if err := a.WatermillRouter.Run(ctx); err != nil {
			a.Zap.Fatal("failed to run watermill router", zap.Error(err))
		}
	}()

	return a.FiberApp.Listen(a.ViperConfig.GetString("address"))
}

func (a *App) GetFiber() *fiber.App {
	return a.FiberApp
}

func (a *App) GetGORM() *gorm.DB {
	return a.GormDb
}

func (a *App) GetRedis() *redis.Client {
	return a.RedisClient
}

func (a *App) stopWatermillRouter() error {
	if !a.WatermillRouter.IsRunning() {
		return nil
	}

	return a.WatermillRouter.Close()
}

func (a *App) Clean(ctx context.Context) error {
	err := a.FiberApp.Shutdown()
	if err != nil {
		return err
	}

	if err = a.stopWatermillRouter(); err != nil {
		return err
	}

	sqlDB, err := a.GormDb.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	err = a.OtelTracer.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetGorm() *gorm.DB {
	return a.GormDb
}
