#!/bin/bash
echo "POST http://localhost:8080/api/v1/tasks/" | vegeta attack -body tasks.json -duration=5s -rate=10 | tee results-POST.bin | vegeta report
echo "POST http://localhost:8080/api/v1/ddn/appportals" | vegeta attack -body appportals.json -duration=5s -rate=10 | tee results-POST.bin | vegeta report
