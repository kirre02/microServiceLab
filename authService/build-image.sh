#!/bin/bash

# Check if Maven wrapper exists 
if [ ! -f "./mvnw" ]; then
    echo "Maven wrapper (mvnw) saknas. Skapar wrapper..."
    mvn wrapper:wrapper
fi

# Grant execution permissions to mvnw if needed
chmod +x ./mvnw

# Build Docker image using Spring Boot's buildpack support
./mvnw spring-boot:build-image -Dspring-boot.build-image.imageName=auth-service:latest

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "Docker image built successfully!"
    echo "You can now run the container using:"
    echo "docker run -p 9090:9090 auth-service:latest"
else
    echo "build failed"
    exit 1
fi

