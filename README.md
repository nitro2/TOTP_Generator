# OTP Generator (Time-based)
## Build
Run make:
```bash
make
```

## Run
`./genotp [YOUR SECRET KEY]`

Example output:
```bash
./genotp GU3TGMZUMRRGE
232470
```

## Best practice
### MacOS
Store your secrect into system:
```
security add-generic-password -a username -s yoursecret -w      
```

Retrieve the pass and pipe to OTP:
```
security find-generic-password -s yoursecret -w | xargs genotp
```
