package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	pgghelpers "github.com/nycmonkey/protoc-gen-gotemplate/helpers"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		// Map to store generated content by filename
		generatedFiles := make(map[string]string)
		// Map to store import path by filename (using the first one encountered)
		generatedImports := make(map[string]protogen.GoImportPath)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			files, err := pgghelpers.GenerateFile(gen, f)
			if err != nil {
				return fmt.Errorf("%s: %w", f.Desc.Path(), err)
			}
			for _, file := range files {
				name := file.GetName()
				if _, ok := generatedFiles[name]; ok {
					generatedFiles[name] += file.GetContent()
				} else {
					generatedFiles[name] = file.GetContent()
					generatedImports[name] = f.GoImportPath
				}
			}
		}

		// Write all generated files
		for name, content := range generatedFiles {
			g := gen.NewGeneratedFile(name, generatedImports[name])
			g.Write([]byte(content))
		}

		return nil
	})
}
