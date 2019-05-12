FROM python:{{ or .version "3.6-slim" }}
COPY . /app
RUN make /app

{{ include "./Dockerfile.deps" -}}

CMD python /app/app.py
