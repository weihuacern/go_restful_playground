#!/bin/bash
#echo "POST http://localhost:8100/api/v1/tasks/" | vegeta attack -body task.json -duration=5s -rate=10 | tee results-POST.bin | vegeta report
echo "POST http://localhost:8100/api/v1/login/" | vegeta attack -body login.json -duration=5s -rate=1 | tee results-POST.bin | vegeta report
