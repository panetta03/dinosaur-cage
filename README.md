# Dinosaur Project

A simple project to manage dinosaurs in a virtual dinosaur park.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Features](#features)
- [Considerations] (#considerations)

## Installation

To get started, follow these steps:

1. Clone the repository:
git clone https://github.com/panetta03/dinosaur-cage
cd dinosaur-cage

2. Install dependencies: go mod tidy

3. Run the application: go run main.go

4. Navigate to Swagger UI http://localhost:8080/swagger/index.html



## Usage

This project allows you to create, manage, and interact with virtual dinosaurs in a park.

## Getting Started

1. Create a new dinosaur.
2. Place the dinosaur in a cage.
3. Interact with the dinosaurs in various ways.

## Configuration

You can configure the project by editing the `config.yml` file. This file contains settings for various aspects of the dinosaur park.

## Features

- Create new dinosaurs of different species.
- Create cages with different attributes.
- Place dinosaurs in different cages.
- Interact with dinosaurs in various ways.
- Manage the dinosaur park efficiently.

## Considerations

The approach to this project to create an extensible way to add new species of dinosaurs along the way. I chose Factory design pattern for this as this provided the best way for future extensibility and mainteance across many species. I chose to build the project using models and controllers to interact with the data. When it came to persistence I ended up with go-memdb. 

Given that I was learning go with this project I wanted to avoid any external dependencies that could make it difficult to use this project. For example I started with a Postgres implementation and then moved on to sqlite. When it became clear that time was factor for this project I ended up with an in memory solution for persistence using go-memdb.

 With Cages I was able to manage to add most of the features of a cage. I abstracted out checks on cages into a separate utils section to avoid too much noise in the controller. Additionally I chose swagger for clean documentation and interaction with the APIs. 
