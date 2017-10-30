# cisco_csr1000v_restApi_lib

Example of usage:

```
import csr_lib

auth = csr_lib.AuthenticationClass('cisco','cisco')
conn = csr_lib.ConnectionClass('192.168.35.11','55443')
token = auth.get_authentication_token(conn)


glob = csr_lib.GlobalClass(conn,auth)
banners = glob.get_banner()

print banners
```
