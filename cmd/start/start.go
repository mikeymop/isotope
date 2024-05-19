package start

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/isotope/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

type Server struct {
	// Config     *Config
	// DB         *database.DB
	// KeyStorage crypto.KeyStorage
	// Keys       *encryption.EncryptionKeys
	// Eventstore *eventstore.Eventstore
	// Queries    *query.Queries
	// AuthzRepo  authz_repo.Repository
	// Storage    static.Storage
	// Commands   *command.Commands
	// Router     *mux.Router
	// TLSConfig  *tls.Config
	Shutdown chan<- os.Signal
}

func New(server chan<- *Server) *cobra.Command {
	start := &cobra.Command{
		Use:   "start",
		Short: "starts isotope",
		Long:  "starts isotope instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			// err := cmd_tls.ModeFromFlag(cmd)
			// if err != nil {
			// 	return err
			// }
			// config := MustNewConfig(viper.GetViper())
			// masterKey, err := key.MasterKey(cmd)
			// if err != nil {
			// 	return err
			// }
			return startIsotope(cmd.Context(), server)
		},
	}

	startFlags(start)

	return start
}

func startIsotope(ctx context.Context, server chan<- *Server) error {
	var port uint16
	port, err := startFlagSet.GetUint16("port")
	if err != nil {
		port = 8080
	}

	shutdown := make(chan os.Signal, 1)
	signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)

	if server != nil {
		server <- &Server{
			Shutdown: shutdown,
		}
		close(server)
	}

	return listen(ctx, port, shutdown)
}

func listen(ctx context.Context, port uint16, shutdown <-chan os.Signal) error {
	e := echo.New()
	e.HideBanner = true
	errCh := make(chan error)

	// e.Any("/api", func(c echo.Context) error {
	// 	if c.Request().Method == "GET" {
	// 		return c.String(http.StatusOK, "Hello GET request")
	// 	}

	// 	if c.Request().Method == "POST" {
	// 		return c.String(http.StatusOK, "Hello POST request")
	// 	}
	// 	return c.String(http.StatusOK, "Hello World")
	// })
	// uiFS, err := fs.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: web.BuildHTTPFS(),
		// Root:       "web",
		HTML5: true,
	}))

	api := e.Group("/api")
	api.Any("", func(c echo.Context) error {
		if c.Request().Method == "GET" {
			return c.String(http.StatusOK, "Hello GET request")
		}

		if c.Request().Method == "POST" {
			return c.String(http.StatusOK, "Hello POST request")
		}
		return c.String(http.StatusOK, "Hello World")
	})

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("error starting server: %w", err)
	case <-shutdown:
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		return shutdownServer(ctx, e)
	case <-ctx.Done():
		return shutdownServer(ctx, e)
	}
}

func shutdownServer(ctx context.Context, server *echo.Echo) error {
	err := server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("could not shutdown gracefully: %w", err)
	}
	fmt.Print("Server shutdown gracefully")
	return nil
}
