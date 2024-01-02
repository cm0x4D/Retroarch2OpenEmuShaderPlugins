package shader

import (
	"fmt"
	"github.com/pierrre/archivefile/zip"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var shaderIncludeLine *regexp.Regexp
var shaderPresetReferenceLine *regexp.Regexp

func init() {
	shaderIncludeLine = regexp.MustCompile(`shader(\d*)\s*=\s*(\S*)`)
	shaderPresetReferenceLine = regexp.MustCompile(`#reference "(.*)"`)
}

type Preset interface {
	FilePath() string
	PresetName() string
	SaveAsOpenEmuPlugin(outputFolder string, keepFolders bool) error
}

type preset struct {
	inputFilePath string
	presetName    string
	presetLines   []string
}

func FindPresets(path string) []Preset {
	var presets []Preset
	_ = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ".slangp" && !strings.HasPrefix(filepath.Base(path), "_") {
			presets = append(presets, &preset{
				inputFilePath: path,
				presetName:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
			})
		}
		return nil
	})
	return presets
}

func NewPreset(path string) (Preset, error) {
	if filepath.Ext(path) != ".slangp" {
		return nil, fmt.Errorf("file %s is not a shader preset", path)
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("file %s is not a shader preset", path)
	}

	return &preset{
		inputFilePath: path,
		presetName:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
	}, nil
}

func (s *preset) FilePath() string {
	return s.inputFilePath
}

func (s *preset) PresetName() string {
	return s.presetName
}

func (s *preset) SaveAsOpenEmuPlugin(outputFolder string, keepFolders bool) error {
	shaderFiles, err := s.parse()
	if err != nil {
		return err
	}

	err = s.createShaderPresetFolder(outputFolder)
	if err != nil {
		return err
	}
	if !keepFolders {
		defer s.removeShaderPresetFolder(outputFolder)
	}

	for _, shaderFile := range shaderFiles {
		err = copyFile(filepath.Join(filepath.Dir(s.inputFilePath), shaderFile), filepath.Join(outputFolder, s.presetName, "shaders", filepath.Base(shaderFile)))
		if err != nil {
			return err
		}
	}

	err = s.savePatchedPreset(filepath.Join(outputFolder, s.presetName, fmt.Sprintf("%s.slangp", s.presetName)))
	if err != nil {
		return err
	}

	return s.compressToOpenEmuPlugin(outputFolder)
}

func (s *preset) parse() ([]string, error) {
	content, err := os.ReadFile(s.inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading preset file %s: %v", s.inputFilePath, err)
	}

	if len(content) == 0 {
		return nil, fmt.Errorf("preset file %s is empty", s.inputFilePath)
	}

	lines := strings.Split(string(content), "\n")

	// TODO: make recursive references working
	for i, _ := range lines {
		line := lines[i]
		referenceMatches := shaderPresetReferenceLine.FindStringSubmatch(line)
		if len(referenceMatches) >= 1 {
			referencedFilePath := referenceMatches[1]
			referencedLines, err := s.loadPresetReference(referencedFilePath)
			if err != nil {
				return nil, err
			}
			lines = append(lines[:i], append(referencedLines, lines[i+1:]...)...)
			i += len(referencedLines)
		}
	}

	var shaderFiles []string
	for _, line := range lines {
		matches := shaderIncludeLine.FindStringSubmatch(line)
		if len(matches) >= 2 {
			shaderFile := strings.Trim(matches[2], `"`)
			shaderFiles = append(shaderFiles, shaderFile)
		}

		if strings.Contains(line, "textures") {
			return nil, fmt.Errorf("shaders with textures are not supported")
		}

		// TODO: remove when recursive references are working
		if strings.Contains(line, "#reference") {
			return nil, fmt.Errorf("shaders with multiple references are not supported")
		}
	}

	if len(shaderFiles) == 0 {
		return nil, fmt.Errorf("preset file %s does not contain any shader", s.inputFilePath)
	}

	s.presetLines = lines
	return shaderFiles, nil
}

func (s *preset) loadPresetReference(relativePath string) ([]string, error) {
	referencedPresetFile, err := os.ReadFile(filepath.Join(filepath.Dir(s.inputFilePath), relativePath))
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(referencedPresetFile), "\n")

	for i, line := range lines {
		matches := shaderIncludeLine.FindStringSubmatch(line)
		if len(matches) >= 2 {
			lines[i] = fmt.Sprintf("shader%s = %s", matches[1], filepath.Join(filepath.Dir(relativePath), filepath.Base(strings.Trim(matches[2], `"`))))
		}
	}

	return lines, nil
}

func (s *preset) createShaderPresetFolder(outputFolder string) error {
	return os.MkdirAll(filepath.Join(outputFolder, s.presetName, "shaders"), 0755)
}

func copyFile(from string, to string) error {
	source, err := os.Open(from)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(to)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}

func (s *preset) savePatchedPreset(to string) error {
	destination, err := os.Create(to)
	if err != nil {
		return err
	}
	defer destination.Close()

	for _, line := range s.presetLines {
		matches := shaderIncludeLine.FindStringSubmatch(line)
		if len(matches) >= 2 {
			shaderNumber := strings.Trim(matches[1], `"`)
			shaderFile := strings.Trim(matches[2], `"`)

			_, err = destination.WriteString(fmt.Sprintf("shader%s = shaders/%s\n", shaderNumber, filepath.Base(shaderFile)))
			if err != nil {
				return err
			}
		} else {
			_, err = destination.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *preset) removeShaderPresetFolder(outputFolder string) {
	_ = os.RemoveAll(filepath.Join(outputFolder, s.presetName))
}

func (s *preset) compressToOpenEmuPlugin(outputFolder string) error {
	return zip.ArchiveFile(filepath.Join(outputFolder, s.presetName)+string(filepath.Separator), filepath.Join(outputFolder, s.presetName+".oeshaderplugin"), nil)
}
