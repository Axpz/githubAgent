from fastapi import FastAPI
from contextlib import asynccontextmanager
import logging


def fake_answer_to_everything_ml_model(x: float):
    return x * 42


ml_models = {}

# 配置日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


# 使用 asynccontextmanager 来创建 lifespan 管理器
@asynccontextmanager
async def lifespan(app: FastAPI):
    logger.info("Lifespan starting...")
    
    # 模拟加载 ML 模型
    logger.info("Loading ML models...")
    ml_models["answer_to_everything"] = fake_answer_to_everything_ml_model
    
    yield  # FastAPI 会在此点暂停，等待请求处理
    
    # yield 之后的代码
    logger.info("Cleaning up ML models...")
    ml_models.clear()  # 清理 ML 模型
    logger.info("Lifespan finished.")


app = FastAPI(lifespan=lifespan)


@app.get("/predict")
async def predict(x: float):
    result = ml_models["answer_to_everything"](x)
    return {"result": result}