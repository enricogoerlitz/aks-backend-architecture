# Basis-Image mit Python 3.12
FROM python:3.12

# Installiere Abhängigkeiten zum Ausführen von Flyway & Python
RUN apt-get update && \
    apt-get install -y curl unzip default-jre && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Flyway installieren
ENV FLYWAY_VERSION=10.18.0
RUN curl -L "https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}-linux-x64.tar.gz" \
    -o flyway.tar.gz && \
    tar -xzf flyway.tar.gz && \
    mv flyway-${FLYWAY_VERSION} /opt/flyway && \
    ln -s /opt/flyway/flyway /usr/local/bin/flyway && \
    rm flyway.tar.gz

# Set working directory
WORKDIR /app

# Python requirements installieren
COPY scripts/requirements.txt ./scripts/requirements.txt
RUN pip install --no-cache-dir -r scripts/requirements.txt

# Restliche App kopieren (Passe ggf. Pfadstruktur an)
COPY . .

# Standard-EntryPoint
# ENTRYPOINT ["python", "scripts/flyway.py", "migrate"]
