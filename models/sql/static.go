package models

var AseanCountries = []Country{
	{Name: "Indonesia"},
	{Name: "Singapore"},
	{Name: "Malaysia"},
	{Name: "Thailand"},
	{Name: "Brunei"},
	{Name: "Philipines"},
	{Name: "Laos"},
	{Name: "Vietnam"},
	{Name: "Cambodia"},
	{Name: "Myanmar"},
}

var Categories = []Category{
	{Name: "Foodie"},
	{Name: "Beer"},
	{Name: "Beverage"},
	{Name: "Gadget"},
	{Name: "Laptop"},
	{Name: "Electronics"},
	{Name: "Men"},
	{Name: "Women"},
	{Name: "Outdoors"},
	{Name: "Health"},
	{Name: "Household"},
	{Name: "Books"},
	{Name: "Tools"},
}

var Payments = []Payment{
	{Name: "BJA", Method: "MBanking"},
	{Name: "BLI", Method: "MBanking"},
	{Name: "BMI", Method: "MBanking"},
	{Name: "Sendiri", Method: "MBanking"},
	{Name: "DANE", Method: "EWallet"},
	{Name: "OPO", Method: "EWallet"},
	{Name: "Gopey", Method: "EWallet"},
	{Name: "Shopipay", Method: "EWallet"},
}

var Shipments = []Shipment{
	{Name: "JNI", Method: "Intercity"},
	{Name: "JNP", Method: "Intercity"},
	{Name: "Bahana Express", Method: "Intercity"},
	{Name: "AnterAe", Method: "Intercity"},
	{Name: "Grap Express", Method: "Intracity"},
	{Name: "GoGo Send", Method: "Intracity"},
	{Name: "FedUp", Method: "International"},
}
