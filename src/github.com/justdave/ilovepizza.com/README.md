As a user I need an online address book exposed as a REST API. I need the data set to include
the following data fields:

First Name, Last Name, Email Address, and Phone Number

I need the api to follow standard rest semantics to support listing entries, showing a specific
single entry, and adding, modifying, and deleting entries.

The code for the address book should include regular go test files that demonstrate how to
exercise all operations of the service.

Finally I need the service to provide endpoints that can export and import the address book
data in a CSV format.

To run this program type go run main.go. To test the functionality use Postman and point the URL to localhost:3000.

	router.HandleFunc("/customer", readAll).Methods("GET")
	router.HandleFunc("/customer/{phone}", read).Methods("GET")
	router.HandleFunc("/customer/{phone}", create).Methods("PUT")
	router.HandleFunc("/customer/{phone}", delete).Methods("DELETE")
	router.HandleFunc("/export", dumpcsv).Methods("GET")
	router.HandleFunc("/import", grabcsv).Methods("GET")
  
  Type the URL of the desired function. 
  
  Demo Data is formatted like so: {Firstname, Lastname, Email, Phone} 
  fghj,fg,hj,1700
  rtyu,rt,yu,1600
  asdf,as,df,1800
  qwer,qw,er,1900
  zxcv,zx,cv,2000
