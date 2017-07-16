#!/bin/sh

curl -H "Authorization: Bearer $WERCKER_TOKEN" \
     -H "Content-Type: application/json" \
     -d "{\"pipelineId\": \"$PIPELINE_ID\", \"branch\": \"werckerize\", \"envVars\": [{\"key\": \"TGM_URL\", \"value\": \"$1\"}]}" \
     https://app.wercker.com/api/v3/runs


     #-d "{\"pipelineId\": \"$PIPELINE_ID\", \"branch\": \"werckerize\", \"envVars\": {\"TGM_URL\": \"$1\"}}" \
     #-d "{\"pipelineId\": \"$PIPELINE_ID\", \"branch\": \"werckerize\", \"envVars\": [{\"key\": \"TGM_URL\", \"value\": \"$1\"}]}" \
