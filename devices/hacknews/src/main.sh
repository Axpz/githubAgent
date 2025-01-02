#!/bin/bash

BASE_DIR=$(dirname $(readlink -f $0))
DAEMON_PATH="$BASE_DIR/main.py"
DAEMON_NAME="hacknewsAgent"
LOG_FILE="$BASE_DIR/logs/$DAEMON_NAME.log"
PID_FILE="$BASE_DIR/run/$DAEMON_NAME.pid"

start() {
    if [ -f $PID_FILE ]; then
        PID=$(cat $PID_FILE)
        if ps -p $PID > /dev/null; then
            echo "$DAEMON_NAME is already running."
            return
        fi
    fi

    echo "Starting $DAEMON_NAME..."
    mkdir -p $BASE_DIR/logs $BASE_DIR/run
    # nohup python3 $DAEMON_PATH > $LOG_FILE 2>&1 &
    python3 $DAEMON_PATH

    echo $! > $PID_FILE
    echo "$DAEMON_NAME started."
}

stop() {
    if [ -f $PID_FILE ]; then
        PID=$(cat $PID_FILE)
        echo "Stopping $DAEMON_NAME..."
        kill $PID
        sleep 2
        if ps -p $PID > /dev/null; then
            echo "Force killing $DAEMON_NAME..."
            kill -9 $PID
        fi
        rm $PID_FILE
        echo "$DAEMON_NAME stopped."
    else
        echo "$DAEMON_NAME is not running."
    fi
}

status() {
    if [ -f $PID_FILE ]; then
        PID=$(cat $PID_FILE)
        if ps -p $PID > /dev/null; then
            echo "$DAEMON_NAME is running."
        else
            echo "$DAEMON_NAME is not running."
        fi
    else
        echo "$DAEMON_NAME is not running."
    fi
}

case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    status)
        status
        ;;
    restart)
        stop
        start
        ;;
    *)
        echo "Usage: $0 {start|stop|status|restart}"
        exit 1
esac
