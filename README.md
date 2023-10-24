# My First Golang Project

This project allows you to create a basic device management application using Go (Golang). The application can be used to register, update, delete, and view devices.

## Requirements

- Go (Golang) should be installed.
- Gorm and Gin libraries should be installed.

## Installation

1. Clone this repository to your local directory.

   ```bash
   git clone https://github.com/muzeyr/go-lang-api-example
   ```

2. Navigate to the project directory.

   ```bash
   cd go-lang-api-example
   ```

3. Create the SQLite database.

   ```bash
   touch test.sqlite
   ```

4. Start the application.

   ```bash
   go run main.go
   ```

5. The application should now be running at [http://localhost:8088](http://localhost:8088).

## Endpoints

- `GET /devices`: To view all devices.
- `GET /devices/:id`: To view a specific device by ID.
- `POST /devices`: To create a new device.
- `PATCH /devices/:id`: To update a specific device.

## Example Usage

### Listing Devices

```bash
curl http://localhost:8088/devices
```

### Viewing a Specific Device

```bash
curl http://localhost:8088/devices/1
```

### Creating a New Device

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "New Device"}' http://localhost:8088/devices
```

### Updating a Device

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"title": "Updated Device"}' http://localhost:8088/devices/1
```