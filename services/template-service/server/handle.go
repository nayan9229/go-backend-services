package server

import (
	"net/http"

	"github.com/nayan9229/go-backend-services/services/template-service/model"
)

func (s *Server) HtmlHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	html := `<<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>API Server - Route Not Found</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f8f9fa;
				color: #343a40;
				text-align: center;
				padding: 50px;
			}
			.container {
				max-width: 600px;
				margin: 0 auto;
				background: white;
				padding: 20px;
				border-radius: 10px;
				box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			}
			h1 {
				font-size: 36px;
				margin-bottom: 20px;
			}
			p {
				font-size: 18px;
				margin-bottom: 30px;
			}
			.home-link {
				display: inline-block;
				padding: 10px 20px;
				font-size: 16px;
				color: white;
				background-color: #007bff;
				text-decoration: none;
				border-radius: 5px;
				transition: background-color 0.3s;
			}
			.home-link:hover {
				background-color: #0056b3;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Welcome to Our API Server</h1>
			<p>The route you have requested is not available at the moment. Please check the URL or refer to our API documentation for the correct endpoints.</p>
			<a href="/" class="home-link">Return to Homepage</a>
		</div>
	</body>
	</html>
	`
	return html, nil
}

func (s *Server) JsonHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	response := model.Response{
		Message: "API server is running good",
	}
	return response, nil
}
