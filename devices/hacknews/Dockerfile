FROM python:3.12-slim

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY ./src /app

RUN pip install --no-cache-dir -r requirements.txt

CMD ["python", "main_hacknews_job.py"]
