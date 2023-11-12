# Automated ETH Token Salary Distribution System

This project is an automated salary distribution system using Ethereum (ETH) tokens. It operates on Infura's Ethereum test networks and utilizes the Go-Ethereum API for blockchain interactions.

## Getting Started

### Prerequisites
- Metamask wallet account.
- Access to Ethereum test networks.

### Initial Setup

#### Set Up Metamask Wallet
1. Create an account in the Metamask wallet.
2. Acquire free ETH for testing purposes on any Ethereum test network.

#### Acquire Free ETH for Your Wallet
- For Sepolia test network, use the faucet: [Sepolia Faucet](https://sepolia-faucet.pk910.de/)
- For Goerli test network, use the faucet: [Goerli Faucet](https://goerli-faucet.pk910.de/)

#### Create a Project on Infura
- Register and create a new project on Infura to access Ethereum test networks and APIs. [Infura](https://infura.io/)

### Installation and Configuration

#### Clone the Repository
Clone the repository using:
```git clone git@github.com:sherpalden/sal-dist.git```


#### Environment Setup
Create a `.env` file in the project root and initialize the following variables:
- `NETWORK_URL`: HTTP endpoint URL of an Ethereum test network.
- `OWNER_PRIVATE_KEY`: Your Metamask wallet's private key.
- `OWNER_PUBLIC_ADDRESS`: Your Metamask wallet's public address.

#### Install Dependencies
Run the following command to download the necessary Go modules:
```go mod download```


## Running the Application

Follow these steps to run the application:

1. **Create Wallet**:
```go run cmd/wallet/wallet.go passphrase=password keypath=./wallet/keystores/employees```


2. **Create Employee**:
```go run cmd/employee/employee.go -id="123" -name="John Doe" -walletaddress="f8fe52bf219d1ec036e7d1c7f29691775b8f7eb9" -salary="50000"```


3. **Balance Inquiry**:
```go run cmd/balance/balance.go```


### Additional Resources

- Tutorial on obtaining free ETH for test networks: [YouTube](https://youtu.be/Ni3XocoNaDI)
- Project tutorial reference video: [YouTube](https://youtu.be/EB0KkSkG5XU?list=PLay9kDOVd_x7hbhssw4pTKZHzzc6OG0e_)
