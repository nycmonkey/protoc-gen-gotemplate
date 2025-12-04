package pgghelpers

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func GenerateFile(gen *protogen.Plugin, file *protogen.File) ([]*pluginpb.CodeGeneratorResponse_File, error) {
	// Parse parameters
	params := ParseParams(gen.Request.GetParameter())

	if params.Debug {
		// Log debug info if needed
	}

	fd := file.Proto

	var filesToGenerate []*pluginpb.CodeGeneratorResponse_File

	if params.All {
		// Generate for the whole file
		encoder := NewGenericTemplateBasedEncoder(params.TemplateDir, fd, params.Debug, params.DestinationDir)
		filesToGenerate = append(filesToGenerate, encoder.Files()...)
	} else if params.FileMode {
		if len(fd.Service) > 0 {
			encoder := NewGenericTemplateBasedEncoder(params.TemplateDir, fd, params.Debug, params.DestinationDir)
			filesToGenerate = append(filesToGenerate, encoder.Files()...)
		}
	} else {
		// Service mode
		for _, service := range fd.Service {
			encoder := NewGenericServiceTemplateBasedEncoder(params.TemplateDir, service, fd, params.Debug, params.DestinationDir)
			filesToGenerate = append(filesToGenerate, encoder.Files()...)
		}
	}

	return filesToGenerate, nil
}
