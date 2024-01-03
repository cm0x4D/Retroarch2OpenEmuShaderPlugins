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
    \
    shaders_slang/anti-aliasing/advanced-aa.slangp \
    shaders_slang/anti-aliasing/fxaa.slangp \
    shaders_slang/anti-aliasing/reverse-aa.slangp \
    \
    shaders_slang/crt/advanced_crt_whkrmrgks0.slangp \
    shaders_slang/crt/crt-1tap.slangp \
    shaders_slang/crt/crt-aperture.slangp \
    shaders_slang/crt/crt-blurPi-sharp.slangp \
    shaders_slang/crt/crt-blurPi-soft.slangp \
    shaders_slang/crt/crt-caligari.slangp \
    shaders_slang/crt/crt-cgwg-fast.slangp \
    shaders_slang/crt/crt-consumer.slangp \
    shaders_slang/crt/crt-Cyclon.slangp \
    shaders_slang/crt/crt-easymode.slangp \
    shaders_slang/crt/crt-frutbunn.slangp \
    shaders_slang/crt/crt-gdv-mini.slangp \
    shaders_slang/crt/crt-gdv-mini-ultra-trinitron.slangp \
    shaders_slang/crt/crt-geom.slangp \
    shaders_slang/crt/crt-geom-mini.slangp \
    shaders_slang/crt/crt-geom-tate.slangp \
    shaders_slang/crt/crt-hyllian-3d.slangp \
    shaders_slang/crt/crt-hyllian-fast.slangp \
    shaders_slang/crt/crt-hyllian-multipass.slangp \
    shaders_slang/crt/crt-interlaced-halation.slangp \
    shaders_slang/crt/crt-lottes.slangp \
    shaders_slang/crt/crt-lottes-fast.slangp \
    shaders_slang/crt/crt-lottes-multipass.slangp \
    shaders_slang/crt/crt-mattias.slangp \
    shaders_slang/crt/crt-nes-mini.slangp \
    shaders_slang/crt/crt-nobody.slangp \
    shaders_slang/crt/crt-pi.slangp \
    shaders_slang/crt/crt-pocket.slangp \
    shaders_slang/crt/crt-simple.slangp \
    shaders_slang/crt/crt-sines.slangp \
    shaders_slang/crt/fake-crt-geom.slangp \
    shaders_slang/crt/fake-crt-geom-potato.slangp \
    shaders_slang/crt/fakelottes.slangp \
    shaders_slang/crt/gizmo-crt.slangp \
    shaders_slang/crt/gizmo-slotmask-crt.slangp \
    shaders_slang/crt/tvout-tweaks.slangp \
    shaders_slang/crt/yee64.slangp \
    shaders_slang/crt/yeetron.slangp \
    shaders_slang/crt/zfast-crt-composite.slangp \
    shaders_slang/crt/zfast-crt-curvature.slangp

build-all:
	@go run main.go -i $(INPUT_FOLDER) -o $(OUTPUT_FOLDER)

build-tested:
	@for TESTED_SHADER_PRESET in $(TESTED_SHADER_PRESETS); do \
		go run main.go -f $$TESTED_SHADER_PRESET -o $(TESTED_FOLDER); \
	done

clean:
	rm -rf $(OUTPUT_FOLDER) $(TESTED_FOLDER)

.PHONY: all build-tested build-all clean