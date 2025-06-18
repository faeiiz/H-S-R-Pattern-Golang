demo/
│
├── cmd/
│   └── main.go                  # App entry point
│
├── internal/
│   ├── handler/                 # HTTP Handlers (controllers)
│   ├── service/                 # Business logic layer
│   ├── repository/              # Data access layer
│   ├── model/                   # Domain models (structs)
│   ├── dto/                     # Request/Response types
│   └── server/                  # Bootstrap logic (e.g., routing)
│
├── test/
│   └── mock/                    # Mocks for testing
│
└── go.mod
