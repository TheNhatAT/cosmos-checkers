# This Dockerfile will: 
# set up a Golang environment(using Golang Docker Image)
# install Ignite
# clone the checkers game repository
# start a checkers blockchain
# expose api endpoints

# Use Golang Docker Image
FROM golang:1.18

# Set working dir
WORKDIR /home/checkers

# Install tools
RUN apt-get install curl git

# Install ignite
RUN curl -L https://get.ignite.com/cli@v0.22.1! | bash

# Copy checker from local to container
COPY . /home/checkers

# Copy Config.yml
ARG configyml
COPY ${configyml} /home/checkers/config.yml

# copy entrypoint script
COPY ./run-checkers.sh /home/checkers
RUN chmod +x ./run-checkers.sh

ENTRYPOINT ["/bin/bash", "-c", "/home/checkers/run-checkers.sh"]
EXPOSE 26657 1317 4500

# build with ./build-images.sh
