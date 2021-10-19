FROM docker.io/library/node:16.11.1-alpine AS frontend
WORKDIR /work

COPY frontend/package.json frontend/yarn.lock frontend/.yarnclean ./
RUN yarn install --frozen-lockfile --non-interactive

COPY frontend/ ./
RUN  yarn run build


FROM docker.io/library/golang:1.16.2 AS backend
WORKDIR /work
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN make kroki


FROM docker.io/busybox:latest
ENV PORT=8443 \
    TLS_CERT=/tls/cert.pem \
    TLS_KEY=/tls/key.pem
CMD [ "/kroki" ]
EXPOSE $PORT

COPY --from=frontend /work/public /public
COPY --from=backend /work/kroki /kroki
USER 10100111001
