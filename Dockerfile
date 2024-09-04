FROM golang:1.22.4

# Install git
RUN apt-get update && apt-get install -y git

# Set the working directory
WORKDIR /api

# Clone the repository
RUN git clone https://github.com/IrvinTM/urlBit.git --depth 1 .

# Build the Go application
RUN go build -o urlBit .

# Command to run the application
CMD ["./urlBit"]