FROM alpine:latest

# Set your working directory (optional)
WORKDIR /app

# Install any additional packages (replace with your needs)
COPY application .
COPY dev.yaml .
RUN chmod +x application

# Define the command to run your application
CMD [ "/app/application" ]