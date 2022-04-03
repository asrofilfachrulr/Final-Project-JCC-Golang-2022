package models

var UsersDummy = []User{
	{FullName: "John Doe", Username: "john", Email: "john@mail.com"},
	{FullName: "Mary Sue", Username: "mary", Email: "mary@mail.com"},
	{FullName: "Xi Ng", Username: "xi", Email: "nihaoma@mail.com"},
	{FullName: "Mark Bob", Username: "mark", Email: "mark@mail.com"},
	{FullName: "Patricia Ng", Username: "ng", Email: "ng@mail.com"},
	{FullName: "Jack Tyler", Username: "jack", Email: "jack@mail.com"},
	{FullName: "Tony Like", Username: "tony", Email: "tony@mail.com"},
	{FullName: "Andy Lim", Username: "andy", Email: "andy@mail.com"},
}

var Merchants = []Merchant{
	{Name: "Jaya Store", Rating: 5},
	{Name: "Sinar Muda", Rating: 4.9},
	{Name: "Java Net Tech", Rating: 4.5},
}

var MerchantAddresses = []MerchantAddress{
	{OfflineStoreAddress: "Jl. Sukarno Hatta 235", City: "Bandung", CountryID: 1},
	{OfflineStoreAddress: "Jl. Kalveri 120", City: "Jakarta", CountryID: 1},
	{OfflineStoreAddress: "Jl. Pattimura 32", City: "Surabaya", CountryID: 1},
}

var Products = []Product{
	{Name: "Indomi Sedap", Desc: "Mie goreng pilihan nomer #1 di Indonesia", Price: 2700, CategoryID: 1},
	{Name: "Indomi Kuah Ayam", Desc: "Menggunakan kaldu ayam asli, Indomi Kuah Ayam siap memulai aktifitas kamu agar semakin berwarna ", Price: 2900, CategoryID: 1},
	{Name: "Levono Thinklad 260x", Desc: "Second like new Ex-Singapore, mulus 98.99%. i5-6200u, RAM DDR4 8GB, SSD SATA 256GB", Price: 3900000, CategoryID: 5},
	{Name: "Oxadon Oye", Desc: "Meredakan gejala flu dan sakil kepala ringan", Price: 1500, CategoryID: 10},
	{Name: "Sumsang Ultra Max 12", Desc: "Six cameras, 16GB RAM, 5000mAh Battery , and SnapNaga gen 1", Price: 12000000, CategoryID: 4},
	{Name: "Kukira kau home", Desc: "Novel best seller dari Mamank Garox ke-12, Menceritakan tentang pahitnya minum obat.", Price: 54000, CategoryID: 12},
}
