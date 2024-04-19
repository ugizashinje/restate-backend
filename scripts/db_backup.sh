#!/bin/bash
pg_dump -U admin -d superset -h localhost -p 5433 >> superset.meta.backup.sql