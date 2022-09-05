# go-calipower
Get available capacity, estimates, and reserves from California ISO.

## Usage
```go
data, err := getPowerData() 
if err != nil {
	...
}
fmt.Printf("Current demand in California: %s", data.CurrentDemand)
```

## Acknowledgements
Data is collected from FlexAlert.org, a website run by the California ISO. Data is meant to be used within the allowance of the California ISO.
