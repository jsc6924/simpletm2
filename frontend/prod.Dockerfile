# Use the official Caddy image
FROM caddy:2.8.4-alpine

# Copy the build output to the Caddy HTML directory
COPY dist /usr/share/caddy

# Copy the Caddyfile to the Caddy configuration directory
COPY Caddyfile /etc/caddy/Caddyfile

# Expose port 80 to the outside world
EXPOSE 80

# By default, Caddy serves files from /usr/share/caddy
# You can also include a Caddyfile to customize the configuration