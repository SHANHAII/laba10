import uvicorn
from fastapi import FastAPI
import asyncio
from contextlib import asynccontextmanager

@asynccontextmanager
async def lifespan(app: FastAPI):
    print("Приложение стартовало")
    yield
    print("Выполняется Graceful Shutdown...")
    await asyncio.sleep(1) # Имитация закрытия соединений
    print("Приложение остановлено")

app = FastAPI(lifespan=lifespan)

@app.get("/")
async def root():
    return {"status": "Python service is live"}

if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=8083)