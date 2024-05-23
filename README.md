# Pokemon Parser

This project is a service for parsing information about Pokémon from a website and providing it in JSON format through an API.

## Installation

1. Clone the repository to your computer:

```bash
git clone https://github.com/senitapqan/ArchivePokemonScraping.git
```

2. Navigate to the project directory:

```bash
cd ArchivePokemonScraping
```

3. Install dependencies:

```bash
go mod tidy
```

## Usage

### Running the Server

Run the server:
```bash
go run cmd/main.go
```

By default, the server will be available at http://localhost:8080.

### Getting a List of All Pokémon

To get a list of all Pokémon, send a GET request to /all:

```http
GET http://localhost:8080/all
```

## Example Response

```json
{
  Data": [
    {
      "name": "Bulbasaur",
      "price": "£63.00£",
      "image_url": "https://scrapeme.live/wp-content/uploads/2018/08/001-350x350.png"
    },
    {
      "name": "Ivysaur",
      "price": "£87.00£",
      "image_url": "https://scrapeme.live/wp-content/uploads/2018/08/002-350x350.png"
    },
    ...
  ]
}
```
