#!/bin/bash
psql -U admin -d superset -h localhost -p 5433 -f superset.meta.backup.sql