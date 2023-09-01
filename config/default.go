package config

//nolint:lll
const Default = `
app:
  name: thanos
  env: debugging
  port: 8080
  graceful-shutdown: 1s
  cors-enabled: false
  request-timeout: 10s
  log-path: ./build/thanos.log
  schedule-lock-ttl: 10s

mysql:
  schema: thanos
  port: 3306
  host: mariadb
  username: maria
  password: VhZHNhM2RhZGRzYWRhc
  conn-max-lifetime: 1s
  max-idle-conns: 2
  max-open-conns: 0

authentication:
  access-expiration-in-minute : 2880
  refresh-expiration-in-minute : 5760
  jwt-secret: "Jwt3ecRETKEY9837"

zap-logger:
  compress: false
  max-age: 30
  max-size: 20
  max-backups: 60
  local-time: false
  log-path: ./logs/thanos.log

hash-id:
  min-length: 30

static:
  static-file-path: "/app/statics/"
  character-file-path: "/app/statics/characters/"
  image-base: "127.0.0.1:8000/static"

`
