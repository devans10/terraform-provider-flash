FROM ubuntu:latest as STAGEONE

# install hugo
ENV HUGO_VERSION=0.54.0
ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz /tmp/
RUN tar -xf /tmp/hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz -C /usr/local/bin/

RUN apt-get update
RUN apt-get install -y python3-pygments

# build site
COPY src /source
RUN hugo --source=/source/ --destination=/public/

FROM nginx:stable-alpine
COPY --from=STAGEONE /public/ /usr/share/nginx/html/
EXPOSE 80
