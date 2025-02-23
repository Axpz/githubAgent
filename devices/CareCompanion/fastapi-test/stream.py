from fastapi import FastAPI
from fastapi.responses import StreamingResponse
import time

app = FastAPI()


async def fake_video_streamer():
    for i in range(20):
        time.sleep(1)
        yield b"some fake video bytes\n"


@app.get("/")
async def main():
    return StreamingResponse(fake_video_streamer())