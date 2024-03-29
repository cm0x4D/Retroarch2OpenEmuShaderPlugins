# Retroarch2OpenEmuShaderPlugins

[![](https://github.com/cm0x4D/Retroarch2OpenEmuShaderPlugins/actions/workflows/latest.yml/badge.svg?branch=main)](https://github.com/cm0x4D/Retroarch2OpenEmuShaderPlugins/actions/workflows/latest.yml)

## Overview

This project aims to bring as many shaders as possible from the current Retroarch slang shader folder to OpenEmu by 
providing a way to automatically convert the existing Retroarch shader presets to OpenEmu shader plugins. Shaders 
enhance the visual experience of retro games by applying various graphical effects, and having my favorite Retroarch 
shaders in OpenEmu was the motivation for this project.

Note that this project is still in its early stages and is not yet ready for general use. The current version can 
convert already many shaders, currently only shaders without textures and other resources are converted. Most shader
plugins work, but some still cause a crash of OpenEmu and I did not have the time to investigate the cause of the
crashes or to test every shader plugin. The current version is only tested on macOS 14.1 (M3) and OpenEmu 2.4.1.

I created a curated list of shaders that are known to work and look good in OpenEmu. You can find the list in the
Makefile in the `TESTED_SHADER_PRESETS` variable. If you want to add a shader to the list, feel free to open a pull
request or an issue. Ideally with a screenshot of the shader in action, so that I can add it to the list of tested 
shaders.

When building the shader plugins, all shaders that can be converted are saved in the folder `plugins`. All shaders that
are known to work and look good in OpenEmu are saved in the folder `plugins_tested`. When building both targets are 
build per default. If you only want to build the tested shaders, you can run `make build-tested` instead of just `make`.

If you want to help with this project, feel free to open an issue or a pull request. I am especially interested in
fixing the remaining crashes and adding support for shaders with textures and other resources.

The shaders included in this repository are taken from Retroarch version **1.16.0**. You can replace the folder 
`shaders_slang` with the one from your Retroarch installation to use newer shaders.

## Install shader plugins

You can download a given **version** or the **latest** (build from main branch) shader plugins from the 
[releases page](https://github.com/cm0x4D/Retroarch2OpenEmuShaderPlugins/releases). Each shader is provided as an
`*.oeshaderplugin` file. To install a shader plugin, just drag the shader plugin into the OpenEmu window. The shader
plugin will be installed automatically. That is how things are supposed to work on a Mac ;-).

## Build

There is no need to build the shaders yourself, you can download the latest build from the releases page. Note that only 
tested shaders are included in the release builds, so you will not get all shaders that are available. If you want
to build the shaders yourself or try untested shader plugins, you can follow the instructions below.

### Requirements

In order to build the OpenEmu shader plugins, you need to have the following tools installed:

- [git](https://git-scm.com/)
- [make](https://www.gnu.org/software/make/)
- [go](https://golang.org/)

### How to build the shader plugins

1. Clone this repository

```bash
    git clone https://github.com/cm0x4D/Retroarch2OpenEmuShaderPlugins.git
```
2. Change into the project directory

```bash
    cd Retroarch2OpenEmuShaderPlugins
```

3. Build the shader plugins

```bash
    make
```

All OpenEmu shader plugins are now located in the `plugins` directory. All tested shaders are located in the 
`plugins_tested` folder.

If you want only to build the tested shaders, you can run `make build-tested` instead of just `make`. If you want to 
build all shaders without the tested shaders, you can run `make build-all`.