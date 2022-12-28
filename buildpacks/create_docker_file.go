package buildpacks

import (
	"os"
	"strings"
)

func CreateDockerfile(Workdir string, Port string, language string, Name string) (err error) {

	rname := strings.ReplaceAll(Name, ".", "_")

	docker_nodejs := `
FROM node:16.3.0-alpine3.13
WORKDIR ` + Workdir + "/" + rname + `
COPY package*.json ./
COPY . .
EXPOSE ` + Port + `
RUN npm install
CMD [ "node", "index.js" ]
	`

	docker_golang := `
FROM golang:latest
WORKDIR ` + Workdir + "/" + rname + `
COPY . .
RUN go get
RUN go build -o main .
EXPOSE ` + Port + `
CMD ["./main"]
	`
	docker_python := `
FROM python:3.9.5-alpine3.13
WORKDIR ` + Workdir + `
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
EXPOSE ` + Port + `
CMD ["python", "main.py"]
	`
	docker_php := `
FROM php:8.0.3-apache
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "php", "-S", "0.0.0.0:80"]
	`
	docker_rust := `
FROM rust:1.51.0-alpine3.13
WORKDIR ` + Workdir + `
COPY . .
RUN cargo build --release
EXPOSE ` + Port + `
CMD [ "./target/release/main" ]
	`

	dockerfile := Workdir + "/" + rname + "/"

	//Create folder for dockerfile

	//Create Docker file using os force create
	file, err := os.Create(dockerfile + "Dockerfile")
	if err != nil {
		return err
	}

	//Switch method based on language
	switch language {
	case "go":
		//Write Dockerfile
		_, err = file.WriteString(docker_golang)
		if err != nil {
			return err
		}

	case "nodejs":
		//Write Dockerfile
		_, err = file.WriteString(docker_nodejs)
		if err != nil {
			return err
		}

	case "python":
		//Write Dockerfile
		_, err = file.WriteString(docker_python)
		if err != nil {
			return err
		}

	case "php":
		//Write Dockerfile
		_, err = file.WriteString(docker_php)
		if err != nil {
			return err
		}

	case "rust":
		//Write Dockerfile
		_, err = file.WriteString(docker_rust)
		if err != nil {
			return err
		}

	case "clojure":
		//Write Dockerfile
		_, err = file.WriteString("FROM clojure:openjdk-11-lein-2.9.5-alpine\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN lein uberjar\n" + "EXPOSE " + Port + "\n" + "CMD [\"java\", \"-jar\", \"target/uberjar/main-0.1.0-SNAPSHOT-standalone.jar\"]")
		if err != nil {
			return err
		}

	case "java":
		//Write Dockerfile
		_, err = file.WriteString("FROM openjdk:11.0.10-jdk-buster\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN javac main.java\n" + "EXPOSE " + Port + "\n" + "CMD [\"java\", \"main\"]")
		if err != nil {
			return err
		}

	case "ruby":
		//Write Dockerfile
		_, err = file.WriteString("FROM ruby:2.7.2-alpine3.13\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN bundle install\n" + "EXPOSE " + Port + "\n" + "CMD [\"ruby\", \"main.rb\"]")
		if err != nil {
			return err
		}

	case "c":
		//Write Dockerfile
		_, err = file.WriteString("FROM gcc:10.2.0-alpine3.13\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN gcc main.c -o main\n" + "EXPOSE " + Port + "\n" + "CMD [\"./main\"]")
		if err != nil {
			return err
		}

	case "csharp":
		//Write Dockerfile
		_, err = file.WriteString("FROM mcr.microsoft.com/dotnet/sdk:5.0-alpine\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN dotnet publish -c Release -o out\n" + "EXPOSE " + Port + "\n" + "CMD [\"dotnet\", \"out/main.dll\"]")
		if err != nil {
			return err
		}

	case "swift":
		//Write Dockerfile
		_, err = file.WriteString("FROM swift:5.3.3-focal\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN swift build -c release\n" + "EXPOSE " + Port + "\n" + "CMD [\"./.build/release/main\"]")
		if err != nil {
			return err
		}

	case "elixir":
		//Write Dockerfile
		_, err = file.WriteString("FROM elixir:1.11.3-alpine\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN mix local.hex --force\n" + "RUN mix local.rebar --force\n" + "RUN mix deps.get\n" + "RUN mix compile\n" + "EXPOSE " + Port + "\n" + "CMD [\"mix\", \"run\", \"--no-halt\"]")
		if err != nil {
			return err
		}

	case "haskell":
		//Write Dockerfile
		_, err = file.WriteString("FROM haskell:8.10.4-alpine\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN stack build\n" + "EXPOSE " + Port + "\n" + "CMD [\"stack\", \"exec\", \"main\"]")
		if err != nil {
			return err
		}

	case "dart":
		//Write Dockerfile
		_, err = file.WriteString("FROM google/dart:2.12.0\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN pub get\n" + "EXPOSE " + Port + "\n" + "CMD [\"dart\", \"main.dart\"]")
		if err != nil {
			return err
		}

	case "kotlin":
		//Write Dockerfile
		_, err = file.WriteString("FROM openjdk:11.0.10-jdk-buster\n" + "WORKDIR /app\n" + "COPY . .\n" + "RUN kotlinc main.kt -include-runtime -d main.jar\n" + "EXPOSE " + Port + "\n" + "CMD [\"java\", \"-jar\", \"main.jar\"]")
		if err != nil {
			return err
		}

	}

	//Close file
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
