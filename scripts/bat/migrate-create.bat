@echo off
if "%1"=="" (
    echo required parameter seq is missing
    echo Usage: migrate-create.bat ^<migration_name^>
    exit /b 1
)

docker compose run --rm affiliate-system-migrate create -ext sql -dir /backend/migrations -seq "%1"