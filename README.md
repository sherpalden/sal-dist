### Automated ETH Token Salary Distribution System
## Steps to run the application
- Create wallet
    ```go run cmd/wallet/wallet.go passphrase=password keypath=./wallet/keystores/employees```
- Create employee
    ```go run cmd/employee/employee.go -id="123" -name="John Doe" -walletaddress="f8fe52bf219d1ec036e7d1c7f29691775b8f7eb9" -salary="50000"```

- Balance inquiry
    ```go run cmd/balance/balance.go```

- Run server to distribute salary
    ```go run cmd/main.go```