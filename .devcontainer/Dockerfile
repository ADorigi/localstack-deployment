# Officail go image 
FROM golang:1.22

# Create a non-root user named vscode
RUN useradd -m vscode

# Install Terraform
ARG TERRAFORM_VERSION=1.9.4
RUN apt-get update && apt-get install -y wget unzip \
    && wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_arm64.zip \
    && unzip terraform_${TERRAFORM_VERSION}_linux_arm64.zip \
    && mv terraform /usr/local/bin/ \
    && rm terraform_${TERRAFORM_VERSION}_linux_arm64.zip \
    && apt-get clean

WORKDIR /workspace

# Debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Expose debugging port for Delve
EXPOSE 2345

# Default command
CMD ["sleep", "infinity"]