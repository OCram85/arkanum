FROM quay.io/linuxserver.io/code-server:4.104.1

#LABEL build_version=""
LABEL maintainer="OCram85"
ARG VERSION
LABEL build_version="${VERSION}"
LABEL org.opencontainers.image.authors="OCram85"
LABEL org.opencontainers.image.vendor="OCram85"

LABEL org.opencontainers.image.title="Arkanum"
LABEL org.opencontainers.image.description="Code-Server container optimized for daily use."
LABEL org.opencontainers.image.licenses="AGPL-3.0"
ARG TAG
LABEL org.opencontainers.image.version="${TAG}"

LABEL org.opencontainers.image.url="https://gitea.ocram85.com/arkanum/arkanum"
LABEL org.opencontainers.image.source="https://gitea.ocram85.com/arkanum/arkanum.git"
LABEL org.opencontainers.image.documentation="https://gitea.ocram85.com/arkanum/arkanum"

#region starship
RUN \
  echo "**** install starship prompt ****" && \
  curl -sS -o /tmp/install.sh https://starship.rs/install.sh && \
  chmod +x /tmp/install.sh && \
  /tmp/install.sh --verbose --force --version latest && \
  rm -f /tmp/install.sh && \
  echo "eval \"\$(starship init bash)\"" >> /etc/bash.bashrc

ENV STARSHIP_CONFIG=/etc/starship.toml
COPY starship.toml /etc/starship.toml
#endregion starship

#region git
ADD gitconfig-system /etc/gitconfig
RUN \
  echo "**** setup git ****" && \
  echo 'source /usr/share/bash-completion/completions/git' >> /etc/bash.bashrc
#endregion git

#region cli
ADD arkanum /usr/bin/
ADD arkanum-completion /etc/bash_completion.d/
RUN \
  chmod +x /usr/bin/arkanum && \
  chmod +x /etc/bash_completion.d/arkanum-completion && \
  echo 'source /etc/bash_completion.d/arkanum-completion' >> /etc/bash.bashrc && \
  touch "$HOME/enable_motd" && \
  echo "if [[ ! -e \"$HOME/data/User/settings.json\" ]]; then arkanum config install-extensions && arkanum config reset-codesettings && \
    echo -e \"🧙 \\e[32markanum\\e[0m: Please reload Arkanum to finalize the setup...\" && read foo; fi" >> /etc/bash.bashrc && \
  echo "if [[ -e \"$HOME/enable_motd\" ]]; then echo -e \"Use 🧙 \\e[32m'arkanum'\\e[0m to install missing runtimes like dotnet or NodeJs.\"; fi" >> /etc/bash.bashrc
#endregion cli

#region firacode

# TODO: validate dir: /lib/vscode/out/vs/code/browser/workbench/workbench.css
#WORKDIR /app/code-server/lib/vscode/out/vs/workbench
WORKDIR /app/code-server/lib/vscode/out/vs/code/browser/workbench
ADD FiraCode/fonts/* ./fonts/
ADD FiraCode/fonts.css ./
RUN cat fonts.css >> workbench.css
#endregion firacode

#region code-server
WORKDIR /
# remove code-server specific files to override with branded values.
# changes product images + app name
RUN \
  rm -rf /app/code-server/src/browser/media && \
  rm -f /etc/s6-overlay/s6-rc.d/svc-code-server/run && \
  echo 'alias summon="code-server"' >> /etc/bash.bashrc

COPY code-server/media /app/code-server/src/browser/media
COPY code-server/root/etc/s6-overlay/s6-rc.d/svc-code-server/run /etc/s6-overlay/s6-rc.d/svc-code-server/run
#endregion code-server

#region add-packages
RUN \
  apt-get update && \
  apt-get install --no-install-recommends -y \
    file \
    make && \
  apt-get clean
#endregion add-packages
