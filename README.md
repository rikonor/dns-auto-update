DNS Auto-Update
---

This tool can be used to automatically update DNS settings.  
It fetches the current public IP from [ifconfig.co](ifconfig.co), and compares it to the currently set DNS IP on the selected provider.  
If a discrepancy is found, it will update the DNS setting.

### Supported providers:
* GoDaddy

### Usage

```
GODADDY_API_KEY=<API_KEY> \
GODADDY_API_SECRET=<API_SECRET> \
dns-auto-update -d "knickknacklabs.com -p "godaddy"
```

**Feel free to add additional providers.**

### Ideas:
* Consul
* Dropbox
