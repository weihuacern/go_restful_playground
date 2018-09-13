FROM alpine:3.8

# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf && mkdir /tmp/authserver

ENV GOLANG_VERSION 1.10.2

# make-sure-R0-is-zero-before-main-on-ppc64le.patch: https://github.com/golang/go/commit/9aea0e89b6df032c29d0add8d69ba2c95f1106d9 (Go 1.9)
COPY . /tmp/authserver/.
#COPY *.patch /go-alpine-patches/
RUN set -eux; \
    apk add --no-cache \
        bash \
        gcc \
        musl \
        musl-dev \
        openssl \
        go \
        make \
        git \
        linux-pam  \
        linux-pam-dev \
    ; \
    export \
# set GOROOT_BOOTSTRAP such that we can actually build Go
        GOROOT_BOOTSTRAP="$(go env GOROOT)" \
# ... and set "cross-building" related vars to the installed system's values so that we create a build targeting the proper arch
# (for example, if our build host is GOARCH=amd64, but our build env/image is GOARCH=386, our build needs GOARCH=386)
        GOOS="$(go env GOOS)" \
        GOARCH="$(go env GOARCH)" \
        GOHOSTOS="$(go env GOHOSTOS)" \
        GOHOSTARCH="$(go env GOHOSTARCH)" \
    ; \
# also explicitly set GO386 and GOARM if appropriate
# https://github.com/docker-library/golang/issues/184
    apkArch="$(apk --print-arch)"; \
    case "$apkArch" in \
        armhf) export GOARM='6' ;; \
        x86) export GO386='387' ;; \
    esac; \
    \
    wget -O go.tgz "https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz"; \
    echo '6264609c6b9cd8ed8e02ca84605d727ce1898d74efa79841660b2e3e985a98bd *go.tgz' | sha256sum -c -; \
    tar -C /usr/local -xzf go.tgz; \
    rm go.tgz; \
    \
    cd /usr/local/go/src; \
    for p in /go-alpine-patches/*.patch; do \
        [ -f "$p" ] || continue; \
        patch -p2 -i "$p"; \
    done; \
    ./make.bash; \
    export PATH="/usr/local/go/bin:/root/go/bin:$PATH"; \
    go version && \
    go get -u golang.org/x/lint/golint && \
    go get -u github.com/gin-gonic/gin && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/mattn/go-sqlite3 && \
    go get -u github.com/lib/pq && \
    go get -u github.com/satori/go.uuid && \
    go get -u github.com/dgrijalva/jwt-go && \
    go get -u github.com/tsenart/vegeta && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \
    mkdir -p /opt/helios/bin &&\
    cd /tmp/authserver && sh create-binary.sh && cp auth-server-binary /opt/helios/bin/ && cd / && rm -rf /tmp/authserver\
    rm -rf /go-alpine-patches;

#dump helios_auth pam module
RUN touch /etc/pam.d/helios_auth && \
    cat /dev/null > /etc/pam.d/helios_auth && \
    echo "# check authorization" >> /etc/pam.d/helios_auth && \
    echo "auth       required     pam_unix.so" >> /etc/pam.d/helios_auth && \
    echo "account    required     pam_unix.so" >> /etc/pam.d/helios_auth && \
    echo "password   requisite    pam_pwquality.so try_first_pass local_users_only retry=3 authtok_type=" >> /etc/pam.d/helios_auth && \
    echo "password   sufficient   pam_unix.so sha512 shadow nullok try_first_pass use_authtok" >> /etc/pam.d/helios_auth && \
    echo "password   required     pam_deny.so" >> /etc/pam.d/helios_auth && \
    echo "session    required     pam_unix.so" >> /etc/pam.d/helios_auth;

USER root
RUN chmod -R 777 /opt/helios/bin
WORKDIR /opt/helios/bin
COPY docker-entrypoint.sh /
RUN chmod -R 777 /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
