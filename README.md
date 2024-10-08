#### Setup and Run the Application

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/keyur0811/go-library-api.git
   cd go-library-api
   ```

2. **Install Dependencies** (if using external libraries like `gorilla/mux`):
   ```bash
   go get -u github.com/gorilla/mux
   ```

3. **Run the Application**:
   ```bash
   go run main.go
   ```

4. **Test the API Endpoints using PowerShell**:

   - **Create a new book (POST request)**:
     ```powershell
     Invoke-RestMethod -Uri http://localhost:8080/books -Method POST -ContentType "application/json" -Body '{"title": "Go Programming", "author": "John Doe"}'
     ```

   - **Get all books (GET request)**:
     ```powershell
     Invoke-RestMethod -Uri http://localhost:8080/books -Method GET
     ```

   - **Get a book by ID (GET request)**:
     ```powershell
     Invoke-RestMethod -Uri http://localhost:8080/books/1 -Method GET
     ```

   - **Update a book (PUT request)**:
     ```powershell
     Invoke-RestMethod -Uri http://localhost:8080/books/1 -Method PUT -ContentType "application/json" -Body '{"title": "Advanced Go Programming", "author": "John Doe", "status": "available"}'
     ```

   - **Delete a book (DELETE request)**:
     ```powershell
     Invoke-RestMethod -Uri http://localhost:8080/books/1 -Method DELETE
     ```
