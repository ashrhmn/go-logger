FROM node:20 as ui-builder
WORKDIR /app
RUN npm install -g pnpm
COPY client/package.json client/pnpm-lock.yaml ./
RUN pnpm install
COPY client .
RUN pnpm run build

FROM golang as app-builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
ENV GO_ENV=production
RUN CGO_ENABLED=0 go build -o /main

FROM scratch
COPY --from=app-builder /main /main
COPY --from=ui-builder /app/dist /views
EXPOSE 8080
ENV GO_ENV=production
ENV MONGO_URI="mongodb://host.docker.internal:27017/test?w=majority"
ENV RMQ_URL="amqp://guest:guest@host.docker.internal:5672/"
ENTRYPOINT [ "/main" ]