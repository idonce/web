package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type page struct {
	File     string           // output filename
	Active   string           // sidebar active id
	Title    string           // <title>
	CSS      string           // extra <style>
	Body     string           // page content
	JS       string           // extra <script>
	Sections []sidebarLink    // sidebar section anchors
}

var pages = []page{
	{"index.html", "home", "idonce — Prove you're human. Nothing else.", landingCSS, landingBody, landingJS, landingSections},
	{"developers.html", "developers", "idonce — Developers", developersCSS, developersBody, developersJS, developerSections},
	{"privacy.html", "privacy", "idonce — Privacy Policy", legalCSS, privacyBody, "", nil},
	{"terms.html", "terms", "idonce — Terms of Service", legalCSS, termsBody, "", nil},
	{"legal-notice.html", "legal-notice", "idonce — Legal Notice", legalCSS, legalNoticeBody, "", nil},
}

func renderPage(p page) string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>` + p.Title + `</title>
` + layoutCSS + p.CSS + `
</head>
<body>
<div class="shell">
` + sidebarHTML(p.Active, p.Sections) + `
<main class="content">
` + p.Body + `
</main>
</div>
` + p.JS + scrollSpyJS + `
</body>
</html>`
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--build" {
		build()
		return
	}
	serve()
}

// --build: generate static HTML files to dist/
func build() {
	dist := "dist"
	os.MkdirAll(dist, 0755)
	os.MkdirAll(filepath.Join(dist, "fonts"), 0755)

	// Render pages
	for _, p := range pages {
		html := renderPage(p)
		path := filepath.Join(dist, p.File)
		if err := os.WriteFile(path, []byte(html), 0644); err != nil {
			log.Fatalf("failed to write %s: %v", path, err)
		}
		fmt.Printf("  %s\n", path)
	}

	// Copy fonts
	fonts, _ := os.ReadDir("fonts")
	for _, f := range fonts {
		if f.IsDir() {
			continue
		}
		data, err := os.ReadFile(filepath.Join("fonts", f.Name()))
		if err != nil {
			continue
		}
		os.WriteFile(filepath.Join(dist, "fonts", f.Name()), data, 0644)
		fmt.Printf("  %s\n", filepath.Join(dist, "fonts", f.Name()))
	}

	// Copy styleguide
	if data, err := os.ReadFile("styleguide.html"); err == nil {
		os.WriteFile(filepath.Join(dist, "styleguide.html"), data, 0644)
		fmt.Printf("  %s\n", filepath.Join(dist, "styleguide.html"))
	}

	fmt.Printf("\nBuild complete: %s/\n", dist)
}

// Default: run dev server
func serve() {
	mux := http.NewServeMux()

	// Pages
	for _, p := range pages {
		var pattern string
		if p.File == "index.html" {
			pattern = "/"
		} else {
			pattern = "/" + p.File[:len(p.File)-5] // strip .html
		}
		html := renderPage(p)
		mux.HandleFunc("GET "+pattern, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(html))
		})
	}
	// /demo redirects to landing
	mux.HandleFunc("GET /demo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(renderPage(pages[0])))
	})

	// Static assets
	mux.HandleFunc("GET /fonts/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))).ServeHTTP(w, r)
	})

	// Health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9091"
	}

	log.Printf("idonce Web (dev) listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
