#!/bin/bash

URL=$1
echo URL=$URL
NAME=${URL##*/}
NAME=${NAME/\#file-1_archive-edn/}
echo NAME=$NAME

curl -H "Authorization: Bearer $WERCKER_TOKEN" \
     -H "Content-Type: application/json" \
     -d "{\"pipelineId\": \"$PIPELINE_ID\", \"branch\": \"werckerize\", \"message\": \"Triggering build for $NAME\", \"envVars\": [{\"key\": \"TGM_URL\", \"value\": \"$URL\"}, {\"key\": \"TGM_NAME\", \"value\": \"$NAME\"}]}" \
     https://app.wercker.com/api/v3/runs


     #-d "{\"pipelineId\": \"$PIPELINE_ID\", \"branch\": \"werckerize\", \"envVars\": {\"TGM_URL\": \"$1\"}}" \
     #-d "{\"pipelineId\": \"$PIPELINE_ID\", \"branch\": \"werckerize\", \"envVars\": [{\"key\": \"TGM_URL\", \"value\": \"$1\"}]}" \
