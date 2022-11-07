FROM linuxserver/code-server:4.7.1

RUN \
  echo "**** install starshipt prompt ****" && \
  curl -sS https://starship.rs/install.sh | sh -s -- -f && \
  echo "eval \"\$(starship init bash)\"" >> /etc/bash.bashrc

ENV STARSHIP_CONFIG=/etc/starship.toml

COPY starship.toml /etc/starship.toml

RUN \
  echo "**** setup git ****" && \
  git config --system credential.helper store && \
  echo 'source /usr/share/bash-completion/completions/git' >> /etc/bash.bashrc

ADD install-devruntime /usr/bin/
RUN \
  chmod +x /usr/bin/install-devruntime && \
  echo "echo \"Use 'install-devruntime' to install missing runtimes like dotnet or NodeJs.\"" >> /etc/bash.bashrc
