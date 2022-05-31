![sastodaraz logo](./frontend/icons/sastodaraz.png)

## What is Sasto Daraz?

Sasto Daraz is a browser extension that helps you find sastodeal alternatives for items on daraz shopping site.

You can download the extension for [microsoft edge](https://microsoftedge.microsoft.com/addons/detail/sastodaraz/iibdhkajglbdaaalmflbkfpdkjgllaek).

> This extension is inspired from [Pasaley](https://github.com/ch0c0l8ra1n/Pasaley) and borrows frontend code.

## How it works?

Everytime you visit a product page in [Daraz website](https://www.daraz.com.np/), the extension calls the backend API to fetch list of similar products and displays it.

The backend is a GO server that scrapes the [SastoDeal website](https://www.sastodeal.com/) and fetches list of alternatives.

## How to run locally?

- Go to browser extensions page (`edge://extensions` for Microsoft edge) and click on `Load Unpacked` button and select frontend folder.
- Open terminal on backend folder and run command `go run main.go`
- Visit [Daraz website](https://www.daraz.com.np/) and open any product page to see extension in action.

## Todo

- [ ] Write tests!
