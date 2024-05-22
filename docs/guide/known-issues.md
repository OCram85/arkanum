# 💣 Known Issues + Limitations

## 🐛 Starship.rs

Starship detects workspaces as active python projects. It always appends the prompt
fragment `via 🐍 (lsiopy)`. For now I disabled the python module in starship.

## 🐛 Default extensions installation timing error

If the automatic installation of the default extension fails, you can always retry he installation with the
following command:

```bash
# restart the installation
arkanum config install-extensions
# Optional: reset the vscode user setting
arkanum config reset-codesettings
# Reload with command F1 + Developer: Reload Window
```
