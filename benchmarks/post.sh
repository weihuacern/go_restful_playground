#!/bin/bash
echo "POST http://localhost:8100/api/v1/tasks/" | vegeta attack -body tasks.json -duration=5s -rate=10 | tee results-POST.bin | vegeta report
