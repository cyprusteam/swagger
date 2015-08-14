package main

import (
  "flag"
  "log"

  "github.com/cyprusteam/swagger/generator"
)

var apiPackage = flag.String("apiPackage", "", "The package that implements the API controllers, relative to $GOPATH/src")
var mainApiFile = flag.String("mainApiFile", "", "The file that contains the general API annotations, relative to $GOPATH/src")
var outputFormat = flag.String("format", "go", "Output format type for the generated files: "+generator.AVAILABLE_FORMATS)
var outputSpec = flag.String("output", "", "Output (path) for the generated file(s)")
var controllerClass = flag.String("controllerClass", "", "Speed up parsing by specifying which receiver objects have the controller methods")
var packageName = flag.String("packageName", "", "The package name docs file")

func main() {
  flag.Parse()

  if *mainApiFile == "" {
    *mainApiFile = *apiPackage + "/main.go"
  }

  if *apiPackage == "" {
    flag.PrintDefaults()
    return
  }

  params := generator.Params{
    ApiPackage:      *apiPackage,
    MainApiFile:     *mainApiFile,
    OutputFormat:    *outputFormat,
    OutputSpec:      *outputSpec,
    ControllerClass: *controllerClass,
    PackageName:     *packageName,
  }

  err := generator.Run(params)
  if err != nil {
    log.Fatal(err.Error())
  }
}
