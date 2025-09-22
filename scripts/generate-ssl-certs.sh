#!/bin/bash

# Script to generate self-signed SSL certificates for development

SSL_DIR="./deployments/docker/nginx/ssl"
CERT_FILE="$SSL_DIR/nginx-selfsigned.crt"
KEY_FILE="$SSL_DIR/nginx-selfsigned.key"

echo "🔒 Generating self-signed SSL certificates for Andino development..."

# Create SSL directory if it doesn't exist
mkdir -p "$SSL_DIR"

# Generate private key and certificate
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
	-keyout "$KEY_FILE" \
	-out "$CERT_FILE" \
	-config <(
		echo '[dn]'
		echo 'CN=andino.local'
		echo '[req]'
		echo 'distinguished_name = dn'
		echo '[extensions]'
		echo 'subjectAltName = DNS:andino.local,DNS:api.andino.local,DNS:docs.andino.local,DNS:swagger.andino.local,DNS:localhost,IP:127.0.0.1'
		echo 'keyUsage = keyEncipherment,dataEncipherment'
		echo 'extendedKeyUsage = serverAuth'
	) \
	-extensions extensions

# Set proper permissions
chmod 600 "$KEY_FILE"
chmod 644 "$CERT_FILE"

echo "✅ SSL certificates generated successfully!"
echo "📄 Certificate: $CERT_FILE"
echo "🔑 Private key: $KEY_FILE"
echo ""
echo "📝 Add these entries to your /etc/hosts file:"
echo "127.0.0.1 andino.local"
echo "127.0.0.1 api.andino.local"
echo "127.0.0.1 docs.andino.local"
echo "127.0.0.1 swagger.andino.local"
