package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"bytes"
	"fmt"
)
// Defining Variables
var AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_DEFAULT_REGION, AWS_EMAIL_ADDRESS string

func preflight_check() {
    if AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID"); len(AWS_ACCESS_KEY_ID) == 0 {
        log.Fatal("AWS_ACCESS_KEY_ID NOT FOUND, please add your AWS_ACCESS_KEY_ID in ENV")
    }
    if AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY"); len(AWS_SECRET_ACCESS_KEY) == 0 {
        log.Fatal("AWS_SECRET_ACCESS_KEY NOT FOUND, please add your AWS_ACCESS_KEY_ID in ENV")
    }
    if AWS_DEFAULT_REGION = os.Getenv("AWS_DEFAULT_REGION"); len(AWS_DEFAULT_REGION) == 0 {
        AWS_DEFAULT_REGION = "ap-southeast-1"
    }
    if AWS_EMAIL_ADDRESS = os.Getenv("AWS_EMAIL_ADDRESS"); len(AWS_EMAIL_ADDRESS) == 0 {
        AWS_EMAIL_ADDRESS = "yourname@youremail.com"
    }
    if err := exec.Command("docker").Run(); err != nil {
		log.Fatal("Docker command not found or not access to docker")
	}
}

func main() {
    preflight_check()

    cmdReturn, err := exec.Command("aws", "ecr", "get-login", "--no-include-email","--region", AWS_DEFAULT_REGION).Output()
    if err != nil {
        log.Printf("Error from aws get-login")
        log.Fatal(err)
    }
    // log.Printf(string(cmdReturn))
    s := strings.Split(strings.TrimSuffix(string(cmdReturn), "\n"), " ")
    aws, token, url := s[3], s[5], s[6]

    cmd := exec.Command("docker", "login", "-u", aws, "-p", token, url)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err = cmd.Run()
    if err != nil {
        log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
    }
    // kubectl delete secret regsecret
    // kubectl create secret docker-registry regsecret --docker-server= --docker-username= --docker-password= --docker-email=vinnson.lee@grabtaxi.com
    exec.Command("kubectl", "delete", "secret", "regsecret").Run()
    cmdReturn, err = exec.Command("kubectl", "create", "secret", "docker-registry", "regsecret", "--docker-server="+ url, "--docker-username="+ aws, "--docker-password="+ token, "--docker-email="+ AWS_EMAIL_ADDRESS).Output()
    log.Printf(string(cmdReturn))
}
