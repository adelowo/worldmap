FROM golang:1.19 as build-env
WORKDIR /go/src/github.com/adelowo/worldmap
#
# Create a netrc file using the credentials specified using --build-arg

# ARG ACCESS_TOKEN="nothing"
# ARG ACCESS_TOKEN_USR="nothing"

# RUN printf "machine github.com\n\
#     login ${ACCESS_TOKEN_USR}\n\
#     password ${ACCESS_TOKEN}\n\
#     \n\
#     machine api.github.com\n\
#     login ${ACCESS_TOKEN_USR}\n\
#     password ${ACCESS_TOKEN}\n"\
#     >> /root/.netrc
# RUN chmod 600 /root/.netrc

COPY ./go.mod /go/src/github.com/adelowo/worldmap
COPY ./go.sum /go/src/github.com/adelowo/worldmap

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
RUN go mod verify
# COPY the source code as the last step
COPY . .


RUN CGO_ENABLED=0
RUN go install ./cmd

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/cmd /
COPY --from=build-env /go/src/github.com/adelowo/worldmap/testdata/map.txt /
ENV MAP_FILE="map.txt"
ENV ALIEN_COUNT=1000
CMD ["/cmd"]

