# HomeList

An experimental list application with Go and HTMX

## Getting Started

### Development Environment

This repo uses Dev Containers to make development faster to set up. You'll need
the following:

- Docker (Docker Desktop is generally the fastest but also can be buggy)
- Docker Compose (Comes preloaded with Docker Desktop)
- Visual Studio Code (or another editor compatible with `.devcontainer`)
- The Dev Containers Remote Extension for VS Code

Once you have your prerequisites setup, you can open the command palette (Ctrl+Shift+P
or Cmd+Shift+P) and choose `Dev Containers: Clone Repository in Container Volume`.
Then use the Github integration and select this repository (`superlinkx/HomeList`).
Dev Containers will walk you through the rest of the setup process and start cloning/building
the dev container.

### Running the app

This project uses `just` to make common tasks easier to navigate. You can see all
commands by running:

```sh
$ just
```

Run the application:

```sh
$ just run
```

Debug the application:

```sh
$ just debug
```

Change what hostURL to use:

```sh
$ just run -hostURL "0.0.0.0:8080"
```

## License

MIT License

Copyright (c) 2023 Alyx Holms

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
