FROM rikonor/scratch

WORKDIR /app
COPY dns-auto-update /app

ENTRYPOINT ["/app/dns-auto-update"]

