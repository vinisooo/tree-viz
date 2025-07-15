package dir

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func WalkPath(path string, outputPath string) {

	currentDir := "root"
	outputContent := "digraph G {\n"
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		deepCount := strings.Count(p, "/")

		spaces := ""
		for i := 0; i <= deepCount; i++ {
			spaces += "─"
		}

		name := strings.Split(p, "/")[deepCount]

		if d.IsDir() {
			fmt.Printf("├%s📂 %s\n", spaces, name)
			outputContent += fmt.Sprintf("\t\"%s\" -> \"%s\" \n", currentDir, name)
			currentDir = name
		} else {
			fmt.Printf("├%s📄 %s\n", spaces, name)
			outputContent += fmt.Sprintf("\t\"%s\" -> \"%s\" \n", currentDir, name)
		}

		return nil
	})

	if err != nil {
		fmt.Errorf("unexpected error on walking path %s \nDetails:%s", path, err.Error())
	}
	outputContent += "}"
	err, dotUrl := writeDotFile(outputContent)

	if err != nil {
		fmt.Errorf("unexpected error on writing dot file %s \nDetails:%s", path, err.Error())
	}

	createImage(dotUrl, outputPath)
}

func writeDotFile(outputContent string) (error, string) {
	fileName := fmt.Sprintf("graph-%s.dot", time.Now().Format("2006-01-0215:04:05"))
	fullPath := filepath.Join("out/", fileName)

	content := []byte(outputContent)
	err := os.WriteFile(fullPath, content, 0644)
	if err != nil {
		return err, ""
	}

	return nil, fullPath
}

func createImage(dotPath string, imageOutputPath string) {
	os.MkdirAll(imageOutputPath, 0755)

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(imageOutputPath, fmt.Sprintf("graph-%s.png", timestamp))

	cmd := exec.Command("dot", "-Tpng", dotPath, "-o", fullPath)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erro ao gerar imagem: %v\n", err)
	} else {
		fmt.Printf("Imagem criada em: %s\n", fullPath)
	}
}
