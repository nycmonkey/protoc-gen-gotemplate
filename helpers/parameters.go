package pgghelpers

import (
	"log"
	"strings"
)

const (
	boolTrue  = "true"
	boolFalse = "false"
)

type Parameters struct {
	TemplateDir       string
	DestinationDir    string
	Debug             bool
	All               bool
	SinglePackageMode bool
	FileMode          bool
}

func ParseParams(parameter string) Parameters {
	var params Parameters

	if parameter != "" {
		for _, param := range strings.Split(parameter, ",") {
			parts := strings.Split(param, "=")
			if len(parts) != 2 {
				log.Printf("Err: invalid parameter: %q", param)
				continue
			}
			switch parts[0] {
			case "template_dir":
				params.TemplateDir = parts[1]
			case "destination_dir":
				params.DestinationDir = parts[1]
			case "single-package-mode":
				switch strings.ToLower(parts[1]) {
				case boolTrue, "t":
					params.SinglePackageMode = true
				case boolFalse, "f":
				default:
					log.Printf("Err: invalid value for single-package-mode: %q", parts[1])
				}
			case "debug":
				switch strings.ToLower(parts[1]) {
				case boolTrue, "t":
					params.Debug = true
				case boolFalse, "f":
				default:
					log.Printf("Err: invalid value for debug: %q", parts[1])
				}
			case "all":
				switch strings.ToLower(parts[1]) {
				case boolTrue, "t":
					params.All = true
				case boolFalse, "f":
				default:
					log.Printf("Err: invalid value for debug: %q", parts[1])
				}
			case "file-mode":
				switch strings.ToLower(parts[1]) {
				case boolTrue, "t":
					params.FileMode = true
				case boolFalse, "f":
				default:
					log.Printf("Err: invalid value for file-mode: %q", parts[1])
				}
			default:
				log.Printf("Err: unknown parameter: %q", param)
			}
		}
	}
	return params
}
