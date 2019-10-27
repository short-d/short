# Build Container
FROM byliuyang/go-builder:latest AS builder
COPY . .
# Install and verify dependencies
RUN go mod download \
	&& go mod verify \
	# Build binary without dwarf debug tables.
	# Also trim paths to provide more reproduceable builds.
	&& go build -trimpath -ldflags="-s -w" -o build/short main.go

# Build Artifact
FROM alpine:latest
ENV DB_HOST=postgres \
	DB_PORT=5432
RUN apk --no-cache add \
	ca-certificates \
	tzdata \
	bash
WORKDIR /app
COPY --from=builder /app/scripts scripts
COPY --from=builder /app/app/adapter/migration migration
COPY --from=builder /app/build/short .

CMD [ "sh", "-c", "./scripts/wait-for-it -s -t 0 $DB_HOST:$DB_PORT -- ./short start --migration migration" ]
