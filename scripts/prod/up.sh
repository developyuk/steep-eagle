#!/usr/bin/env bash
set -x

MYPWD=$(pwd)

cd $MYPWD/web/ops
npm run build
cd $MYPWD/web/tutor
npm run build
mv $MYPWD/api-v2/app.yaml $MYPWD/api-v2/app/
mv $MYPWD/api-v2/requirements.txt $MYPWD/api-v2/app/

cd $MYPWD
gcloud app deploy $MYPWD/web/ops/app.yaml $MYPWD/web/tutor/app.yaml $MYPWD/api-v2/app/app.yaml