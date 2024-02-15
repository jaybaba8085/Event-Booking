# Event Booking - APIs [ Go Lang Project]
# Creator- Sharginjay Patel
This document serves as a guide for the GO Land Event Booking API, a REST API built with the Go programming language. 

## Project Description

The GO Land Event Booking API facilitates event booking by offering functionalities for:

* **Viewing a list of available events**
* **Creating new events** (requires authentication)
* **Updating existing events** (requires authentication and can only be done by the event creator)
* **Deleting events** (requires authentication and can only be done by the event creator)
* **Creating new users**
* **Authenticating users** and generating JSON Web Tokens (JWTs) for subsequent requests
* **Registering users for events** (requires authentication)
* **Cancelling registrations** (requires authentication)

## API Endpoints

| Method | Endpoint | Description | Authentication |
|---|---|---|---|
| GET | /events | Retrieves a list of all available events | Not required |
| GET | /events/<id> | Retrieves details of a specific event | Not required |
| POST | /events | Creates a new event | Required |
| PUT | /events/<id> | Updates an existing event | Required (Only the event creator can update the event) |
| DELETE | /events/<id> | Deletes an event | Required (Only the event creator can delete the event) |
| POST | /signup | Creates a new user | Not required |
| POST | /login | Authenticates a user and generates a JWT | Not required |
| POST | /events/<id>/register | Registers a user for an event | Required |
| DELETE | /events/<id>/register | Cancels a user's registration for an event | Required |

## Authentication

Endpoints requiring authentication use JWTs for authorization. Users can obtain a JWT by logging in via the `/login` endpoint. The obtained JWT must be included in the Authorization header of subsequent requests that require authentication.

## Running the API

To run the API locally, you will need Go installed on your system. Once you have Go installed, you can clone this repository and run the following command in your terminal:

