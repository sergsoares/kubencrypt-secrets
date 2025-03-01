package main

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"filippo.io/age"
	flags "github.com/jessevdk/go-flags"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Options struct {
	Password   string `short:"p" long:"password" env:"PASSWORD"`
	Kubeconfig string `long:"kubeconfig" env:"KUBECONFIG"`
}

func parseOptions() Options {
	var options Options
	_, err := flags.Parse(&options)
	if err != nil {
		panic(err)
	}

	if options.Kubeconfig == "" {
		home := homedir.HomeDir()
		options.Kubeconfig = filepath.Join(home, clientcmd.RecommendedHomeDir, clientcmd.RecommendedFileName)
		//Log: debug kubeconfig path used
	}

	if !filepath.IsAbs(options.Kubeconfig) {
		panic("Kubeconfig isn't a absolute path")
	}

	if len(options.Password) == 0 {
		panic("--password can't be empty.")
	}

	return options
}

func main() {
	options := parseOptions()

	config, _ := clientcmd.BuildConfigFromFlags("", options.Kubeconfig)
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	secrets, err := client.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d secrets to be saved \n", len(secrets.Items))

	currentTime := time.Now().Format("20060102150405")
	secretsFolder := "secrets-" + currentTime

	err = os.Mkdir(secretsFolder, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	archive, err := os.Create(secretsFolder + ".zip")
	if err != nil {
		panic(err.Error())
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	for _, v := range secrets.Items {
		res, _ := json.Marshal(v)

		var out bytes.Buffer
		err := json.Indent(&out, res, "", "\t")

		if err != nil {
			panic(err.Error())
		}

		filePath := secretsFolder + "/" + v.ObjectMeta.Namespace + "-" + v.ObjectMeta.Name + ".json"
		err = os.WriteFile(filePath, out.Bytes(), 0644)
		if err != nil {
			panic(err.Error())
		}

		zipFile, err := zipWriter.Create(filePath)
		if err != nil {
			panic(err.Error())
		}

		_, err = zipFile.Write(out.Bytes())
		if err != nil {
			panic(err.Error())
		}
	}

	err = zipWriter.Close()
	if err != nil {
		panic(err.Error())
	}

	recipient, err := age.NewScryptRecipient(options.Password)
	if err != nil {
		panic(err.Error())
	}

	recipient.SetWorkFactor(2)

	outputFile, _ := os.Create(secretsFolder + ".zip.age")
	ageWriter, err := age.Encrypt(outputFile, recipient)
	if err != nil {
		panic(err.Error())
	}
	defer ageWriter.Close()

	inputFile, err := os.Open(secretsFolder + ".zip")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	_, err = io.Copy(ageWriter, inputFile)
	if err != nil {
		panic(err)
	}
}
