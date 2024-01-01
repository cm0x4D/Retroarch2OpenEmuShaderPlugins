# Retroarch2OpenEmuShaderPlugins

## Overview

This project aims to bring as many shaders as possible from the current Retroarch slang shader folder to OpenEmu by 
providing a way to automatically convert the existing Retroarch shader presets to OpenEmu shader plugins. Shaders 
enhance the visual experience of retro games by applying various graphical effects, and having my favorite Retroarch 
shaders in OpenEmu was the motivation for this project.

Note that this project is still in its early stages and is not yet ready for general use. The current version can 
convert already many shaders, currently only shaders without textures and other resources are converted. Most shader
plugins work, but some still cause a crash of OpenEmu and I did not have the time to investigate the cause of the
crashes or to test every shader plugin. The current version is only tested on macOS 14.1 (M3) and OpenEmu 2.4.1.

If you want to help with this project, feel free to open an issue or a pull request. I am especially interested in
fixing the remaining crashes and adding support for shaders with textures and other resources.

The shaders included in this repository are taken from Retroarch version 1.16.0. You can replace the folder 
`shaders_slang` with the one from your Retroarch installation to use newer shaders.

## Requirements

In order to build the OpenEmu shader plugins, you need to have the following tools installed:

- [git](https://git-scm.com/)
- [make](https://www.gnu.org/software/make/)
- [go](https://golang.org/)

## How to build the shader plugins

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

The shader plugins are now located in the `plugins` directory.
