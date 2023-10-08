module qat-computer

go 1.19

require (
	github.com/stretchr/testify v1.8.4
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Packages system

replace qat-computer/docs => ../docs

replace qat-computer/helpers => ../helpers

replace qat-computer/logger => ../logger

replace qat-computer/on_go => ../on_go

replace qat-computer/utils => ../utils
