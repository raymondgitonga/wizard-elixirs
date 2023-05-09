# Wizard-Client

<p style="text-align:center;"><img src="resources/harry-gopher.png" alt="Gopher" width="250" height="350"></p>


Wizard-Client is a CLI app that allows users to create magical elixirs based on the ingredients they have at their disposal. Simply input a list of ingredients, and the app will return the elixirs you can make with those ingredients.

## Overview

### Project structure
```
.
|-- Makefile
|-- README.md
|-- cmd
|   `-- root.go
|-- demo.gif
|-- go.mod
|-- go.sum
|-- internal
|   |-- elixircreator
|   |   |-- elixircreator.go
|   |   |-- elixircreator_test.go
|   |   |-- prompt.go
|   |   -- prompt_mock.go
|   |-- wizardclient
|       |-- client.go
|       |-- client_mock.go
|       |-- client_test.go
|-- main.go
```

- The main functionality lies in the `elixircreator` package. `elixircreator.go` is responsible for coordinating user interactions  and the creation of elixirs using the provided ingredients. `prompt.go` is an abstraction over the prompt survey library, that collects users responses.
- The `wizardclient `package holds the HTTP client for fetching ingredient and elixir data
- Both packages make use of interfaces to make the code more flexible, allowing for easier testing and the ability to swap out implementations when needed.
- The design pattern used in the provided code is the Dependency Injection pattern. Specific dependencies are injected into structs that need them rather than being hard-coded.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)

## Installation

To install the app, ensure that you have [Go](https://golang.org/) installed on your system. Then, clone this repository:

```bash
git clone https://github.com/raymondgitonga/wizard-client.git

```

Next, navigate to the project directory:
`cd wizard-client`

### Technologies Used
1. Go as the programming language
2. [Cobra](https://github.com/spf13/cobra) library to bootstrap the CLI
3. [Go-Survey](https://github.com/AlecAivazis/survey/v2) library to make the CLI interactive

## Usage

To use the Wizard-Client app, run the following command:

```bash
make elixir
```

You'll be prompted to confirm if you're ready to make elixirs. After confirmation, choose your ingredients from the list provided. The app will then display the elixirs you can make with your chosen ingredients in JSON format.

![](resources/demo.gif)
## Testing
To run tests for the project, execute the following command:
```bash
make tests
```

## Author

- Raymond Gitonga




