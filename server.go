package flightlesssomething

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/ravener/discord-oauth2"
	openai "github.com/sashabaranov/go-openai"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

var (
	// GORM database object
	db *gorm.DB

	// Discord conf object
	discordConf *oauth2.Config

	// Benchmarks directory
	benchmarksDir string

	// OpenAI
	openaiClient *openai.Client
	openaiModel  string
)

func Start(c *Config, version string) {
	// Setup OpenAI client //

	if c.OpenAIApiKey != "" {
		openaiClientConf := openai.DefaultConfig(c.OpenAIApiKey)
		openaiClientConf.BaseURL = c.OpenAIURL
		openaiClient = openai.NewClientWithConfig(openaiClientConf)
		openaiModel = c.OpenAIModel
	}

	// Setup data dir //

	_, err := os.Stat(c.DataDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(c.DataDir, 0755)
		if err != nil {
			panic("Failed to create data dir: " + err.Error())
		}
	} else if err != nil {
		panic("Failed to check data dir: " + err.Error())
	}

	benchmarksDir = filepath.Join(c.DataDir, "benchmarks")
	_, err = os.Stat(benchmarksDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(benchmarksDir, 0755)
		if err != nil {
			panic("Failed to create benchmarks dir: " + err.Error())
		}
	} else if err != nil {
		panic("Failed to check benchmarks dir: " + err.Error())
	}

	// Setup Discord OAuth2 //

	discordConf = &oauth2.Config{
		Endpoint:     discord.Endpoint,
		Scopes:       []string{discord.ScopeIdentify},
		RedirectURL:  c.DiscordRedirectURL,
		ClientID:     c.DiscordClientID,
		ClientSecret: c.DiscordClientSecret,
	}

	// Setup gorm (database) //

	db, err = gorm.Open(sqlite.Open(filepath.Join(c.DataDir, "database.db")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	store := gormsessions.NewStore(db, true, []byte(c.SessionSecret))
	db.AutoMigrate(&User{}, &Benchmark{})

	// Setup gin //

	if version == "" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))

	// Create a new FuncMap and add the version function
	funcMap := template.FuncMap{
		"version": func() string {
			if version == "" {
				return "dev"
			}
			return version
		},
	}

	// Create a new template, apply the function map, and parse the templates
	tmpl := template.New("").Funcs(funcMap)
	tmpl, err = tmpl.ParseFS(templatesFS, "templates/*.tmpl")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	// Set the HTML template for Gin
	r.SetHTMLTemplate(tmpl)

	// Serve static files
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatalf("Failed to create static file system: %v", err)
	}
	r.StaticFS("/static/", http.FS(staticFS))

	r.GET("/", func(c *gin.Context) { c.Redirect(http.StatusTemporaryRedirect, "/benchmarks") })

	r.GET("/benchmarks", getBenchmarks)

	r.GET("/benchmark", getBenchmarkCreate)
	r.POST("/benchmark", postBenchmarkCreate)
	r.GET("/benchmark/:id", getBenchmark)
	r.GET("/benchmark/:id/edit", getBenchmarkEdit)
	r.POST("/benchmark/:id/edit", postBenchmarkEdit)
	r.DELETE("/benchmark/:id", deleteBenchmark)
	r.GET("/benchmark/:id/download", getBenchmarkDownload)

	r.GET("/user/:id", getUser)

	r.GET("/login", getLogin)
	r.GET("/login/callback", getLoginCallback)
	r.GET("/logout", getLogout)

	r.Run(c.Bind)
}
