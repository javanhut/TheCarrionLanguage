# Use the official Go image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Carrion source code into the container
COPY . .

# Run the setup script
RUN chmod +x setup.sh
CMD ["./setup.sh"]
