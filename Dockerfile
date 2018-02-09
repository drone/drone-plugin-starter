FROM alpine:3.2

ARG GITSHA="VCS ref not found"
ARG BUILDDATE="Build date not found"
LABEL org.label-schema.url="https://github.com/gjtempleton/drone-pypi" \
      org.label-schema.name="Drone-Pypi" \
      org.label-schema.license="Apache-2.0" \
      org.label-schema.vcs-url="https://github.com/gjtempleton/drone-pypi" \
      org.label-schema.schema-version="1.0" \
      org.label-schema.description="Simple Drone CI plugin written for and tested on version 0.8+ to publish pypi packages" \
      org.label-schema.vcs-ref=$GITSHA \
      org.label-schema.build-date=$BUILDDATE

RUN apk add -U \
	ca-certificates \
	python3 \
 && rm -rf /var/cache/apk/* \
 && pip3 install --no-cache-dir --upgrade \
	pip \
	setuptools \
	wheel

ADD drone-pypi /bin/
ENTRYPOINT ["/bin/drone-pypi"]
