FROM golang:1.14.2

# setting project direactory
WORKDIR /app

# copy all files under porject directory
COPY . .

# build go files
RUN go build -o main .

# run built files
CMD ["./main"]