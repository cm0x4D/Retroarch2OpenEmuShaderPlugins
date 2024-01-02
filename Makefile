all: build-tested build-all

OUTPUT_FOLDER := ./plugins
INPUT_FOLDER := ./shaders_slang
TESTED_FOLDER := ./plugins_tested

TESTED_SHADER_PRESETS = \
	shaders_slang/handheld/lcd-grid.slangp \
	shaders_slang/handheld/lcd-grid-v2.slangp \
	shaders_slang/handheld/lcd-grid-v2-gba-color.slangp \
	shaders_slang/handheld/lcd-grid-v2-gba-color-motionblur.slangp \
	shaders_slang/handheld/lcd-grid-v2-gbc-color.slangp \
	shaders_slang/handheld/lcd-grid-v2-gbc-color-motionblur.slangp \
	shaders_slang/handheld/lcd-grid-v2-nds-color.slangp \
    shaders_slang/handheld/lcd-grid-v2-nds-color-motionblur.slangp \
    shaders_slang/handheld/lcd-grid-v2-psp-color.slangp \
    shaders_slang/handheld/lcd-grid-v2-psp-color-motionblur.slangp \
    shaders_slang/handheld/zfast-lcd.slangp \
    shaders_slang/anti-aliasing/advanced-aa.slangp \
    shaders_slang/anti-aliasing/fxaa.slangp \
    shaders_slang/anti-aliasing/reverse-aa.slangp

build-all:
	@go run main.go -i $(INPUT_FOLDER) -o $(OUTPUT_FOLDER)

build-tested:
	@for TESTED_SHADER_PRESET in $(TESTED_SHADER_PRESETS); do \
		go run main.go -f $$TESTED_SHADER_PRESET -o $(TESTED_FOLDER); \
	done

clean:
	rm -rf $(OUTPUT_FOLDER) $(TESTED_FOLDER)

.PHONY: all build-tested build-all clean