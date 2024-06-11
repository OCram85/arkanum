# VSCode

## Default Settings

Arkanum tests for existing VSCode user config. If not preset, it sets the default config with these values:

|               Key                | Value                                                                                   | Description                                                                                              |
| :------------------------------: | :-------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
|    `window.menuBarVisibility`    | **compact**                                                                             | Uses compact main menu bar in hamburger style.                                                           |
|      `workbench.colorTheme`      | **One Dark Pro Darker**                                                                 | Enables default color theme.                                                                             |
|      `workbench.iconTheme`       | **vscode-icons**                                                                        | Enables default icon theme for file tree.                                                                |
|       `editor.fontFamily`        | **'FiraCode', 'FiraCode Nerd Font', 'FiraCode NF', Consolas, 'Courier New', monospace** | Enables included FiraCode font for all possible variations in the file editor.                           |
| `terminal.integrated.fontFamily` | **'FiraCode Mono', 'FiraCode Nerd Font Mono', 'FiraCode NFM', Consolas, monospace**     | Enables included FiraCode fonts in terminal views. Uses mono variation to enable showing icons.          |
|      `editor.fontLigatures`      | **true**                                                                                | Enables font ligatures supported in FiraCode: `->`, `---`, `!=` ...                                      |
|      `editor.formatOnSave`       | **true**                                                                                | Enables format on save features based on language config.                                                |
|     `extensions.autoUpdate`      | **false**                                                                               | Disables automatic update for installed extensions. Prevents running into VSCode compatibility problems. |
|        `git.confirmSync`         | **false**                                                                               | Disables notification popup for git sync action.                                                         |
|    `telemetry.telemetryLevel`    | **off**                                                                                 | Disables sending telemetry data for VSCode and GitLense Extension.                                       |

## Extensions

Arkanum includes the following VSCode extensions from the Open VSX Registry:

### GitLense

Adds advanced git features with the [GitLense](https://open-vsx.org/extension/eamodio/gitlens) extension.

### One Dark Pro

Adds [One Dark Pro](https://open-vsx.org/extension/zhuangtongfa/material-theme) theme as default color theme.

### vscode-icons

Adds [vscode-icons](https://open-vsx.org/extension/vscode-icons-team/vscode-icons) extension as default icon theme.

## Referenced Source Files

::: code-group

<<< @/../Dockerfile#cli{Dockerfile}

<<< @/../arkanum#code-settings{4-15 bash:line-numbers}
:::
