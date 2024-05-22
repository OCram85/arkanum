# FiraCode

## Referenced Source Files

::: code-group
<<< @/../Dockerfile#firacode{Dockerfile}
<<< @/../FiraCode/fonts.css{css}
:::

We add custom fonts on build time to the Arkanum docker image. Therefore we copy the font files and patch the
_CodeServer / Openvscode-server_ `workbench.css` files. So we can use host and consume additional fonts without
relying on external CDNs.
