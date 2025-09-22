# Makefile updates for Nginx
# Add these to your existing Makefile

# Nginx specific commands
.PHONY: nginx-setup nginx-ssl nginx-test nginx-reload nginx-logs

# Setup Nginx configuration
nginx-setup:
	@echo "🌐 Setting up Nginx configuration..."
	mkdir -p deployments/docker/nginx/{conf.d,ssl,logs,html}
	@echo "✅ Nginx directories created"

# Generate SSL certificates for development
nginx-ssl:
	@echo "🔒 Generating SSL certificates..."
	chmod +x scripts/generate-ssl-certs.sh
	./scripts/generate-ssl-certs.sh

# Test Nginx configuration
nginx-test:
	@echo "🧪 Testing Nginx configuration..."
	cd deployments/docker && docker-compose exec nginx nginx -t

# Reload Nginx configuration
nginx-reload:
	@echo "🔄 Reloading Nginx configuration..."
	cd deployments/docker && docker-compose exec nginx nginx -s reload

# Show Nginx logs
nginx-logs:
	@echo "📜 Showing Nginx logs..."
	cd deployments/docker && docker-compose logs -f nginx

# Setup hosts file entries (requires sudo)
nginx-hosts:
	@echo "📝 Adding entries to /etc/hosts (requires sudo)..."
	@echo "127.0.0.1 andino.local" | sudo tee -a /etc/hosts
	@echo "127.0.0.1 api.andino.local" | sudo tee -a /etc/hosts
	@echo "127.0.0.1 docs.andino.local" | sudo tee -a /etc/hosts
	@echo "127.0.0.1 swagger.andino.local" | sudo tee -a /etc/hosts
	@echo "✅ Hosts entries added"

# Full Nginx setup
nginx-init: nginx-setup nginx-ssl
	@echo "✅ Nginx setup complete!"
	@echo ""
	@echo "Next steps:"
	@echo "1. Run 'make nginx-hosts' to setup local domains"
	@echo "2. Run 'make docker-up-proxy' to start with Nginx"
	@echo "3. Visit https://andino.local"

# Update help command
help:
	@echo "🏔️  Andino - Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  run           - Run development server"
	@echo "  dev           - Run with hot reload (air)"
	@echo "  build         - Build binary"
	@echo "  test          - Run tests"
	@echo "  clean         - Clean build artifacts"
	@echo ""
	@echo "Docker:"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-up     - Start Docker containers"
	@echo "  docker-stop   - Stop Docker containers"
	@echo "  docker-logs   - Show Docker logs"
	@echo "  docker-up-tools - Start with management tools"
	@echo "  docker-up-proxy - Start with Nginx proxy"
	@echo "  docker-up-full  - Start full stack"
	@echo ""
	@echo "Nginx:"
	@echo "  nginx-init    - Complete Nginx setup"
	@echo "  nginx-ssl     - Generate SSL certificates"
	@echo "  nginx-test    - Test Nginx config"
	@echo "  nginx-reload  - Reload Nginx config"
	@echo "  nginx-hosts   - Setup local domains"
	@echo ""
	@echo "Environment:"
	@echo "  env-setup     - Setup .env file"
	@echo "  swagger-gen   - Generate Swagger docs"
	@echo "  install-tools - Install dev tools"
	@echo "  tidy          - Tidy dependencies"
