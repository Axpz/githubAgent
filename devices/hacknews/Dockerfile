FROM python:3.12.6

WORKDIR /app

COPY ./src /app

RUN apt-get update && apt-get install -y --no-install-recommends build-essential vim curl \
    && pip install --no-cache-dir -r /app/requirements.txt \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

CMD ["python", "main_hacknews_job.py"]

