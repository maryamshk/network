package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract for managing IOCs (Indicators of Compromise)
type SmartContract struct {
	contractapi.Contract
}

// IOC represents a submitted Indicator of Compromise
type IOC struct {
	ID        string `json:"id"`        // Unique IOC ID
	Reporter  string `json:"reporter"`  // Organization that reported the IOC
	Details   string `json:"details"`   // Description of the threat
	IPFSHash  string `json:"ipfsHash"`  // IPFS Link for additional threat data
	Validated bool   `json:"validated"` // Status of IOC validation
	Validator string `json:"validator"` // Validator who approved/rejected the IOC
}

// SubmitIOC allows organizations to report new threats
func (s *SmartContract) SubmitIOC(ctx contractapi.TransactionContextInterface, id string, reporter string, details string, ipfsHash string) error {
	// Check if the IOC already exists
	existingIOC, _ := ctx.GetStub().GetState(id)
	if existingIOC != nil {
		return fmt.Errorf("IOC with ID %s already exists", id)
	}

	// Create a new IOC entry
	ioc := IOC{
		ID:        id,
		Reporter:  reporter,
		Details:   details,
		IPFSHash:  ipfsHash,
		Validated: false,
		Validator: "",
	}

	iocAsBytes, _ := json.Marshal(ioc)
	return ctx.GetStub().PutState(id, iocAsBytes)
}

// ValidateIOC allows authorized validators to approve threats
func (s *SmartContract) ValidateIOC(ctx contractapi.TransactionContextInterface, id string, validator string) error {
	iocAsBytes, err := ctx.GetStub().GetState(id)
	if err != nil || iocAsBytes == nil {
		return fmt.Errorf("IOC with ID %s not found", id)
	}

	var ioc IOC
	_ = json.Unmarshal(iocAsBytes, &ioc)

	// Ensure the IOC isn't already validated
	if ioc.Validated {
		return fmt.Errorf("IOC %s has already been validated", id)
	}

	// Update validation status
	ioc.Validated = true
	ioc.Validator = validator

	updatedIOCAsBytes, _ := json.Marshal(ioc)
	return ctx.GetStub().PutState(id, updatedIOCAsBytes)
}

// QueryIOC fetches a specific IOC by ID
func (s *SmartContract) QueryIOC(ctx contractapi.TransactionContextInterface, id string) (*IOC, error) {
	iocAsBytes, err := ctx.GetStub().GetState(id)
	if err != nil || iocAsBytes == nil {
		return nil, fmt.Errorf("IOC with ID %s not found", id)
	}

	var ioc IOC
	_ = json.Unmarshal(iocAsBytes, &ioc)
	return &ioc, nil
}

// QueryAllIOCs fetches all IOCs in the ledger
func (s *SmartContract) QueryAllIOCs(ctx contractapi.TransactionContextInterface) ([]*IOC, error) {
	iterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	var iocs []*IOC
	for iterator.HasNext() {
		queryResponse, err := iterator.Next()
		if err != nil {
			return nil, err
		}

		var ioc IOC
		_ = json.Unmarshal(queryResponse.Value, &ioc)
		iocs = append(iocs, &ioc)
	}

	return iocs, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating smart contract: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}
}
