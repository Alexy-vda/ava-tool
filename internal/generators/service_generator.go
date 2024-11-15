package generators

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// Embedding all template files
//
//go:embed templates/main.go.tmpl
var mainTemplate string

//go:embed templates/main-db.go.tmpl
var mainDBTemplate string

//go:embed templates/global.go.tmpl
var globalTemplate string

//go:embed templates/health.go.tmpl
var healthTemplate string

//go:embed templates/origin.go.tmpl
var originTemplate string

//go:embed templates/router.go.tmpl
var routerHelperTemplate string

//go:embed templates/request.go.tmpl
var requestTemplate string

//go:embed templates/middleware.go.tmpl
var middlewareTemplate string

//go:embed templates/server.go.tmpl
var serverTemplate string

//go:embed templates/init.go.tmpl
var initTemplate string

//go:embed templates/migration.go.tmpl
var migrationTemplate string

//go:embed templates/base.go.tmpl
var baseTemplate string

//go:embed templates/config.go.tmpl
var configTemplate string

//go:embed templates/.env.tmpl
var envTemplate string

//go:embed templates/.gitignore.tmpl
var gitignoreTemplate string

//go:embed templates/Dockerfile.tmpl
var dockerfileTemplate string

//go:embed templates/docker-compose.yaml.tmpl
var dockerComposeTemplate string

//go:embed templates/docker-compose-db.yaml.tmpl
var dockerComposeDBTemplate string

// GenerateService crée la structure des fichiers pour un nouveau service
func GenerateService(name string, includeDB bool, port string) {
	baseDir := name
	directories := []string{
		"cmd/server", "internal/configs", "internal/api/handlers",
		"internal/api/routers", "internal/server", "internal/services", "internal/helpers",
	}
	if includeDB {
		directories = append(directories, "internal/db/", "internal/models")
	}
	for _, dir := range directories {
		os.MkdirAll(filepath.Join(baseDir, dir), os.ModePerm)
	}

	// Définition des fichiers à générer et des templates associés
	files := map[string]string{
		filepath.Join(baseDir, "cmd/server/main.go"):              selectTemplate(includeDB, mainTemplate, mainDBTemplate),
		filepath.Join(baseDir, "internal/api/routers/global.go"):  globalTemplate,
		filepath.Join(baseDir, "internal/api/handlers/health.go"): healthTemplate,
		filepath.Join(baseDir, "internal/configs/config.go"):       configTemplate,
		filepath.Join(baseDir, "internal/helpers/router.go"):      routerHelperTemplate,
		filepath.Join(baseDir, "internal/helpers/request.go"):     requestTemplate,
		filepath.Join(baseDir, "internal/helpers/origin.go"):      originTemplate,
		filepath.Join(baseDir, "internal/server/middleware.go"):   middlewareTemplate,
		filepath.Join(baseDir, "internal/server/server.go"):       serverTemplate,
		filepath.Join(baseDir, ".env"):                            envTemplate,
		filepath.Join(baseDir, ".gitignore"):                      gitignoreTemplate,
		filepath.Join(baseDir, "Dockerfile"):                      dockerfileTemplate,
		filepath.Join(baseDir, "docker-compose.yaml"):             selectTemplate(includeDB, dockerComposeTemplate, dockerComposeDBTemplate),
	}

	// Ajout des fichiers spécifiques à la DB si nécessaire
	if includeDB {
		files[filepath.Join(baseDir, "internal/db/migration.go")] = migrationTemplate
		files[filepath.Join(baseDir, "internal/db/init.go")] = initTemplate
		files[filepath.Join(baseDir, "internal/models/base.go")] = baseTemplate
	}

	// Génération des fichiers
	for path, tmplContent := range files {
		generateFileFromTemplate(path, tmplContent, map[string]string{
			"ServiceName": name,
			"Port":        port,
		})
	}
	fmt.Printf("Service %s created successfully\n", name)
}

// selectTemplate sélectionne le template selon la présence ou non de la DB
func selectTemplate(includeDB bool, templateNoDB, templateWithDB string) string {
	if includeDB {
		return templateWithDB
	}
	return templateNoDB
}

// generateFileFromTemplate génère un fichier à partir d'un contenu de template et d'une map de données
func generateFileFromTemplate(outputPath, tmplContent string, data map[string]string) {
	tmpl, err := template.New("template").Parse(tmplContent)
	if err != nil {
		fmt.Printf("Failed to parse template: %v\n", err)
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", outputPath, err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Failed to execute template for %s: %v\n", outputPath, err)
	}
}
