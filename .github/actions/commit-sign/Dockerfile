FROM alpine/git:latest
RUN mkdir /ghcli
WORKDIR /ghcli
RUN wget https://github.com/cli/cli/releases/download/v1.4.0/gh_1.4.0_linux_386.tar.gz -O ghcli.tar.gz
RUN tar --strip-components=1 -xf ghcli.tar.gz

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY entrypoint.sh /home/entrypoint.sh
RUN chmod +x /home/entrypoint.sh
# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/home/entrypoint.sh"]
