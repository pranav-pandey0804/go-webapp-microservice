# Go Gin Web Application

This is a simple web application built using the [Gin](https://github.com/gin-gonic/gin) web framework in Go. The application displays a list of articles on the home page and supports responding with HTML, JSON, or XML based on the request headers.

## Project Structure

- `main.go`: The entry point of the application. It initializes the Gin router, loads the HTML templates, and defines the route for the index page.
- `models.article.go`: Defines the `article` struct and provides an in-memory list of articles along with a function to retrieve them.
- `handlers.article.go`: Contains the route handler function to display the list of articles on the index page and a render function to respond in different formats.
- `templates/`: Contains HTML template files for rendering the web pages.
  - `header.html`: The header section of the web pages, including the page title and Bootstrap CSS/JS.
  - `menu.html`: The navigation menu.
  - `footer.html`: The footer section of the web pages.
  - `index.html`: The main content of the home page, displaying the list of articles.
- `models.article_test.go`: Unit tests for the article model.
- `handlers.article_test.go`: Unit tests for the index route handler, including tests for JSON and XML responses.
- `common_test.go`: Helper functions for setting up and processing HTTP requests in tests.

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/pranav-pandey0804/go-webapp-microservice.git
    cd webapp-go-gin
    ```

2. **Install dependencies**:
    Ensure you have Go installed. If not, download and install it from [golang.org](https://golang.org/dl/).

3. **Run the application**:
    ```bash
    go run main.go
    ```

4. **Open your web browser** and navigate to `http://localhost:8080` to view the application.

## Running Tests

To run the unit tests, use the following command:
```bash
go test -v
```