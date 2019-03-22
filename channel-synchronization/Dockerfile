FROM golang:1.8
ENV BAR production
# Create app directory
WORKDIR /usr/src/app
COPY . .
CMD ["go", "run", "app.go"]