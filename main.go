package main

import (
	"fmt"
	"github.com/devfile/library/pkg/devfile/generator"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"log"
	"os"
)

func main() {

	fmt.Println(generator.IngressSpecParams{})
	fmt.Println(common.DevfileOptions{})

	settings := cli.New()

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), "default", os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalf("Ошибка инициализации Helm: %v", err)
	}

	chartPath := "./my-chart"

	_, err := loader.Load(chartPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки чарта: %v", err)
	}

	values := map[string]interface{}{
		"replicaCount": 2,
		"image": map[string]interface{}{
			"repository": "nginx",
			"tag":        "1.21",
		},
	}

	client := action.NewInstall(actionConfig)
	client.ReleaseName = "my-release"
	client.Namespace = "default"

	fmt.Printf("Релиз установлен: %s\n", client.ReleaseName)
	fmt.Printf("Namespace: %s\n", client.Namespace)
	fmt.Printf("Values: %s\n", values)
}
