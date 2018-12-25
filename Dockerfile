FROM golang:1.11 as backend
RUN mkdir -p /go/src/github.com/shiguanghuxian/etcd-manage
ADD . /go/src/github.com/shiguanghuxian/etcd-manage
WORKDIR /go/src/github.com/shiguanghuxian/etcd-manage
RUN CGO_ENABLED=0 go build

# FROM node:8 as frontend
# RUN mkdir /app
# ADD static /app
# WORKDIR /app
# RUN npm install -g cnpm  --registry=https://registry.npm.taobao.org
# RUN cnpm install -g @vue/cli
# RUN cnpm install && npm run build

FROM alpine:latest
RUN mkdir -p /app/static/dist /app/conf
COPY --from=backend /go/src/github.com/shiguanghuxian/etcd-manage/etcd-manage /app
# COPY --from=frontend /app/dist /app/static/dist
COPY static/dist /app/static/dist
COPY conf/config.default.ini /app/conf
EXPOSE 9090
WORKDIR /app
CMD ["./etcd-manage"]
