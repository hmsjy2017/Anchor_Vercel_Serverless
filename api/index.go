package handler

import (
	"github.com/AH-dark/Anchor/bootstrap"
	"github.com/AH-dark/Anchor/routers"
	"github.com/AH-dark/Anchor/pkg/utils"
	"net/http"
	"os"
)

func init() {
	// Set up a dummy config.yaml path for Vercel environment
	// In a real scenario, you might want to load config from environment variables or a remote source
	_ = os.Setenv("ANCHOR_CONFIG_PATH", "/tmp/config.yaml")
	bootstrap.Init("/tmp/config.yaml") // Initialize with a dummy path, actual config should be handled via env vars
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router := routers.InitRouter()
	router.ServeHTTP(w, r)
}

func main() {
	// This main function is not executed in Vercel Serverless environment
	// The Handler function is the entry point
	utils.Log().Info("Vercel Serverless function started.")
}


