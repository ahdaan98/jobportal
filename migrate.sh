#!/bin/bash

NAMESPACE="jobportal"
LOCAL_PORT=5433
POD_NAME=$(kubectl get pods -n $NAMESPACE -l database=postgres -o jsonpath="{.items[0].metadata.name}")

start_port_forward() {
    kubectl port-forward -n $NAMESPACE $POD_NAME $LOCAL_PORT:5432 &
    PF_PID=$!
    sleep 5
}

stop_port_forward() {
    kill $PF_PID
}

if [ $# -ne 1 ]; then
    echo "Usage: $0 [up|down]"
    exit 1
fi

start_port_forward

if [ "$1" == "up" ]; then
    migrate -database "postgres://postgres:2211@localhost:$LOCAL_PORT/jobportal?sslmode=disable" -path db/migrations/ up
elif [ "$1" == "down" ]; then
    migrate -database "postgres://postgres:2211@localhost:$LOCAL_PORT/jobportal?sslmode=disable" -path db/migrations/ down
else
    echo "Invalid option: $1. Use 'up' or 'down'."
    stop_port_forward
    exit 1
fi

stop_port_forward

echo "Migration $1 completed."