$TTL    3600 ; 1 hour
@       IN      SOA     localhost. localhost. (
                            1           ; Serial
                         3600           ; Refresh [1h]
                          600           ; Retry   [10m]
                        86400           ; Expire  [1d]
                          600 )         ; Negative Cache TTL [1h]

$TTL 0  ; 0 seconds
                        NS      cdn.site.com.
                        A       127.0.0.1