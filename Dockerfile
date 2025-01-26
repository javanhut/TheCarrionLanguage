# Use the official Go image
FROM golang:latest

# 1) Set environment variables explicitly
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# (Optional) Add a label to identify the image
LABEL name="CarrionLanguage"

# 2) Set our working directory
WORKDIR /app

# 3) Copy everything into /app
COPY . .

RUN mkdir carrion_files

COPY ./src/examples/ ./carrion_files

# 4) Make sure your scripts are executable
RUN chmod +x docker/docker-setup.sh docker/docker-install.sh

# 5) Debug step: confirm Go is on PATH
RUN go version

# 6) Use bash for all RUN commands, in case your script needs bash features
SHELL ["/bin/bash", "-c"]

# 7) Run docker-setup.sh, which should now detect Go
RUN ./docker/docker-setup.sh

# 8) Run Bash by default in the final container
CMD ["bash"]

