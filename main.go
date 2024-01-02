package main

import (
	"Retrorach2OpenEmuShaders/shader"
	"flag"
	"fmt"
	"os"
)

func main() {
	shaderPresetFile := flag.String("f", "", "Convert a single shader preset to an OpenEmu shader plugin")
	inputFolder := flag.String("i", "", "Find all shaders in the folder and convert them to OpenEmu shader plugins")
	outputFolder := flag.String("o", "", "Output folder for the OpenEmu shader plugins, defaults to 'out'")
	keepFolders := flag.Bool("k", false, "Keep the folders after converting them to OpenEmu shader plugins")

	flag.Parse()

	if *outputFolder == "" {
		*outputFolder = "out"
	}

	if *inputFolder != "" {
		presets := shader.FindPresets(*inputFolder)
		for _, preset := range presets {
			err := preset.SaveAsOpenEmuPlugin(*outputFolder, *keepFolders)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error saving preset %s: %v\n", preset.PresetName(), err)
			} else {
				fmt.Printf("Saved preset: %s\n", preset.PresetName())
			}
		}
	} else if *shaderPresetFile != "" {
		preset, err := shader.NewPreset(*shaderPresetFile)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error loading preset %s: %v\n", *shaderPresetFile, err)
			return
		}

		err = preset.SaveAsOpenEmuPlugin(*outputFolder, *keepFolders)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error saving preset %s: %v\n", preset.PresetName(), err)
		} else {
			fmt.Printf("Saved preset: %s\n", preset.PresetName())
		}
	} else {
		flag.Usage()
	}
}
