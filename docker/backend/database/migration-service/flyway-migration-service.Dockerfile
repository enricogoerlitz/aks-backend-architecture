# Basis-Image mit Python 3.12 + OpenJDK (ARM64-kompatibel)
FROM python:3.12-slim

RUN apt-get update && \
    apt-get install -y curl unzip openjdk-17-jre-headless && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Flyway JAR runterladen
ENV FLYWAY_VERSION=10.18.0
RUN mkdir -p /opt/flyway && \
    curl -L "https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}.tar.gz" \
    -o flyway.tar.gz && \
    tar -xzf flyway.tar.gz && \
    mv flyway-${FLYWAY_VERSION}/* /opt/flyway && \
    ln -s /opt/flyway/flyway /usr/local/bin/flyway && \
    rm -rf flyway.tar.gz

# Set working directory
WORKDIR /app

COPY scripts/requirements.txt ./scripts/requirements.txt
RUN pip install --no-cache-dir -r scripts/requirements.txt

COPY . .

# ENTRYPOINT ["python", "scripts/flyway.py", "migrate"]
