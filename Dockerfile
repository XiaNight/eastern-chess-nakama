FROM heroiclabs/nakama-pluginbuilder:3.11.0 AS builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /backend
COPY . .

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

FROM heroiclabs/nakama:3.11.0

COPY --from=builder /backend/backend.so /nakama/data/modules
COPY --from=builder /backend/*.yml /nakama/data/
