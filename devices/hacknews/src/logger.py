from loguru import logger
import sys

# Define log format
LOG_FORMAT = "{time:YYYY-MM-DD HH:mm:ss} | {level} | {module}:{function}:{line} - {message}"

# Remove default logger configuration
logger.remove()

# Configure console logging (stdout and stderr)
logger.add(sys.stdout, level="DEBUG", format=LOG_FORMAT, colorize=True)
logger.add(sys.stderr, level="ERROR", format=LOG_FORMAT, colorize=True)

# Configure file logging with rotation (1 MB file size)
logger.add("logs/app.log", rotation="5 MB", level="DEBUG", format=LOG_FORMAT)

LOG = logger
log = logger

# Expose `logger` for direct use
__all__ = ["LOG", "log"]
