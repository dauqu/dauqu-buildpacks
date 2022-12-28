package buildpacks

import (
	"os"
	"strings"
)

func CreateDockerfile(Workdir string, Port string, language string, Name string) (err error) {

	rname := strings.ReplaceAll(Name, ".", "_")

	docker_static := `
FROM nginx:1.19.6-alpine
WORKDIR ` + Workdir + "/" + rname + `
COPY . .
EXPOSE ` + Port + `
CMD ["nginx", "-g", "daemon off;"]
	`

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
FROM golang:alpine as builder
WORKDIR ` + Workdir + "/" + rname + `
COPY . .
RUN go get
RUN go build -o main .

FROM alpine
WORKDIR /app
COPY --from=builder ` + Workdir + "/" + rname + ` /app/
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
	docker_clojure := `
FROM clojure:openjdk-11-lein-2.9.5-alpine
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "lein", "run" ]
	`
	docker_ruby := `
FROM ruby:2.7.2-alpine3.13
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "ruby", "main.rb" ]
	`
	docker_java := `
FROM openjdk:11.0.10-jdk
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "java", "main.java" ]
	`
	docker_c := `
FROM gcc:10.2.0
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "gcc", "main.c" ]
	`
	docker_csharp := `
FROM mcr.microsoft.com/dotnet/sdk:5.0
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "dotnet", "run" ]
	`
	docker_swift := `
FROM swift:5.3.3
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "swift", "main.swift" ]
	`
	docker_elixir := `
FROM elixir:1.11.3-alpine
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "elixir", "main.exs" ]
	`
	docker_haskell := `
FROM haskell:8.10.4
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "ghc", "main.hs" ]
	`
	docker_perl := `
FROM perl:5.32.1
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "perl", "main.pl" ]
	`
	docker_dart := `
FROM google/dart:2.12.0
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "dart", "main.dart" ]
	`
	docker_lua := `
FROM lua:5.4.1-alpine
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "lua", "main.lua" ]
	`
	docker_r := `
FROM rocker/r-ver:4.0.3
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "Rscript", "main.R" ]
	`
	docker_kotlin := `
FROM openjdk:11.0.10-jdk
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "kotlinc", "main.kt" ]
	`
	docker_scala := `
FROM openjdk:11.0.10-jdk
WORKDIR ` + Workdir + `
COPY . .
EXPOSE ` + Port + `
CMD [ "scalac", "main.scala" ]
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
	case "static":
		//Write Dockerfile
		_, err = file.WriteString(docker_static)
		if err != nil {
			return err
		}

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
		_, err = file.WriteString(docker_clojure)
		if err != nil {
			return err
		}

	case "java":
		//Write Dockerfile
		_, err = file.WriteString(docker_java)
		if err != nil {
			return err
		}

	case "ruby":
		//Write Dockerfile
		_, err = file.WriteString(docker_ruby)
		if err != nil {
			return err
		}

	case "c":
		//Write Dockerfile
		_, err = file.WriteString(docker_c)
		if err != nil {
			return err
		}

	case "csharp":
		//Write Dockerfile
		_, err = file.WriteString(docker_csharp)
		if err != nil {
			return err
		}

	case "swift":
		//Write Dockerfile
		_, err = file.WriteString(docker_swift)
		if err != nil {
			return err
		}

	case "elixir":
		//Write Dockerfile
		_, err = file.WriteString(docker_elixir)
		if err != nil {
			return err
		}

	case "haskell":
		//Write Dockerfile
		_, err = file.WriteString(docker_haskell)
		if err != nil {
			return err
		}

	case "dart":
		//Write Dockerfile
		_, err = file.WriteString(docker_dart)
		if err != nil {
			return err
		}

	case "kotlin":
		//Write Dockerfile
		_, err = file.WriteString(docker_kotlin)
		if err != nil {
			return err
		}

	case "scala":
		//Write Dockerfile
		_, err = file.WriteString(docker_scala)
		if err != nil {
			return err
		}

	case "perl":
		//Write Dockerfile
		_, err = file.WriteString(docker_perl)
		if err != nil {
			return err
		}

	case "lua":
		//Write Dockerfile
		_, err = file.WriteString(docker_lua)
		if err != nil {
			return err
		}

	case "r":
		//Write Dockerfile
		_, err = file.WriteString(docker_r)
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
