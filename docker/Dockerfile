FROM ubuntu:23.04

ENV USER=appuser
ENV GROUP=appgroup
ENV UID=12345
ENV GID=23456

COPY src/ /

# Update packages and install CURL
RUN apt update \
    && apt install -y curl

RUN curl https://get.kubevious.io/cli.sh -o install-kubevious.sh \ 
    && chmod +x install-kubevious.sh \ 
    && ./install-kubevious.sh \
    && rm install-kubevious.sh

RUN groupadd --gid ${GID} ${GROUP} \
    && useradd --uid ${UID} --gid ${GID} --shell /bin/bash --create-home ${USER}

USER "$USER"

ENTRYPOINT ["/entrypoint.sh"]