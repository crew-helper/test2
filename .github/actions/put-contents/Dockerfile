FROM alpine/git:latest
RUN mkdir /ghcli
RUN wget http://dl-cdn.alpinelinux.org/alpine/edge/community/x86_64/github-cli-1.3.1-r0.apk
RUN apk update
RUN apk add --allow-untrusted ghcli/github-cli-1.3.1-r0.apk

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY entrypoint.sh /home/entrypoint.sh
RUN chmod +x /home/entrypoint.sh
# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/home/entrypoint.sh"]
