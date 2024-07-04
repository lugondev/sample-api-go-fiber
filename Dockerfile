# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/* \

WORKDIR /app
# Copy the binary to the production image from the builder stage.
#COPY ./.env ./.env
COPY ./build/bin ./bin

#RUN ls -alh .
# Run the web service on container startup.

CMD ["./bin"]
