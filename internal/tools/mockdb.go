package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"noah": {
		AuthToken: "root",
		Username:  "noah",
	},
	"sander": {
		AuthToken: "crazylongpasswordthateveniwouldforgetit",
		Username:  "sander",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"noah": {
		Coins:    100,
		Username: "noah",
	},
	"sander": {
		Coins:    727,
		Username: "sander",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
