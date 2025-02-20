from fastapi import FastAPI, Request

app = FastAPI()


@app.middleware("http")
async def log_requests(request: Request, call_next):
    print(f"Received request: {request.method} {request.url}")
    print(f"Headers: {request.headers}")
    print(f"Query: {request.query_params}")
    body = await request.body()
    print(f"Body: {body}")
    response = await call_next(request)
    return response


@app.get("/")
@app.post("/")
async def read_root():
    return {"message": "Request received and logged!"}


if __name__ == '__main__':
    import uvicorn

    uvicorn.run(app, host="0.0.0.0", port=5000)
