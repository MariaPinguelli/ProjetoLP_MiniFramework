package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"path"
)

func main() {
	var result map[string]interface{}
	
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao ler o path do programa em execução:", err)
		return
	}

	config_filepath := path.Join(cwd, "config.json")
	jsonFile, err := os.Open(config_filepath)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo json: %s", err)
		return
	}
	
	byteValue, _ := ioutil.ReadAll(jsonFile)

    json.Unmarshal([]byte(byteValue), &result)

	programName := result["filepath"].(string)

	cmd := exec.Command("go", "run", programName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Erro ao executar o programa %s: %v\n", programName, err)
	}
}
