package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
)

const genGoDocURL = "https://protobuf.dev/reference/go/go-generated"
const grpcDocURL = "https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		showVersion()
		os.Exit(0)
	}
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Fprintf(os.Stdout, "See "+genGoDocURL+" for usage information.\n")
		os.Exit(0)
	}

	var (
		flags        flag.FlagSet
		plugins      = flags.String("plugins", "", "deprecated option")
		importPrefix = flags.String("import_prefix", "", "prefix to prepend to import paths")
		//experimentalStripNonFunctionalCodegen = flags.Bool("experimental_strip_nonfunctional_codegen", false, "experimental_strip_nonfunctional_codegen true means that the plugin will not emit certain parts of the generated code in order to make it possible to compare a proto2/proto3 file with its equivalent (according to proto spec) editions file. Primarily, this is the encoded descriptor.")
	)

	opts := &protogen.Options{
		ParamFunc: flags.Set,
		ImportRewriteFunc: func(importPath protogen.GoImportPath) protogen.GoImportPath {
			switch importPath {
			case "context", "fmt", "math":
				return importPath
			}
			if *importPrefix != "" {
				return protogen.GoImportPath(*importPrefix) + importPath
			}
			return importPath
		},
	}

	err := run(opts, func(gen *protogen.Plugin) error {
		if *plugins != "" {
			reason := fmt.Sprintf("protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC\n\nSee %s for more information.", grpcDocURL)
			return errors.New(reason)
		}

		for _, f := range gen.Files {
			if f.Generate {
				gengo.GenerateFile(gen, f)
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		gen.SupportedEditionsMinimum = gengo.SupportedEditionsMinimum
		gen.SupportedEditionsMaximum = gengo.SupportedEditionsMaximum
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}
