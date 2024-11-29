package main

import (
	"fmt"
)

var username, pass, role, productName string
var productID, qty int
var price float64

type User struct {
	ID       int
	Username string
	Password string
	Role     string
	Approved bool
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Transaction struct {
	ID       int
	Buyer    User
	Product  Product
	Quantity int
	Total    float64
}

type Store struct {
	Owners       []User
	Products     []Product
	Users        []User
	Transactions []Transaction
}

func main() {
	var choice string
	var store Store

	// Inisialisasi User toko
	store.Users = append(store.Users, User{ID: 1, Username: "admin", Password: "admin", Role: "admin", Approved: true})

	for {
		fmt.Println("-Welcome to Onlenkeun-")
		fmt.Println("-----------------------------")
		fmt.Println("Belanja lebih mudah dengan Onlenkeun!")
		fmt.Println("-----------------------------")
		fmt.Println("Menu: ")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("-----------------------------")
		fmt.Print("Pilih Menu: ")

		fmt.Scan(&choice)
		switch choice {
		case "1":
			// Register user
			fmt.Print("Masukkan Username yang diinginkan: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan Password yang dinginkan: ")
			fmt.Scan(&pass)
			fmt.Print("Masukkan role anda (buyer/owner): ")
			fmt.Scan(&role)
			store.RegisterUser(username, pass, role)
		case "2":
			// Login user
			fmt.Print("Masukkan Username anda: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan Password anda: ")
			fmt.Scan(&pass)
			user, found := store.LoginUser(username, pass)
			if found {
				store.UserMenu(user)
			} else {
				fmt.Println("Pastikan Username dan Password anda benar serta akun telah di approve admin.")
			}
		case "3":
			fmt.Println("Sampai jumpa lagi!!!!")
			return
		default:
			fmt.Println("Invalid.")
		}
	}
}

func (s *Store) RegisterUser(username, password, role string) {
	id := len(s.Users) + 1
	user := User{ID: id, Username: username, Password: password, Role: role, Approved: false}
	s.Users = append(s.Users, user)
	fmt.Println("User", username, "registered. Harap menunggu persetujuan admin:)")
}

func (s *Store) LoginUser(username, password string) (User, bool) {
	for _, user := range s.Users {
		if user.Username == username && user.Password == password && user.Approved {
			return user, true
		}
	}
	return User{}, false
}

func (s *Store) UserMenu(user User) {
	switch user.Role {
	case "admin":
		s.AdminMenu()
	case "owner":
		s.OwnerMenu(user)
	case "buyer":
		s.BuyerMenu(user)
	}
}

func (s *Store) AdminMenu() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Check Registered user")
		fmt.Println("2. Approve User")
		fmt.Println("3. Logout")
		fmt.Print("Pilih menu: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			s.AdminListUsers()
		case 2:
			fmt.Print("Masukkan ID user yang ingin di approve: ")
			var userID int
			fmt.Scan(&userID)
			s.ApproveUser(userID)
		case 3:
			fmt.Println("Bye Admin:)")
			return
		default:
			fmt.Println("Invalid.")
		}
	}
}

func (s *Store) OwnerMenu(owner User) {
	for {
		fmt.Println("Owner Menu:")
		fmt.Println("1. Add Product")
		fmt.Println("2. Edit Product")
		fmt.Println("3. Delete Product")
		fmt.Println("4. View Products")
		fmt.Println("5. View Transactions")
		fmt.Println("6. Logout")
		fmt.Print("Pilih menu: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan nama produk: ")
			fmt.Scan(&productName)
			fmt.Print("Masukkan harga produk: ")
			fmt.Scan(&price)
			s.AddProduct(productName, price)
		case 2:
			fmt.Print("Masukkan ID produk yang ingin di edit: ")
			fmt.Scan(&productID)
			fmt.Print("Masukkan nama produk yang baru: ")
			fmt.Scan(&productName)
			fmt.Print("Masukkan harga produk yang baru: ")
			fmt.Scan(&price)
			s.EditProduct(productID, productName, price)
		case 3:
			fmt.Print("Masukkan ID produk yang ingin dihapus: ")
			fmt.Scan(&productID)
			s.DeleteProduct(productID)
		case 4:
			s.ListProducts()
		case 5:
			s.ListTransactions()
		case 6:
			fmt.Println("Sampai jumpa lagi!!!")
			return
		default:
			fmt.Println("Invalid.")
		}
	}
}

