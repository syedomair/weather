FROM golang:1.8.3 AS build-env

ADD . /src 
 
WORKDIR /app

COPY --from=build-env /src/build/weather /app/
RUN cd /app && mkdir config
COPY --from=build-env /src/config/config_test.yml /app/config

ENTRYPOINT ./weather -config="config/config_test.yml"
 
EXPOSE 8088
