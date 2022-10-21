FROM linuxserver/code-server:4.7.1

RUN \
  echo "**** install starshipt prompt ****" && \
  curl -sS https://starship.rs/install.sh | sh && \
  echo "eval \"\$(starship init bash)\"" >> /etc/bash.bashrc

ENV STARSHIP_CONFIG=/etc/starship.tom

COPY staship.toml /etc/starship.toml

RUN \
  "**** setup git ****" && \
  git config --system credential.helper store

RUN \
  echo "**** install dev runtimes ****" && \
  curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash - && \
  apt-get update && \
  apt-get install -y \
    nodejs \
    golang-go &&\
  echo "**** clean up ****" && \
  apt-get clean && \
  rm -rf \
    /config/* \
    /tmp/* \
    /var/lib/apt/lists/* \
    /var/tmp/*
