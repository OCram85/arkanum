FROM quay.io/linuxserver.io/code-server:4.8.3

RUN \
  echo "**** install starship prompt ****" && \
  curl -sS https://starship.rs/install.sh | sh -s -- -f && \
  echo "eval \"\$(starship init bash)\"" >> /etc/bash.bashrc

ENV STARSHIP_CONFIG=/etc/starship.toml

COPY starship.toml /etc/starship.toml

ADD gitconfig-system /etc/gitconfig
RUN \
  echo "**** setup git ****" && \
  # using prepared systemwide config file instead.
  #git config --system credential.helper store && \
  echo 'source /usr/share/bash-completion/completions/git' >> /etc/bash.bashrc

ADD arkanum /usr/bin/
ADD arkanum-completion /etc/bash_completion.d/
RUN \
  chmod +x /usr/bin/arkanum && \
  chmod +x /etc/bash_completion.d/arkanum-completion && \
  echo 'source /etc/bash_completion.d/arkanum-completion' >> /etc/bash.bashrc && \
  touch "$HOME/enable_motd" && \
  echo "if [[ ! -e \"$HOME/data/User/settings.json\" ]]; then arkanum --install-extensions && arkanum --reset-codesetting && \
    echo -e \"🧙 \\e[32markanum\\e[0m: Please reload Arkanum to finalize the setup...\" && read foo; fi" >> /etc/bash.bashrc && \
  echo "if [[ -e \"$HOME/enable_motd\" ]]; then echo -e \"Use 🧙 \\e[32m'arkanum'\\e[0m to install missing runtimes like dotnet or NodeJs.\"; fi" >> /etc/bash.bashrc

WORKDIR /app/code-server/lib/vscode/out/vs/workbench
ADD FiraCode/fonts/* ./fonts/
ADD FiraCode/fonts.css ./
RUN cat fonts.css >> workbench.web.main.css
