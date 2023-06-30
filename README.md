# Flight Provider API

This is the Flight Provider API that allows you to manage flights, aircrafts, cities, and days with a flight.

## Table of Contents

- [Routes](#routes)
  - [Flights](#flights)
    - [`GET /flights`](#get-flights)
    - [`GET /flights/:id`](#get-flight)
    - [`PATCH /flights/:id`](#patch-flight)
  - [Aircrafts](#aircrafts)
    - [`GET /aircrafts`](#get-aircrafts)
  - [Cities](#cities)
    - [`GET /cities`](#get-cities)
  - [Days with Flight](#days-with-flight)
    - [`GET /days-with-flight`](#get-days-with-flight)
- [Libraries and Dependencies Used](#libraries-and-dependencies-used)
- [Installation](#installation)

## Routes

### Flights

#### `GET /flights`

Get a list of flights based on parameters

**Request:**

| Parameter     | Type   | Required | Description                  |
| ------------- | ------ | -------- | ---------------------------- |
| `source`      | string | yes      | The source city name         |
| `destination` | string | yes      | The destination city name    |
| `departing`   | string | yes      | The date of departure (YYYY-MM-DD) |

**Response (200 OK):**

```json
[
  {
    "id": "ead35522-0f32-4ad4-8e7a-8506440ed9d2",
    "flight_number": "TD339",
    "source": "tehran",
    "destination": "sari"
    // more details
  },
  // more flights
]
```

#### `GET /flights/:id`

Find a flight by id

**Request:**

| Parameter | Type  | Required | Description           |
| --------- | ----- | -------- | --------------------- |
| `id`      | uuid  | yes      | The ID of the flight to retrieve |

**Response (200 OK):**

```json
{
  "id": "ead35522-0f32-4ad4-8e7a-8506440ed9d2",
  "flight_number": "TD339",
  "source": "tehran",
  "destination": "sari"
  // more details
}
```

#### `PATCH /flights/:id`

Update a flight by id

**Request:**

| Parameter | Type   | Required | Description                                               |
| --------- | ------ | -------- | --------------------------------------------------------- |
| `id`      | uuid   | yes      | The ID of the flight to update                            |
| `action`  | string | yes      | The type of action. Possible values: `["reserv", "cancel"]` |
| `count`   | int    | yes      | The number of flights to reserve or cancel                 |

**Response (200 OK):**

```json
true
```

### Aircrafts

#### `GET /aircrafts`

Get a list of aircrafts

**Request:**

*No request parameters*

**Response (200 OK):**

```json
[
  "airbus a319",
  "airbus a321",
  "airbus a310"
  // more aircrafts
]
```

### Cities

#### `GET /cities`

Get a list of cities

**Request:**

*No request parameters*

**Response (200 OK):**

```json
[
  "tehran",
  "karaj",
  "ahvaz"
  // more cities
]
```

### Days with Flight

#### `GET /days-with-flight`

Get a list of days with a flight

**Request:**

*No request parameters*

**Response (200 OK):**

```json
[
  "2021-10-01",
  "2021-10-02"
  // more dates
]

## Libraries and Dependencies Used

The most important libraries and dependencies used in the `letsgo-flight-provider` module are:

- `github.com/labstack/echo/v4`: A web framework for building RESTful APIs.
- `gorm.io/gorm`: An ORM (Object Relational Mapping) library for working with databases.
- `gorm.io/driver/postgres`: The PostgreSQL driver for GORM.
- `github.com/google/uuid`: A library for generating UUIDs (Universally Unique Identifiers).

These libraries are essential for the main functionality of the module, which is to provide a flight provider service via a RESTful API. The `echo` framework is used to handle incoming HTTP requests and send back responses, while `gorm` and its PostgreSQL driver are used for database access and operations. Finally, the `uuid` library is used for generating unique identifiers for flights.

## Installation

To install and run the Flight Provider API, please follow these steps:

1. Make sure you have Go installed on your system.

2. Clone the repository:

3. Change into the project directory:
   ```
   cd letsgo-flight-provider
   ```

4. Install the required libraries and dependencies:
   ```
   go mod tidy
   ```

5. Rename the `.env_sample` file to `.env`, and update it with your database information.

6. Run the project using `go run`:
   ```
   go run .
   ```

7. The service will start on port 8080, and you can access it through `http://localhost:8080`.


## Database Initialization

The `flights.sql` file included in the project contains mock data and creates a table for storing flight information. However, please note that if the date of the data in this file is in the past, the application will not display the outdated information.

To initialize the database with the data from the `flights.sql` file, including deleting the existing `flights` table, replacing its data with the new data from the `flights.sql` file, and then running the program again, you can follow these steps:

1. Ensure that your PostgreSQL database is set up and running.

2. Open the `flights.sql` file in a text editor.

3. Delete the old SQL statements and replace them with the new SQL statements.

4. Connect to your PostgreSQL database using a PostgreSQL client (e.g., `psql` command line tool).

5. Execute the following command in the PostgreSQL client to delete the existing `flights` table:

   ```sql
   DROP TABLE IF EXISTS flights;
   ```

6. Run the program again.

Once you have completed these steps, your database will be initialized with the new data from the `flights.sql` file. The Flight Provider API will then be able to access and display the updated flight information stored in the table. You can run the program again to see the changes taking effect.