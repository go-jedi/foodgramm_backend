# API Documentation for Telegram Web Application

## Overview

This API provides functionality for a Telegram web application, allowing users to interact with various services such as authentication, user management, product management, recipe generation, and payment processing. The API is designed to be REST and follows the OpenAPI (Swagger) 2.0 specification.

### Base URL
The base URL for all API endpoints is:

`http://localhost:50050/v1`

### Authentication
Most endpoints require authentication via a Bearer token. The token should be included in the `Authorization` header as follows:

Authorization: Bearer <token>

### Endpoints

#### Health Check

`Ping`
* GET /ping
    * Description: Check server status.
    * Response:
        *    200 OK: Returns {"message": "pong"}.

### Authentication

#### Check User Token

* POST /v1/auth/check
    * Description: Check if the provided Telegram ID and token are valid.
    * Request Body: auth.CheckDTO
    * Response:
      * 200 OK: Returns auth.CheckResponse.
      * 400 Bad Request: Returns auth.ErrorResponse.
      * 500 Internal Server Error: Returns auth.ErrorResponse.

#### Refresh User Token

* POST /v1/auth/refresh
  * Description: Refresh the access token using the provided Telegram ID and refresh token.
  * Request Body: auth.RefreshDTO
  * Response:
    * 200 OK: Returns auth.RefreshResponse.
    * 400 Bad Request: Returns auth.ErrorResponse.
    * 500 Internal Server Error: Returns auth.ErrorResponse.

#### Sign In User

* POST /v1/auth/signin
  * Description: Sign in a user using their Telegram ID, username, first name, and last name.
  * Request Body: auth.SignInDTO
  * Response:
    * 200 OK: Returns auth.SignInResp.
    * 400 Bad Request: Returns auth.ErrorResponse.
    * 500 Internal Server Error: Returns auth.ErrorResponse.

### Payment

#### Create Payment Link

* POST /v1/payment/link
  * Description: Creates a payment link based on the provided Telegram ID and payment type.
  * Request Body: payment.CreateDTO
  * Response:
    * 200 OK: Returns the payment link as a string.
    * 400 Bad Request: Returns payment.ErrorResponse.
    * 404 Not Found: Returns payment.ErrorResponse.
    * 500 Internal Server Error: Returns payment.ErrorResponse.

#### Check Payment Status via WebSocket

* GET /ws/v1/payment/check/{telegramID}
  * Description: Establishes a WebSocket connection to check the payment status for a given Telegram ID.
  * Response:
    * 200 OK: Returns a boolean indicating the payment status.
    * 400 Bad Request: Returns payment.ErrorResponse.
    * 500 Internal Server Error: Returns payment.ErrorResponse.

### Product

#### Add Excluded Products by Telegram ID

* POST /v1/product/exclude/telegram/id
  * Description: Add excluded products for a user by their unique Telegram ID.
  * Request Body: product.AddExcludeProductsByTelegramIDDTO
  * Response:
    * 200 OK: Returns product.UserExcludedProducts.
    * 400 Bad Request: Returns product.ErrorResponse.
    * 500 Internal Server Error: Returns product.ErrorResponse.

#### Get Excluded Products by Telegram ID

* GET /v1/product/exclude/telegram/{telegramID}
  * Description: Get excluded products for a user by their unique Telegram ID.
  * Response:
    * 200 OK: Returns product.UserExcludedProducts.
    * 400 Bad Request: Returns product.ErrorResponse.
    * 500 Internal Server Error: Returns product.ErrorResponse.

#### Delete Excluded Products by Telegram ID

* DELETE /v1/product/exclude/telegram/{telegramID}
  * Description: Delete excluded products for a user identified by their Telegram ID.
  * Response:
    * 200 OK: Returns product.UserExcludedProducts.
    * 400 Bad Request: Returns product.ErrorResponse.
    * 500 Internal Server Error: Returns product.ErrorResponse.

### Recipe

#### Get Free Recipes by Telegram ID

* GET /v1/recipe/free/telegram/{telegramID}
  * Description: Retrieve the free recipes information for a user identified by their Telegram ID.
  * Response:
    * 200 OK: Returns recipe.UserFreeRecipes.
    * 400 Bad Request: Returns recipe.ErrorResponse.
    * 500 Internal Server Error: Returns recipe.ErrorResponse.

#### Generate a New Recipe

* POST /v1/recipe/generate
  * Description: Generates a new recipe based on the provided parameters.
  * Request Body: recipe.GenerateRecipeDTO
  * Response:
    * 200 OK: Returns recipe.Recipes.
    * 400 Bad Request: Returns recipe.ErrorResponse.
    * 404 Not Found: Returns recipe.ErrorResponse.
    * 500 Internal Server Error: Returns recipe.ErrorResponse.

#### Get Recipes by Telegram ID

* GET /v1/recipe/telegram/{telegramID}
  * Description: Retrieve recipes for a user by their Telegram ID.
  * Response:
    * 200 OK: Returns an array of recipe.Recipes.
    * 400 Bad Request: Returns recipe.ErrorResponse.
    * 404 Not Found: Returns recipe.ErrorResponse.
    * 500 Internal Server Error: Returns recipe.ErrorResponse.

#### Get Recipe of the Day

* GET /v1/recipe_of_days
  * Description: Retrieve the current recipe of the day.
  * Response:
    * 200 OK: Returns recipeofdays.Recipe.
    * 500 Internal Server Error: Returns recipeofdays.ErrorResponse.

### Subscription

#### Check if Subscription Exists by Telegram ID

* GET /v1/subscription/exists/telegram/{telegramID}
  * Description: Checks whether a subscription exists for a user with the specified Telegram ID.
  * Response:
    * 200 OK: Returns a boolean indicating if the subscription exists.
    * 400 Bad Request: Returns subscription.ErrorResponse.
    * 404 Not Found: Returns subscription.ErrorResponse.
    * 500 Internal Server Error: Returns subscription.ErrorResponse.

#### Get Subscription by Telegram ID

* GET /v1/subscription/telegram/{telegramID}
  * Description: Retrieves subscription details for a user with the specified Telegram ID.
  * Response:
    * 200 OK: Returns subscription.Subscription.
    * 400 Bad Request: Returns subscription.ErrorResponse.
    * 404 Not Found: Returns subscription.ErrorResponse.
    * 500 Internal Server Error: Returns subscription.ErrorResponse.

### User

#### Create a New User

* POST /v1/user
  * Description: Create a new user with the provided details.
  * Request Body: user.CreateDTO
  * Response:
    * 201 Created: Returns user.User.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 409 Conflict: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Update a User

* PUT /v1/user
  * Description: Update a user's information.
  * Request Body: user.UpdateDTO
  * Response:
    * 200 OK: Returns user.User.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Get All Users

* GET /v1/user/all
  * Description: Retrieve a list of all users.
  * Response:
    * 200 OK: Returns an array of user.User.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Check if a User Exists

* POST /v1/user/exists
  * Description: Check if a user exists by Telegram ID and Username.
  * Request Body: user.ExistsDTO
  * Response:
    * 200 OK: Returns a boolean indicating if the user exists.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Check if a User Exists by ID

* GET /v1/user/exists/id/{userID}
  * Description: Check if a user exists by their unique identifier.
  * Response:
    * 200 OK: Returns a boolean indicating if the user exists.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Check if a User Exists by Telegram ID

* GET /v1/user/exists/telegram/{telegramID}
  * Description: Check if a user exists by their unique Telegram ID.
  * Response:
    * 200 OK: Returns a boolean indicating if the user exists.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Get a User by ID

* GET /v1/user/id/{userID}
  * Description: Retrieve a user by their unique identifier.
    * Response:
      * 200 OK: Returns user.User.
      * 400 Bad Request: Returns user.ErrorResponse.
      * 500 Internal Server Error: Returns user.ErrorResponse.

#### Delete a User by ID

* DELETE /v1/user/id/{userID}
  * Description: Delete a user by their unique identifier.
  * Response:
    * 200 OK: Returns the ID of the deleted user.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Get a User by Telegram ID

* GET /v1/user/telegram/{telegramID}
  * Description: Retrieve a user by their unique Telegram ID.
  * Response:
    * 200 OK: Returns user.User.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Delete a User by Telegram ID

* DELETE /v1/user/telegram/{telegramID}
  * Description: Delete a user by their unique Telegram ID.
  * Response:
    * 200 OK: Returns the deleted Telegram ID.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.

#### Get a List of Users with Pagination

* POST /v1/user/list
  * Description: Retrieve a list of users with pagination based on page and size.
  * Request Body: user.ListDTO
  * Response:
    * 200 OK: Returns user.ListResponseSwagger.
    * 400 Bad Request: Returns user.ErrorResponse.
    * 500 Internal Server Error: Returns user.ErrorResponse.