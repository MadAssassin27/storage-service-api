<h3 align="center">storage-service-api</h3>

  <p align="center">
    API for authorization, upload file, and file listing
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project

API for authorization, upload file (image and video), and file listing with Go

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

Before we start to run or test the API, first we need to know how to install it.

### Installation

1. Install Go
2. Clone the repo
   ```sh
   git clone https://github.com/MadAssassin27/storage-service-api.git
   ```
3. Install Go packages
   ```sh
   go get github.com/dgrijalva/jwt-go
   ```
   ```sh
   go get github.com/go-playground/validator/v10
   ```
   ```sh
   go get github.com/gofiber/fiber/v2
   ```
   ```sh
   go get gorm.io/driver/sqlite
   ```
   ```sh
   go get gorm.io/gorm
   ```
4. To simplfy the interview process I use sqlite for the database, so we need to have sqlite installed on our local system. In real development practice, I have no issue using more versatile database such as postgre or mysql
5. For testing I use Postman, you can use both the web version or the desktop version

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- Run Go API -->
## Run Go API

It’s time for us to run the API

1.	Open your terminal and execute:
```sh
go run main.go
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Set up Postman -->
## Set up Postman
You can set your own Postman collection or you can access mine in the link below:
```sh
https://www.postman.com/material-observer-24642556/workspace/madassassin-workspace/collection/19104806-775e2dcf-c356-4d59-86f0-949f65c29ce3?action=share&creator=19104806
```


1.	Open Postman and create a new collection
2.	Add four request <br/>
    •	Login: set to POST <br/>
    •	Upload: set to Post <br/>
    •	File List: set to GET <br/>
    •	Get All User: set to GET <br/>

3.	I’m running the API on port 8000, so we need to set the URL to: <br/>
    http://127.0.0.1:8000 <br/>
    With addition to 4 request: <br/>
      •	Login: http://127.0.0.1:8000/login <br/>
      •	Upload: http://127.0.0.1:8000/upload <br/>
      •	File List: http://127.0.0.1:8000/items <br/>
      •	Get All User: http://127.0.0.1:8000/users <br/>
    Notes: To simplify the routing process I don’t use API versioning, but in production I experienced on using API versioning

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- API Testing -->
## API Testing
After we set up postman and run the API from our terminal, it’s time to test this API

# Users
I have prepared 4 users for us to try <br/>
  •	Username: admin, Password: admin <br/>
    Role: Admin, Authorize to upload file and retrieve file list <br/>
    <br/>
  •	Username: admin1, Password: admin1 <br/>
    Role: Admin, Authorize to upload file and retrieve file list <br/>
    <br/>
  •	Username: user, Password: user, Role: user <br/>
  <br/>
  •	Username: user1, Password: user1, Role: user <br/>
  
# Login
1.	Open login request and enter JSON Body:
```sh
   {
    "username": "admin",
    "password": "admin"
}
```

2.	Send and we will get a token response
3.	Copy and paste the token as x-token value on the Headers KEY to other request to run the authorization mechanism

Do it to the other three request and now you are accessing the API as Admin

# Upload
This API has Upload function to upload image and video file

1.	Open upload request and now we send data via Form File <br/>
2.	Fill the name value and click select files to choose a file you want to upload <br/>
3.	Now click send and it should have a response like this <br/>
4.	We can open the picture/video by clicking the url to open it on postman, or you can copy the url and open it on your web browser <br/>

# File List
Yes we can open the file that we just uploaded, but some user can also retrieve the list of file uploaded by other user
1.	Open file list request and click send as admin, and we should get the file list


<p align="right">(<a href="#readme-top">back to top</a>)</p>
