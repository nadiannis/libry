<div align="center">
  <br>
  <h1>Libry</h1>
  <p>📖 A simple library CLI app 📖</p>
  <br>
</div>

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Entities](#entities)
- [Usage](#usage)

## Description

[`^ back to top ^`](#table-of-contents)

**Libry** is a simple CLI app that allows people to borrow books from a library. User of the app can be created when borrowing a book. The app is made with Go. It is created as a submission for the first-week exam in the Go phase of the Backend Development Training.

## Features

[`^ back to top ^`](#table-of-contents)

- View list of users & their currently borrowed books.
- View list of books.
- View list of borrowed books.
- Borrow a book.
- Return a book.
- Show all commands.
- Quit the app.

## Entities

[`^ back to top ^`](#table-of-contents)

There are 4 entities: **Book**, **User**, **Borrow**, & **BorrowStatus**.

**Book**

- id: `string`
- title: `string`
- author: `string`

**User**

- id: `string`
- username: `string`
- books: `[]*Book`

**Borrow**

- id: `string`
- book_id: `string`
- user_id: `string`
- start_date: `timestamp`
- end_date: `timestamp`
- BorrowStatus

**BorrowStatus**

- status: `string`

## Usage

[`^ back to top ^`](#table-of-contents)

- Start the app.

  ```bash
  go run ./cmd
  ```

- Enter a command.

  Choose one of these commands:

  **View list of users & their currently borrowed books**

  ```bash
  \lu
  ```

  **View list of books**

  ```bash
  \lb
  ```

  **View list of borrowed books**

  ```bash
  \lbb
  ```

  **Borrow a book**

  ```bash
  \b
  ```

  **Return a book**

  ```bash
  \r
  ```

  **Show all commands**

  ```bash
  \c
  ```

  **Quit the app**

  ```bash
  \q
  ```
