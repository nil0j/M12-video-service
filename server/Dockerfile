FROM golang:1.24
WORKDIR /program

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# COPY init.sql .

EXPOSE 8080
CMD ["go", "run", "./"]

# CMD ["/program/jirafeitor"]

