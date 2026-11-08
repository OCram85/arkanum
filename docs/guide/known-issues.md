# 💣 Known Issues + Limitations

## 🐛 Starship.rs

Starship detects workspaces as active python projects. It always appends the prompt
fragment `via 🐍 (lsiopy)`. For now I disabled the python module in starship.

## 🐛 First boot setup of default extensions

The container does not install the default extensions automatically. Open a new
terminal and run the following commands once to finish the setup:

```bash
# reset the vscode user settings
arkanum config reset-codesettings
# install the recommended default extensions
arkanum config install-extensions
# Reload with command F1 + Developer: Reload Window
```
