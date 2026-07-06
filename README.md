# Overview
This is a tool which can parse the documentation for Jomini engine games. These are Victoria 3, Crusader Kings 3 and Europa Universalis V.

## How To Build
First download and install the Go SDK:
- https://go.dev/doc/install

Next, open the project folder in a terminal (e.g. cmd) and run the following command:
```
go build
```

That is it. There should be an executable in the project folder now.

## Setup
Generate the documentation files for the game you want. They can be generated with these two in-game console commands:
- `script_docs`
- `dump_data_types`

The generated files can be found in `<Documents>/Paradox Interactive/<Game>/`. These are the relevant files:
- `docs/custom_localization.log`
- `docs/effects.log`
- `docs/event_targets.log`
- `docs/modifiers.log`
- `docs/on_actions.log`
- `docs/triggers.log`
- `logs/event_scopes.log`
- `data_types/data_types_common.log`
- `data_types/data_types_gui.log`
- `data_types/data_types_internalclausewitzgui.log`
- `data_types/data_types_script.log`
- `data_types/data_types_uncategorized.log`

Put the generated documentation into the `docs/new` folder.

The `docs/old` is only needed when you want to generate a digest and should contain the old versions documentation to generate a diff of both versions.

## Usage
Run the application in the command line with one of the following commands:
- `digest` to generate a modding digest
- `cwt` to generate CWT configuration files
- `npp` to generate Notepad++ language files
- `json` to generate machine-readable JSON files

Example: `.\pdx-documentation-manager.exe json`