func (s *Store) BuyerMenu(buyer User) {
	for {
		fmt.Println("Buyer Menu:")
		fmt.Println("1. View Products")
		fmt.Println("2. Buy Product")
		fmt.Println("3. Logout")
		fmt.Print("Pilih menu: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			s.ListProducts()
		case 2:
			fmt.Print("Masukkan ID produk yang ingin di checkout: ")
			fmt.Scan(&productID)
			fmt.Print("Masukkan jumlah yang diinginkan: ")
			fmt.Scan(&qty)
			s.BuyProduct(buyer, productID, qty)
		case 3:
			fmt.Println("Sampai jumpa lagi!!!")
			return
		default:
			fmt.Println("Invalid.")
		}
	}
}

func (s *Store) ListProducts() {
	fmt.Println("Produk yang tersedia:")
	for _, product := range s.Products {
		fmt.Printf("ID: %d, Name: %s, Price: Rp.%.2f\n", product.ID, product.Name, product.Price)
	}
}

func (s *Store) AddProduct(name string, price float64) {
	id := len(s.Products) + 1
	product := Product{ID: id, Name: name, Price: price}
	s.Products = append(s.Products, product)
	fmt.Println("Product", name, "added.")
}

func (s *Store) EditProduct(productID int, name string, price float64) {
	for i, product := range s.Products {
		if product.ID == productID {
			s.Products[i].Name = name
			s.Products[i].Price = price
			fmt.Println("Product", name, "edited.")
			return
		}
	}
	fmt.Println("Produk tidak ditemukan.")
}

func (s *Store) DeleteProduct(productID int) {
	for i, product := range s.Products {
		if product.ID == productID {
			s.Products = append(s.Products[:i], s.Products[i+1:]...)
			fmt.Println("Product", product.Name, "deleted.")
			// Update IDs for remaining products
			for j := i; j < len(s.Products); j++ {
				s.Products[j].ID--
			}
			return

		}
	}
	fmt.Println("Produk tidak ditemukan.")
}

func (s *Store) BuyProduct(buyer User, productID, quantity int) {
	for _, product := range s.Products {
		if product.ID == productID {
			total := product.Price * float64(quantity)
			transactionID := len(s.Transactions) + 1
			transaction := Transaction{
				ID:       transactionID,
				Buyer:    buyer,
				Product:  product,
				Quantity: quantity,
				Total:    total,
			}
			s.Transactions = append(s.Transactions, transaction)
			fmt.Println("Transaction successful.")
			fmt.Println("Terima Kasih telah berbelanja di Onlenkeun!")
			return
		}
	}
	fmt.Println("Produk tidak ditemukan.")
}

func (s *Store) ListTransactions() {
	fmt.Println("Transactions:")
	for _, transaction := range s.Transactions {
		fmt.Printf("ID: %d, Buyer: %s, Product: %s, Quantity: %d, Total: %.2f\n",
			transaction.ID, transaction.Buyer.Username, transaction.Product.Name,
			transaction.Quantity, transaction.Total)
	}
}

func (s *Store) AdminListUsers() {
	fmt.Println("Registered Users:")
	for _, user := range s.Users {
		fmt.Printf("ID: %d, Username: %s, Role: %s, Approved: %t\n",
			user.ID, user.Username, user.Role, user.Approved)
	}
}

func (s *Store) ApproveUser(userID int) {
	for i, user := range s.Users {
		if user.ID == userID {
			s.Users[i].Approved = true
			fmt.Println("User", user.Username, "approved.")
			return
		}
	}
	fmt.Println("User tidak ditemukan.")
}
