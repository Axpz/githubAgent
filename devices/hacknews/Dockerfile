# 构建阶段
FROM python:3.12.6-slim AS builder

WORKDIR /app

# 安装必要的工具和依赖
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential git curl \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# 拷贝源码和安装 Python 依赖到自定义路径
COPY ./src /app
RUN pip install --no-cache-dir --target=/app/vendor -r /app/requirements.txt

# 运行阶段
FROM python:3.12.6-slim

WORKDIR /app

# 安装运行环境所需的 git 工具
RUN apt-get update && apt-get install -y --no-install-recommends git ssh \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# 复制源码和依赖
COPY --from=builder /app /app

# 设置 PYTHONPATH，确保可以加载自定义路径的包
ENV PYTHONPATH=/app/vendor

# 运行应用
CMD ["python", "main_hacknews_job.py"]

