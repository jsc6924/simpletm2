# copy conf file to /etc/nginx/sites-available
```
sudo cp nginx/sites-available/* /etc/nginx/sites-available
sudo ln -s /etc/nginx/sites-available/simpletm2.conf /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx

sudo certbot --nginx -d simpletm2.jscrosoft.com -d simpletm-backend.jscrosoft.com
sudo certbot renew --dry-run
```
