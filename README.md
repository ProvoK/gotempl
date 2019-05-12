# Gotempl

## Rationale
Just an script for dealing with modular Dockerfile creation, in absence of a "INCLUDE" directive.

It's as of now a very untested, yolo-crafted script born in a rainy mood Sunday afternoon.
If proven useful, I'll be spent time to refine it.

## Usage


### Template file

Any text file containing go template syntax.
A special function `include` permits to specify as argument a path to a file that will be merged into the template.
#### Local path resolving

When a template file is passed, all `include` directives will be resolved based on the folder the template is located.

Instead, when reading from Stdin, all `include` directives will use the current directory for resolving relative paths.

### Options

```
-t/-template: path to the template file. If missing, reads from Stdin.
-o/-output: path for the rendered template. If missing, prints to Stdout.
-var=VARNAME=VARVALUE: variables that will be used when rendering the template. Optional. Default VARVALUE is an empty string.
```



### Example

Dockerfile
```
FROM python:{{ or .version "3.6-slim" }}
COPY . /app
RUN make /app

{{ include "./Dockerfile.deps" -}}

CMD python /app/app.py
```

Dockerfile.deps
```
RUN pip install -r requirements.txt
```

#### Basic
```
gotempl -t Dockerfile
```

is equivalent to

```
cat Dockerfile | gotempl
```

outputs
```
FROM python:3.6-slim
COPY . /app
RUN make /app

RUN pip install -r requirements.txt
CMD python /app/app.py
```

#### Vars

```
gotempl -t Dockerfile -var=version=3.7-slim
```
outputs
```
FROM python:3.7-slim
COPY . /app
RUN make /app

RUN pip install -r requirements.txt
CMD python /app/app.py
```


## Build
```
go build -ldflags "-s -w"

```
