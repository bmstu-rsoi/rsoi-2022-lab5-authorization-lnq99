#!/usr/bin/env bash
set -e

export VARIANT="v1"
export SCRIPT_PATH=/docker-entrypoint-initdb.d/
export PGPASSWORD=postgres
psql -f "$SCRIPT_PATH/scripts/db-$VARIANT.sql"

psql -U program -d postgres -f "$SCRIPT_PATH/scripts/db-$VARIANT-create-tables-and-seed.sql"
