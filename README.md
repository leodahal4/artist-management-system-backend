# Project Name

The Artist-Management-System-Backend project is a backend application designed
to manage artists and their music. It provides a platform where artists can be
registered and their information can be stored. The system includes a
superadmin role with full privileges, and additional users can register and
gain similar privileges after being verified by the existing superadmin. The
core models of the system are the 'Artist' model, which holds various details
about artists, and the 'Music' model, which stores information about the music
associated with each artist. This project forms the backend foundation for an
artist management system, with the frontend portion to be developed in a
separate repository.

## Table of Contents

-   [Overview](#overview)
-   [Installation](#installation)

## Overview

The Artist-Management-System-Backend project is a backend application that
serves as the foundation for an artist management system. It provides a
platform for managing artists and their music, allowing users to register,
store artist information, and manage music details associated with each artist.

The key features of the Artist-Management-System-Backend include:

-   **User Roles**: The system includes a superadmin role with full privileges,
    while other users can register and gain similar privileges after verification
    by the existing superadmin.

-   **Artist Model**: The application includes an 'Artist' model, which allows
    for storing and managing various details about artists such as their name,
    biography, contact information, and more.

-   **Music Model**: The 'Music' model enables the storage and management of
    music information associated with each artist, including details like song
    titles, albums, release dates, genres, and more.

By providing a robust backend infrastructure, this project aims to facilitate
the management of artists and their music, laying the groundwork for a
comprehensive artist management system. The backend will work in conjunction
with a separate frontend repository, which will provide the user interface for
interacting with the system.

## Installation

To install and run the Artist-Management-System-Backend project, follow these steps:

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/leodahal4/artist-management-system-backend.git
    ```

2. Navigate to the project directory:

```bash
cd Artist-Management-System-Backend
```

3. Modify the environment variables:

Rename the .env.example file to .env.
Open the .env file and update the configuration values according to your
environment. This may include database connection details, API keys, or any
other necessary settings.

4. CREATE the enums on the database
   For this you can run the file which is present inside the cli folder,

```bash
go run cmd/cli/enums.go
```

5. Start the server

Option 1: Run the pre-built binary (for Linux):
Inside the project folder, you can find a binary file,
Run the binary file, residing on the project root folder

Option 2: Build and run using Go:
Ensure you have Go (Golang) installed on your system.

Execute the following command to run the project:
`bash go run cmd/app/main.go `

The Artist-Management-System-Backend should now be up and running on your local machine.
