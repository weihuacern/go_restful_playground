#!/bin/bash
#echo "POST http://localhost:8100/api/v1/tasks/" | vegeta attack -body task.json -duration=5s -rate=10 | tee results-POST.bin | vegeta report
echo "POST http://192.168.7.140:8100/api/v1/login/" | vegeta attack -body login.json -duration=2s -rate=5 | tee results-POST.bin | vegeta report
